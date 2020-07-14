package AssignmentReport

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import Parties_PartyIDs "grgrbll/fix/Shared/Parties_PartyIDs"

import SecAltIDGrp_SecurityAltID "grgrbll/fix/Shared/SecAltIDGrp_SecurityAltID"

import EvntGrp_Events "grgrbll/fix/Shared/EvntGrp_Events"

import InstrumentParties_InstrumentParties "grgrbll/fix/Shared/InstrumentParties_InstrumentParties"

import InstrmtLegGrp_Legs "grgrbll/fix/Shared/InstrmtLegGrp_Legs"

import UndInstrmtGrp_Underlyings "grgrbll/fix/Shared/UndInstrmtGrp_Underlyings"

import PositionQty_Positions "grgrbll/fix/Shared/PositionQty_Positions"

import PositionAmountData_PosAmt "grgrbll/fix/Shared/PositionAmountData_PosAmt"


type assignmentReport_RegexValidator struct {
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
    AsgnRptID *(regexp.Regexp)
    TotNumAssignmentReports *(regexp.Regexp)
    LastRptRequested *(regexp.Regexp)
    Account *(regexp.Regexp)
    AccountType *(regexp.Regexp)
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
    Currency *(regexp.Regexp)
    ThresholdAmount *(regexp.Regexp)
    SettlPrice *(regexp.Regexp)
    SettlPriceType *(regexp.Regexp)
    UnderlyingSettlPrice *(regexp.Regexp)
    ExpireDate *(regexp.Regexp)
    AssignmentMethod *(regexp.Regexp)
    AssignmentUnit *(regexp.Regexp)
    OpenInterest *(regexp.Regexp)
    ExerciseMethod *(regexp.Regexp)
    SettlSessID *(regexp.Regexp)
    SettlSessSubID *(regexp.Regexp)
    ClearingBusinessDate *(regexp.Regexp)
    Text *(regexp.Regexp)
    EncodedTextLen *(regexp.Regexp)
    EncodedText *(regexp.Regexp)
    PriorSettlPrice *(regexp.Regexp)
    SignatureLength *(regexp.Regexp)
    Signature *(regexp.Regexp)
    CheckSum *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myAssignmentReport_RegexValidator assignmentReport_RegexValidator

func init() {
    myAssignmentReport_RegexValidator
    myAssignmentReport_RegexValidator.BeginString = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.BodyLength = regexp.MustCompile(`\d+`)
    myAssignmentReport_RegexValidator.MsgType = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.SenderCompID = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.TargetCompID = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.OnBehalfOfCompID = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.DeliverToCompID = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.SecureDataLen = regexp.MustCompile(`\d+`)
    myAssignmentReport_RegexValidator.SecureData = regexp.MustCompile(`.*`)
    myAssignmentReport_RegexValidator.MsgSeqNum = regexp.MustCompile(`\d+`)
    myAssignmentReport_RegexValidator.SenderSubID = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.SenderLocationID = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.TargetSubID = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.TargetLocationID = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.OnBehalfOfSubID = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.OnBehalfOfLocationID = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.DeliverToSubID = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.DeliverToLocationID = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.PossDupFlag = regexp.MustCompile(`[YN]`)
    myAssignmentReport_RegexValidator.PossResend = regexp.MustCompile(`[YN]`)
    myAssignmentReport_RegexValidator.SendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myAssignmentReport_RegexValidator.OrigSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myAssignmentReport_RegexValidator.XmlDataLen = regexp.MustCompile(`\d+`)
    myAssignmentReport_RegexValidator.XmlData = regexp.MustCompile(`.*`)
    myAssignmentReport_RegexValidator.MessageEncoding = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.LastMsgSeqNumProcessed = regexp.MustCompile(`\d+`)
    myAssignmentReport_RegexValidator.ApplVerID = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.CstmApplVerID = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.AsgnRptID = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.TotNumAssignmentReports = regexp.MustCompile(`-?\d+`)
    myAssignmentReport_RegexValidator.LastRptRequested = regexp.MustCompile(`[YN]`)
    myAssignmentReport_RegexValidator.Account = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.AccountType = regexp.MustCompile(`-?\d+`)
    myAssignmentReport_RegexValidator.Symbol = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.SymbolSfx = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.SecurityID = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.SecurityIDSource = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.Product = regexp.MustCompile(`-?\d+`)
    myAssignmentReport_RegexValidator.CFICode = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.SecurityType = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.SecuritySubType = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.MaturityMonthYear = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myAssignmentReport_RegexValidator.MaturityDate = regexp.MustCompile(`[0-9]{8}`)
    myAssignmentReport_RegexValidator.CouponPaymentDate = regexp.MustCompile(`[0-9]{8}`)
    myAssignmentReport_RegexValidator.IssueDate = regexp.MustCompile(`[0-9]{8}`)
    myAssignmentReport_RegexValidator.RepoCollateralSecurityType = regexp.MustCompile(`-?\d+`)
    myAssignmentReport_RegexValidator.RepurchaseTerm = regexp.MustCompile(`-?\d+`)
    myAssignmentReport_RegexValidator.RepurchaseRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAssignmentReport_RegexValidator.Factor = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAssignmentReport_RegexValidator.CreditRating = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.InstrRegistry = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.CountryOfIssue = regexp.MustCompile(`[A-Z]{2}`)
    myAssignmentReport_RegexValidator.StateOrProvinceOfIssue = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.LocaleOfIssue = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.RedemptionDate = regexp.MustCompile(`[0-9]{8}`)
    myAssignmentReport_RegexValidator.StrikePrice = regexp.MustCompile(``)
    myAssignmentReport_RegexValidator.StrikeCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myAssignmentReport_RegexValidator.OptAttribute = regexp.MustCompile(`[^|]`)
    myAssignmentReport_RegexValidator.ContractMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAssignmentReport_RegexValidator.CouponRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAssignmentReport_RegexValidator.SecurityExchange = regexp.MustCompile(`[A-Z0-9]{4}`)
    myAssignmentReport_RegexValidator.Issuer = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.EncodedIssuerLen = regexp.MustCompile(`\d+`)
    myAssignmentReport_RegexValidator.EncodedIssuer = regexp.MustCompile(`.*`)
    myAssignmentReport_RegexValidator.SecurityDesc = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.EncodedSecurityDescLen = regexp.MustCompile(`\d+`)
    myAssignmentReport_RegexValidator.EncodedSecurityDesc = regexp.MustCompile(`.*`)
    myAssignmentReport_RegexValidator.Pool = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.ContractSettlMonth = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myAssignmentReport_RegexValidator.CPProgram = regexp.MustCompile(`-?\d+`)
    myAssignmentReport_RegexValidator.CPRegType = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.DatedDate = regexp.MustCompile(`[0-9]{8}`)
    myAssignmentReport_RegexValidator.InterestAccrualDate = regexp.MustCompile(`[0-9]{8}`)
    myAssignmentReport_RegexValidator.SecurityStatus = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.SettleOnOpenFlag = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.InstrmtAssignmentMethod = regexp.MustCompile(`[^|]`)
    myAssignmentReport_RegexValidator.StrikeMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAssignmentReport_RegexValidator.StrikeValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAssignmentReport_RegexValidator.MinPriceIncrement = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAssignmentReport_RegexValidator.PositionLimit = regexp.MustCompile(`-?\d+`)
    myAssignmentReport_RegexValidator.NTPositionLimit = regexp.MustCompile(`-?\d+`)
    myAssignmentReport_RegexValidator.UnitofMeasure = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.TimeUnit = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.MaturityTime = regexp.MustCompile(`\d{2}:\d{2}(:\d{2})?(Z|[\+-]\d{2}(:\d{2})?)`)
    myAssignmentReport_RegexValidator.Currency = regexp.MustCompile(`[A-Z]{3}`)
    myAssignmentReport_RegexValidator.ThresholdAmount = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAssignmentReport_RegexValidator.SettlPrice = regexp.MustCompile(``)
    myAssignmentReport_RegexValidator.SettlPriceType = regexp.MustCompile(`-?\d+`)
    myAssignmentReport_RegexValidator.UnderlyingSettlPrice = regexp.MustCompile(``)
    myAssignmentReport_RegexValidator.ExpireDate = regexp.MustCompile(`[0-9]{8}`)
    myAssignmentReport_RegexValidator.AssignmentMethod = regexp.MustCompile(`[^|]`)
    myAssignmentReport_RegexValidator.AssignmentUnit = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAssignmentReport_RegexValidator.OpenInterest = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAssignmentReport_RegexValidator.ExerciseMethod = regexp.MustCompile(`[^|]`)
    myAssignmentReport_RegexValidator.SettlSessID = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.SettlSessSubID = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.ClearingBusinessDate = regexp.MustCompile(`[0-9]{8}`)
    myAssignmentReport_RegexValidator.Text = regexp.MustCompile(`[^|]*`)
    myAssignmentReport_RegexValidator.EncodedTextLen = regexp.MustCompile(`\d+`)
    myAssignmentReport_RegexValidator.EncodedText = regexp.MustCompile(`.*`)
    myAssignmentReport_RegexValidator.PriorSettlPrice = regexp.MustCompile(``)
    myAssignmentReport_RegexValidator.SignatureLength = regexp.MustCompile(`\d+`)
    myAssignmentReport_RegexValidator.Signature = regexp.MustCompile(`.*`)
    myAssignmentReport_RegexValidator.CheckSum = regexp.MustCompile(`[^|]*`)
}



func (m *AssignmentReport) HasRequiredFields() bool {
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
    
    
    
    valid = valid && m.AsgnRptID.HasValue()
    
    
    
    for _, g := range(m.PartyIDs) {
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
    
    
    
    for _, g := range(m.Legs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Underlyings) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Positions) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.PosAmt) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.ClearingBusinessDate.HasValue()
    
    
    
    valid = valid && m.CheckSum.HasValue()
    
    
    
    
    
    return valid
}



