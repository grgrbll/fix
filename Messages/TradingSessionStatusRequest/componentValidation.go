package TradingSessionStatusRequest

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"


type tradingSessionStatusRequest_RegexValidator struct {
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
    TradSesReqID *(regexp.Regexp)
    TradingSessionID *(regexp.Regexp)
    TradingSessionSubID *(regexp.Regexp)
    TradSesMethod *(regexp.Regexp)
    TradSesMode *(regexp.Regexp)
    SubscriptionRequestType *(regexp.Regexp)
    SecurityExchange *(regexp.Regexp)
    SignatureLength *(regexp.Regexp)
    Signature *(regexp.Regexp)
    CheckSum *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myTradingSessionStatusRequest_RegexValidator tradingSessionStatusRequest_RegexValidator

func init() {
    myTradingSessionStatusRequest_RegexValidator
    myTradingSessionStatusRequest_RegexValidator.BeginString = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatusRequest_RegexValidator.BodyLength = regexp.MustCompile(`\d+`)
    myTradingSessionStatusRequest_RegexValidator.MsgType = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatusRequest_RegexValidator.SenderCompID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatusRequest_RegexValidator.TargetCompID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatusRequest_RegexValidator.OnBehalfOfCompID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatusRequest_RegexValidator.DeliverToCompID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatusRequest_RegexValidator.SecureDataLen = regexp.MustCompile(`\d+`)
    myTradingSessionStatusRequest_RegexValidator.SecureData = regexp.MustCompile(`.*`)
    myTradingSessionStatusRequest_RegexValidator.MsgSeqNum = regexp.MustCompile(`\d+`)
    myTradingSessionStatusRequest_RegexValidator.SenderSubID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatusRequest_RegexValidator.SenderLocationID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatusRequest_RegexValidator.TargetSubID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatusRequest_RegexValidator.TargetLocationID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatusRequest_RegexValidator.OnBehalfOfSubID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatusRequest_RegexValidator.OnBehalfOfLocationID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatusRequest_RegexValidator.DeliverToSubID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatusRequest_RegexValidator.DeliverToLocationID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatusRequest_RegexValidator.PossDupFlag = regexp.MustCompile(`[YN]`)
    myTradingSessionStatusRequest_RegexValidator.PossResend = regexp.MustCompile(`[YN]`)
    myTradingSessionStatusRequest_RegexValidator.SendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myTradingSessionStatusRequest_RegexValidator.OrigSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myTradingSessionStatusRequest_RegexValidator.XmlDataLen = regexp.MustCompile(`\d+`)
    myTradingSessionStatusRequest_RegexValidator.XmlData = regexp.MustCompile(`.*`)
    myTradingSessionStatusRequest_RegexValidator.MessageEncoding = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatusRequest_RegexValidator.LastMsgSeqNumProcessed = regexp.MustCompile(`\d+`)
    myTradingSessionStatusRequest_RegexValidator.ApplVerID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatusRequest_RegexValidator.CstmApplVerID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatusRequest_RegexValidator.TradSesReqID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatusRequest_RegexValidator.TradingSessionID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatusRequest_RegexValidator.TradingSessionSubID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatusRequest_RegexValidator.TradSesMethod = regexp.MustCompile(`-?\d+`)
    myTradingSessionStatusRequest_RegexValidator.TradSesMode = regexp.MustCompile(`-?\d+`)
    myTradingSessionStatusRequest_RegexValidator.SubscriptionRequestType = regexp.MustCompile(`[^|]`)
    myTradingSessionStatusRequest_RegexValidator.SecurityExchange = regexp.MustCompile(`[A-Z0-9]{4}`)
    myTradingSessionStatusRequest_RegexValidator.SignatureLength = regexp.MustCompile(`\d+`)
    myTradingSessionStatusRequest_RegexValidator.Signature = regexp.MustCompile(`.*`)
    myTradingSessionStatusRequest_RegexValidator.CheckSum = regexp.MustCompile(`[^|]*`)
}



func (m *TradingSessionStatusRequest) HasRequiredFields() bool {
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
    
    
    
    valid = valid && m.TradSesReqID.HasValue()
    
    
    
    valid = valid && m.SubscriptionRequestType.HasValue()
    
    
    
    valid = valid && m.CheckSum.HasValue()
    
    
    
    
    
    return valid
}




