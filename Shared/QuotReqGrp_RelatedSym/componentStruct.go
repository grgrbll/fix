package QuotReqGrp_RelatedSym

import "bufio"
import "grgrbll/Fix/FixDef/Types"

// Autogenerated at , do not edit



import SecAltIDGrp_SecurityAltID "grgrbll/fix/Shared/SecAltIDGrp_SecurityAltID"

import EvntGrp_Events "grgrbll/fix/Shared/EvntGrp_Events"

import InstrumentParties_InstrumentParties "grgrbll/fix/Shared/InstrumentParties_InstrumentParties"

import UndInstrmtGrp_Underlyings "grgrbll/fix/Shared/UndInstrmtGrp_Underlyings"

import Stipulations_Stipulations "grgrbll/fix/Shared/Stipulations_Stipulations"

import QuotReqLegsGrp_Legs "grgrbll/fix/Shared/QuotReqLegsGrp_Legs"

import QuotQualGrp_QuoteQualifiers "grgrbll/fix/Shared/QuotQualGrp_QuoteQualifiers"

import Parties_PartyIDs "grgrbll/fix/Shared/Parties_PartyIDs"


type QuotReqGrp_RelatedSym struct {
    Symbol Types.String `id:"55", required:"N" `
    SymbolSfx Types.String `id:"65", required:"N" `
    SecurityID Types.String `id:"48", required:"N" `
    SecurityIDSource Types.String `id:"22", required:"N" `
    SecurityAltID []SecAltIDGrp_SecurityAltID.SecAltIDGrp_SecurityAltID `counter_name:"NoSecurityAltID", counter_id:"454"`
    Product Types.Int `id:"460", required:"N" `
    CFICode Types.String `id:"461", required:"N" `
    SecurityType Types.String `id:"167", required:"N" `
    SecuritySubType Types.String `id:"762", required:"N" `
    MaturityMonthYear Types.MonthYear `id:"200", required:"N" `
    MaturityDate Types.LocalMktDate `id:"541", required:"N" `
    CouponPaymentDate Types.LocalMktDate `id:"224", required:"N" `
    IssueDate Types.LocalMktDate `id:"225", required:"N" `
    RepoCollateralSecurityType Types.Int `id:"239", required:"N" `
    RepurchaseTerm Types.Int `id:"226", required:"N" `
    RepurchaseRate Types.Percentage `id:"227", required:"N" `
    Factor Types.Float `id:"228", required:"N" `
    CreditRating Types.String `id:"255", required:"N" `
    InstrRegistry Types.String `id:"543", required:"N" `
    CountryOfIssue Types.Country `id:"470", required:"N" `
    StateOrProvinceOfIssue Types.String `id:"471", required:"N" `
    LocaleOfIssue Types.String `id:"472", required:"N" `
    RedemptionDate Types.LocalMktDate `id:"240", required:"N" `
    StrikePrice Types.Price `id:"202", required:"N" `
    StrikeCurrency Types.Currency `id:"947", required:"N" `
    OptAttribute Types.Char `id:"206", required:"N" `
    ContractMultiplier Types.Float `id:"231", required:"N" `
    CouponRate Types.Percentage `id:"223", required:"N" `
    SecurityExchange Types.Exchange `id:"207", required:"N" `
    Issuer Types.String `id:"106", required:"N" `
    EncodedIssuerLen Types.Length `id:"348", required:"N" `
    EncodedIssuer Types.Data `id:"349", required:"N" `
    SecurityDesc Types.String `id:"107", required:"N" `
    EncodedSecurityDescLen Types.Length `id:"350", required:"N" `
    EncodedSecurityDesc Types.Data `id:"351", required:"N" `
    Pool Types.String `id:"691", required:"N" `
    ContractSettlMonth Types.MonthYear `id:"667", required:"N" `
    CPProgram Types.Int `id:"875", required:"N" `
    CPRegType Types.String `id:"876", required:"N" `
    Events []EvntGrp_Events.EvntGrp_Events `counter_name:"NoEvents", counter_id:"864"`
    DatedDate Types.LocalMktDate `id:"873", required:"N" `
    InterestAccrualDate Types.LocalMktDate `id:"874", required:"N" `
    SecurityStatus Types.String `id:"965", required:"N" `
    SettleOnOpenFlag Types.String `id:"966", required:"N" `
    InstrmtAssignmentMethod Types.Char `id:"1049", required:"N" `
    StrikeMultiplier Types.Float `id:"967", required:"N" `
    StrikeValue Types.Float `id:"968", required:"N" `
    MinPriceIncrement Types.Float `id:"969", required:"N" `
    PositionLimit Types.Int `id:"970", required:"N" `
    NTPositionLimit Types.Int `id:"971", required:"N" `
    InstrumentParties []InstrumentParties_InstrumentParties.InstrumentParties_InstrumentParties `counter_name:"NoInstrumentParties", counter_id:"1018"`
    UnitofMeasure Types.String `id:"996", required:"N" `
    TimeUnit Types.String `id:"997", required:"N" `
    MaturityTime Types.TZTimeOnly `id:"1079", required:"N" `
    AgreementDesc Types.String `id:"913", required:"N" `
    AgreementID Types.String `id:"914", required:"N" `
    AgreementDate Types.LocalMktDate `id:"915", required:"N" `
    AgreementCurrency Types.Currency `id:"918", required:"N" `
    TerminationType Types.Int `id:"788", required:"N" `
    StartDate Types.LocalMktDate `id:"916", required:"N" `
    EndDate Types.LocalMktDate `id:"917", required:"N" `
    DeliveryType Types.Int `id:"919", required:"N" `
    MarginRatio Types.Percentage `id:"898", required:"N" `
    Underlyings []UndInstrmtGrp_Underlyings.UndInstrmtGrp_Underlyings `counter_name:"NoUnderlyings", counter_id:"711"`
    PrevClosePx Types.Price `id:"140", required:"N" `
    QuoteRequestType Types.Int `id:"303", required:"N" `
    QuoteType Types.Int `id:"537", required:"N" `
    TradingSessionID Types.String `id:"336", required:"N" `
    TradingSessionSubID Types.String `id:"625", required:"N" `
    TradeOriginationDate Types.LocalMktDate `id:"229", required:"N" `
    Side Types.Char `id:"54", required:"N" `
    QtyType Types.Int `id:"854", required:"N" `
    OrderQty Types.Qty `id:"38", required:"N" `
    CashOrderQty Types.Qty `id:"152", required:"N" `
    OrderPercent Types.Percentage `id:"516", required:"N" `
    RoundingDirection Types.Char `id:"468", required:"N" `
    RoundingModulus Types.Float `id:"469", required:"N" `
    SettlType Types.String `id:"63", required:"N" `
    SettlDate Types.LocalMktDate `id:"64", required:"N" `
    SettlDate2 Types.LocalMktDate `id:"193", required:"N" `
    OrderQty2 Types.Qty `id:"192", required:"N" `
    Currency Types.Currency `id:"15", required:"N" `
    Stipulations []Stipulations_Stipulations.Stipulations_Stipulations `counter_name:"NoStipulations", counter_id:"232"`
    Account Types.String `id:"1", required:"N" `
    AcctIDSource Types.Int `id:"660", required:"N" `
    AccountType Types.Int `id:"581", required:"N" `
    Legs []QuotReqLegsGrp_Legs.QuotReqLegsGrp_Legs `counter_name:"NoLegs", counter_id:"555"`
    QuoteQualifiers []QuotQualGrp_QuoteQualifiers.QuotQualGrp_QuoteQualifiers `counter_name:"NoQuoteQualifiers", counter_id:"735"`
    QuotePriceType Types.Int `id:"692", required:"N" `
    OrdType Types.Char `id:"40", required:"N" `
    ValidUntilTime Types.UTCTimestamp `id:"62", required:"N" `
    ExpireTime Types.UTCTimestamp `id:"126", required:"N" `
    TransactTime Types.UTCTimestamp `id:"60", required:"N" `
    Spread Types.PriceOffset `id:"218", required:"N" `
    BenchmarkCurveCurrency Types.Currency `id:"220", required:"N" `
    BenchmarkCurveName Types.String `id:"221", required:"N" `
    BenchmarkCurvePoint Types.String `id:"222", required:"N" `
    BenchmarkPrice Types.Price `id:"662", required:"N" `
    BenchmarkPriceType Types.Int `id:"663", required:"N" `
    BenchmarkSecurityID Types.String `id:"699", required:"N" `
    BenchmarkSecurityIDSource Types.String `id:"761", required:"N" `
    PriceType Types.Int `id:"423", required:"N" `
    Price Types.Price `id:"44", required:"N" `
    Price2 Types.Price `id:"640", required:"N" `
    YieldType Types.String `id:"235", required:"N" `
    Yield Types.Percentage `id:"236", required:"N" `
    YieldCalcDate Types.LocalMktDate `id:"701", required:"N" `
    YieldRedemptionDate Types.LocalMktDate `id:"696", required:"N" `
    YieldRedemptionPrice Types.Price `id:"697", required:"N" `
    YieldRedemptionPriceType Types.Int `id:"698", required:"N" `
    PartyIDs []Parties_PartyIDs.Parties_PartyIDs `counter_name:"NoPartyIDs", counter_id:"453"`
    _controlBlock fixMessageControlBlock
}
