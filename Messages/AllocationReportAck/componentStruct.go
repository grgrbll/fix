package AllocationReportAck

import "bufio"
import "grgrbll/Fix/FixDef/Types"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import Parties_PartyIDs "grgrbll/fix/Shared/Parties_PartyIDs"

import AllocAckGrp_Allocs "grgrbll/fix/Shared/AllocAckGrp_Allocs"


type AllocationReportAck struct {
    BeginString Types.String `id:"8", required:"Y" `
    BodyLength Types.Length `id:"9", required:"Y" `
    MsgType Types.String `id:"35", required:"Y" `
    SenderCompID Types.String `id:"49", required:"Y" `
    TargetCompID Types.String `id:"56", required:"Y" `
    OnBehalfOfCompID Types.String `id:"115", required:"N" `
    DeliverToCompID Types.String `id:"128", required:"N" `
    SecureDataLen Types.Length `id:"90", required:"N" `
    SecureData Types.Data `id:"91", required:"N" `
    MsgSeqNum Types.SeqNum `id:"34", required:"Y" `
    SenderSubID Types.String `id:"50", required:"N" `
    SenderLocationID Types.String `id:"142", required:"N" `
    TargetSubID Types.String `id:"57", required:"N" `
    TargetLocationID Types.String `id:"143", required:"N" `
    OnBehalfOfSubID Types.String `id:"116", required:"N" `
    OnBehalfOfLocationID Types.String `id:"144", required:"N" `
    DeliverToSubID Types.String `id:"129", required:"N" `
    DeliverToLocationID Types.String `id:"145", required:"N" `
    PossDupFlag Types.Boolean `id:"43", required:"N" `
    PossResend Types.Boolean `id:"97", required:"N" `
    SendingTime Types.UTCTimestamp `id:"52", required:"Y" `
    OrigSendingTime Types.UTCTimestamp `id:"122", required:"N" `
    XmlDataLen Types.Length `id:"212", required:"N" `
    XmlData Types.Data `id:"213", required:"N" `
    MessageEncoding Types.String `id:"347", required:"N" `
    LastMsgSeqNumProcessed Types.SeqNum `id:"369", required:"N" `
    Hops []HopGrp_Hops.HopGrp_Hops `counter_name:"NoHops", counter_id:"627"`
    ApplVerID Types.String `id:"1128", required:"N" `
    CstmApplVerID Types.String `id:"1129", required:"N" `
    AllocReportID Types.String `id:"755", required:"Y" `
    AllocID Types.String `id:"70", required:"Y" `
    PartyIDs []Parties_PartyIDs.Parties_PartyIDs `counter_name:"NoPartyIDs", counter_id:"453"`
    SecondaryAllocID Types.String `id:"793", required:"N" `
    TradeDate Types.LocalMktDate `id:"75", required:"N" `
    TransactTime Types.UTCTimestamp `id:"60", required:"N" `
    AllocStatus Types.Int `id:"87", required:"N" `
    AllocRejCode Types.Int `id:"88", required:"N" `
    AllocReportType Types.Int `id:"794", required:"N" `
    AllocIntermedReqType Types.Int `id:"808", required:"N" `
    MatchStatus Types.Char `id:"573", required:"N" `
    Product Types.Int `id:"460", required:"N" `
    SecurityType Types.String `id:"167", required:"N" `
    Text Types.String `id:"58", required:"N" `
    EncodedTextLen Types.Length `id:"354", required:"N" `
    EncodedText Types.Data `id:"355", required:"N" `
    Allocs []AllocAckGrp_Allocs.AllocAckGrp_Allocs `counter_name:"NoAllocs", counter_id:"78"`
    ClearingBusinessDate Types.LocalMktDate `id:"715", required:"N" `
    AvgPxIndicator Types.Int `id:"819", required:"N" `
    Quantity Types.Qty `id:"53", required:"N" `
    AllocTransType Types.Char `id:"71", required:"N" `
    SignatureLength Types.Length `id:"93", required:"N" `
    Signature Types.Data `id:"89", required:"N" `
    CheckSum Types.String `id:"10", required:"Y" `
    _controlBlock fixMessageControlBlock
}