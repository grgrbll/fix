package MarketDataRequest

import "bufio"
import "grgrbll/Fix/FixDef/Types"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import MDReqGrp_MDEntryTypes "grgrbll/fix/Shared/MDReqGrp_MDEntryTypes"

import InstrmtMDReqGrp_RelatedSym "grgrbll/fix/Shared/InstrmtMDReqGrp_RelatedSym"

import TrdgSesGrp_TradingSessions "grgrbll/fix/Shared/TrdgSesGrp_TradingSessions"


type MarketDataRequest struct {
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
    MDReqID Types.String `id:"262", required:"Y" `
    SubscriptionRequestType Types.Char `id:"263", required:"Y" `
    MarketDepth Types.Int `id:"264", required:"Y" `
    MDUpdateType Types.Int `id:"265", required:"N" `
    AggregatedBook Types.Boolean `id:"266", required:"N" `
    OpenCloseSettlFlag Types.MultipleCharValue `id:"286", required:"N" `
    Scope Types.MultipleCharValue `id:"546", required:"N" `
    MDImplicitDelete Types.Boolean `id:"547", required:"N" `
    MDEntryTypes []MDReqGrp_MDEntryTypes.MDReqGrp_MDEntryTypes `counter_name:"NoMDEntryTypes", counter_id:"267"`
    RelatedSym []InstrmtMDReqGrp_RelatedSym.InstrmtMDReqGrp_RelatedSym `counter_name:"NoRelatedSym", counter_id:"146"`
    TradingSessions []TrdgSesGrp_TradingSessions.TrdgSesGrp_TradingSessions `counter_name:"NoTradingSessions", counter_id:"386"`
    ApplQueueAction Types.Int `id:"815", required:"N" `
    ApplQueueMax Types.Int `id:"812", required:"N" `
    MDQuoteType Types.Int `id:"1070", required:"N" `
    SignatureLength Types.Length `id:"93", required:"N" `
    Signature Types.Data `id:"89", required:"N" `
    CheckSum Types.String `id:"10", required:"Y" `
    _controlBlock fixMessageControlBlock
}
