package SettlPtysSubGrp_SettlPartySubIDs

import "bufio"
import "regexp"

// Autogenerated at , do not edit




type settlPtysSubGrp_SettlPartySubIDs_RegexValidator struct {
    SettlPartySubID *(regexp.Regexp)
    SettlPartySubIDType *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var mySettlPtysSubGrp_SettlPartySubIDs_RegexValidator settlPtysSubGrp_SettlPartySubIDs_RegexValidator

func init() {
    mySettlPtysSubGrp_SettlPartySubIDs_RegexValidator
    mySettlPtysSubGrp_SettlPartySubIDs_RegexValidator.SettlPartySubID = regexp.MustCompile(`[^|]*`)
    mySettlPtysSubGrp_SettlPartySubIDs_RegexValidator.SettlPartySubIDType = regexp.MustCompile(`-?\d+`)
}



func (m *SettlPtysSubGrp_SettlPartySubIDs) HasRequiredFields() bool {
    valid := true
    
    

    
    
    
    
    return valid
}




