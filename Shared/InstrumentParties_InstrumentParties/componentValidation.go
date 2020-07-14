package InstrumentParties_InstrumentParties

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import InstrumentPtysSubGrp_InstrumentPartySubIDs "grgrbll/fix/Shared/InstrumentPtysSubGrp_InstrumentPartySubIDs"


type instrumentParties_InstrumentParties_RegexValidator struct {
    InstrumentPartyID *(regexp.Regexp)
    InstrumentPartyIDSource *(regexp.Regexp)
    InstrumentPartyRole *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myInstrumentParties_InstrumentParties_RegexValidator instrumentParties_InstrumentParties_RegexValidator

func init() {
    myInstrumentParties_InstrumentParties_RegexValidator
    myInstrumentParties_InstrumentParties_RegexValidator.InstrumentPartyID = regexp.MustCompile(`[^|]*`)
    myInstrumentParties_InstrumentParties_RegexValidator.InstrumentPartyIDSource = regexp.MustCompile(`[^|]`)
    myInstrumentParties_InstrumentParties_RegexValidator.InstrumentPartyRole = regexp.MustCompile(`-?\d+`)
}



func (m *InstrumentParties_InstrumentParties) HasRequiredFields() bool {
    valid := true
    
    

    
    
    for _, g := range(m.InstrumentPartySubIDs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    
    
    return valid
}



