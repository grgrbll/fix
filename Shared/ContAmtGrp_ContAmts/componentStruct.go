package ContAmtGrp_ContAmts

import "bufio"
import "grgrbll/Fix/FixDef/Types"

// Autogenerated at , do not edit




type ContAmtGrp_ContAmts struct {
    ContAmtType Types.Int `id:"519", required:"N" `
    ContAmtValue Types.Float `id:"520", required:"N" `
    ContAmtCurr Types.Currency `id:"521", required:"N" `
    _controlBlock fixMessageControlBlock
}
