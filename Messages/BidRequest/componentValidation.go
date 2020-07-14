package BidRequest

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import BidDescReqGrp_BidDescriptors "grgrbll/fix/Shared/BidDescReqGrp_BidDescriptors"

import BidCompReqGrp_BidComponents "grgrbll/fix/Shared/BidCompReqGrp_BidComponents"


type bidRequest_RegexValidator struct {
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
    BidID *(regexp.Regexp)
    ClientBidID *(regexp.Regexp)
    BidRequestTransType *(regexp.Regexp)
    ListName *(regexp.Regexp)
    TotNoRelatedSym *(regexp.Regexp)
    BidType *(regexp.Regexp)
    NumTickets *(regexp.Regexp)
    Currency *(regexp.Regexp)
    SideValue1 *(regexp.Regexp)
    SideValue2 *(regexp.Regexp)
    LiquidityIndType *(regexp.Regexp)
    WtAverageLiquidity *(regexp.Regexp)
    ExchangeForPhysical *(regexp.Regexp)
    OutMainCntryUIndex *(regexp.Regexp)
    CrossPercent *(regexp.Regexp)
    ProgRptReqs *(regexp.Regexp)
    ProgPeriodInterval *(regexp.Regexp)
    IncTaxInd *(regexp.Regexp)
    ForexReq *(regexp.Regexp)
    NumBidders *(regexp.Regexp)
    TradeDate *(regexp.Regexp)
    BidTradeType *(regexp.Regexp)
    BasisPxType *(regexp.Regexp)
    StrikeTime *(regexp.Regexp)
    Text *(regexp.Regexp)
    EncodedTextLen *(regexp.Regexp)
    EncodedText *(regexp.Regexp)
    SignatureLength *(regexp.Regexp)
    Signature *(regexp.Regexp)
    CheckSum *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myBidRequest_RegexValidator bidRequest_RegexValidator

func init() {
    myBidRequest_RegexValidator
    myBidRequest_RegexValidator.BeginString = regexp.MustCompile(`[^|]*`)
    myBidRequest_RegexValidator.BodyLength = regexp.MustCompile(`\d+`)
    myBidRequest_RegexValidator.MsgType = regexp.MustCompile(`[^|]*`)
    myBidRequest_RegexValidator.SenderCompID = regexp.MustCompile(`[^|]*`)
    myBidRequest_RegexValidator.TargetCompID = regexp.MustCompile(`[^|]*`)
    myBidRequest_RegexValidator.OnBehalfOfCompID = regexp.MustCompile(`[^|]*`)
    myBidRequest_RegexValidator.DeliverToCompID = regexp.MustCompile(`[^|]*`)
    myBidRequest_RegexValidator.SecureDataLen = regexp.MustCompile(`\d+`)
    myBidRequest_RegexValidator.SecureData = regexp.MustCompile(`.*`)
    myBidRequest_RegexValidator.MsgSeqNum = regexp.MustCompile(`\d+`)
    myBidRequest_RegexValidator.SenderSubID = regexp.MustCompile(`[^|]*`)
    myBidRequest_RegexValidator.SenderLocationID = regexp.MustCompile(`[^|]*`)
    myBidRequest_RegexValidator.TargetSubID = regexp.MustCompile(`[^|]*`)
    myBidRequest_RegexValidator.TargetLocationID = regexp.MustCompile(`[^|]*`)
    myBidRequest_RegexValidator.OnBehalfOfSubID = regexp.MustCompile(`[^|]*`)
    myBidRequest_RegexValidator.OnBehalfOfLocationID = regexp.MustCompile(`[^|]*`)
    myBidRequest_RegexValidator.DeliverToSubID = regexp.MustCompile(`[^|]*`)
    myBidRequest_RegexValidator.DeliverToLocationID = regexp.MustCompile(`[^|]*`)
    myBidRequest_RegexValidator.PossDupFlag = regexp.MustCompile(`[YN]`)
    myBidRequest_RegexValidator.PossResend = regexp.MustCompile(`[YN]`)
    myBidRequest_RegexValidator.SendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myBidRequest_RegexValidator.OrigSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myBidRequest_RegexValidator.XmlDataLen = regexp.MustCompile(`\d+`)
    myBidRequest_RegexValidator.XmlData = regexp.MustCompile(`.*`)
    myBidRequest_RegexValidator.MessageEncoding = regexp.MustCompile(`[^|]*`)
    myBidRequest_RegexValidator.LastMsgSeqNumProcessed = regexp.MustCompile(`\d+`)
    myBidRequest_RegexValidator.ApplVerID = regexp.MustCompile(`[^|]*`)
    myBidRequest_RegexValidator.CstmApplVerID = regexp.MustCompile(`[^|]*`)
    myBidRequest_RegexValidator.BidID = regexp.MustCompile(`[^|]*`)
    myBidRequest_RegexValidator.ClientBidID = regexp.MustCompile(`[^|]*`)
    myBidRequest_RegexValidator.BidRequestTransType = regexp.MustCompile(`[^|]`)
    myBidRequest_RegexValidator.ListName = regexp.MustCompile(`[^|]*`)
    myBidRequest_RegexValidator.TotNoRelatedSym = regexp.MustCompile(`-?\d+`)
    myBidRequest_RegexValidator.BidType = regexp.MustCompile(`-?\d+`)
    myBidRequest_RegexValidator.NumTickets = regexp.MustCompile(`-?\d+`)
    myBidRequest_RegexValidator.Currency = regexp.MustCompile(`[A-Z]{3}`)
    myBidRequest_RegexValidator.SideValue1 = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myBidRequest_RegexValidator.SideValue2 = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myBidRequest_RegexValidator.LiquidityIndType = regexp.MustCompile(`-?\d+`)
    myBidRequest_RegexValidator.WtAverageLiquidity = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myBidRequest_RegexValidator.ExchangeForPhysical = regexp.MustCompile(`[YN]`)
    myBidRequest_RegexValidator.OutMainCntryUIndex = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myBidRequest_RegexValidator.CrossPercent = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myBidRequest_RegexValidator.ProgRptReqs = regexp.MustCompile(`-?\d+`)
    myBidRequest_RegexValidator.ProgPeriodInterval = regexp.MustCompile(`-?\d+`)
    myBidRequest_RegexValidator.IncTaxInd = regexp.MustCompile(`-?\d+`)
    myBidRequest_RegexValidator.ForexReq = regexp.MustCompile(`[YN]`)
    myBidRequest_RegexValidator.NumBidders = regexp.MustCompile(`-?\d+`)
    myBidRequest_RegexValidator.TradeDate = regexp.MustCompile(`[0-9]{8}`)
    myBidRequest_RegexValidator.BidTradeType = regexp.MustCompile(`[^|]`)
    myBidRequest_RegexValidator.BasisPxType = regexp.MustCompile(`[^|]`)
    myBidRequest_RegexValidator.StrikeTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myBidRequest_RegexValidator.Text = regexp.MustCompile(`[^|]*`)
    myBidRequest_RegexValidator.EncodedTextLen = regexp.MustCompile(`\d+`)
    myBidRequest_RegexValidator.EncodedText = regexp.MustCompile(`.*`)
    myBidRequest_RegexValidator.SignatureLength = regexp.MustCompile(`\d+`)
    myBidRequest_RegexValidator.Signature = regexp.MustCompile(`.*`)
    myBidRequest_RegexValidator.CheckSum = regexp.MustCompile(`[^|]*`)
}



func (m *BidRequest) HasRequiredFields() bool {
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
    
    
    
    valid = valid && m.ClientBidID.HasValue()
    
    
    
    valid = valid && m.BidRequestTransType.HasValue()
    
    
    
    valid = valid && m.TotNoRelatedSym.HasValue()
    
    
    
    valid = valid && m.BidType.HasValue()
    
    
    
    for _, g := range(m.BidDescriptors) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.BidComponents) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.BidTradeType.HasValue()
    
    
    
    valid = valid && m.BasisPxType.HasValue()
    
    
    
    valid = valid && m.CheckSum.HasValue()
    
    
    
    
    
    return valid
}




