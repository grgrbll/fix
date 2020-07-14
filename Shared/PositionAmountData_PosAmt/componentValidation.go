package PositionAmountData_PosAmt

import "bufio"
import "regexp"

// Autogenerated at , do not edit




type positionAmountData_PosAmt_RegexValidator struct {
    PosAmtType *(regexp.Regexp)
    PosAmt *(regexp.Regexp)
    PositionCurrency *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myPositionAmountData_PosAmt_RegexValidator positionAmountData_PosAmt_RegexValidator

func init() {
    myPositionAmountData_PosAmt_RegexValidator
    myPositionAmountData_PosAmt_RegexValidator.PosAmtType = regexp.MustCompile(`[^|]*`)
    myPositionAmountData_PosAmt_RegexValidator.PosAmt = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myPositionAmountData_PosAmt_RegexValidator.PositionCurrency = regexp.MustCompile(`[^|]*`)
}



func (m *PositionAmountData_PosAmt) HasRequiredFields() bool {
    valid := true
    
    

    
    
    
    
    return valid
}



