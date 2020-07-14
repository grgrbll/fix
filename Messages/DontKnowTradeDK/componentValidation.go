package DontKnowTradeDK

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import SecAltIDGrp_SecurityAltID "grgrbll/fix/Shared/SecAltIDGrp_SecurityAltID"

import EvntGrp_Events "grgrbll/fix/Shared/EvntGrp_Events"

import InstrumentParties_InstrumentParties "grgrbll/fix/Shared/InstrumentParties_InstrumentParties"

import UndInstrmtGrp_Underlyings "grgrbll/fix/Shared/UndInstrmtGrp_Underlyings"

import InstrmtLegGrp_Legs "grgrbll/fix/Shared/InstrmtLegGrp_Legs"


type dontKnowTradeDK_RegexValidator struct {
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
    Text *(regexp.Regexp)
    EncodedTextLen *(regexp.Regexp)
    EncodedText *(regexp.Regexp)
    SignatureLength *(regexp.Regexp)
    Signature *(regexp.Regexp)
    CheckSum *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myDontKnowTradeDK_RegexValidator dontKnowTradeDK_RegexValidator

func init() {
    myDontKnowTradeDK_RegexValidator
    myDontKnowTradeDK_RegexValidator.BeginString = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.BodyLength = regexp.MustCompile(`\d+`)
    myDontKnowTradeDK_RegexValidator.MsgType = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.SenderCompID = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.TargetCompID = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.OnBehalfOfCompID = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.DeliverToCompID = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.SecureDataLen = regexp.MustCompile(`\d+`)
    myDontKnowTradeDK_RegexValidator.SecureData = regexp.MustCompile(`.*`)
    myDontKnowTradeDK_RegexValidator.MsgSeqNum = regexp.MustCompile(`\d+`)
    myDontKnowTradeDK_RegexValidator.SenderSubID = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.SenderLocationID = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.TargetSubID = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.TargetLocationID = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.OnBehalfOfSubID = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.OnBehalfOfLocationID = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.DeliverToSubID = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.DeliverToLocationID = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.PossDupFlag = regexp.MustCompile(`[YN]`)
    myDontKnowTradeDK_RegexValidator.PossResend = regexp.MustCompile(`[YN]`)
    myDontKnowTradeDK_RegexValidator.SendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myDontKnowTradeDK_RegexValidator.OrigSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myDontKnowTradeDK_RegexValidator.XmlDataLen = regexp.MustCompile(`\d+`)
    myDontKnowTradeDK_RegexValidator.XmlData = regexp.MustCompile(`.*`)
    myDontKnowTradeDK_RegexValidator.MessageEncoding = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.LastMsgSeqNumProcessed = regexp.MustCompile(`\d+`)
    myDontKnowTradeDK_RegexValidator.ApplVerID = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.CstmApplVerID = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.OrderID = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.SecondaryOrderID = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.ExecID = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.DKReason = regexp.MustCompile(`[^|]`)
    myDontKnowTradeDK_RegexValidator.Symbol = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.SymbolSfx = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.SecurityID = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.SecurityIDSource = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.Product = regexp.MustCompile(`-?\d+`)
    myDontKnowTradeDK_RegexValidator.CFICode = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.SecurityType = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.SecuritySubType = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.MaturityMonthYear = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myDontKnowTradeDK_RegexValidator.MaturityDate = regexp.MustCompile(`[0-9]{8}`)
    myDontKnowTradeDK_RegexValidator.CouponPaymentDate = regexp.MustCompile(`[0-9]{8}`)
    myDontKnowTradeDK_RegexValidator.IssueDate = regexp.MustCompile(`[0-9]{8}`)
    myDontKnowTradeDK_RegexValidator.RepoCollateralSecurityType = regexp.MustCompile(`-?\d+`)
    myDontKnowTradeDK_RegexValidator.RepurchaseTerm = regexp.MustCompile(`-?\d+`)
    myDontKnowTradeDK_RegexValidator.RepurchaseRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myDontKnowTradeDK_RegexValidator.Factor = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myDontKnowTradeDK_RegexValidator.CreditRating = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.InstrRegistry = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.CountryOfIssue = regexp.MustCompile(`[A-Z]{2}`)
    myDontKnowTradeDK_RegexValidator.StateOrProvinceOfIssue = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.LocaleOfIssue = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.RedemptionDate = regexp.MustCompile(`[0-9]{8}`)
    myDontKnowTradeDK_RegexValidator.StrikePrice = regexp.MustCompile(``)
    myDontKnowTradeDK_RegexValidator.StrikeCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myDontKnowTradeDK_RegexValidator.OptAttribute = regexp.MustCompile(`[^|]`)
    myDontKnowTradeDK_RegexValidator.ContractMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myDontKnowTradeDK_RegexValidator.CouponRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myDontKnowTradeDK_RegexValidator.SecurityExchange = regexp.MustCompile(`[A-Z0-9]{4}`)
    myDontKnowTradeDK_RegexValidator.Issuer = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.EncodedIssuerLen = regexp.MustCompile(`\d+`)
    myDontKnowTradeDK_RegexValidator.EncodedIssuer = regexp.MustCompile(`.*`)
    myDontKnowTradeDK_RegexValidator.SecurityDesc = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.EncodedSecurityDescLen = regexp.MustCompile(`\d+`)
    myDontKnowTradeDK_RegexValidator.EncodedSecurityDesc = regexp.MustCompile(`.*`)
    myDontKnowTradeDK_RegexValidator.Pool = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.ContractSettlMonth = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myDontKnowTradeDK_RegexValidator.CPProgram = regexp.MustCompile(`-?\d+`)
    myDontKnowTradeDK_RegexValidator.CPRegType = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.DatedDate = regexp.MustCompile(`[0-9]{8}`)
    myDontKnowTradeDK_RegexValidator.InterestAccrualDate = regexp.MustCompile(`[0-9]{8}`)
    myDontKnowTradeDK_RegexValidator.SecurityStatus = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.SettleOnOpenFlag = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.InstrmtAssignmentMethod = regexp.MustCompile(`[^|]`)
    myDontKnowTradeDK_RegexValidator.StrikeMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myDontKnowTradeDK_RegexValidator.StrikeValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myDontKnowTradeDK_RegexValidator.MinPriceIncrement = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myDontKnowTradeDK_RegexValidator.PositionLimit = regexp.MustCompile(`-?\d+`)
    myDontKnowTradeDK_RegexValidator.NTPositionLimit = regexp.MustCompile(`-?\d+`)
    myDontKnowTradeDK_RegexValidator.UnitofMeasure = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.TimeUnit = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.MaturityTime = regexp.MustCompile(`\d{2}:\d{2}(:\d{2})?(Z|[\+-]\d{2}(:\d{2})?)`)
    myDontKnowTradeDK_RegexValidator.Side = regexp.MustCompile(`[^|]`)
    myDontKnowTradeDK_RegexValidator.OrderQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myDontKnowTradeDK_RegexValidator.CashOrderQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myDontKnowTradeDK_RegexValidator.OrderPercent = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myDontKnowTradeDK_RegexValidator.RoundingDirection = regexp.MustCompile(`[^|]`)
    myDontKnowTradeDK_RegexValidator.RoundingModulus = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myDontKnowTradeDK_RegexValidator.LastQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myDontKnowTradeDK_RegexValidator.LastPx = regexp.MustCompile(``)
    myDontKnowTradeDK_RegexValidator.Text = regexp.MustCompile(`[^|]*`)
    myDontKnowTradeDK_RegexValidator.EncodedTextLen = regexp.MustCompile(`\d+`)
    myDontKnowTradeDK_RegexValidator.EncodedText = regexp.MustCompile(`.*`)
    myDontKnowTradeDK_RegexValidator.SignatureLength = regexp.MustCompile(`\d+`)
    myDontKnowTradeDK_RegexValidator.Signature = regexp.MustCompile(`.*`)
    myDontKnowTradeDK_RegexValidator.CheckSum = regexp.MustCompile(`[^|]*`)
}



func (m *DontKnowTradeDK) HasRequiredFields() bool {
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
    
    
    
    valid = valid && m.ExecID.HasValue()
    
    
    
    valid = valid && m.DKReason.HasValue()
    
    
    
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




