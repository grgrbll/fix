package TrdRegTimestamps_TrdRegTimestamps

import "bufio"
import "regexp"

// Autogenerated at , do not edit




type trdRegTimestamps_TrdRegTimestamps_RegexValidator struct {
    TrdRegTimestamp *(regexp.Regexp)
    TrdRegTimestampType *(regexp.Regexp)
    TrdRegTimestampOrigin *(regexp.Regexp)
    DeskType *(regexp.Regexp)
    DeskTypeSource *(regexp.Regexp)
    DeskOrderHandlingInst *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myTrdRegTimestamps_TrdRegTimestamps_RegexValidator trdRegTimestamps_TrdRegTimestamps_RegexValidator

func init() {
    myTrdRegTimestamps_TrdRegTimestamps_RegexValidator
    myTrdRegTimestamps_TrdRegTimestamps_RegexValidator.TrdRegTimestamp = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myTrdRegTimestamps_TrdRegTimestamps_RegexValidator.TrdRegTimestampType = regexp.MustCompile(`-?\d+`)
    myTrdRegTimestamps_TrdRegTimestamps_RegexValidator.TrdRegTimestampOrigin = regexp.MustCompile(`[^|]*`)
    myTrdRegTimestamps_TrdRegTimestamps_RegexValidator.DeskType = regexp.MustCompile(`[^|]*`)
    myTrdRegTimestamps_TrdRegTimestamps_RegexValidator.DeskTypeSource = regexp.MustCompile(`-?\d+`)
    myTrdRegTimestamps_TrdRegTimestamps_RegexValidator.DeskOrderHandlingInst = regexp.MustCompile(`[ ]?(\w+[ ])+\w+?`)
}



func (m *TrdRegTimestamps_TrdRegTimestamps) HasRequiredFields() bool {
    valid := true
    
    

    
    
    
    
    return valid
}



