package QuoteStatusRequest

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import SecAltIDGrp_SecurityAltID "grgrbll/fix/Shared/SecAltIDGrp_SecurityAltID"

import EvntGrp_Events "grgrbll/fix/Shared/EvntGrp_Events"

import InstrumentParties_InstrumentParties "grgrbll/fix/Shared/InstrumentParties_InstrumentParties"

import UndInstrmtGrp_Underlyings "grgrbll/fix/Shared/UndInstrmtGrp_Underlyings"

import InstrmtLegGrp_Legs "grgrbll/fix/Shared/InstrmtLegGrp_Legs"

import Parties_PartyIDs "grgrbll/fix/Shared/Parties_PartyIDs"


type quoteStatusRequest_RegexValidator struct {
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
    QuoteStatusReqID *(regexp.Regexp)
    QuoteID *(regexp.Regexp)
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
    Account *(regexp.Regexp)
    AcctIDSource *(regexp.Regexp)
    AccountType *(regexp.Regexp)
    TradingSessionID *(regexp.Regexp)
    TradingSessionSubID *(regexp.Regexp)
    SubscriptionRequestType *(regexp.Regexp)
    SignatureLength *(regexp.Regexp)
    Signature *(regexp.Regexp)
    CheckSum *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myQuoteStatusRequest_RegexValidator quoteStatusRequest_RegexValidator

func init() {
    myQuoteStatusRequest_RegexValidator
    myQuoteStatusRequest_RegexValidator.BeginString = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.BodyLength = regexp.MustCompile(`\d+`)
    myQuoteStatusRequest_RegexValidator.MsgType = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.SenderCompID = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.TargetCompID = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.OnBehalfOfCompID = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.DeliverToCompID = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.SecureDataLen = regexp.MustCompile(`\d+`)
    myQuoteStatusRequest_RegexValidator.SecureData = regexp.MustCompile(`.*`)
    myQuoteStatusRequest_RegexValidator.MsgSeqNum = regexp.MustCompile(`\d+`)
    myQuoteStatusRequest_RegexValidator.SenderSubID = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.SenderLocationID = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.TargetSubID = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.TargetLocationID = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.OnBehalfOfSubID = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.OnBehalfOfLocationID = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.DeliverToSubID = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.DeliverToLocationID = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.PossDupFlag = regexp.MustCompile(`[YN]`)
    myQuoteStatusRequest_RegexValidator.PossResend = regexp.MustCompile(`[YN]`)
    myQuoteStatusRequest_RegexValidator.SendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myQuoteStatusRequest_RegexValidator.OrigSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myQuoteStatusRequest_RegexValidator.XmlDataLen = regexp.MustCompile(`\d+`)
    myQuoteStatusRequest_RegexValidator.XmlData = regexp.MustCompile(`.*`)
    myQuoteStatusRequest_RegexValidator.MessageEncoding = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.LastMsgSeqNumProcessed = regexp.MustCompile(`\d+`)
    myQuoteStatusRequest_RegexValidator.ApplVerID = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.CstmApplVerID = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.QuoteStatusReqID = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.QuoteID = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.Symbol = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.SymbolSfx = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.SecurityID = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.SecurityIDSource = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.Product = regexp.MustCompile(`-?\d+`)
    myQuoteStatusRequest_RegexValidator.CFICode = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.SecurityType = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.SecuritySubType = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.MaturityMonthYear = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myQuoteStatusRequest_RegexValidator.MaturityDate = regexp.MustCompile(`[0-9]{8}`)
    myQuoteStatusRequest_RegexValidator.CouponPaymentDate = regexp.MustCompile(`[0-9]{8}`)
    myQuoteStatusRequest_RegexValidator.IssueDate = regexp.MustCompile(`[0-9]{8}`)
    myQuoteStatusRequest_RegexValidator.RepoCollateralSecurityType = regexp.MustCompile(`-?\d+`)
    myQuoteStatusRequest_RegexValidator.RepurchaseTerm = regexp.MustCompile(`-?\d+`)
    myQuoteStatusRequest_RegexValidator.RepurchaseRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuoteStatusRequest_RegexValidator.Factor = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuoteStatusRequest_RegexValidator.CreditRating = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.InstrRegistry = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.CountryOfIssue = regexp.MustCompile(`[A-Z]{2}`)
    myQuoteStatusRequest_RegexValidator.StateOrProvinceOfIssue = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.LocaleOfIssue = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.RedemptionDate = regexp.MustCompile(`[0-9]{8}`)
    myQuoteStatusRequest_RegexValidator.StrikePrice = regexp.MustCompile(``)
    myQuoteStatusRequest_RegexValidator.StrikeCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myQuoteStatusRequest_RegexValidator.OptAttribute = regexp.MustCompile(`[^|]`)
    myQuoteStatusRequest_RegexValidator.ContractMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuoteStatusRequest_RegexValidator.CouponRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuoteStatusRequest_RegexValidator.SecurityExchange = regexp.MustCompile(`[A-Z0-9]{4}`)
    myQuoteStatusRequest_RegexValidator.Issuer = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.EncodedIssuerLen = regexp.MustCompile(`\d+`)
    myQuoteStatusRequest_RegexValidator.EncodedIssuer = regexp.MustCompile(`.*`)
    myQuoteStatusRequest_RegexValidator.SecurityDesc = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.EncodedSecurityDescLen = regexp.MustCompile(`\d+`)
    myQuoteStatusRequest_RegexValidator.EncodedSecurityDesc = regexp.MustCompile(`.*`)
    myQuoteStatusRequest_RegexValidator.Pool = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.ContractSettlMonth = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myQuoteStatusRequest_RegexValidator.CPProgram = regexp.MustCompile(`-?\d+`)
    myQuoteStatusRequest_RegexValidator.CPRegType = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.DatedDate = regexp.MustCompile(`[0-9]{8}`)
    myQuoteStatusRequest_RegexValidator.InterestAccrualDate = regexp.MustCompile(`[0-9]{8}`)
    myQuoteStatusRequest_RegexValidator.SecurityStatus = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.SettleOnOpenFlag = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.InstrmtAssignmentMethod = regexp.MustCompile(`[^|]`)
    myQuoteStatusRequest_RegexValidator.StrikeMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuoteStatusRequest_RegexValidator.StrikeValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuoteStatusRequest_RegexValidator.MinPriceIncrement = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuoteStatusRequest_RegexValidator.PositionLimit = regexp.MustCompile(`-?\d+`)
    myQuoteStatusRequest_RegexValidator.NTPositionLimit = regexp.MustCompile(`-?\d+`)
    myQuoteStatusRequest_RegexValidator.UnitofMeasure = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.TimeUnit = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.MaturityTime = regexp.MustCompile(`\d{2}:\d{2}(:\d{2})?(Z|[\+-]\d{2}(:\d{2})?)`)
    myQuoteStatusRequest_RegexValidator.AgreementDesc = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.AgreementID = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.AgreementDate = regexp.MustCompile(`[0-9]{8}`)
    myQuoteStatusRequest_RegexValidator.AgreementCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myQuoteStatusRequest_RegexValidator.TerminationType = regexp.MustCompile(`-?\d+`)
    myQuoteStatusRequest_RegexValidator.StartDate = regexp.MustCompile(`[0-9]{8}`)
    myQuoteStatusRequest_RegexValidator.EndDate = regexp.MustCompile(`[0-9]{8}`)
    myQuoteStatusRequest_RegexValidator.DeliveryType = regexp.MustCompile(`-?\d+`)
    myQuoteStatusRequest_RegexValidator.MarginRatio = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuoteStatusRequest_RegexValidator.Account = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.AcctIDSource = regexp.MustCompile(`-?\d+`)
    myQuoteStatusRequest_RegexValidator.AccountType = regexp.MustCompile(`-?\d+`)
    myQuoteStatusRequest_RegexValidator.TradingSessionID = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.TradingSessionSubID = regexp.MustCompile(`[^|]*`)
    myQuoteStatusRequest_RegexValidator.SubscriptionRequestType = regexp.MustCompile(`[^|]`)
    myQuoteStatusRequest_RegexValidator.SignatureLength = regexp.MustCompile(`\d+`)
    myQuoteStatusRequest_RegexValidator.Signature = regexp.MustCompile(`.*`)
    myQuoteStatusRequest_RegexValidator.CheckSum = regexp.MustCompile(`[^|]*`)
}



func (m *QuoteStatusRequest) HasRequiredFields() bool {
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
    
    
    
    for _, g := range(m.Underlyings) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Legs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.PartyIDs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.CheckSum.HasValue()
    
    
    
    
    
    return valid
}



