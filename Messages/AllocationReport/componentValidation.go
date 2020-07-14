package AllocationReport

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import OrdAllocGrp_Orders "grgrbll/fix/Shared/OrdAllocGrp_Orders"

import ExecAllocGrp_Execs "grgrbll/fix/Shared/ExecAllocGrp_Execs"

import SecAltIDGrp_SecurityAltID "grgrbll/fix/Shared/SecAltIDGrp_SecurityAltID"

import EvntGrp_Events "grgrbll/fix/Shared/EvntGrp_Events"

import InstrumentParties_InstrumentParties "grgrbll/fix/Shared/InstrumentParties_InstrumentParties"

import AttrbGrp_InstrAttrib "grgrbll/fix/Shared/AttrbGrp_InstrAttrib"

import UndInstrmtGrp_Underlyings "grgrbll/fix/Shared/UndInstrmtGrp_Underlyings"

import InstrmtLegGrp_Legs "grgrbll/fix/Shared/InstrmtLegGrp_Legs"

import Parties_PartyIDs "grgrbll/fix/Shared/Parties_PartyIDs"

import Stipulations_Stipulations "grgrbll/fix/Shared/Stipulations_Stipulations"

import AllocGrp_Allocs "grgrbll/fix/Shared/AllocGrp_Allocs"

import PositionAmountData_PosAmt "grgrbll/fix/Shared/PositionAmountData_PosAmt"


type allocationReport_RegexValidator struct {
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
    AllocReportID *(regexp.Regexp)
    AllocID *(regexp.Regexp)
    AllocTransType *(regexp.Regexp)
    AllocReportRefID *(regexp.Regexp)
    AllocCancReplaceReason *(regexp.Regexp)
    SecondaryAllocID *(regexp.Regexp)
    AllocReportType *(regexp.Regexp)
    AllocStatus *(regexp.Regexp)
    AllocRejCode *(regexp.Regexp)
    RefAllocID *(regexp.Regexp)
    AllocIntermedReqType *(regexp.Regexp)
    AllocLinkID *(regexp.Regexp)
    AllocLinkType *(regexp.Regexp)
    BookingRefID *(regexp.Regexp)
    AllocNoOrdersType *(regexp.Regexp)
    PreviouslyReported *(regexp.Regexp)
    ReversalIndicator *(regexp.Regexp)
    MatchType *(regexp.Regexp)
    Side *(regexp.Regexp)
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
    DeliveryForm *(regexp.Regexp)
    PctAtRisk *(regexp.Regexp)
    AgreementDesc *(regexp.Regexp)
    AgreementID *(regexp.Regexp)
    AgreementDate *(regexp.Regexp)
    AgreementCurrency *(regexp.Regexp)
    TerminationType *(regexp.Regexp)
    StartDate *(regexp.Regexp)
    EndDate *(regexp.Regexp)
    DeliveryType *(regexp.Regexp)
    MarginRatio *(regexp.Regexp)
    Quantity *(regexp.Regexp)
    QtyType *(regexp.Regexp)
    LastMkt *(regexp.Regexp)
    TradeOriginationDate *(regexp.Regexp)
    TradingSessionID *(regexp.Regexp)
    TradingSessionSubID *(regexp.Regexp)
    PriceType *(regexp.Regexp)
    AvgPx *(regexp.Regexp)
    AvgParPx *(regexp.Regexp)
    Spread *(regexp.Regexp)
    BenchmarkCurveCurrency *(regexp.Regexp)
    BenchmarkCurveName *(regexp.Regexp)
    BenchmarkCurvePoint *(regexp.Regexp)
    BenchmarkPrice *(regexp.Regexp)
    BenchmarkPriceType *(regexp.Regexp)
    BenchmarkSecurityID *(regexp.Regexp)
    BenchmarkSecurityIDSource *(regexp.Regexp)
    Currency *(regexp.Regexp)
    AvgPxPrecision *(regexp.Regexp)
    TradeDate *(regexp.Regexp)
    TransactTime *(regexp.Regexp)
    SettlType *(regexp.Regexp)
    SettlDate *(regexp.Regexp)
    BookingType *(regexp.Regexp)
    GrossTradeAmt *(regexp.Regexp)
    Concession *(regexp.Regexp)
    TotalTakedown *(regexp.Regexp)
    NetMoney *(regexp.Regexp)
    PositionEffect *(regexp.Regexp)
    AutoAcceptIndicator *(regexp.Regexp)
    Text *(regexp.Regexp)
    EncodedTextLen *(regexp.Regexp)
    EncodedText *(regexp.Regexp)
    NumDaysInterest *(regexp.Regexp)
    AccruedInterestRate *(regexp.Regexp)
    AccruedInterestAmt *(regexp.Regexp)
    TotalAccruedInterestAmt *(regexp.Regexp)
    InterestAtMaturity *(regexp.Regexp)
    EndAccruedInterestAmt *(regexp.Regexp)
    StartCash *(regexp.Regexp)
    EndCash *(regexp.Regexp)
    LegalConfirm *(regexp.Regexp)
    YieldType *(regexp.Regexp)
    Yield *(regexp.Regexp)
    YieldCalcDate *(regexp.Regexp)
    YieldRedemptionDate *(regexp.Regexp)
    YieldRedemptionPrice *(regexp.Regexp)
    YieldRedemptionPriceType *(regexp.Regexp)
    TotNoAllocs *(regexp.Regexp)
    LastFragment *(regexp.Regexp)
    ClearingBusinessDate *(regexp.Regexp)
    TrdType *(regexp.Regexp)
    TrdSubType *(regexp.Regexp)
    MultiLegReportingType *(regexp.Regexp)
    CustOrderCapacity *(regexp.Regexp)
    TradeInputSource *(regexp.Regexp)
    RndPx *(regexp.Regexp)
    MessageEventSource *(regexp.Regexp)
    TradeInputDevice *(regexp.Regexp)
    AvgPxIndicator *(regexp.Regexp)
    SignatureLength *(regexp.Regexp)
    Signature *(regexp.Regexp)
    CheckSum *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myAllocationReport_RegexValidator allocationReport_RegexValidator

func init() {
    myAllocationReport_RegexValidator
    myAllocationReport_RegexValidator.BeginString = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.BodyLength = regexp.MustCompile(`\d+`)
    myAllocationReport_RegexValidator.MsgType = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.SenderCompID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.TargetCompID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.OnBehalfOfCompID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.DeliverToCompID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.SecureDataLen = regexp.MustCompile(`\d+`)
    myAllocationReport_RegexValidator.SecureData = regexp.MustCompile(`.*`)
    myAllocationReport_RegexValidator.MsgSeqNum = regexp.MustCompile(`\d+`)
    myAllocationReport_RegexValidator.SenderSubID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.SenderLocationID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.TargetSubID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.TargetLocationID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.OnBehalfOfSubID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.OnBehalfOfLocationID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.DeliverToSubID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.DeliverToLocationID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.PossDupFlag = regexp.MustCompile(`[YN]`)
    myAllocationReport_RegexValidator.PossResend = regexp.MustCompile(`[YN]`)
    myAllocationReport_RegexValidator.SendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myAllocationReport_RegexValidator.OrigSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myAllocationReport_RegexValidator.XmlDataLen = regexp.MustCompile(`\d+`)
    myAllocationReport_RegexValidator.XmlData = regexp.MustCompile(`.*`)
    myAllocationReport_RegexValidator.MessageEncoding = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.LastMsgSeqNumProcessed = regexp.MustCompile(`\d+`)
    myAllocationReport_RegexValidator.ApplVerID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.CstmApplVerID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.AllocReportID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.AllocID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.AllocTransType = regexp.MustCompile(`[^|]`)
    myAllocationReport_RegexValidator.AllocReportRefID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.AllocCancReplaceReason = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.SecondaryAllocID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.AllocReportType = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.AllocStatus = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.AllocRejCode = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.RefAllocID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.AllocIntermedReqType = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.AllocLinkID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.AllocLinkType = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.BookingRefID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.AllocNoOrdersType = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.PreviouslyReported = regexp.MustCompile(`[YN]`)
    myAllocationReport_RegexValidator.ReversalIndicator = regexp.MustCompile(`[YN]`)
    myAllocationReport_RegexValidator.MatchType = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.Side = regexp.MustCompile(`[^|]`)
    myAllocationReport_RegexValidator.Symbol = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.SymbolSfx = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.SecurityID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.SecurityIDSource = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.Product = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.CFICode = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.SecurityType = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.SecuritySubType = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.MaturityMonthYear = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myAllocationReport_RegexValidator.MaturityDate = regexp.MustCompile(`[0-9]{8}`)
    myAllocationReport_RegexValidator.CouponPaymentDate = regexp.MustCompile(`[0-9]{8}`)
    myAllocationReport_RegexValidator.IssueDate = regexp.MustCompile(`[0-9]{8}`)
    myAllocationReport_RegexValidator.RepoCollateralSecurityType = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.RepurchaseTerm = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.RepurchaseRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAllocationReport_RegexValidator.Factor = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAllocationReport_RegexValidator.CreditRating = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.InstrRegistry = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.CountryOfIssue = regexp.MustCompile(`[A-Z]{2}`)
    myAllocationReport_RegexValidator.StateOrProvinceOfIssue = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.LocaleOfIssue = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.RedemptionDate = regexp.MustCompile(`[0-9]{8}`)
    myAllocationReport_RegexValidator.StrikePrice = regexp.MustCompile(``)
    myAllocationReport_RegexValidator.StrikeCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myAllocationReport_RegexValidator.OptAttribute = regexp.MustCompile(`[^|]`)
    myAllocationReport_RegexValidator.ContractMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAllocationReport_RegexValidator.CouponRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAllocationReport_RegexValidator.SecurityExchange = regexp.MustCompile(`[A-Z0-9]{4}`)
    myAllocationReport_RegexValidator.Issuer = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.EncodedIssuerLen = regexp.MustCompile(`\d+`)
    myAllocationReport_RegexValidator.EncodedIssuer = regexp.MustCompile(`.*`)
    myAllocationReport_RegexValidator.SecurityDesc = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.EncodedSecurityDescLen = regexp.MustCompile(`\d+`)
    myAllocationReport_RegexValidator.EncodedSecurityDesc = regexp.MustCompile(`.*`)
    myAllocationReport_RegexValidator.Pool = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.ContractSettlMonth = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myAllocationReport_RegexValidator.CPProgram = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.CPRegType = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.DatedDate = regexp.MustCompile(`[0-9]{8}`)
    myAllocationReport_RegexValidator.InterestAccrualDate = regexp.MustCompile(`[0-9]{8}`)
    myAllocationReport_RegexValidator.SecurityStatus = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.SettleOnOpenFlag = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.InstrmtAssignmentMethod = regexp.MustCompile(`[^|]`)
    myAllocationReport_RegexValidator.StrikeMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAllocationReport_RegexValidator.StrikeValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAllocationReport_RegexValidator.MinPriceIncrement = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAllocationReport_RegexValidator.PositionLimit = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.NTPositionLimit = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.UnitofMeasure = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.TimeUnit = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.MaturityTime = regexp.MustCompile(`\d{2}:\d{2}(:\d{2})?(Z|[\+-]\d{2}(:\d{2})?)`)
    myAllocationReport_RegexValidator.DeliveryForm = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.PctAtRisk = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAllocationReport_RegexValidator.AgreementDesc = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.AgreementID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.AgreementDate = regexp.MustCompile(`[0-9]{8}`)
    myAllocationReport_RegexValidator.AgreementCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myAllocationReport_RegexValidator.TerminationType = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.StartDate = regexp.MustCompile(`[0-9]{8}`)
    myAllocationReport_RegexValidator.EndDate = regexp.MustCompile(`[0-9]{8}`)
    myAllocationReport_RegexValidator.DeliveryType = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.MarginRatio = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAllocationReport_RegexValidator.Quantity = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAllocationReport_RegexValidator.QtyType = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.LastMkt = regexp.MustCompile(`[A-Z0-9]{4}`)
    myAllocationReport_RegexValidator.TradeOriginationDate = regexp.MustCompile(`[0-9]{8}`)
    myAllocationReport_RegexValidator.TradingSessionID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.TradingSessionSubID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.PriceType = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.AvgPx = regexp.MustCompile(``)
    myAllocationReport_RegexValidator.AvgParPx = regexp.MustCompile(``)
    myAllocationReport_RegexValidator.Spread = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAllocationReport_RegexValidator.BenchmarkCurveCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myAllocationReport_RegexValidator.BenchmarkCurveName = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.BenchmarkCurvePoint = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.BenchmarkPrice = regexp.MustCompile(``)
    myAllocationReport_RegexValidator.BenchmarkPriceType = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.BenchmarkSecurityID = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.BenchmarkSecurityIDSource = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.Currency = regexp.MustCompile(`[A-Z]{3}`)
    myAllocationReport_RegexValidator.AvgPxPrecision = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.TradeDate = regexp.MustCompile(`[0-9]{8}`)
    myAllocationReport_RegexValidator.TransactTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myAllocationReport_RegexValidator.SettlType = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.SettlDate = regexp.MustCompile(`[0-9]{8}`)
    myAllocationReport_RegexValidator.BookingType = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.GrossTradeAmt = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAllocationReport_RegexValidator.Concession = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAllocationReport_RegexValidator.TotalTakedown = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAllocationReport_RegexValidator.NetMoney = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAllocationReport_RegexValidator.PositionEffect = regexp.MustCompile(`[^|]`)
    myAllocationReport_RegexValidator.AutoAcceptIndicator = regexp.MustCompile(`[YN]`)
    myAllocationReport_RegexValidator.Text = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.EncodedTextLen = regexp.MustCompile(`\d+`)
    myAllocationReport_RegexValidator.EncodedText = regexp.MustCompile(`.*`)
    myAllocationReport_RegexValidator.NumDaysInterest = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.AccruedInterestRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAllocationReport_RegexValidator.AccruedInterestAmt = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAllocationReport_RegexValidator.TotalAccruedInterestAmt = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAllocationReport_RegexValidator.InterestAtMaturity = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAllocationReport_RegexValidator.EndAccruedInterestAmt = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAllocationReport_RegexValidator.StartCash = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAllocationReport_RegexValidator.EndCash = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAllocationReport_RegexValidator.LegalConfirm = regexp.MustCompile(`[YN]`)
    myAllocationReport_RegexValidator.YieldType = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.Yield = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myAllocationReport_RegexValidator.YieldCalcDate = regexp.MustCompile(`[0-9]{8}`)
    myAllocationReport_RegexValidator.YieldRedemptionDate = regexp.MustCompile(`[0-9]{8}`)
    myAllocationReport_RegexValidator.YieldRedemptionPrice = regexp.MustCompile(``)
    myAllocationReport_RegexValidator.YieldRedemptionPriceType = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.TotNoAllocs = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.LastFragment = regexp.MustCompile(`[YN]`)
    myAllocationReport_RegexValidator.ClearingBusinessDate = regexp.MustCompile(`[0-9]{8}`)
    myAllocationReport_RegexValidator.TrdType = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.TrdSubType = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.MultiLegReportingType = regexp.MustCompile(`[^|]`)
    myAllocationReport_RegexValidator.CustOrderCapacity = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.TradeInputSource = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.RndPx = regexp.MustCompile(``)
    myAllocationReport_RegexValidator.MessageEventSource = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.TradeInputDevice = regexp.MustCompile(`[^|]*`)
    myAllocationReport_RegexValidator.AvgPxIndicator = regexp.MustCompile(`-?\d+`)
    myAllocationReport_RegexValidator.SignatureLength = regexp.MustCompile(`\d+`)
    myAllocationReport_RegexValidator.Signature = regexp.MustCompile(`.*`)
    myAllocationReport_RegexValidator.CheckSum = regexp.MustCompile(`[^|]*`)
}



func (m *AllocationReport) HasRequiredFields() bool {
    valid := true
    
    isSetExecAllocGrp := false
    isValidExecAllocGrp := true

    
    
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
    
    
    
    valid = valid && m.AllocReportID.HasValue()
    
    
    
    valid = valid && m.AllocTransType.HasValue()
    
    
    
    valid = valid && m.AllocReportType.HasValue()
    
    
    
    valid = valid && m.AllocStatus.HasValue()
    
    
    
    for _, g := range(m.Orders) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    isSetExecAllocGrp = isSetExecAllocGrp || len(m.Execs) > 0
    for _, g := range(m.Execs) {
        isValidExecAllocGrp = isValidExecAllocGrp && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.Side.HasValue()
    
    
    
    for _, g := range(m.SecurityAltID) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Events) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.InstrumentParties) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.InstrAttrib) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Underlyings) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Legs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.Quantity.HasValue()
    
    
    
    valid = valid && m.AvgPx.HasValue()
    
    
    
    for _, g := range(m.PartyIDs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.TradeDate.HasValue()
    
    
    
    for _, g := range(m.Stipulations) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Allocs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.PosAmt) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.CheckSum.HasValue()
    
    
    
    valid = valid && (isSetExecAllocGrp == false || isValidExecAllocGrp)
    
    return valid
}



