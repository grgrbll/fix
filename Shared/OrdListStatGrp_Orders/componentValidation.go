package OrdListStatGrp_Orders

import "bufio"
import "regexp"

// Autogenerated at , do not edit




type ordListStatGrp_Orders_RegexValidator struct {
    ClOrdID *(regexp.Regexp)
    SecondaryClOrdID *(regexp.Regexp)
    CumQty *(regexp.Regexp)
    OrdStatus *(regexp.Regexp)
    WorkingIndicator *(regexp.Regexp)
    LeavesQty *(regexp.Regexp)
    CxlQty *(regexp.Regexp)
    AvgPx *(regexp.Regexp)
    OrdRejReason *(regexp.Regexp)
    Text *(regexp.Regexp)
    EncodedTextLen *(regexp.Regexp)
    EncodedText *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myOrdListStatGrp_Orders_RegexValidator ordListStatGrp_Orders_RegexValidator

func init() {
    myOrdListStatGrp_Orders_RegexValidator
    myOrdListStatGrp_Orders_RegexValidator.ClOrdID = regexp.MustCompile(`[^|]*`)
    myOrdListStatGrp_Orders_RegexValidator.SecondaryClOrdID = regexp.MustCompile(`[^|]*`)
    myOrdListStatGrp_Orders_RegexValidator.CumQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myOrdListStatGrp_Orders_RegexValidator.OrdStatus = regexp.MustCompile(`[^|]`)
    myOrdListStatGrp_Orders_RegexValidator.WorkingIndicator = regexp.MustCompile(`[YN]`)
    myOrdListStatGrp_Orders_RegexValidator.LeavesQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myOrdListStatGrp_Orders_RegexValidator.CxlQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myOrdListStatGrp_Orders_RegexValidator.AvgPx = regexp.MustCompile(``)
    myOrdListStatGrp_Orders_RegexValidator.OrdRejReason = regexp.MustCompile(`-?\d+`)
    myOrdListStatGrp_Orders_RegexValidator.Text = regexp.MustCompile(`[^|]*`)
    myOrdListStatGrp_Orders_RegexValidator.EncodedTextLen = regexp.MustCompile(`\d+`)
    myOrdListStatGrp_Orders_RegexValidator.EncodedText = regexp.MustCompile(`.*`)
}



func (m *OrdListStatGrp_Orders) HasRequiredFields() bool {
    valid := true
    
    

    
    
    valid = valid && m.ClOrdID.HasValue()
    
    
    
    valid = valid && m.CumQty.HasValue()
    
    
    
    valid = valid && m.OrdStatus.HasValue()
    
    
    
    valid = valid && m.LeavesQty.HasValue()
    
    
    
    valid = valid && m.CxlQty.HasValue()
    
    
    
    valid = valid && m.AvgPx.HasValue()
    
    
    
    
    
    return valid
}




