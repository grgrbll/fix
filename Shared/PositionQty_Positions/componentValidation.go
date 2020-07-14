package PositionQty_Positions

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import NestedParties_NestedPartyIDs "grgrbll/fix/Shared/NestedParties_NestedPartyIDs"


type positionQty_Positions_RegexValidator struct {
    PosType *(regexp.Regexp)
    LongQty *(regexp.Regexp)
    ShortQty *(regexp.Regexp)
    PosQtyStatus *(regexp.Regexp)
    QuantityDate *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myPositionQty_Positions_RegexValidator positionQty_Positions_RegexValidator

func init() {
    myPositionQty_Positions_RegexValidator
    myPositionQty_Positions_RegexValidator.PosType = regexp.MustCompile(`[^|]*`)
    myPositionQty_Positions_RegexValidator.LongQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myPositionQty_Positions_RegexValidator.ShortQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myPositionQty_Positions_RegexValidator.PosQtyStatus = regexp.MustCompile(`-?\d+`)
    myPositionQty_Positions_RegexValidator.QuantityDate = regexp.MustCompile(`[0-9]{8}`)
}



func (m *PositionQty_Positions) HasRequiredFields() bool {
    valid := true
    
    

    
    
    for _, g := range(m.NestedPartyIDs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    
    
    return valid
}




