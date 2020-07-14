import yaml
import jinja2 
import argparse
import os
import copy
from pathlib import Path
import shutil

parser = argparse.ArgumentParser(description="Convert a yaml fix messages file into golang code")
parser.add_argument("config_file", help="path to the config file")
parser.add_argument("input_folder", help="path to a folder of yaml fix definitions")
parser.add_argument("output", help="go file output")
args = parser.parse_args()

print("Generating files")

config = {}
with open(args.config_file,'r') as file:
    cfg = yaml.load(file, Loader=yaml.FullLoader)
    config.update(cfg)
    
createdirs = [config['message_output_package'],
              config['shared_component_output_package'],
              config['factory_output_package']]

for d in createdirs:
    shutil.rmtree(os.path.join(".", d), ignore_errors=True)
    Path(os.path.join(".", d)).mkdir()


    
definitions = {}
files = os.listdir(args.input_folder)
for filename in files:
    if filename == "config.yaml":
        continue
    if filename.endswith(".yaml"):
        with open(os.path.join(args.input_folder, filename),'r') as file:
            print(os.path.join(args.input_folder, filename))
            fd = yaml.load(file, Loader=yaml.FullLoader)
            
            definitions.update(fd)

def processRegex(r, substitutions):
    ret = r
    subStart = ret.find('{{')
    subEnd = (ret.find('}}', subStart + 2) if subStart != -1 else -1)
    while subStart != -1:
        subTag = ret[subStart+2:subEnd]
        subVal = substitutions.get(subTag,"")
        ret = ret[:subStart] + subVal + (ret[subEnd+2:] if len(ret) > subEnd+2 else "")
        subStart = ret.find('{{')
        subEnd = (ret.find('}}', subStart + 2) if subStart != -1 else -1)
    return ret
        
def isBasic(field):
    return field['TYPE'] == 'BASIC'
    
def isComponent(field):
    return field['TYPE'] == 'COMPONENT'
    
def isRepeat(field):
    return field['TYPE'] == 'REPEAT'
   
def isRequired(field):
    return field.get('required', False)
    
def toPrivate(s):
    return s[0].lower() + s[1:]

def isMessage(component):
    return 'message_type' in component and component['message_type'] != ''
    
processedComponents = set()

def processComponent(name):
    if name in processedComponents:
        return
    
    definitions[name]['validationGroups'] = []
    definitions[name]['component_imports'] = []
    msgType = definitions[name].get('message_type',"")

    definitions[name]['packagepath'] = ""
    
    if len(definitions[name].get('message_type',"")) > 0:
        definitions[name]['generate'] = True
        definitions[name]['packagepath'] = os.path.join(config['message_output_package'], name)
    else:
        definitions[name]['packagepath'] = os.path.join(config['shared_component_output_package'],  name)
    definitions[name]['import_path'] = os.path.join(config['module_name'], definitions[name]['packagepath'])
    
    newFields = {}
    for fieldName, fieldData in definitions[name]["fields"].items():
        if "component_type" in fieldData:
            processComponent(fieldData['component_type'])
            
            flatFields = copy.deepcopy(definitions[fieldData['component_type']]['fields'])
            definitions[name]['component_imports'].extend(definitions[fieldData['component_type']]['component_imports'])

            

            useValidationGroup = False
            if not isRequired(fieldData):
                for fn, fd in flatFields.items():
                    if isRequired(fd):
                        fd['validationGroup'] = fieldName
                        useValidationGroup = True
                        
            if useValidationGroup:
                definitions[name]["validationGroups"].append(fieldName)

            newFields.update(flatFields)
             
        elif "counter_name" in fieldData:
            newFields[fieldName] = fieldData
            newFields[fieldName]["TYPE"] = "REPEAT"
            
            processComponent(fieldData['repeating_component'])
            definitions[fieldData['repeating_component']]['generate'] = True
            importpath = definitions[fieldData["repeating_component"]]["import_path"]
            im = f'{ fieldData["repeating_component"] } "{ importpath }"' 
            definitions[name]["component_imports"].append(im)
            
            print(name, im)
        else:
            newFields[fieldName] = fieldData
            newFields[fieldName]["TYPE"] = "BASIC"
            
        if "validation_regex" in fieldData:
            fieldData['validation_regex'] = processRegex(fieldData['validation_regex'], config['fix_type_validation'])
            
    definitions[name]['fields'] = newFields
    
    processedComponents.add(name)
    
    
for name in definitions.keys():
    processComponent(name)

names = list(definitions.keys())
for name in names:
    if not definitions[name].get('generate', False):
        definitions.pop(name, None)
        

env = jinja2.Environment(loader=jinja2.FileSystemLoader(searchpath=os.path.join(".",config['template_input_dir'])))
env.globals['isBasic'] = isBasic
env.globals['isComponent'] = isComponent
env.globals.update(isRepeat=isRepeat)
env.globals['isRequired'] = isRequired
env.globals['toPrivate'] = toPrivate
env.globals['isMessage'] = isMessage

writepath = os.path.join(args.output, "base.go")
with open(writepath,"w+") as outputFile:
    tp = env.get_template("fixBase.txt")
    out = tp.render(definitions=definitions, config=config)
    outputFile.write(out)
    outputFile.close()
 
writepath = os.path.join(config['factory_output_package'], "factory.go")
with open(writepath,"w+") as outputFile:
    tp = env.get_template("factory.go")
    out = tp.render(definitions=definitions, config=config)
    print(out)
    outputFile.write(out)
    outputFile.close()
    
names = list(definitions.keys())


messageTemplates = ["componentMarshalling.go", "componentStruct.go","componentValidation.go" ]
for template in messageTemplates:
    tp = env.get_template(template)
    for name in names:
        if  definitions[name].get('generate', False):
                
            writepath = os.path.join(".", definitions[name]['packagepath'])
            Path(writepath).mkdir(parents=True, exist_ok=True)
            writepath = os.path.join(writepath, template)
            out = tp.render(definitions={name: definitions[name]}, config=config)
            
            package = f"package {name}"
            imports = 'import "bufio"\n' 
            if out.find("strconv") > 0:
                imports += 'import "strconv"\n'
            if out.find("Types") > 0:
                imports += 'import "grgrbll/Fix/FixDef/Types"\n'
            if out.find("regexp") > 0:
                imports += 'import "regexp"\n' 
            if out.find("fmt") > 0:
                imports += 'import "fmt"\n' 
            with open(writepath,"w+") as outputFile:
                outputFile.write(f"package {name}")
                outputFile.write(f"{package}\n\n{imports}\n")
                outputFile.write(out)
                outputFile.close()

    
    