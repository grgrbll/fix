package CompIDStatGrp_CompIDs

import "bufio"
import "grgrbll/Fix/FixDef/Types"

// Autogenerated at , do not edit




type CompIDStatGrp_CompIDs struct {
    RefCompID Types.String `id:"930", required:"Y" `
    RefSubID Types.String `id:"931", required:"N" `
    LocationID Types.String `id:"283", required:"N" `
    DeskID Types.String `id:"284", required:"N" `
    StatusValue Types.Int `id:"928", required:"Y" `
    StatusText Types.String `id:"929", required:"N" `
    _controlBlock fixMessageControlBlock
}