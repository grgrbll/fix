package NestedParties2_Nested2PartyIDs

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import NstdPtys2SubGrp_Nested2PartySubIDs "grgrbll/fix/Shared/NstdPtys2SubGrp_Nested2PartySubIDs"


type nestedParties2_Nested2PartyIDs_RegexValidator struct {
    Nested2PartyID *(regexp.Regexp)
    Nested2PartyIDSource *(regexp.Regexp)
    Nested2PartyRole *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myNestedParties2_Nested2PartyIDs_RegexValidator nestedParties2_Nested2PartyIDs_RegexValidator

func init() {
    myNestedParties2_Nested2PartyIDs_RegexValidator
    myNestedParties2_Nested2PartyIDs_RegexValidator.Nested2PartyID = regexp.MustCompile(`[^|]*`)
    myNestedParties2_Nested2PartyIDs_RegexValidator.Nested2PartyIDSource = regexp.MustCompile(`[^|]`)
    myNestedParties2_Nested2PartyIDs_RegexValidator.Nested2PartyRole = regexp.MustCompile(`-?\d+`)
}



func (m *NestedParties2_Nested2PartyIDs) HasRequiredFields() bool {
    valid := true
    
    

    
    
    for _, g := range(m.Nested2PartySubIDs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    
    
    return valid
}




