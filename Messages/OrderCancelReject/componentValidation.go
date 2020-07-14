package OrderCancelReject

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"


type orderCancelReject_RegexValidator struct {
    BeginString *(regexp.Regexp)
    BodyLength *(regexp.Regexp)
    MsgType *(regexp.Regexp)
    SenderCompID *(regexp.Regexp)
    TargetCompID *(regexp.Regexp)
    OnBehalfOfCompID *(regexp.Regexp)
    DeliverToCompID *(regexp.Regexp)
    SecureDataLen *(regexp.Regexp)
    SecureData *(regexp.Regexp)
    MsgSeqNum *(regexp.Regexp)
    SenderSubID *(regexp.Regexp)
    SenderLocationID *(regexp.Regexp)
    TargetSubID *(regexp.Regexp)
    TargetLocationID *(regexp.Regexp)
    OnBehalfOfSubID *(regexp.Regexp)
    OnBehalfOfLocationID *(regexp.Regexp)
    DeliverToSubID *(regexp.Regexp)
    DeliverToLocationID *(regexp.Regexp)
    PossDupFlag *(regexp.Regexp)
    PossResend *(regexp.Regexp)
    SendingTime *(regexp.Regexp)
    OrigSendingTime *(regexp.Regexp)
    XmlDataLen *(regexp.Regexp)
    XmlData *(regexp.Regexp)
    MessageEncoding *(regexp.Regexp)
    LastMsgSeqNumProcessed *(regexp.Regexp)
    ApplVerID *(regexp.Regexp)
    CstmApplVerID *(regexp.Regexp)
    OrderID *(regexp.Regexp)
    SecondaryOrderID *(regexp.Regexp)
    SecondaryClOrdID *(regexp.Regexp)
    ClOrdID *(regexp.Regexp)
    ClOrdLinkID *(regexp.Regexp)
    OrigClOrdID *(regexp.Regexp)
    OrdStatus *(regexp.Regexp)
    WorkingIndicator *(regexp.Regexp)
    OrigOrdModTime *(regexp.Regexp)
    ListID *(regexp.Regexp)
    Account *(regexp.Regexp)
    AcctIDSource *(regexp.Regexp)
    AccountType *(regexp.Regexp)
    TradeOriginationDate *(regexp.Regexp)
    TradeDate *(regexp.Regexp)
    TransactTime *(regexp.Regexp)
    CxlRejResponseTo *(regexp.Regexp)
    CxlRejReason *(regexp.Regexp)
    Text *(regexp.Regexp)
    EncodedTextLen *(regexp.Regexp)
    EncodedText *(regexp.Regexp)
    SignatureLength *(regexp.Regexp)
    Signature *(regexp.Regexp)
    CheckSum *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myOrderCancelReject_RegexValidator orderCancelReject_RegexValidator

func init() {
    myOrderCancelReject_RegexValidator
    myOrderCancelReject_RegexValidator.BeginString = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.BodyLength = regexp.MustCompile(`\d+`)
    myOrderCancelReject_RegexValidator.MsgType = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.SenderCompID = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.TargetCompID = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.OnBehalfOfCompID = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.DeliverToCompID = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.SecureDataLen = regexp.MustCompile(`\d+`)
    myOrderCancelReject_RegexValidator.SecureData = regexp.MustCompile(`.*`)
    myOrderCancelReject_RegexValidator.MsgSeqNum = regexp.MustCompile(`\d+`)
    myOrderCancelReject_RegexValidator.SenderSubID = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.SenderLocationID = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.TargetSubID = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.TargetLocationID = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.OnBehalfOfSubID = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.OnBehalfOfLocationID = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.DeliverToSubID = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.DeliverToLocationID = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.PossDupFlag = regexp.MustCompile(`[YN]`)
    myOrderCancelReject_RegexValidator.PossResend = regexp.MustCompile(`[YN]`)
    myOrderCancelReject_RegexValidator.SendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myOrderCancelReject_RegexValidator.OrigSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myOrderCancelReject_RegexValidator.XmlDataLen = regexp.MustCompile(`\d+`)
    myOrderCancelReject_RegexValidator.XmlData = regexp.MustCompile(`.*`)
    myOrderCancelReject_RegexValidator.MessageEncoding = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.LastMsgSeqNumProcessed = regexp.MustCompile(`\d+`)
    myOrderCancelReject_RegexValidator.ApplVerID = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.CstmApplVerID = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.OrderID = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.SecondaryOrderID = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.SecondaryClOrdID = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.ClOrdID = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.ClOrdLinkID = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.OrigClOrdID = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.OrdStatus = regexp.MustCompile(`[^|]`)
    myOrderCancelReject_RegexValidator.WorkingIndicator = regexp.MustCompile(`[YN]`)
    myOrderCancelReject_RegexValidator.OrigOrdModTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myOrderCancelReject_RegexValidator.ListID = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.Account = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.AcctIDSource = regexp.MustCompile(`-?\d+`)
    myOrderCancelReject_RegexValidator.AccountType = regexp.MustCompile(`-?\d+`)
    myOrderCancelReject_RegexValidator.TradeOriginationDate = regexp.MustCompile(`[0-9]{8}`)
    myOrderCancelReject_RegexValidator.TradeDate = regexp.MustCompile(`[0-9]{8}`)
    myOrderCancelReject_RegexValidator.TransactTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myOrderCancelReject_RegexValidator.CxlRejResponseTo = regexp.MustCompile(`[^|]`)
    myOrderCancelReject_RegexValidator.CxlRejReason = regexp.MustCompile(`-?\d+`)
    myOrderCancelReject_RegexValidator.Text = regexp.MustCompile(`[^|]*`)
    myOrderCancelReject_RegexValidator.EncodedTextLen = regexp.MustCompile(`\d+`)
    myOrderCancelReject_RegexValidator.EncodedText = regexp.MustCompile(`.*`)
    myOrderCancelReject_RegexValidator.SignatureLength = regexp.MustCompile(`\d+`)
    myOrderCancelReject_RegexValidator.Signature = regexp.MustCompile(`.*`)
    myOrderCancelReject_RegexValidator.CheckSum = regexp.MustCompile(`[^|]*`)
}



func (m *OrderCancelReject) HasRequiredFields() bool {
    valid := true
    
    

    
    
    valid = valid && m.BeginString.HasValue()
    
    
    
    valid = valid && m.BodyLength.HasValue()
    
    
    
    valid = valid && m.MsgType.HasValue()
    
    
    
    valid = valid && m.SenderCompID.HasValue()
    
    
    
    valid = valid && m.TargetCompID.HasValue()
    
    
    
    valid = valid && m.MsgSeqNum.HasValue()
    
    
    
    valid = valid && m.SendingTime.HasValue()
    
    
    
    for _, g := range(m.Hops) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.OrderID.HasValue()
    
    
    
    valid = valid && m.ClOrdID.HasValue()
    
    
    
    valid = valid && m.OrigClOrdID.HasValue()
    
    
    
    valid = valid && m.OrdStatus.HasValue()
    
    
    
    valid = valid && m.CxlRejResponseTo.HasValue()
    
    
    
    valid = valid && m.CheckSum.HasValue()
    
    
    
    
    
    return valid
}




