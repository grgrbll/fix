package ExecutionAcknowledgement

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import SecAltIDGrp_SecurityAltID "grgrbll/fix/Shared/SecAltIDGrp_SecurityAltID"

import EvntGrp_Events "grgrbll/fix/Shared/EvntGrp_Events"

import InstrumentParties_InstrumentParties "grgrbll/fix/Shared/InstrumentParties_InstrumentParties"

import UndInstrmtGrp_Underlyings "grgrbll/fix/Shared/UndInstrmtGrp_Underlyings"

import InstrmtLegGrp_Legs "grgrbll/fix/Shared/InstrmtLegGrp_Legs"


type executionAcknowledgement_RegexValidator struct {
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
    ClOrdID *(regexp.Regexp)
    ExecAckStatus *(regexp.Regexp)
    ExecID *(regexp.Regexp)
    DKReason *(regexp.Regexp)
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
    Side *(regexp.Regexp)
    OrderQty *(regexp.Regexp)
    CashOrderQty *(regexp.Regexp)
    OrderPercent *(regexp.Regexp)
    RoundingDirection *(regexp.Regexp)
    RoundingModulus *(regexp.Regexp)
    LastQty *(regexp.Regexp)
    LastPx *(regexp.Regexp)
    PriceType *(regexp.Regexp)
    LastParPx *(regexp.Regexp)
    CumQty *(regexp.Regexp)
    AvgPx *(regexp.Regexp)
    Text *(regexp.Regexp)
    EncodedTextLen *(regexp.Regexp)
    EncodedText *(regexp.Regexp)
    SignatureLength *(regexp.Regexp)
    Signature *(regexp.Regexp)
    CheckSum *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myExecutionAcknowledgement_RegexValidator executionAcknowledgement_RegexValidator

func init() {
    myExecutionAcknowledgement_RegexValidator
    myExecutionAcknowledgement_RegexValidator.BeginString = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.BodyLength = regexp.MustCompile(`\d+`)
    myExecutionAcknowledgement_RegexValidator.MsgType = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.SenderCompID = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.TargetCompID = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.OnBehalfOfCompID = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.DeliverToCompID = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.SecureDataLen = regexp.MustCompile(`\d+`)
    myExecutionAcknowledgement_RegexValidator.SecureData = regexp.MustCompile(`.*`)
    myExecutionAcknowledgement_RegexValidator.MsgSeqNum = regexp.MustCompile(`\d+`)
    myExecutionAcknowledgement_RegexValidator.SenderSubID = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.SenderLocationID = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.TargetSubID = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.TargetLocationID = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.OnBehalfOfSubID = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.OnBehalfOfLocationID = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.DeliverToSubID = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.DeliverToLocationID = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.PossDupFlag = regexp.MustCompile(`[YN]`)
    myExecutionAcknowledgement_RegexValidator.PossResend = regexp.MustCompile(`[YN]`)
    myExecutionAcknowledgement_RegexValidator.SendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myExecutionAcknowledgement_RegexValidator.OrigSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myExecutionAcknowledgement_RegexValidator.XmlDataLen = regexp.MustCompile(`\d+`)
    myExecutionAcknowledgement_RegexValidator.XmlData = regexp.MustCompile(`.*`)
    myExecutionAcknowledgement_RegexValidator.MessageEncoding = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.LastMsgSeqNumProcessed = regexp.MustCompile(`\d+`)
    myExecutionAcknowledgement_RegexValidator.ApplVerID = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.CstmApplVerID = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.OrderID = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.SecondaryOrderID = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.ClOrdID = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.ExecAckStatus = regexp.MustCompile(`[^|]`)
    myExecutionAcknowledgement_RegexValidator.ExecID = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.DKReason = regexp.MustCompile(`[^|]`)
    myExecutionAcknowledgement_RegexValidator.Symbol = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.SymbolSfx = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.SecurityID = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.SecurityIDSource = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.Product = regexp.MustCompile(`-?\d+`)
    myExecutionAcknowledgement_RegexValidator.CFICode = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.SecurityType = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.SecuritySubType = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.MaturityMonthYear = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myExecutionAcknowledgement_RegexValidator.MaturityDate = regexp.MustCompile(`[0-9]{8}`)
    myExecutionAcknowledgement_RegexValidator.CouponPaymentDate = regexp.MustCompile(`[0-9]{8}`)
    myExecutionAcknowledgement_RegexValidator.IssueDate = regexp.MustCompile(`[0-9]{8}`)
    myExecutionAcknowledgement_RegexValidator.RepoCollateralSecurityType = regexp.MustCompile(`-?\d+`)
    myExecutionAcknowledgement_RegexValidator.RepurchaseTerm = regexp.MustCompile(`-?\d+`)
    myExecutionAcknowledgement_RegexValidator.RepurchaseRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myExecutionAcknowledgement_RegexValidator.Factor = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myExecutionAcknowledgement_RegexValidator.CreditRating = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.InstrRegistry = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.CountryOfIssue = regexp.MustCompile(`[A-Z]{2}`)
    myExecutionAcknowledgement_RegexValidator.StateOrProvinceOfIssue = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.LocaleOfIssue = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.RedemptionDate = regexp.MustCompile(`[0-9]{8}`)
    myExecutionAcknowledgement_RegexValidator.StrikePrice = regexp.MustCompile(``)
    myExecutionAcknowledgement_RegexValidator.StrikeCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myExecutionAcknowledgement_RegexValidator.OptAttribute = regexp.MustCompile(`[^|]`)
    myExecutionAcknowledgement_RegexValidator.ContractMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myExecutionAcknowledgement_RegexValidator.CouponRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myExecutionAcknowledgement_RegexValidator.SecurityExchange = regexp.MustCompile(`[A-Z0-9]{4}`)
    myExecutionAcknowledgement_RegexValidator.Issuer = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.EncodedIssuerLen = regexp.MustCompile(`\d+`)
    myExecutionAcknowledgement_RegexValidator.EncodedIssuer = regexp.MustCompile(`.*`)
    myExecutionAcknowledgement_RegexValidator.SecurityDesc = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.EncodedSecurityDescLen = regexp.MustCompile(`\d+`)
    myExecutionAcknowledgement_RegexValidator.EncodedSecurityDesc = regexp.MustCompile(`.*`)
    myExecutionAcknowledgement_RegexValidator.Pool = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.ContractSettlMonth = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myExecutionAcknowledgement_RegexValidator.CPProgram = regexp.MustCompile(`-?\d+`)
    myExecutionAcknowledgement_RegexValidator.CPRegType = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.DatedDate = regexp.MustCompile(`[0-9]{8}`)
    myExecutionAcknowledgement_RegexValidator.InterestAccrualDate = regexp.MustCompile(`[0-9]{8}`)
    myExecutionAcknowledgement_RegexValidator.SecurityStatus = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.SettleOnOpenFlag = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.InstrmtAssignmentMethod = regexp.MustCompile(`[^|]`)
    myExecutionAcknowledgement_RegexValidator.StrikeMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myExecutionAcknowledgement_RegexValidator.StrikeValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myExecutionAcknowledgement_RegexValidator.MinPriceIncrement = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myExecutionAcknowledgement_RegexValidator.PositionLimit = regexp.MustCompile(`-?\d+`)
    myExecutionAcknowledgement_RegexValidator.NTPositionLimit = regexp.MustCompile(`-?\d+`)
    myExecutionAcknowledgement_RegexValidator.UnitofMeasure = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.TimeUnit = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.MaturityTime = regexp.MustCompile(`\d{2}:\d{2}(:\d{2})?(Z|[\+-]\d{2}(:\d{2})?)`)
    myExecutionAcknowledgement_RegexValidator.Side = regexp.MustCompile(`[^|]`)
    myExecutionAcknowledgement_RegexValidator.OrderQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myExecutionAcknowledgement_RegexValidator.CashOrderQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myExecutionAcknowledgement_RegexValidator.OrderPercent = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myExecutionAcknowledgement_RegexValidator.RoundingDirection = regexp.MustCompile(`[^|]`)
    myExecutionAcknowledgement_RegexValidator.RoundingModulus = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myExecutionAcknowledgement_RegexValidator.LastQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myExecutionAcknowledgement_RegexValidator.LastPx = regexp.MustCompile(``)
    myExecutionAcknowledgement_RegexValidator.PriceType = regexp.MustCompile(`-?\d+`)
    myExecutionAcknowledgement_RegexValidator.LastParPx = regexp.MustCompile(``)
    myExecutionAcknowledgement_RegexValidator.CumQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myExecutionAcknowledgement_RegexValidator.AvgPx = regexp.MustCompile(``)
    myExecutionAcknowledgement_RegexValidator.Text = regexp.MustCompile(`[^|]*`)
    myExecutionAcknowledgement_RegexValidator.EncodedTextLen = regexp.MustCompile(`\d+`)
    myExecutionAcknowledgement_RegexValidator.EncodedText = regexp.MustCompile(`.*`)
    myExecutionAcknowledgement_RegexValidator.SignatureLength = regexp.MustCompile(`\d+`)
    myExecutionAcknowledgement_RegexValidator.Signature = regexp.MustCompile(`.*`)
    myExecutionAcknowledgement_RegexValidator.CheckSum = regexp.MustCompile(`[^|]*`)
}



func (m *ExecutionAcknowledgement) HasRequiredFields() bool {
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
    
    
    
    valid = valid && m.ExecAckStatus.HasValue()
    
    
    
    valid = valid && m.ExecID.HasValue()
    
    
    
    for _, g := range(m.SecurityAltID) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Events) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.InstrumentParties) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Underlyings) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Legs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.Side.HasValue()
    
    
    
    valid = valid && m.CheckSum.HasValue()
    
    
    
    
    
    return valid
}




