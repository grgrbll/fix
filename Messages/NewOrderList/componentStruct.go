package NewOrderList

import "bufio"
import "grgrbll/Fix/FixDef/Types"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import ListOrdGrp_Orders "grgrbll/fix/Shared/ListOrdGrp_Orders"

import RootParties_RootPartyIDs "grgrbll/fix/Shared/RootParties_RootPartyIDs"


type NewOrderList struct {
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
    ListID Types.String `id:"66", required:"Y" `
    BidID Types.String `id:"390", required:"N" `
    ClientBidID Types.String `id:"391", required:"N" `
    ProgRptReqs Types.Int `id:"414", required:"N" `
    BidType Types.Int `id:"394", required:"Y" `
    ProgPeriodInterval Types.Int `id:"415", required:"N" `
    CancellationRights Types.Char `id:"480", required:"N" `
    MoneyLaunderingStatus Types.Char `id:"481", required:"N" `
    RegistID Types.String `id:"513", required:"N" `
    ListExecInstType Types.Char `id:"433", required:"N" `
    ListExecInst Types.String `id:"69", required:"N" `
    EncodedListExecInstLen Types.Length `id:"352", required:"N" `
    EncodedListExecInst Types.Data `id:"353", required:"N" `
    AllowableOneSidednessPct Types.Percentage `id:"765", required:"N" `
    AllowableOneSidednessValue Types.Amt `id:"766", required:"N" `
    AllowableOneSidednessCurr Types.Currency `id:"767", required:"N" `
    TotNoOrders Types.Int `id:"68", required:"Y" `
    LastFragment Types.Boolean `id:"893", required:"N" `
    Orders []ListOrdGrp_Orders.ListOrdGrp_Orders `counter_name:"NoOrders", counter_id:"73"`
    RootPartyIDs []RootParties_RootPartyIDs.RootParties_RootPartyIDs `counter_name:"NoRootPartyIDs", counter_id:"1116"`
    SignatureLength Types.Length `id:"93", required:"N" `
    Signature Types.Data `id:"89", required:"N" `
    CheckSum Types.String `id:"10", required:"Y" `
    _controlBlock fixMessageControlBlock
}
