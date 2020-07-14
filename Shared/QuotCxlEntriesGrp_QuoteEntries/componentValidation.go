package QuotCxlEntriesGrp_QuoteEntries

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import SecAltIDGrp_SecurityAltID "grgrbll/fix/Shared/SecAltIDGrp_SecurityAltID"

import EvntGrp_Events "grgrbll/fix/Shared/EvntGrp_Events"

import InstrumentParties_InstrumentParties "grgrbll/fix/Shared/InstrumentParties_InstrumentParties"

import UndInstrmtGrp_Underlyings "grgrbll/fix/Shared/UndInstrmtGrp_Underlyings"

import InstrmtLegGrp_Legs "grgrbll/fix/Shared/InstrmtLegGrp_Legs"


type quotCxlEntriesGrp_QuoteEntries_RegexValidator struct {
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
    _controlBlock fixMessageControlBlock
}


var myQuotCxlEntriesGrp_QuoteEntries_RegexValidator quotCxlEntriesGrp_QuoteEntries_RegexValidator

func init() {
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.Symbol = regexp.MustCompile(`[^|]*`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.SymbolSfx = regexp.MustCompile(`[^|]*`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.SecurityID = regexp.MustCompile(`[^|]*`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.SecurityIDSource = regexp.MustCompile(`[^|]*`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.Product = regexp.MustCompile(`-?\d+`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.CFICode = regexp.MustCompile(`[^|]*`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.SecurityType = regexp.MustCompile(`[^|]*`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.SecuritySubType = regexp.MustCompile(`[^|]*`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.MaturityMonthYear = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.MaturityDate = regexp.MustCompile(`[0-9]{8}`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.CouponPaymentDate = regexp.MustCompile(`[0-9]{8}`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.IssueDate = regexp.MustCompile(`[0-9]{8}`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.RepoCollateralSecurityType = regexp.MustCompile(`-?\d+`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.RepurchaseTerm = regexp.MustCompile(`-?\d+`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.RepurchaseRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.Factor = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.CreditRating = regexp.MustCompile(`[^|]*`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.InstrRegistry = regexp.MustCompile(`[^|]*`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.CountryOfIssue = regexp.MustCompile(`[A-Z]{2}`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.StateOrProvinceOfIssue = regexp.MustCompile(`[^|]*`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.LocaleOfIssue = regexp.MustCompile(`[^|]*`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.RedemptionDate = regexp.MustCompile(`[0-9]{8}`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.StrikePrice = regexp.MustCompile(``)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.StrikeCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.OptAttribute = regexp.MustCompile(`[^|]`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.ContractMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.CouponRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.SecurityExchange = regexp.MustCompile(`[A-Z0-9]{4}`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.Issuer = regexp.MustCompile(`[^|]*`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.EncodedIssuerLen = regexp.MustCompile(`\d+`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.EncodedIssuer = regexp.MustCompile(`.*`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.SecurityDesc = regexp.MustCompile(`[^|]*`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.EncodedSecurityDescLen = regexp.MustCompile(`\d+`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.EncodedSecurityDesc = regexp.MustCompile(`.*`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.Pool = regexp.MustCompile(`[^|]*`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.ContractSettlMonth = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.CPProgram = regexp.MustCompile(`-?\d+`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.CPRegType = regexp.MustCompile(`[^|]*`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.DatedDate = regexp.MustCompile(`[0-9]{8}`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.InterestAccrualDate = regexp.MustCompile(`[0-9]{8}`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.SecurityStatus = regexp.MustCompile(`[^|]*`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.SettleOnOpenFlag = regexp.MustCompile(`[^|]*`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.InstrmtAssignmentMethod = regexp.MustCompile(`[^|]`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.StrikeMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.StrikeValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.MinPriceIncrement = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.PositionLimit = regexp.MustCompile(`-?\d+`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.NTPositionLimit = regexp.MustCompile(`-?\d+`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.UnitofMeasure = regexp.MustCompile(`[^|]*`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.TimeUnit = regexp.MustCompile(`[^|]*`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.MaturityTime = regexp.MustCompile(`\d{2}:\d{2}(:\d{2})?(Z|[\+-]\d{2}(:\d{2})?)`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.AgreementDesc = regexp.MustCompile(`[^|]*`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.AgreementID = regexp.MustCompile(`[^|]*`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.AgreementDate = regexp.MustCompile(`[0-9]{8}`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.AgreementCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.TerminationType = regexp.MustCompile(`-?\d+`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.StartDate = regexp.MustCompile(`[0-9]{8}`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.EndDate = regexp.MustCompile(`[0-9]{8}`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.DeliveryType = regexp.MustCompile(`-?\d+`)
    myQuotCxlEntriesGrp_QuoteEntries_RegexValidator.MarginRatio = regexp.MustCompile(`-?\d*(\.\d*)?`)
}



func (m *QuotCxlEntriesGrp_QuoteEntries) HasRequiredFields() bool {
    valid := true
    
    

    
    
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
    
    
    
    
    
    return valid
}




