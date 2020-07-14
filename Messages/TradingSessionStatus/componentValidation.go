package TradingSessionStatus

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import SecAltIDGrp_SecurityAltID "grgrbll/fix/Shared/SecAltIDGrp_SecurityAltID"

import EvntGrp_Events "grgrbll/fix/Shared/EvntGrp_Events"

import InstrumentParties_InstrumentParties "grgrbll/fix/Shared/InstrumentParties_InstrumentParties"


type tradingSessionStatus_RegexValidator struct {
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
    UnsolicitedIndicator *(regexp.Regexp)
    TradSesStatus *(regexp.Regexp)
    TradSesStatusRejReason *(regexp.Regexp)
    TradSesStartTime *(regexp.Regexp)
    TradSesOpenTime *(regexp.Regexp)
    TradSesPreCloseTime *(regexp.Regexp)
    TradSesCloseTime *(regexp.Regexp)
    TradSesEndTime *(regexp.Regexp)
    TotalVolumeTraded *(regexp.Regexp)
    Text *(regexp.Regexp)
    EncodedTextLen *(regexp.Regexp)
    EncodedText *(regexp.Regexp)
    Symbol *(regexp.Regexp)
    SymbolSfx *(regexp.Regexp)
    SecurityID *(regexp.Regexp)
    SecurityIDSource *(regexp.Regexp)
    Product *(regexp.Regexp)
    CFICode *(regexp.Regexp)
    SecurityType *(regexp.Regexp)
    SecuritySubType *(regexp.Regexp)
    MaturityMonthYear *(regexp.Regexp)
    MaturityDate *(regexp.Regexp)
    CouponPaymentDate *(regexp.Regexp)
    IssueDate *(regexp.Regexp)
    RepoCollateralSecurityType *(regexp.Regexp)
    RepurchaseTerm *(regexp.Regexp)
    RepurchaseRate *(regexp.Regexp)
    Factor *(regexp.Regexp)
    CreditRating *(regexp.Regexp)
    InstrRegistry *(regexp.Regexp)
    CountryOfIssue *(regexp.Regexp)
    StateOrProvinceOfIssue *(regexp.Regexp)
    LocaleOfIssue *(regexp.Regexp)
    RedemptionDate *(regexp.Regexp)
    StrikePrice *(regexp.Regexp)
    StrikeCurrency *(regexp.Regexp)
    OptAttribute *(regexp.Regexp)
    ContractMultiplier *(regexp.Regexp)
    CouponRate *(regexp.Regexp)
    SecurityExchange *(regexp.Regexp)
    Issuer *(regexp.Regexp)
    EncodedIssuerLen *(regexp.Regexp)
    EncodedIssuer *(regexp.Regexp)
    SecurityDesc *(regexp.Regexp)
    EncodedSecurityDescLen *(regexp.Regexp)
    EncodedSecurityDesc *(regexp.Regexp)
    Pool *(regexp.Regexp)
    ContractSettlMonth *(regexp.Regexp)
    CPProgram *(regexp.Regexp)
    CPRegType *(regexp.Regexp)
    DatedDate *(regexp.Regexp)
    InterestAccrualDate *(regexp.Regexp)
    SecurityStatus *(regexp.Regexp)
    SettleOnOpenFlag *(regexp.Regexp)
    InstrmtAssignmentMethod *(regexp.Regexp)
    StrikeMultiplier *(regexp.Regexp)
    StrikeValue *(regexp.Regexp)
    MinPriceIncrement *(regexp.Regexp)
    PositionLimit *(regexp.Regexp)
    NTPositionLimit *(regexp.Regexp)
    UnitofMeasure *(regexp.Regexp)
    TimeUnit *(regexp.Regexp)
    MaturityTime *(regexp.Regexp)
    SignatureLength *(regexp.Regexp)
    Signature *(regexp.Regexp)
    CheckSum *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myTradingSessionStatus_RegexValidator tradingSessionStatus_RegexValidator

func init() {
    myTradingSessionStatus_RegexValidator
    myTradingSessionStatus_RegexValidator.BeginString = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.BodyLength = regexp.MustCompile(`\d+`)
    myTradingSessionStatus_RegexValidator.MsgType = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.SenderCompID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.TargetCompID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.OnBehalfOfCompID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.DeliverToCompID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.SecureDataLen = regexp.MustCompile(`\d+`)
    myTradingSessionStatus_RegexValidator.SecureData = regexp.MustCompile(`.*`)
    myTradingSessionStatus_RegexValidator.MsgSeqNum = regexp.MustCompile(`\d+`)
    myTradingSessionStatus_RegexValidator.SenderSubID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.SenderLocationID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.TargetSubID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.TargetLocationID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.OnBehalfOfSubID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.OnBehalfOfLocationID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.DeliverToSubID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.DeliverToLocationID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.PossDupFlag = regexp.MustCompile(`[YN]`)
    myTradingSessionStatus_RegexValidator.PossResend = regexp.MustCompile(`[YN]`)
    myTradingSessionStatus_RegexValidator.SendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myTradingSessionStatus_RegexValidator.OrigSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myTradingSessionStatus_RegexValidator.XmlDataLen = regexp.MustCompile(`\d+`)
    myTradingSessionStatus_RegexValidator.XmlData = regexp.MustCompile(`.*`)
    myTradingSessionStatus_RegexValidator.MessageEncoding = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.LastMsgSeqNumProcessed = regexp.MustCompile(`\d+`)
    myTradingSessionStatus_RegexValidator.ApplVerID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.CstmApplVerID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.TradSesReqID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.TradingSessionID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.TradingSessionSubID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.TradSesMethod = regexp.MustCompile(`-?\d+`)
    myTradingSessionStatus_RegexValidator.TradSesMode = regexp.MustCompile(`-?\d+`)
    myTradingSessionStatus_RegexValidator.UnsolicitedIndicator = regexp.MustCompile(`[YN]`)
    myTradingSessionStatus_RegexValidator.TradSesStatus = regexp.MustCompile(`-?\d+`)
    myTradingSessionStatus_RegexValidator.TradSesStatusRejReason = regexp.MustCompile(`-?\d+`)
    myTradingSessionStatus_RegexValidator.TradSesStartTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myTradingSessionStatus_RegexValidator.TradSesOpenTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myTradingSessionStatus_RegexValidator.TradSesPreCloseTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myTradingSessionStatus_RegexValidator.TradSesCloseTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myTradingSessionStatus_RegexValidator.TradSesEndTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myTradingSessionStatus_RegexValidator.TotalVolumeTraded = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradingSessionStatus_RegexValidator.Text = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.EncodedTextLen = regexp.MustCompile(`\d+`)
    myTradingSessionStatus_RegexValidator.EncodedText = regexp.MustCompile(`.*`)
    myTradingSessionStatus_RegexValidator.Symbol = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.SymbolSfx = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.SecurityID = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.SecurityIDSource = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.Product = regexp.MustCompile(`-?\d+`)
    myTradingSessionStatus_RegexValidator.CFICode = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.SecurityType = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.SecuritySubType = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.MaturityMonthYear = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myTradingSessionStatus_RegexValidator.MaturityDate = regexp.MustCompile(`[0-9]{8}`)
    myTradingSessionStatus_RegexValidator.CouponPaymentDate = regexp.MustCompile(`[0-9]{8}`)
    myTradingSessionStatus_RegexValidator.IssueDate = regexp.MustCompile(`[0-9]{8}`)
    myTradingSessionStatus_RegexValidator.RepoCollateralSecurityType = regexp.MustCompile(`-?\d+`)
    myTradingSessionStatus_RegexValidator.RepurchaseTerm = regexp.MustCompile(`-?\d+`)
    myTradingSessionStatus_RegexValidator.RepurchaseRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradingSessionStatus_RegexValidator.Factor = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradingSessionStatus_RegexValidator.CreditRating = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.InstrRegistry = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.CountryOfIssue = regexp.MustCompile(`[A-Z]{2}`)
    myTradingSessionStatus_RegexValidator.StateOrProvinceOfIssue = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.LocaleOfIssue = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.RedemptionDate = regexp.MustCompile(`[0-9]{8}`)
    myTradingSessionStatus_RegexValidator.StrikePrice = regexp.MustCompile(``)
    myTradingSessionStatus_RegexValidator.StrikeCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myTradingSessionStatus_RegexValidator.OptAttribute = regexp.MustCompile(`[^|]`)
    myTradingSessionStatus_RegexValidator.ContractMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradingSessionStatus_RegexValidator.CouponRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradingSessionStatus_RegexValidator.SecurityExchange = regexp.MustCompile(`[A-Z0-9]{4}`)
    myTradingSessionStatus_RegexValidator.Issuer = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.EncodedIssuerLen = regexp.MustCompile(`\d+`)
    myTradingSessionStatus_RegexValidator.EncodedIssuer = regexp.MustCompile(`.*`)
    myTradingSessionStatus_RegexValidator.SecurityDesc = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.EncodedSecurityDescLen = regexp.MustCompile(`\d+`)
    myTradingSessionStatus_RegexValidator.EncodedSecurityDesc = regexp.MustCompile(`.*`)
    myTradingSessionStatus_RegexValidator.Pool = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.ContractSettlMonth = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myTradingSessionStatus_RegexValidator.CPProgram = regexp.MustCompile(`-?\d+`)
    myTradingSessionStatus_RegexValidator.CPRegType = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.DatedDate = regexp.MustCompile(`[0-9]{8}`)
    myTradingSessionStatus_RegexValidator.InterestAccrualDate = regexp.MustCompile(`[0-9]{8}`)
    myTradingSessionStatus_RegexValidator.SecurityStatus = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.SettleOnOpenFlag = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.InstrmtAssignmentMethod = regexp.MustCompile(`[^|]`)
    myTradingSessionStatus_RegexValidator.StrikeMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradingSessionStatus_RegexValidator.StrikeValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradingSessionStatus_RegexValidator.MinPriceIncrement = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradingSessionStatus_RegexValidator.PositionLimit = regexp.MustCompile(`-?\d+`)
    myTradingSessionStatus_RegexValidator.NTPositionLimit = regexp.MustCompile(`-?\d+`)
    myTradingSessionStatus_RegexValidator.UnitofMeasure = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.TimeUnit = regexp.MustCompile(`[^|]*`)
    myTradingSessionStatus_RegexValidator.MaturityTime = regexp.MustCompile(`\d{2}:\d{2}(:\d{2})?(Z|[\+-]\d{2}(:\d{2})?)`)
    myTradingSessionStatus_RegexValidator.SignatureLength = regexp.MustCompile(`\d+`)
    myTradingSessionStatus_RegexValidator.Signature = regexp.MustCompile(`.*`)
    myTradingSessionStatus_RegexValidator.CheckSum = regexp.MustCompile(`[^|]*`)
}



func (m *TradingSessionStatus) HasRequiredFields() bool {
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
    
    
    
    valid = valid && m.TradingSessionID.HasValue()
    
    
    
    valid = valid && m.TradSesStatus.HasValue()
    
    
    
    for _, g := range(m.SecurityAltID) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Events) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.InstrumentParties) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.CheckSum.HasValue()
    
    
    
    
    
    return valid
}




