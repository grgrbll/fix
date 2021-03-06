package CrossOrderCancelReplaceRequest

import "bufio"
import "grgrbll/Fix/FixDef/Types"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import SideCrossOrdModGrp_Sides "grgrbll/fix/Shared/SideCrossOrdModGrp_Sides"

import SecAltIDGrp_SecurityAltID "grgrbll/fix/Shared/SecAltIDGrp_SecurityAltID"

import EvntGrp_Events "grgrbll/fix/Shared/EvntGrp_Events"

import InstrumentParties_InstrumentParties "grgrbll/fix/Shared/InstrumentParties_InstrumentParties"

import UndInstrmtGrp_Underlyings "grgrbll/fix/Shared/UndInstrmtGrp_Underlyings"

import InstrmtLegGrp_Legs "grgrbll/fix/Shared/InstrmtLegGrp_Legs"

import TrdgSesGrp_TradingSessions "grgrbll/fix/Shared/TrdgSesGrp_TradingSessions"

import Stipulations_Stipulations "grgrbll/fix/Shared/Stipulations_Stipulations"

import StrategyParametersGrp_StrategyParameters "grgrbll/fix/Shared/StrategyParametersGrp_StrategyParameters"

import RootParties_RootPartyIDs "grgrbll/fix/Shared/RootParties_RootPartyIDs"


type CrossOrderCancelReplaceRequest struct {
    BeginString Types.String `id:"8", required:"Y" `
    BodyLength Types.Length `id:"9", required:"Y" `
    MsgType Types.String `id:"35", required:"Y" `
    SenderCompID Types.String `id:"49", required:"Y" `
    TargetCompID Types.String `id:"56", required:"Y" `
    OnBehalfOfCompID Types.String `id:"115", required:"N" `
    DeliverToCompID Types.String `id:"128", required:"N" `
    SecureDataLen Types.Length `id:"90", required:"N" `
    SecureData Types.Data `id:"91", required:"N" `
    MsgSeqNum Types.SeqNum `id:"34", required:"Y" `
    SenderSubID Types.String `id:"50", required:"N" `
    SenderLocationID Types.String `id:"142", required:"N" `
    TargetSubID Types.String `id:"57", required:"N" `
    TargetLocationID Types.String `id:"143", required:"N" `
    OnBehalfOfSubID Types.String `id:"116", required:"N" `
    OnBehalfOfLocationID Types.String `id:"144", required:"N" `
    DeliverToSubID Types.String `id:"129", required:"N" `
    DeliverToLocationID Types.String `id:"145", required:"N" `
    PossDupFlag Types.Boolean `id:"43", required:"N" `
    PossResend Types.Boolean `id:"97", required:"N" `
    SendingTime Types.UTCTimestamp `id:"52", required:"Y" `
    OrigSendingTime Types.UTCTimestamp `id:"122", required:"N" `
    XmlDataLen Types.Length `id:"212", required:"N" `
    XmlData Types.Data `id:"213", required:"N" `
    MessageEncoding Types.String `id:"347", required:"N" `
    LastMsgSeqNumProcessed Types.SeqNum `id:"369", required:"N" `
    Hops []HopGrp_Hops.HopGrp_Hops `counter_name:"NoHops", counter_id:"627"`
    ApplVerID Types.String `id:"1128", required:"N" `
    CstmApplVerID Types.String `id:"1129", required:"N" `
    OrderID Types.String `id:"37", required:"N" `
    CrossID Types.String `id:"548", required:"Y" `
    OrigCrossID Types.String `id:"551", required:"Y" `
    CrossType Types.Int `id:"549", required:"Y" `
    CrossPrioritization Types.Int `id:"550", required:"Y" `
    Sides []SideCrossOrdModGrp_Sides.SideCrossOrdModGrp_Sides `counter_name:"NoSides", counter_id:"552"`
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
    Underlyings []UndInstrmtGrp_Underlyings.UndInstrmtGrp_Underlyings `counter_name:"NoUnderlyings", counter_id:"711"`
    Legs []InstrmtLegGrp_Legs.InstrmtLegGrp_Legs `counter_name:"NoLegs", counter_id:"555"`
    SettlType Types.String `id:"63", required:"N" `
    SettlDate Types.LocalMktDate `id:"64", required:"N" `
    HandlInst Types.Char `id:"21", required:"N" `
    ExecInst Types.MultipleCharValue `id:"18", required:"N" `
    MinQty Types.Qty `id:"110", required:"N" `
    MaxFloor Types.Qty `id:"111", required:"N" `
    ExDestination Types.Exchange `id:"100", required:"N" `
    TradingSessions []TrdgSesGrp_TradingSessions.TrdgSesGrp_TradingSessions `counter_name:"NoTradingSessions", counter_id:"386"`
    ProcessCode Types.Char `id:"81", required:"N" `
    PrevClosePx Types.Price `id:"140", required:"N" `
    LocateReqd Types.Boolean `id:"114", required:"N" `
    TransactTime Types.UTCTimestamp `id:"60", required:"Y" `
    Stipulations []Stipulations_Stipulations.Stipulations_Stipulations `counter_name:"NoStipulations", counter_id:"232"`
    OrdType Types.Char `id:"40", required:"Y" `
    PriceType Types.Int `id:"423", required:"N" `
    Price Types.Price `id:"44", required:"N" `
    StopPx Types.Price `id:"99", required:"N" `
    Spread Types.PriceOffset `id:"218", required:"N" `
    BenchmarkCurveCurrency Types.Currency `id:"220", required:"N" `
    BenchmarkCurveName Types.String `id:"221", required:"N" `
    BenchmarkCurvePoint Types.String `id:"222", required:"N" `
    BenchmarkPrice Types.Price `id:"662", required:"N" `
    BenchmarkPriceType Types.Int `id:"663", required:"N" `
    BenchmarkSecurityID Types.String `id:"699", required:"N" `
    BenchmarkSecurityIDSource Types.String `id:"761", required:"N" `
    YieldType Types.String `id:"235", required:"N" `
    Yield Types.Percentage `id:"236", required:"N" `
    YieldCalcDate Types.LocalMktDate `id:"701", required:"N" `
    YieldRedemptionDate Types.LocalMktDate `id:"696", required:"N" `
    YieldRedemptionPrice Types.Price `id:"697", required:"N" `
    YieldRedemptionPriceType Types.Int `id:"698", required:"N" `
    Currency Types.Currency `id:"15", required:"N" `
    ComplianceID Types.String `id:"376", required:"N" `
    IOIID Types.String `id:"23", required:"N" `
    QuoteID Types.String `id:"117", required:"N" `
    TimeInForce Types.Char `id:"59", required:"N" `
    EffectiveTime Types.UTCTimestamp `id:"168", required:"N" `
    ExpireDate Types.LocalMktDate `id:"432", required:"N" `
    ExpireTime Types.UTCTimestamp `id:"126", required:"N" `
    GTBookingInst Types.Int `id:"427", required:"N" `
    MaxShow Types.Qty `id:"210", required:"N" `
    PegOffsetValue Types.Float `id:"211", required:"N" `
    PegMoveType Types.Int `id:"835", required:"N" `
    PegOffsetType Types.Int `id:"836", required:"N" `
    PegLimitType Types.Int `id:"837", required:"N" `
    PegRoundDirection Types.Int `id:"838", required:"N" `
    PegScope Types.Int `id:"840", required:"N" `
    PegPriceType Types.Int `id:"1094", required:"N" `
    PegSecurityIDSource Types.String `id:"1096", required:"N" `
    PegSecurityID Types.String `id:"1097", required:"N" `
    PegSymbol Types.String `id:"1098", required:"N" `
    PegSecurityDesc Types.String `id:"1099", required:"N" `
    DiscretionInst Types.Char `id:"388", required:"N" `
    DiscretionOffsetValue Types.Float `id:"389", required:"N" `
    DiscretionMoveType Types.Int `id:"841", required:"N" `
    DiscretionOffsetType Types.Int `id:"842", required:"N" `
    DiscretionLimitType Types.Int `id:"843", required:"N" `
    DiscretionRoundDirection Types.Int `id:"844", required:"N" `
    DiscretionScope Types.Int `id:"846", required:"N" `
    TargetStrategy Types.Int `id:"847", required:"N" `
    TargetStrategyParameters Types.String `id:"848", required:"N" `
    ParticipationRate Types.Percentage `id:"849", required:"N" `
    CancellationRights Types.Char `id:"480", required:"N" `
    MoneyLaunderingStatus Types.Char `id:"481", required:"N" `
    RegistID Types.String `id:"513", required:"N" `
    Designation Types.String `id:"494", required:"N" `
    StrategyParameters []StrategyParametersGrp_StrategyParameters.StrategyParametersGrp_StrategyParameters `counter_name:"NoStrategyParameters", counter_id:"957"`
    HostCrossID Types.String `id:"961", required:"N" `
    TransBkdTime Types.UTCTimestamp `id:"483", required:"N" `
    RootPartyIDs []RootParties_RootPartyIDs.RootParties_RootPartyIDs `counter_name:"NoRootPartyIDs", counter_id:"1116"`
    MatchIncrement Types.Qty `id:"1089", required:"N" `
    MaxPriceLevels Types.Int `id:"1090", required:"N" `
    SecondaryDisplayQty Types.Qty `id:"1082", required:"N" `
    DisplayWhen Types.Char `id:"1083", required:"N" `
    DisplayMethod Types.Char `id:"1084", required:"N" `
    DisplayLowQty Types.Qty `id:"1085", required:"N" `
    DisplayHighQty Types.Qty `id:"1086", required:"N" `
    DisplayMinIncr Types.Qty `id:"1087", required:"N" `
    RefreshQty Types.Qty `id:"1088", required:"N" `
    DisplayQty Types.Qty `id:"1138", required:"N" `
    PriceProtectionScope Types.Char `id:"1092", required:"N" `
    TriggerType Types.Char `id:"1100", required:"N" `
    TriggerAction Types.Char `id:"1101", required:"N" `
    TriggerPrice Types.Price `id:"1102", required:"N" `
    TriggerSymbol Types.String `id:"1103", required:"N" `
    TriggerSecurityID Types.String `id:"1104", required:"N" `
    TriggerSecurityIDSource Types.String `id:"1105", required:"N" `
    TriggerSecurityDesc Types.String `id:"1106", required:"N" `
    TriggerPriceType Types.Char `id:"1107", required:"N" `
    TriggerPriceTypeScope Types.Char `id:"1108", required:"N" `
    TriggerPriceDirection Types.Char `id:"1109", required:"N" `
    TriggerNewPrice Types.Price `id:"1110", required:"N" `
    TriggerOrderType Types.Char `id:"1111", required:"N" `
    TriggerNewQty Types.Qty `id:"1112", required:"N" `
    TriggerTradingSessionID Types.String `id:"1113", required:"N" `
    TriggerTradingSessionSubID Types.String `id:"1114", required:"N" `
    ExDestinationIDSource Types.Char `id:"1133", required:"N" `
    SignatureLength Types.Length `id:"93", required:"N" `
    Signature Types.Data `id:"89", required:"N" `
    CheckSum Types.String `id:"10", required:"Y" `
    _controlBlock fixMessageControlBlock
}
