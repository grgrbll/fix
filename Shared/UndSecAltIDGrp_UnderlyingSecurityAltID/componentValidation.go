package UndSecAltIDGrp_UnderlyingSecurityAltID

import "bufio"
import "regexp"

// Autogenerated at , do not edit




type undSecAltIDGrp_UnderlyingSecurityAltID_RegexValidator struct {
    UnderlyingSecurityAltID *(regexp.Regexp)
    UnderlyingSecurityAltIDSource *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myUndSecAltIDGrp_UnderlyingSecurityAltID_RegexValidator undSecAltIDGrp_UnderlyingSecurityAltID_RegexValidator

func init() {
    myUndSecAltIDGrp_UnderlyingSecurityAltID_RegexValidator
    myUndSecAltIDGrp_UnderlyingSecurityAltID_RegexValidator.UnderlyingSecurityAltID = regexp.MustCompile(`[^|]*`)
    myUndSecAltIDGrp_UnderlyingSecurityAltID_RegexValidator.UnderlyingSecurityAltIDSource = regexp.MustCompile(`[^|]*`)
}



func (m *UndSecAltIDGrp_UnderlyingSecurityAltID) HasRequiredFields() bool {
    valid := true
    
    

    
    
    
    
    return valid
}




