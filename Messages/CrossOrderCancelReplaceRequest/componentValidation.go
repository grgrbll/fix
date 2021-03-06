package CrossOrderCancelReplaceRequest

import "bufio"
import "regexp"

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


type crossOrderCancelReplaceRequest_RegexValidator struct {
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
    CrossID *(regexp.Regexp)
    OrigCrossID *(regexp.Regexp)
    CrossType *(regexp.Regexp)
    CrossPrioritization *(regexp.Regexp)
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
    SettlType *(regexp.Regexp)
    SettlDate *(regexp.Regexp)
    HandlInst *(regexp.Regexp)
    ExecInst *(regexp.Regexp)
    MinQty *(regexp.Regexp)
    MaxFloor *(regexp.Regexp)
    ExDestination *(regexp.Regexp)
    ProcessCode *(regexp.Regexp)
    PrevClosePx *(regexp.Regexp)
    LocateReqd *(regexp.Regexp)
    TransactTime *(regexp.Regexp)
    OrdType *(regexp.Regexp)
    PriceType *(regexp.Regexp)
    Price *(regexp.Regexp)
    StopPx *(regexp.Regexp)
    Spread *(regexp.Regexp)
    BenchmarkCurveCurrency *(regexp.Regexp)
    BenchmarkCurveName *(regexp.Regexp)
    BenchmarkCurvePoint *(regexp.Regexp)
    BenchmarkPrice *(regexp.Regexp)
    BenchmarkPriceType *(regexp.Regexp)
    BenchmarkSecurityID *(regexp.Regexp)
    BenchmarkSecurityIDSource *(regexp.Regexp)
    YieldType *(regexp.Regexp)
    Yield *(regexp.Regexp)
    YieldCalcDate *(regexp.Regexp)
    YieldRedemptionDate *(regexp.Regexp)
    YieldRedemptionPrice *(regexp.Regexp)
    YieldRedemptionPriceType *(regexp.Regexp)
    Currency *(regexp.Regexp)
    ComplianceID *(regexp.Regexp)
    IOIID *(regexp.Regexp)
    QuoteID *(regexp.Regexp)
    TimeInForce *(regexp.Regexp)
    EffectiveTime *(regexp.Regexp)
    ExpireDate *(regexp.Regexp)
    ExpireTime *(regexp.Regexp)
    GTBookingInst *(regexp.Regexp)
    MaxShow *(regexp.Regexp)
    PegOffsetValue *(regexp.Regexp)
    PegMoveType *(regexp.Regexp)
    PegOffsetType *(regexp.Regexp)
    PegLimitType *(regexp.Regexp)
    PegRoundDirection *(regexp.Regexp)
    PegScope *(regexp.Regexp)
    PegPriceType *(regexp.Regexp)
    PegSecurityIDSource *(regexp.Regexp)
    PegSecurityID *(regexp.Regexp)
    PegSymbol *(regexp.Regexp)
    PegSecurityDesc *(regexp.Regexp)
    DiscretionInst *(regexp.Regexp)
    DiscretionOffsetValue *(regexp.Regexp)
    DiscretionMoveType *(regexp.Regexp)
    DiscretionOffsetType *(regexp.Regexp)
    DiscretionLimitType *(regexp.Regexp)
    DiscretionRoundDirection *(regexp.Regexp)
    DiscretionScope *(regexp.Regexp)
    TargetStrategy *(regexp.Regexp)
    TargetStrategyParameters *(regexp.Regexp)
    ParticipationRate *(regexp.Regexp)
    CancellationRights *(regexp.Regexp)
    MoneyLaunderingStatus *(regexp.Regexp)
    RegistID *(regexp.Regexp)
    Designation *(regexp.Regexp)
    HostCrossID *(regexp.Regexp)
    TransBkdTime *(regexp.Regexp)
    MatchIncrement *(regexp.Regexp)
    MaxPriceLevels *(regexp.Regexp)
    SecondaryDisplayQty *(regexp.Regexp)
    DisplayWhen *(regexp.Regexp)
    DisplayMethod *(regexp.Regexp)
    DisplayLowQty *(regexp.Regexp)
    DisplayHighQty *(regexp.Regexp)
    DisplayMinIncr *(regexp.Regexp)
    RefreshQty *(regexp.Regexp)
    DisplayQty *(regexp.Regexp)
    PriceProtectionScope *(regexp.Regexp)
    TriggerType *(regexp.Regexp)
    TriggerAction *(regexp.Regexp)
    TriggerPrice *(regexp.Regexp)
    TriggerSymbol *(regexp.Regexp)
    TriggerSecurityID *(regexp.Regexp)
    TriggerSecurityIDSource *(regexp.Regexp)
    TriggerSecurityDesc *(regexp.Regexp)
    TriggerPriceType *(regexp.Regexp)
    TriggerPriceTypeScope *(regexp.Regexp)
    TriggerPriceDirection *(regexp.Regexp)
    TriggerNewPrice *(regexp.Regexp)
    TriggerOrderType *(regexp.Regexp)
    TriggerNewQty *(regexp.Regexp)
    TriggerTradingSessionID *(regexp.Regexp)
    TriggerTradingSessionSubID *(regexp.Regexp)
    ExDestinationIDSource *(regexp.Regexp)
    SignatureLength *(regexp.Regexp)
    Signature *(regexp.Regexp)
    CheckSum *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myCrossOrderCancelReplaceRequest_RegexValidator crossOrderCancelReplaceRequest_RegexValidator

func init() {
    myCrossOrderCancelReplaceRequest_RegexValidator
    myCrossOrderCancelReplaceRequest_RegexValidator.BeginString = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.BodyLength = regexp.MustCompile(`\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.MsgType = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.SenderCompID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.TargetCompID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.OnBehalfOfCompID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.DeliverToCompID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.SecureDataLen = regexp.MustCompile(`\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.SecureData = regexp.MustCompile(`.*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.MsgSeqNum = regexp.MustCompile(`\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.SenderSubID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.SenderLocationID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.TargetSubID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.TargetLocationID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.OnBehalfOfSubID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.OnBehalfOfLocationID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.DeliverToSubID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.DeliverToLocationID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.PossDupFlag = regexp.MustCompile(`[YN]`)
    myCrossOrderCancelReplaceRequest_RegexValidator.PossResend = regexp.MustCompile(`[YN]`)
    myCrossOrderCancelReplaceRequest_RegexValidator.SendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.OrigSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.XmlDataLen = regexp.MustCompile(`\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.XmlData = regexp.MustCompile(`.*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.MessageEncoding = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.LastMsgSeqNumProcessed = regexp.MustCompile(`\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.ApplVerID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.CstmApplVerID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.OrderID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.CrossID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.OrigCrossID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.CrossType = regexp.MustCompile(`-?\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.CrossPrioritization = regexp.MustCompile(`-?\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.Symbol = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.SymbolSfx = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.SecurityID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.SecurityIDSource = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.Product = regexp.MustCompile(`-?\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.CFICode = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.SecurityType = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.SecuritySubType = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.MaturityMonthYear = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.MaturityDate = regexp.MustCompile(`[0-9]{8}`)
    myCrossOrderCancelReplaceRequest_RegexValidator.CouponPaymentDate = regexp.MustCompile(`[0-9]{8}`)
    myCrossOrderCancelReplaceRequest_RegexValidator.IssueDate = regexp.MustCompile(`[0-9]{8}`)
    myCrossOrderCancelReplaceRequest_RegexValidator.RepoCollateralSecurityType = regexp.MustCompile(`-?\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.RepurchaseTerm = regexp.MustCompile(`-?\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.RepurchaseRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.Factor = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.CreditRating = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.InstrRegistry = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.CountryOfIssue = regexp.MustCompile(`[A-Z]{2}`)
    myCrossOrderCancelReplaceRequest_RegexValidator.StateOrProvinceOfIssue = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.LocaleOfIssue = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.RedemptionDate = regexp.MustCompile(`[0-9]{8}`)
    myCrossOrderCancelReplaceRequest_RegexValidator.StrikePrice = regexp.MustCompile(``)
    myCrossOrderCancelReplaceRequest_RegexValidator.StrikeCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myCrossOrderCancelReplaceRequest_RegexValidator.OptAttribute = regexp.MustCompile(`[^|]`)
    myCrossOrderCancelReplaceRequest_RegexValidator.ContractMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.CouponRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.SecurityExchange = regexp.MustCompile(`[A-Z0-9]{4}`)
    myCrossOrderCancelReplaceRequest_RegexValidator.Issuer = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.EncodedIssuerLen = regexp.MustCompile(`\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.EncodedIssuer = regexp.MustCompile(`.*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.SecurityDesc = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.EncodedSecurityDescLen = regexp.MustCompile(`\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.EncodedSecurityDesc = regexp.MustCompile(`.*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.Pool = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.ContractSettlMonth = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.CPProgram = regexp.MustCompile(`-?\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.CPRegType = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.DatedDate = regexp.MustCompile(`[0-9]{8}`)
    myCrossOrderCancelReplaceRequest_RegexValidator.InterestAccrualDate = regexp.MustCompile(`[0-9]{8}`)
    myCrossOrderCancelReplaceRequest_RegexValidator.SecurityStatus = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.SettleOnOpenFlag = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.InstrmtAssignmentMethod = regexp.MustCompile(`[^|]`)
    myCrossOrderCancelReplaceRequest_RegexValidator.StrikeMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.StrikeValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.MinPriceIncrement = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.PositionLimit = regexp.MustCompile(`-?\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.NTPositionLimit = regexp.MustCompile(`-?\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.UnitofMeasure = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.TimeUnit = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.MaturityTime = regexp.MustCompile(`\d{2}:\d{2}(:\d{2})?(Z|[\+-]\d{2}(:\d{2})?)`)
    myCrossOrderCancelReplaceRequest_RegexValidator.SettlType = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.SettlDate = regexp.MustCompile(`[0-9]{8}`)
    myCrossOrderCancelReplaceRequest_RegexValidator.HandlInst = regexp.MustCompile(`[^|]`)
    myCrossOrderCancelReplaceRequest_RegexValidator.ExecInst = regexp.MustCompile(`[ ]?(\w[ ])+\w?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.MinQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.MaxFloor = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.ExDestination = regexp.MustCompile(`[A-Z0-9]{4}`)
    myCrossOrderCancelReplaceRequest_RegexValidator.ProcessCode = regexp.MustCompile(`[^|]`)
    myCrossOrderCancelReplaceRequest_RegexValidator.PrevClosePx = regexp.MustCompile(``)
    myCrossOrderCancelReplaceRequest_RegexValidator.LocateReqd = regexp.MustCompile(`[YN]`)
    myCrossOrderCancelReplaceRequest_RegexValidator.TransactTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.OrdType = regexp.MustCompile(`[^|]`)
    myCrossOrderCancelReplaceRequest_RegexValidator.PriceType = regexp.MustCompile(`-?\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.Price = regexp.MustCompile(``)
    myCrossOrderCancelReplaceRequest_RegexValidator.StopPx = regexp.MustCompile(``)
    myCrossOrderCancelReplaceRequest_RegexValidator.Spread = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.BenchmarkCurveCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myCrossOrderCancelReplaceRequest_RegexValidator.BenchmarkCurveName = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.BenchmarkCurvePoint = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.BenchmarkPrice = regexp.MustCompile(``)
    myCrossOrderCancelReplaceRequest_RegexValidator.BenchmarkPriceType = regexp.MustCompile(`-?\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.BenchmarkSecurityID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.BenchmarkSecurityIDSource = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.YieldType = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.Yield = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.YieldCalcDate = regexp.MustCompile(`[0-9]{8}`)
    myCrossOrderCancelReplaceRequest_RegexValidator.YieldRedemptionDate = regexp.MustCompile(`[0-9]{8}`)
    myCrossOrderCancelReplaceRequest_RegexValidator.YieldRedemptionPrice = regexp.MustCompile(``)
    myCrossOrderCancelReplaceRequest_RegexValidator.YieldRedemptionPriceType = regexp.MustCompile(`-?\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.Currency = regexp.MustCompile(`[A-Z]{3}`)
    myCrossOrderCancelReplaceRequest_RegexValidator.ComplianceID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.IOIID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.QuoteID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.TimeInForce = regexp.MustCompile(`[^|]`)
    myCrossOrderCancelReplaceRequest_RegexValidator.EffectiveTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.ExpireDate = regexp.MustCompile(`[0-9]{8}`)
    myCrossOrderCancelReplaceRequest_RegexValidator.ExpireTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.GTBookingInst = regexp.MustCompile(`-?\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.MaxShow = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.PegOffsetValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.PegMoveType = regexp.MustCompile(`-?\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.PegOffsetType = regexp.MustCompile(`-?\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.PegLimitType = regexp.MustCompile(`-?\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.PegRoundDirection = regexp.MustCompile(`-?\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.PegScope = regexp.MustCompile(`-?\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.PegPriceType = regexp.MustCompile(`-?\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.PegSecurityIDSource = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.PegSecurityID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.PegSymbol = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.PegSecurityDesc = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.DiscretionInst = regexp.MustCompile(`[^|]`)
    myCrossOrderCancelReplaceRequest_RegexValidator.DiscretionOffsetValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.DiscretionMoveType = regexp.MustCompile(`-?\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.DiscretionOffsetType = regexp.MustCompile(`-?\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.DiscretionLimitType = regexp.MustCompile(`-?\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.DiscretionRoundDirection = regexp.MustCompile(`-?\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.DiscretionScope = regexp.MustCompile(`-?\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.TargetStrategy = regexp.MustCompile(`-?\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.TargetStrategyParameters = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.ParticipationRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.CancellationRights = regexp.MustCompile(`[^|]`)
    myCrossOrderCancelReplaceRequest_RegexValidator.MoneyLaunderingStatus = regexp.MustCompile(`[^|]`)
    myCrossOrderCancelReplaceRequest_RegexValidator.RegistID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.Designation = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.HostCrossID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.TransBkdTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.MatchIncrement = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.MaxPriceLevels = regexp.MustCompile(`-?\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.SecondaryDisplayQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.DisplayWhen = regexp.MustCompile(`[^|]`)
    myCrossOrderCancelReplaceRequest_RegexValidator.DisplayMethod = regexp.MustCompile(`[^|]`)
    myCrossOrderCancelReplaceRequest_RegexValidator.DisplayLowQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.DisplayHighQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.DisplayMinIncr = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.RefreshQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.DisplayQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.PriceProtectionScope = regexp.MustCompile(`[^|]`)
    myCrossOrderCancelReplaceRequest_RegexValidator.TriggerType = regexp.MustCompile(`[^|]`)
    myCrossOrderCancelReplaceRequest_RegexValidator.TriggerAction = regexp.MustCompile(`[^|]`)
    myCrossOrderCancelReplaceRequest_RegexValidator.TriggerPrice = regexp.MustCompile(``)
    myCrossOrderCancelReplaceRequest_RegexValidator.TriggerSymbol = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.TriggerSecurityID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.TriggerSecurityIDSource = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.TriggerSecurityDesc = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.TriggerPriceType = regexp.MustCompile(`[^|]`)
    myCrossOrderCancelReplaceRequest_RegexValidator.TriggerPriceTypeScope = regexp.MustCompile(`[^|]`)
    myCrossOrderCancelReplaceRequest_RegexValidator.TriggerPriceDirection = regexp.MustCompile(`[^|]`)
    myCrossOrderCancelReplaceRequest_RegexValidator.TriggerNewPrice = regexp.MustCompile(``)
    myCrossOrderCancelReplaceRequest_RegexValidator.TriggerOrderType = regexp.MustCompile(`[^|]`)
    myCrossOrderCancelReplaceRequest_RegexValidator.TriggerNewQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myCrossOrderCancelReplaceRequest_RegexValidator.TriggerTradingSessionID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.TriggerTradingSessionSubID = regexp.MustCompile(`[^|]*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.ExDestinationIDSource = regexp.MustCompile(`[^|]`)
    myCrossOrderCancelReplaceRequest_RegexValidator.SignatureLength = regexp.MustCompile(`\d+`)
    myCrossOrderCancelReplaceRequest_RegexValidator.Signature = regexp.MustCompile(`.*`)
    myCrossOrderCancelReplaceRequest_RegexValidator.CheckSum = regexp.MustCompile(`[^|]*`)
}



func (m *CrossOrderCancelReplaceRequest) HasRequiredFields() bool {
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
    
    
    
    valid = valid && m.CrossID.HasValue()
    
    
    
    valid = valid && m.OrigCrossID.HasValue()
    
    
    
    valid = valid && m.CrossType.HasValue()
    
    
    
    valid = valid && m.CrossPrioritization.HasValue()
    
    
    
    for _, g := range(m.Sides) {
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
    
    
    
    for _, g := range(m.TradingSessions) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.TransactTime.HasValue()
    
    
    
    for _, g := range(m.Stipulations) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.OrdType.HasValue()
    
    
    
    for _, g := range(m.StrategyParameters) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.RootPartyIDs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.CheckSum.HasValue()
    
    
    
    
    
    return valid
}




