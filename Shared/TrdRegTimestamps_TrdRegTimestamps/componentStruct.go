package TrdRegTimestamps_TrdRegTimestamps

import "bufio"
import "grgrbll/Fix/FixDef/Types"

// Autogenerated at , do not edit




type TrdRegTimestamps_TrdRegTimestamps struct {
    TrdRegTimestamp Types.UTCTimestamp `id:"769", required:"N" `
    TrdRegTimestampType Types.Int `id:"770", required:"N" `
    TrdRegTimestampOrigin Types.String `id:"771", required:"N" `
    DeskType Types.String `id:"1033", required:"N" `
    DeskTypeSource Types.Int `id:"1034", required:"N" `
    DeskOrderHandlingInst Types.MultipleStringValue `id:"1035", required:"N" `
    _controlBlock fixMessageControlBlock
}
