package HopGrp_Hops

import "bufio"
import "regexp"

// Autogenerated at , do not edit




type hopGrp_Hops_RegexValidator struct {
    HopCompID *(regexp.Regexp)
    HopSendingTime *(regexp.Regexp)
    HopRefID *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myHopGrp_Hops_RegexValidator hopGrp_Hops_RegexValidator

func init() {
    myHopGrp_Hops_RegexValidator
    myHopGrp_Hops_RegexValidator.HopCompID = regexp.MustCompile(`[^|]*`)
    myHopGrp_Hops_RegexValidator.HopSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myHopGrp_Hops_RegexValidator.HopRefID = regexp.MustCompile(`\d+`)
}



func (m *HopGrp_Hops) HasRequiredFields() bool {
    valid := true
    
    

    
    
    
    
    return valid
}




