package RootSubParties_RootPartySubIDs

import "bufio"
import "grgrbll/Fix/FixDef/Types"

// Autogenerated at , do not edit




type RootSubParties_RootPartySubIDs struct {
    RootPartySubID Types.String `id:"1121", required:"N" `
    RootPartySubIDType Types.Int `id:"1122", required:"N" `
    _controlBlock fixMessageControlBlock
}
