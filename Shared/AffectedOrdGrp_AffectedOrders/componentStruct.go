package AffectedOrdGrp_AffectedOrders

import "bufio"
import "grgrbll/Fix/FixDef/Types"

// Autogenerated at , do not edit




type AffectedOrdGrp_AffectedOrders struct {
    OrigClOrdID Types.String `id:"41", required:"N" `
    AffectedOrderID Types.String `id:"535", required:"N" `
    AffectedSecondaryOrderID Types.String `id:"536", required:"N" `
    _controlBlock fixMessageControlBlock
}
