package Stipulations_Stipulations

import "bufio"
import "regexp"

// Autogenerated at , do not edit




type stipulations_Stipulations_RegexValidator struct {
    StipulationType *(regexp.Regexp)
    StipulationValue *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myStipulations_Stipulations_RegexValidator stipulations_Stipulations_RegexValidator

func init() {
    myStipulations_Stipulations_RegexValidator
    myStipulations_Stipulations_RegexValidator.StipulationType = regexp.MustCompile(`[^|]*`)
    myStipulations_Stipulations_RegexValidator.StipulationValue = regexp.MustCompile(`[^|]*`)
}



func (m *Stipulations_Stipulations) HasRequiredFields() bool {
    valid := true
    
    

    
    
    
    
    return valid
}




