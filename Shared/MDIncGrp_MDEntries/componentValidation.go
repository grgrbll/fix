package MDIncGrp_MDEntries

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import SecAltIDGrp_SecurityAltID "grgrbll/fix/Shared/SecAltIDGrp_SecurityAltID"

import EvntGrp_Events "grgrbll/fix/Shared/EvntGrp_Events"

import InstrumentParties_InstrumentParties "grgrbll/fix/Shared/InstrumentParties_InstrumentParties"

import UndInstrmtGrp_Underlyings "grgrbll/fix/Shared/UndInstrmtGrp_Underlyings"

import InstrmtLegGrp_Legs "grgrbll/fix/Shared/InstrmtLegGrp_Legs"

import Parties_PartyIDs "grgrbll/fix/Shared/Parties_PartyIDs"


type mDIncGrp_MDEntries_RegexValidator struct {
    MDUpdateAction *(regexp.Regexp)
    DeleteReason *(regexp.Regexp)
    MDEntryType *(regexp.Regexp)
    MDEntryID *(regexp.Regexp)
    MDEntryRefID *(regexp.Regexp)
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
    FinancialStatus *(regexp.Regexp)
    CorporateAction *(regexp.Regexp)
    MDEntryPx *(regexp.Regexp)
    Currency *(regexp.Regexp)
    MDEntrySize *(regexp.Regexp)
    MDEntryDate *(regexp.Regexp)
    MDEntryTime *(regexp.Regexp)
    TickDirection *(regexp.Regexp)
    MDMkt *(regexp.Regexp)
    TradingSessionID *(regexp.Regexp)
    TradingSessionSubID *(regexp.Regexp)
    QuoteCondition *(regexp.Regexp)
    TradeCondition *(regexp.Regexp)
    MDEntryOriginator *(regexp.Regexp)
    LocationID *(regexp.Regexp)
    DeskID *(regexp.Regexp)
    OpenCloseSettlFlag *(regexp.Regexp)
    TimeInForce *(regexp.Regexp)
    ExpireDate *(regexp.Regexp)
    ExpireTime *(regexp.Regexp)
    MinQty *(regexp.Regexp)
    ExecInst *(regexp.Regexp)
    SellerDays *(regexp.Regexp)
    OrderID *(regexp.Regexp)
    QuoteEntryID *(regexp.Regexp)
    MDEntryBuyer *(regexp.Regexp)
    MDEntrySeller *(regexp.Regexp)
    NumberOfOrders *(regexp.Regexp)
    MDEntryPositionNo *(regexp.Regexp)
    Scope *(regexp.Regexp)
    PriceDelta *(regexp.Regexp)
    NetChgPrevDay *(regexp.Regexp)
    Text *(regexp.Regexp)
    EncodedTextLen *(regexp.Regexp)
    EncodedText *(regexp.Regexp)
    OrderCapacity *(regexp.Regexp)
    MDOriginType *(regexp.Regexp)
    HighPx *(regexp.Regexp)
    LowPx *(regexp.Regexp)
    TradeVolume *(regexp.Regexp)
    SettlType *(regexp.Regexp)
    SettlDate *(regexp.Regexp)
    MDQuoteType *(regexp.Regexp)
    RptSeq *(regexp.Regexp)
    DealingCapacity *(regexp.Regexp)
    MDEntrySpotRate *(regexp.Regexp)
    MDEntryForwardPoints *(regexp.Regexp)
    MDPriceLevel *(regexp.Regexp)
    SecondaryOrderID *(regexp.Regexp)
    OrdType *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myMDIncGrp_MDEntries_RegexValidator mDIncGrp_MDEntries_RegexValidator

func init() {
    myMDIncGrp_MDEntries_RegexValidator
    myMDIncGrp_MDEntries_RegexValidator.MDUpdateAction = regexp.MustCompile(`[^|]`)
    myMDIncGrp_MDEntries_RegexValidator.DeleteReason = regexp.MustCompile(`[^|]`)
    myMDIncGrp_MDEntries_RegexValidator.MDEntryType = regexp.MustCompile(`[^|]`)
    myMDIncGrp_MDEntries_RegexValidator.MDEntryID = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.MDEntryRefID = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.Symbol = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.SymbolSfx = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.SecurityID = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.SecurityIDSource = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.Product = regexp.MustCompile(`-?\d+`)
    myMDIncGrp_MDEntries_RegexValidator.CFICode = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.SecurityType = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.SecuritySubType = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.MaturityMonthYear = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myMDIncGrp_MDEntries_RegexValidator.MaturityDate = regexp.MustCompile(`[0-9]{8}`)
    myMDIncGrp_MDEntries_RegexValidator.CouponPaymentDate = regexp.MustCompile(`[0-9]{8}`)
    myMDIncGrp_MDEntries_RegexValidator.IssueDate = regexp.MustCompile(`[0-9]{8}`)
    myMDIncGrp_MDEntries_RegexValidator.RepoCollateralSecurityType = regexp.MustCompile(`-?\d+`)
    myMDIncGrp_MDEntries_RegexValidator.RepurchaseTerm = regexp.MustCompile(`-?\d+`)
    myMDIncGrp_MDEntries_RegexValidator.RepurchaseRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myMDIncGrp_MDEntries_RegexValidator.Factor = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myMDIncGrp_MDEntries_RegexValidator.CreditRating = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.InstrRegistry = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.CountryOfIssue = regexp.MustCompile(`[A-Z]{2}`)
    myMDIncGrp_MDEntries_RegexValidator.StateOrProvinceOfIssue = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.LocaleOfIssue = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.RedemptionDate = regexp.MustCompile(`[0-9]{8}`)
    myMDIncGrp_MDEntries_RegexValidator.StrikePrice = regexp.MustCompile(``)
    myMDIncGrp_MDEntries_RegexValidator.StrikeCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myMDIncGrp_MDEntries_RegexValidator.OptAttribute = regexp.MustCompile(`[^|]`)
    myMDIncGrp_MDEntries_RegexValidator.ContractMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myMDIncGrp_MDEntries_RegexValidator.CouponRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myMDIncGrp_MDEntries_RegexValidator.SecurityExchange = regexp.MustCompile(`[A-Z0-9]{4}`)
    myMDIncGrp_MDEntries_RegexValidator.Issuer = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.EncodedIssuerLen = regexp.MustCompile(`\d+`)
    myMDIncGrp_MDEntries_RegexValidator.EncodedIssuer = regexp.MustCompile(`.*`)
    myMDIncGrp_MDEntries_RegexValidator.SecurityDesc = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.EncodedSecurityDescLen = regexp.MustCompile(`\d+`)
    myMDIncGrp_MDEntries_RegexValidator.EncodedSecurityDesc = regexp.MustCompile(`.*`)
    myMDIncGrp_MDEntries_RegexValidator.Pool = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.ContractSettlMonth = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myMDIncGrp_MDEntries_RegexValidator.CPProgram = regexp.MustCompile(`-?\d+`)
    myMDIncGrp_MDEntries_RegexValidator.CPRegType = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.DatedDate = regexp.MustCompile(`[0-9]{8}`)
    myMDIncGrp_MDEntries_RegexValidator.InterestAccrualDate = regexp.MustCompile(`[0-9]{8}`)
    myMDIncGrp_MDEntries_RegexValidator.SecurityStatus = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.SettleOnOpenFlag = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.InstrmtAssignmentMethod = regexp.MustCompile(`[^|]`)
    myMDIncGrp_MDEntries_RegexValidator.StrikeMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myMDIncGrp_MDEntries_RegexValidator.StrikeValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myMDIncGrp_MDEntries_RegexValidator.MinPriceIncrement = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myMDIncGrp_MDEntries_RegexValidator.PositionLimit = regexp.MustCompile(`-?\d+`)
    myMDIncGrp_MDEntries_RegexValidator.NTPositionLimit = regexp.MustCompile(`-?\d+`)
    myMDIncGrp_MDEntries_RegexValidator.UnitofMeasure = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.TimeUnit = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.MaturityTime = regexp.MustCompile(`\d{2}:\d{2}(:\d{2})?(Z|[\+-]\d{2}(:\d{2})?)`)
    myMDIncGrp_MDEntries_RegexValidator.FinancialStatus = regexp.MustCompile(`[ ]?(\w[ ])+\w?`)
    myMDIncGrp_MDEntries_RegexValidator.CorporateAction = regexp.MustCompile(`[ ]?(\w[ ])+\w?`)
    myMDIncGrp_MDEntries_RegexValidator.MDEntryPx = regexp.MustCompile(``)
    myMDIncGrp_MDEntries_RegexValidator.Currency = regexp.MustCompile(`[A-Z]{3}`)
    myMDIncGrp_MDEntries_RegexValidator.MDEntrySize = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myMDIncGrp_MDEntries_RegexValidator.MDEntryDate = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])`)
    myMDIncGrp_MDEntries_RegexValidator.MDEntryTime = regexp.MustCompile(`([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myMDIncGrp_MDEntries_RegexValidator.TickDirection = regexp.MustCompile(`[^|]`)
    myMDIncGrp_MDEntries_RegexValidator.MDMkt = regexp.MustCompile(`[A-Z0-9]{4}`)
    myMDIncGrp_MDEntries_RegexValidator.TradingSessionID = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.TradingSessionSubID = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.QuoteCondition = regexp.MustCompile(`[ ]?(\w+[ ])+\w+?`)
    myMDIncGrp_MDEntries_RegexValidator.TradeCondition = regexp.MustCompile(`[ ]?(\w+[ ])+\w+?`)
    myMDIncGrp_MDEntries_RegexValidator.MDEntryOriginator = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.LocationID = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.DeskID = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.OpenCloseSettlFlag = regexp.MustCompile(`[ ]?(\w[ ])+\w?`)
    myMDIncGrp_MDEntries_RegexValidator.TimeInForce = regexp.MustCompile(`[^|]`)
    myMDIncGrp_MDEntries_RegexValidator.ExpireDate = regexp.MustCompile(`[0-9]{8}`)
    myMDIncGrp_MDEntries_RegexValidator.ExpireTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myMDIncGrp_MDEntries_RegexValidator.MinQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myMDIncGrp_MDEntries_RegexValidator.ExecInst = regexp.MustCompile(`[ ]?(\w[ ])+\w?`)
    myMDIncGrp_MDEntries_RegexValidator.SellerDays = regexp.MustCompile(`-?\d+`)
    myMDIncGrp_MDEntries_RegexValidator.OrderID = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.QuoteEntryID = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.MDEntryBuyer = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.MDEntrySeller = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.NumberOfOrders = regexp.MustCompile(`-?\d+`)
    myMDIncGrp_MDEntries_RegexValidator.MDEntryPositionNo = regexp.MustCompile(`-?\d+`)
    myMDIncGrp_MDEntries_RegexValidator.Scope = regexp.MustCompile(`[ ]?(\w[ ])+\w?`)
    myMDIncGrp_MDEntries_RegexValidator.PriceDelta = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myMDIncGrp_MDEntries_RegexValidator.NetChgPrevDay = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myMDIncGrp_MDEntries_RegexValidator.Text = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.EncodedTextLen = regexp.MustCompile(`\d+`)
    myMDIncGrp_MDEntries_RegexValidator.EncodedText = regexp.MustCompile(`.*`)
    myMDIncGrp_MDEntries_RegexValidator.OrderCapacity = regexp.MustCompile(`[^|]`)
    myMDIncGrp_MDEntries_RegexValidator.MDOriginType = regexp.MustCompile(`-?\d+`)
    myMDIncGrp_MDEntries_RegexValidator.HighPx = regexp.MustCompile(``)
    myMDIncGrp_MDEntries_RegexValidator.LowPx = regexp.MustCompile(``)
    myMDIncGrp_MDEntries_RegexValidator.TradeVolume = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myMDIncGrp_MDEntries_RegexValidator.SettlType = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.SettlDate = regexp.MustCompile(`[0-9]{8}`)
    myMDIncGrp_MDEntries_RegexValidator.MDQuoteType = regexp.MustCompile(`-?\d+`)
    myMDIncGrp_MDEntries_RegexValidator.RptSeq = regexp.MustCompile(`-?\d+`)
    myMDIncGrp_MDEntries_RegexValidator.DealingCapacity = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myMDIncGrp_MDEntries_RegexValidator.MDEntrySpotRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myMDIncGrp_MDEntries_RegexValidator.MDEntryForwardPoints = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myMDIncGrp_MDEntries_RegexValidator.MDPriceLevel = regexp.MustCompile(`-?\d+`)
    myMDIncGrp_MDEntries_RegexValidator.SecondaryOrderID = regexp.MustCompile(`[^|]*`)
    myMDIncGrp_MDEntries_RegexValidator.OrdType = regexp.MustCompile(`[^|]`)
}



func (m *MDIncGrp_MDEntries) HasRequiredFields() bool {
    valid := true
    
    

    
    
    valid = valid && m.MDUpdateAction.HasValue()
    
    
    
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
    
    
    
    
    
    return valid
}




