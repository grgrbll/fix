package MarketDataIncrementalRefresh

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import MDIncGrp_MDEntries "grgrbll/fix/Shared/MDIncGrp_MDEntries"

import RoutingGrp_RoutingIDs "grgrbll/fix/Shared/RoutingGrp_RoutingIDs"


type marketDataIncrementalRefresh_RegexValidator struct {
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
    MDReqID *(regexp.Regexp)
    ApplQueueDepth *(regexp.Regexp)
    ApplQueueResolution *(regexp.Regexp)
    MDBookType *(regexp.Regexp)
    MDFeedType *(regexp.Regexp)
    TradeDate *(regexp.Regexp)
    SignatureLength *(regexp.Regexp)
    Signature *(regexp.Regexp)
    CheckSum *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myMarketDataIncrementalRefresh_RegexValidator marketDataIncrementalRefresh_RegexValidator

func init() {
    myMarketDataIncrementalRefresh_RegexValidator
    myMarketDataIncrementalRefresh_RegexValidator.BeginString = regexp.MustCompile(`[^|]*`)
    myMarketDataIncrementalRefresh_RegexValidator.BodyLength = regexp.MustCompile(`\d+`)
    myMarketDataIncrementalRefresh_RegexValidator.MsgType = regexp.MustCompile(`[^|]*`)
    myMarketDataIncrementalRefresh_RegexValidator.SenderCompID = regexp.MustCompile(`[^|]*`)
    myMarketDataIncrementalRefresh_RegexValidator.TargetCompID = regexp.MustCompile(`[^|]*`)
    myMarketDataIncrementalRefresh_RegexValidator.OnBehalfOfCompID = regexp.MustCompile(`[^|]*`)
    myMarketDataIncrementalRefresh_RegexValidator.DeliverToCompID = regexp.MustCompile(`[^|]*`)
    myMarketDataIncrementalRefresh_RegexValidator.SecureDataLen = regexp.MustCompile(`\d+`)
    myMarketDataIncrementalRefresh_RegexValidator.SecureData = regexp.MustCompile(`.*`)
    myMarketDataIncrementalRefresh_RegexValidator.MsgSeqNum = regexp.MustCompile(`\d+`)
    myMarketDataIncrementalRefresh_RegexValidator.SenderSubID = regexp.MustCompile(`[^|]*`)
    myMarketDataIncrementalRefresh_RegexValidator.SenderLocationID = regexp.MustCompile(`[^|]*`)
    myMarketDataIncrementalRefresh_RegexValidator.TargetSubID = regexp.MustCompile(`[^|]*`)
    myMarketDataIncrementalRefresh_RegexValidator.TargetLocationID = regexp.MustCompile(`[^|]*`)
    myMarketDataIncrementalRefresh_RegexValidator.OnBehalfOfSubID = regexp.MustCompile(`[^|]*`)
    myMarketDataIncrementalRefresh_RegexValidator.OnBehalfOfLocationID = regexp.MustCompile(`[^|]*`)
    myMarketDataIncrementalRefresh_RegexValidator.DeliverToSubID = regexp.MustCompile(`[^|]*`)
    myMarketDataIncrementalRefresh_RegexValidator.DeliverToLocationID = regexp.MustCompile(`[^|]*`)
    myMarketDataIncrementalRefresh_RegexValidator.PossDupFlag = regexp.MustCompile(`[YN]`)
    myMarketDataIncrementalRefresh_RegexValidator.PossResend = regexp.MustCompile(`[YN]`)
    myMarketDataIncrementalRefresh_RegexValidator.SendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myMarketDataIncrementalRefresh_RegexValidator.OrigSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myMarketDataIncrementalRefresh_RegexValidator.XmlDataLen = regexp.MustCompile(`\d+`)
    myMarketDataIncrementalRefresh_RegexValidator.XmlData = regexp.MustCompile(`.*`)
    myMarketDataIncrementalRefresh_RegexValidator.MessageEncoding = regexp.MustCompile(`[^|]*`)
    myMarketDataIncrementalRefresh_RegexValidator.LastMsgSeqNumProcessed = regexp.MustCompile(`\d+`)
    myMarketDataIncrementalRefresh_RegexValidator.ApplVerID = regexp.MustCompile(`[^|]*`)
    myMarketDataIncrementalRefresh_RegexValidator.CstmApplVerID = regexp.MustCompile(`[^|]*`)
    myMarketDataIncrementalRefresh_RegexValidator.MDReqID = regexp.MustCompile(`[^|]*`)
    myMarketDataIncrementalRefresh_RegexValidator.ApplQueueDepth = regexp.MustCompile(`-?\d+`)
    myMarketDataIncrementalRefresh_RegexValidator.ApplQueueResolution = regexp.MustCompile(`-?\d+`)
    myMarketDataIncrementalRefresh_RegexValidator.MDBookType = regexp.MustCompile(`-?\d+`)
    myMarketDataIncrementalRefresh_RegexValidator.MDFeedType = regexp.MustCompile(`[^|]*`)
    myMarketDataIncrementalRefresh_RegexValidator.TradeDate = regexp.MustCompile(`[0-9]{8}`)
    myMarketDataIncrementalRefresh_RegexValidator.SignatureLength = regexp.MustCompile(`\d+`)
    myMarketDataIncrementalRefresh_RegexValidator.Signature = regexp.MustCompile(`.*`)
    myMarketDataIncrementalRefresh_RegexValidator.CheckSum = regexp.MustCompile(`[^|]*`)
}



func (m *MarketDataIncrementalRefresh) HasRequiredFields() bool {
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
    
    
    
    for _, g := range(m.MDEntries) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.RoutingIDs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.CheckSum.HasValue()
    
    
    
    
    
    return valid
}




