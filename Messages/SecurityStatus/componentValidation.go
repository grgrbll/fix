package SecurityStatus

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import SecAltIDGrp_SecurityAltID "grgrbll/fix/Shared/SecAltIDGrp_SecurityAltID"

import EvntGrp_Events "grgrbll/fix/Shared/EvntGrp_Events"

import InstrumentParties_InstrumentParties "grgrbll/fix/Shared/InstrumentParties_InstrumentParties"

import AttrbGrp_InstrAttrib "grgrbll/fix/Shared/AttrbGrp_InstrAttrib"

import UndInstrmtGrp_Underlyings "grgrbll/fix/Shared/UndInstrmtGrp_Underlyings"

import InstrmtLegGrp_Legs "grgrbll/fix/Shared/InstrmtLegGrp_Legs"


type securityStatus_RegexValidator struct {
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
    SecurityStatusReqID *(regexp.Regexp)
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
    DeliveryForm *(regexp.Regexp)
    PctAtRisk *(regexp.Regexp)
    Currency *(regexp.Regexp)
    TradingSessionID *(regexp.Regexp)
    TradingSessionSubID *(regexp.Regexp)
    UnsolicitedIndicator *(regexp.Regexp)
    SecurityTradingStatus *(regexp.Regexp)
    FinancialStatus *(regexp.Regexp)
    CorporateAction *(regexp.Regexp)
    HaltReason *(regexp.Regexp)
    InViewOfCommon *(regexp.Regexp)
    DueToRelated *(regexp.Regexp)
    BuyVolume *(regexp.Regexp)
    SellVolume *(regexp.Regexp)
    HighPx *(regexp.Regexp)
    LowPx *(regexp.Regexp)
    LastPx *(regexp.Regexp)
    TransactTime *(regexp.Regexp)
    Adjustment *(regexp.Regexp)
    Text *(regexp.Regexp)
    EncodedTextLen *(regexp.Regexp)
    EncodedText *(regexp.Regexp)
    FirstPx *(regexp.Regexp)
    SignatureLength *(regexp.Regexp)
    Signature *(regexp.Regexp)
    CheckSum *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var mySecurityStatus_RegexValidator securityStatus_RegexValidator

func init() {
    mySecurityStatus_RegexValidator
    mySecurityStatus_RegexValidator.BeginString = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.BodyLength = regexp.MustCompile(`\d+`)
    mySecurityStatus_RegexValidator.MsgType = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.SenderCompID = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.TargetCompID = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.OnBehalfOfCompID = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.DeliverToCompID = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.SecureDataLen = regexp.MustCompile(`\d+`)
    mySecurityStatus_RegexValidator.SecureData = regexp.MustCompile(`.*`)
    mySecurityStatus_RegexValidator.MsgSeqNum = regexp.MustCompile(`\d+`)
    mySecurityStatus_RegexValidator.SenderSubID = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.SenderLocationID = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.TargetSubID = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.TargetLocationID = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.OnBehalfOfSubID = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.OnBehalfOfLocationID = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.DeliverToSubID = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.DeliverToLocationID = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.PossDupFlag = regexp.MustCompile(`[YN]`)
    mySecurityStatus_RegexValidator.PossResend = regexp.MustCompile(`[YN]`)
    mySecurityStatus_RegexValidator.SendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    mySecurityStatus_RegexValidator.OrigSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    mySecurityStatus_RegexValidator.XmlDataLen = regexp.MustCompile(`\d+`)
    mySecurityStatus_RegexValidator.XmlData = regexp.MustCompile(`.*`)
    mySecurityStatus_RegexValidator.MessageEncoding = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.LastMsgSeqNumProcessed = regexp.MustCompile(`\d+`)
    mySecurityStatus_RegexValidator.ApplVerID = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.CstmApplVerID = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.SecurityStatusReqID = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.Symbol = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.SymbolSfx = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.SecurityID = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.SecurityIDSource = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.Product = regexp.MustCompile(`-?\d+`)
    mySecurityStatus_RegexValidator.CFICode = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.SecurityType = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.SecuritySubType = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.MaturityMonthYear = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    mySecurityStatus_RegexValidator.MaturityDate = regexp.MustCompile(`[0-9]{8}`)
    mySecurityStatus_RegexValidator.CouponPaymentDate = regexp.MustCompile(`[0-9]{8}`)
    mySecurityStatus_RegexValidator.IssueDate = regexp.MustCompile(`[0-9]{8}`)
    mySecurityStatus_RegexValidator.RepoCollateralSecurityType = regexp.MustCompile(`-?\d+`)
    mySecurityStatus_RegexValidator.RepurchaseTerm = regexp.MustCompile(`-?\d+`)
    mySecurityStatus_RegexValidator.RepurchaseRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    mySecurityStatus_RegexValidator.Factor = regexp.MustCompile(`-?\d*(\.\d*)?`)
    mySecurityStatus_RegexValidator.CreditRating = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.InstrRegistry = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.CountryOfIssue = regexp.MustCompile(`[A-Z]{2}`)
    mySecurityStatus_RegexValidator.StateOrProvinceOfIssue = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.LocaleOfIssue = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.RedemptionDate = regexp.MustCompile(`[0-9]{8}`)
    mySecurityStatus_RegexValidator.StrikePrice = regexp.MustCompile(``)
    mySecurityStatus_RegexValidator.StrikeCurrency = regexp.MustCompile(`[A-Z]{3}`)
    mySecurityStatus_RegexValidator.OptAttribute = regexp.MustCompile(`[^|]`)
    mySecurityStatus_RegexValidator.ContractMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    mySecurityStatus_RegexValidator.CouponRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    mySecurityStatus_RegexValidator.SecurityExchange = regexp.MustCompile(`[A-Z0-9]{4}`)
    mySecurityStatus_RegexValidator.Issuer = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.EncodedIssuerLen = regexp.MustCompile(`\d+`)
    mySecurityStatus_RegexValidator.EncodedIssuer = regexp.MustCompile(`.*`)
    mySecurityStatus_RegexValidator.SecurityDesc = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.EncodedSecurityDescLen = regexp.MustCompile(`\d+`)
    mySecurityStatus_RegexValidator.EncodedSecurityDesc = regexp.MustCompile(`.*`)
    mySecurityStatus_RegexValidator.Pool = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.ContractSettlMonth = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    mySecurityStatus_RegexValidator.CPProgram = regexp.MustCompile(`-?\d+`)
    mySecurityStatus_RegexValidator.CPRegType = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.DatedDate = regexp.MustCompile(`[0-9]{8}`)
    mySecurityStatus_RegexValidator.InterestAccrualDate = regexp.MustCompile(`[0-9]{8}`)
    mySecurityStatus_RegexValidator.SecurityStatus = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.SettleOnOpenFlag = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.InstrmtAssignmentMethod = regexp.MustCompile(`[^|]`)
    mySecurityStatus_RegexValidator.StrikeMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    mySecurityStatus_RegexValidator.StrikeValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    mySecurityStatus_RegexValidator.MinPriceIncrement = regexp.MustCompile(`-?\d*(\.\d*)?`)
    mySecurityStatus_RegexValidator.PositionLimit = regexp.MustCompile(`-?\d+`)
    mySecurityStatus_RegexValidator.NTPositionLimit = regexp.MustCompile(`-?\d+`)
    mySecurityStatus_RegexValidator.UnitofMeasure = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.TimeUnit = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.MaturityTime = regexp.MustCompile(`\d{2}:\d{2}(:\d{2})?(Z|[\+-]\d{2}(:\d{2})?)`)
    mySecurityStatus_RegexValidator.DeliveryForm = regexp.MustCompile(`-?\d+`)
    mySecurityStatus_RegexValidator.PctAtRisk = regexp.MustCompile(`-?\d*(\.\d*)?`)
    mySecurityStatus_RegexValidator.Currency = regexp.MustCompile(`[A-Z]{3}`)
    mySecurityStatus_RegexValidator.TradingSessionID = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.TradingSessionSubID = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.UnsolicitedIndicator = regexp.MustCompile(`[YN]`)
    mySecurityStatus_RegexValidator.SecurityTradingStatus = regexp.MustCompile(`-?\d+`)
    mySecurityStatus_RegexValidator.FinancialStatus = regexp.MustCompile(`[ ]?(\w[ ])+\w?`)
    mySecurityStatus_RegexValidator.CorporateAction = regexp.MustCompile(`[ ]?(\w[ ])+\w?`)
    mySecurityStatus_RegexValidator.HaltReason = regexp.MustCompile(`[^|]`)
    mySecurityStatus_RegexValidator.InViewOfCommon = regexp.MustCompile(`[YN]`)
    mySecurityStatus_RegexValidator.DueToRelated = regexp.MustCompile(`[YN]`)
    mySecurityStatus_RegexValidator.BuyVolume = regexp.MustCompile(`-?\d*(\.\d*)?`)
    mySecurityStatus_RegexValidator.SellVolume = regexp.MustCompile(`-?\d*(\.\d*)?`)
    mySecurityStatus_RegexValidator.HighPx = regexp.MustCompile(``)
    mySecurityStatus_RegexValidator.LowPx = regexp.MustCompile(``)
    mySecurityStatus_RegexValidator.LastPx = regexp.MustCompile(``)
    mySecurityStatus_RegexValidator.TransactTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    mySecurityStatus_RegexValidator.Adjustment = regexp.MustCompile(`-?\d+`)
    mySecurityStatus_RegexValidator.Text = regexp.MustCompile(`[^|]*`)
    mySecurityStatus_RegexValidator.EncodedTextLen = regexp.MustCompile(`\d+`)
    mySecurityStatus_RegexValidator.EncodedText = regexp.MustCompile(`.*`)
    mySecurityStatus_RegexValidator.FirstPx = regexp.MustCompile(``)
    mySecurityStatus_RegexValidator.SignatureLength = regexp.MustCompile(`\d+`)
    mySecurityStatus_RegexValidator.Signature = regexp.MustCompile(`.*`)
    mySecurityStatus_RegexValidator.CheckSum = regexp.MustCompile(`[^|]*`)
}



func (m *SecurityStatus) HasRequiredFields() bool {
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
    
    
    
    for _, g := range(m.SecurityAltID) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Events) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.InstrumentParties) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.InstrAttrib) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Underlyings) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Legs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.CheckSum.HasValue()
    
    
    
    
    
    return valid
}




