package CompIDReqGrp_CompIDs

import "bufio"
import "regexp"

// Autogenerated at , do not edit




type compIDReqGrp_CompIDs_RegexValidator struct {
    RefCompID *(regexp.Regexp)
    RefSubID *(regexp.Regexp)
    LocationID *(regexp.Regexp)
    DeskID *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myCompIDReqGrp_CompIDs_RegexValidator compIDReqGrp_CompIDs_RegexValidator

func init() {
    myCompIDReqGrp_CompIDs_RegexValidator
    myCompIDReqGrp_CompIDs_RegexValidator.RefCompID = regexp.MustCompile(`[^|]*`)
    myCompIDReqGrp_CompIDs_RegexValidator.RefSubID = regexp.MustCompile(`[^|]*`)
    myCompIDReqGrp_CompIDs_RegexValidator.LocationID = regexp.MustCompile(`[^|]*`)
    myCompIDReqGrp_CompIDs_RegexValidator.DeskID = regexp.MustCompile(`[^|]*`)
}



func (m *CompIDReqGrp_CompIDs) HasRequiredFields() bool {
    valid := true
    
    

    
    
    
    
    return valid
}




