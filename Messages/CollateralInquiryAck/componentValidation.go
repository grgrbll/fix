package CollateralInquiryAck

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import CollInqQualGrp_CollInquiryQualifier "grgrbll/fix/Shared/CollInqQualGrp_CollInquiryQualifier"

import Parties_PartyIDs "grgrbll/fix/Shared/Parties_PartyIDs"

import ExecCollGrp_Execs "grgrbll/fix/Shared/ExecCollGrp_Execs"

import TrdCollGrp_Trades "grgrbll/fix/Shared/TrdCollGrp_Trades"

import SecAltIDGrp_SecurityAltID "grgrbll/fix/Shared/SecAltIDGrp_SecurityAltID"

import EvntGrp_Events "grgrbll/fix/Shared/EvntGrp_Events"

import InstrumentParties_InstrumentParties "grgrbll/fix/Shared/InstrumentParties_InstrumentParties"

import InstrmtLegGrp_Legs "grgrbll/fix/Shared/InstrmtLegGrp_Legs"

import UndInstrmtGrp_Underlyings "grgrbll/fix/Shared/UndInstrmtGrp_Underlyings"


type collateralInquiryAck_RegexValidator struct {
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
    CollInquiryID *(regexp.Regexp)
    CollInquiryStatus *(regexp.Regexp)
    CollInquiryResult *(regexp.Regexp)
    TotNumReports *(regexp.Regexp)
    Account *(regexp.Regexp)
    AccountType *(regexp.Regexp)
    ClOrdID *(regexp.Regexp)
    OrderID *(regexp.Regexp)
    SecondaryOrderID *(regexp.Regexp)
    SecondaryClOrdID *(regexp.Regexp)
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
    AgreementDesc *(regexp.Regexp)
    AgreementID *(regexp.Regexp)
    AgreementDate *(regexp.Regexp)
    AgreementCurrency *(regexp.Regexp)
    TerminationType *(regexp.Regexp)
    StartDate *(regexp.Regexp)
    EndDate *(regexp.Regexp)
    DeliveryType *(regexp.Regexp)
    MarginRatio *(regexp.Regexp)
    SettlDate *(regexp.Regexp)
    Quantity *(regexp.Regexp)
    QtyType *(regexp.Regexp)
    Currency *(regexp.Regexp)
    TradingSessionID *(regexp.Regexp)
    TradingSessionSubID *(regexp.Regexp)
    SettlSessID *(regexp.Regexp)
    SettlSessSubID *(regexp.Regexp)
    ClearingBusinessDate *(regexp.Regexp)
    ResponseTransportType *(regexp.Regexp)
    ResponseDestination *(regexp.Regexp)
    Text *(regexp.Regexp)
    EncodedTextLen *(regexp.Regexp)
    EncodedText *(regexp.Regexp)
    SignatureLength *(regexp.Regexp)
    Signature *(regexp.Regexp)
    CheckSum *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myCollateralInquiryAck_RegexValidator collateralInquiryAck_RegexValidator

func init() {
    myCollateralInquiryAck_RegexValidator
    myCollateralInquiryAck_RegexValidator.BeginString = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.BodyLength = regexp.MustCompile(`\d+`)
    myCollateralInquiryAck_RegexValidator.MsgType = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.SenderCompID = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.TargetCompID = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.OnBehalfOfCompID = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.DeliverToCompID = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.SecureDataLen = regexp.MustCompile(`\d+`)
    myCollateralInquiryAck_RegexValidator.SecureData = regexp.MustCompile(`.*`)
    myCollateralInquiryAck_RegexValidator.MsgSeqNum = regexp.MustCompile(`\d+`)
    myCollateralInquiryAck_RegexValidator.SenderSubID = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.SenderLocationID = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.TargetSubID = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.TargetLocationID = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.OnBehalfOfSubID = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.OnBehalfOfLocationID = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.DeliverToSubID = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.DeliverToLocationID = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.PossDupFlag = regexp.MustCompile(`[YN]`)
    myCollateralInquiryAck_RegexValidator.PossResend = regexp.MustCompile(`[YN]`)
    myCollateralInquiryAck_RegexValidator.SendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myCollateralInquiryAck_RegexValidator.OrigSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myCollateralInquiryAck_RegexValidator.XmlDataLen = regexp.MustCompile(`\d+`)
    myCollateralInquiryAck_RegexValidator.XmlData = regexp.MustCompile(`.*`)
    myCollateralInquiryAck_RegexValidator.MessageEncoding = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.LastMsgSeqNumProcessed = regexp.MustCompile(`\d+`)
    myCollateralInquiryAck_RegexValidator.ApplVerID = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.CstmApplVerID = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.CollInquiryID = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.CollInquiryStatus = regexp.MustCompile(`-?\d+`)
    myCollateralInquiryAck_RegexValidator.CollInquiryResult = regexp.MustCompile(`-?\d+`)
    myCollateralInquiryAck_RegexValidator.TotNumReports = regexp.MustCompile(`-?\d+`)
    myCollateralInquiryAck_RegexValidator.Account = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.AccountType = regexp.MustCompile(`-?\d+`)
    myCollateralInquiryAck_RegexValidator.ClOrdID = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.OrderID = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.SecondaryOrderID = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.SecondaryClOrdID = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.Symbol = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.SymbolSfx = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.SecurityID = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.SecurityIDSource = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.Product = regexp.MustCompile(`-?\d+`)
    myCollateralInquiryAck_RegexValidator.CFICode = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.SecurityType = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.SecuritySubType = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.MaturityMonthYear = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myCollateralInquiryAck_RegexValidator.MaturityDate = regexp.MustCompile(`[0-9]{8}`)
    myCollateralInquiryAck_RegexValidator.CouponPaymentDate = regexp.MustCompile(`[0-9]{8}`)
    myCollateralInquiryAck_RegexValidator.IssueDate = regexp.MustCompile(`[0-9]{8}`)
    myCollateralInquiryAck_RegexValidator.RepoCollateralSecurityType = regexp.MustCompile(`-?\d+`)
    myCollateralInquiryAck_RegexValidator.RepurchaseTerm = regexp.MustCompile(`-?\d+`)
    myCollateralInquiryAck_RegexValidator.RepurchaseRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCollateralInquiryAck_RegexValidator.Factor = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCollateralInquiryAck_RegexValidator.CreditRating = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.InstrRegistry = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.CountryOfIssue = regexp.MustCompile(`[A-Z]{2}`)
    myCollateralInquiryAck_RegexValidator.StateOrProvinceOfIssue = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.LocaleOfIssue = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.RedemptionDate = regexp.MustCompile(`[0-9]{8}`)
    myCollateralInquiryAck_RegexValidator.StrikePrice = regexp.MustCompile(``)
    myCollateralInquiryAck_RegexValidator.StrikeCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myCollateralInquiryAck_RegexValidator.OptAttribute = regexp.MustCompile(`[^|]`)
    myCollateralInquiryAck_RegexValidator.ContractMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCollateralInquiryAck_RegexValidator.CouponRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCollateralInquiryAck_RegexValidator.SecurityExchange = regexp.MustCompile(`[A-Z0-9]{4}`)
    myCollateralInquiryAck_RegexValidator.Issuer = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.EncodedIssuerLen = regexp.MustCompile(`\d+`)
    myCollateralInquiryAck_RegexValidator.EncodedIssuer = regexp.MustCompile(`.*`)
    myCollateralInquiryAck_RegexValidator.SecurityDesc = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.EncodedSecurityDescLen = regexp.MustCompile(`\d+`)
    myCollateralInquiryAck_RegexValidator.EncodedSecurityDesc = regexp.MustCompile(`.*`)
    myCollateralInquiryAck_RegexValidator.Pool = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.ContractSettlMonth = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myCollateralInquiryAck_RegexValidator.CPProgram = regexp.MustCompile(`-?\d+`)
    myCollateralInquiryAck_RegexValidator.CPRegType = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.DatedDate = regexp.MustCompile(`[0-9]{8}`)
    myCollateralInquiryAck_RegexValidator.InterestAccrualDate = regexp.MustCompile(`[0-9]{8}`)
    myCollateralInquiryAck_RegexValidator.SecurityStatus = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.SettleOnOpenFlag = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.InstrmtAssignmentMethod = regexp.MustCompile(`[^|]`)
    myCollateralInquiryAck_RegexValidator.StrikeMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCollateralInquiryAck_RegexValidator.StrikeValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCollateralInquiryAck_RegexValidator.MinPriceIncrement = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCollateralInquiryAck_RegexValidator.PositionLimit = regexp.MustCompile(`-?\d+`)
    myCollateralInquiryAck_RegexValidator.NTPositionLimit = regexp.MustCompile(`-?\d+`)
    myCollateralInquiryAck_RegexValidator.UnitofMeasure = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.TimeUnit = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.MaturityTime = regexp.MustCompile(`\d{2}:\d{2}(:\d{2})?(Z|[\+-]\d{2}(:\d{2})?)`)
    myCollateralInquiryAck_RegexValidator.AgreementDesc = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.AgreementID = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.AgreementDate = regexp.MustCompile(`[0-9]{8}`)
    myCollateralInquiryAck_RegexValidator.AgreementCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myCollateralInquiryAck_RegexValidator.TerminationType = regexp.MustCompile(`-?\d+`)
    myCollateralInquiryAck_RegexValidator.StartDate = regexp.MustCompile(`[0-9]{8}`)
    myCollateralInquiryAck_RegexValidator.EndDate = regexp.MustCompile(`[0-9]{8}`)
    myCollateralInquiryAck_RegexValidator.DeliveryType = regexp.MustCompile(`-?\d+`)
    myCollateralInquiryAck_RegexValidator.MarginRatio = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCollateralInquiryAck_RegexValidator.SettlDate = regexp.MustCompile(`[0-9]{8}`)
    myCollateralInquiryAck_RegexValidator.Quantity = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCollateralInquiryAck_RegexValidator.QtyType = regexp.MustCompile(`-?\d+`)
    myCollateralInquiryAck_RegexValidator.Currency = regexp.MustCompile(`[A-Z]{3}`)
    myCollateralInquiryAck_RegexValidator.TradingSessionID = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.TradingSessionSubID = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.SettlSessID = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.SettlSessSubID = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.ClearingBusinessDate = regexp.MustCompile(`[0-9]{8}`)
    myCollateralInquiryAck_RegexValidator.ResponseTransportType = regexp.MustCompile(`-?\d+`)
    myCollateralInquiryAck_RegexValidator.ResponseDestination = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.Text = regexp.MustCompile(`[^|]*`)
    myCollateralInquiryAck_RegexValidator.EncodedTextLen = regexp.MustCompile(`\d+`)
    myCollateralInquiryAck_RegexValidator.EncodedText = regexp.MustCompile(`.*`)
    myCollateralInquiryAck_RegexValidator.SignatureLength = regexp.MustCompile(`\d+`)
    myCollateralInquiryAck_RegexValidator.Signature = regexp.MustCompile(`.*`)
    myCollateralInquiryAck_RegexValidator.CheckSum = regexp.MustCompile(`[^|]*`)
}



func (m *CollateralInquiryAck) HasRequiredFields() bool {
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
    
    
    
    valid = valid && m.CollInquiryID.HasValue()
    
    
    
    valid = valid && m.CollInquiryStatus.HasValue()
    
    
    
    for _, g := range(m.CollInquiryQualifier) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.PartyIDs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Execs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Trades) {
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
    
    
    
    valid = valid && m.CheckSum.HasValue()
    
    
    
    
    
    return valid
}



