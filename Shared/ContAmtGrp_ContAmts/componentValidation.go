package ContAmtGrp_ContAmts

import "bufio"
import "regexp"

// Autogenerated at , do not edit




type contAmtGrp_ContAmts_RegexValidator struct {
    ContAmtType *(regexp.Regexp)
    ContAmtValue *(regexp.Regexp)
    ContAmtCurr *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myContAmtGrp_ContAmts_RegexValidator contAmtGrp_ContAmts_RegexValidator

func init() {
    myContAmtGrp_ContAmts_RegexValidator
    myContAmtGrp_ContAmts_RegexValidator.ContAmtType = regexp.MustCompile(`-?\d+`)
    myContAmtGrp_ContAmts_RegexValidator.ContAmtValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myContAmtGrp_ContAmts_RegexValidator.ContAmtCurr = regexp.MustCompile(`[A-Z]{3}`)
}



func (m *ContAmtGrp_ContAmts) HasRequiredFields() bool {
    valid := true
    
    

    
    
    
    
    return valid
}



