package TradeCaptureReportAck

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import SecAltIDGrp_SecurityAltID "grgrbll/fix/Shared/SecAltIDGrp_SecurityAltID"

import EvntGrp_Events "grgrbll/fix/Shared/EvntGrp_Events"

import InstrumentParties_InstrumentParties "grgrbll/fix/Shared/InstrumentParties_InstrumentParties"

import TrdRegTimestamps_TrdRegTimestamps "grgrbll/fix/Shared/TrdRegTimestamps_TrdRegTimestamps"

import TrdInstrmtLegGrp_Legs "grgrbll/fix/Shared/TrdInstrmtLegGrp_Legs"

import PositionAmountData_PosAmt "grgrbll/fix/Shared/PositionAmountData_PosAmt"

import TrdCapRptAckSideGrp_Sides "grgrbll/fix/Shared/TrdCapRptAckSideGrp_Sides"

import RootParties_RootPartyIDs "grgrbll/fix/Shared/RootParties_RootPartyIDs"

import UndInstrmtGrp_Underlyings "grgrbll/fix/Shared/UndInstrmtGrp_Underlyings"


type tradeCaptureReportAck_RegexValidator struct {
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
    TradeReportID *(regexp.Regexp)
    TradeReportTransType *(regexp.Regexp)
    TradeReportType *(regexp.Regexp)
    TrdType *(regexp.Regexp)
    TrdSubType *(regexp.Regexp)
    SecondaryTrdType *(regexp.Regexp)
    TransferReason *(regexp.Regexp)
    ExecType *(regexp.Regexp)
    TradeReportRefID *(regexp.Regexp)
    SecondaryTradeReportRefID *(regexp.Regexp)
    TrdRptStatus *(regexp.Regexp)
    TradeReportRejectReason *(regexp.Regexp)
    SecondaryTradeReportID *(regexp.Regexp)
    SubscriptionRequestType *(regexp.Regexp)
    TradeLinkID *(regexp.Regexp)
    TrdMatchID *(regexp.Regexp)
    ExecID *(regexp.Regexp)
    SecondaryExecID *(regexp.Regexp)
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
    TransactTime *(regexp.Regexp)
    ResponseTransportType *(regexp.Regexp)
    ResponseDestination *(regexp.Regexp)
    Text *(regexp.Regexp)
    EncodedTextLen *(regexp.Regexp)
    EncodedText *(regexp.Regexp)
    ClearingFeeIndicator *(regexp.Regexp)
    OrdStatus *(regexp.Regexp)
    ExecRestatementReason *(regexp.Regexp)
    PreviouslyReported *(regexp.Regexp)
    PriceType *(regexp.Regexp)
    UnderlyingTradingSessionID *(regexp.Regexp)
    QtyType *(regexp.Regexp)
    UnderlyingTradingSessionSubID *(regexp.Regexp)
    LastQty *(regexp.Regexp)
    LastPx *(regexp.Regexp)
    LastParPx *(regexp.Regexp)
    LastSpotRate *(regexp.Regexp)
    LastForwardPoints *(regexp.Regexp)
    LastMkt *(regexp.Regexp)
    TradeDate *(regexp.Regexp)
    ClearingBusinessDate *(regexp.Regexp)
    AvgPx *(regexp.Regexp)
    AvgPxIndicator *(regexp.Regexp)
    MultiLegReportingType *(regexp.Regexp)
    TradeLegRefID *(regexp.Regexp)
    SettlType *(regexp.Regexp)
    MatchStatus *(regexp.Regexp)
    MatchType *(regexp.Regexp)
    CopyMsgIndicator *(regexp.Regexp)
    PublishTrdIndicator *(regexp.Regexp)
    ShortSaleReason *(regexp.Regexp)
    SettlDate *(regexp.Regexp)
    SettlSessID *(regexp.Regexp)
    SettlSessSubID *(regexp.Regexp)
    TierCode *(regexp.Regexp)
    MessageEventSource *(regexp.Regexp)
    LastUpdateTime *(regexp.Regexp)
    RndPx *(regexp.Regexp)
    AsOfIndicator *(regexp.Regexp)
    TradeID *(regexp.Regexp)
    SecondaryTradeID *(regexp.Regexp)
    FirmTradeID *(regexp.Regexp)
    SecondaryFirmTradeID *(regexp.Regexp)
    CalculatedCcyLastQty *(regexp.Regexp)
    LastSwapPoints *(regexp.Regexp)
    GrossTradeAmt *(regexp.Regexp)
    TradeHandlingInstr *(regexp.Regexp)
    OrigTradeHandlingInstr *(regexp.Regexp)
    OrigTradeDate *(regexp.Regexp)
    OrigTradeID *(regexp.Regexp)
    OrigSecondaryTradeID *(regexp.Regexp)
    RptSys *(regexp.Regexp)
    SignatureLength *(regexp.Regexp)
    Signature *(regexp.Regexp)
    CheckSum *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myTradeCaptureReportAck_RegexValidator tradeCaptureReportAck_RegexValidator

func init() {
    myTradeCaptureReportAck_RegexValidator
    myTradeCaptureReportAck_RegexValidator.BeginString = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.BodyLength = regexp.MustCompile(`\d+`)
    myTradeCaptureReportAck_RegexValidator.MsgType = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.SenderCompID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.TargetCompID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.OnBehalfOfCompID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.DeliverToCompID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.SecureDataLen = regexp.MustCompile(`\d+`)
    myTradeCaptureReportAck_RegexValidator.SecureData = regexp.MustCompile(`.*`)
    myTradeCaptureReportAck_RegexValidator.MsgSeqNum = regexp.MustCompile(`\d+`)
    myTradeCaptureReportAck_RegexValidator.SenderSubID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.SenderLocationID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.TargetSubID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.TargetLocationID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.OnBehalfOfSubID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.OnBehalfOfLocationID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.DeliverToSubID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.DeliverToLocationID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.PossDupFlag = regexp.MustCompile(`[YN]`)
    myTradeCaptureReportAck_RegexValidator.PossResend = regexp.MustCompile(`[YN]`)
    myTradeCaptureReportAck_RegexValidator.SendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myTradeCaptureReportAck_RegexValidator.OrigSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myTradeCaptureReportAck_RegexValidator.XmlDataLen = regexp.MustCompile(`\d+`)
    myTradeCaptureReportAck_RegexValidator.XmlData = regexp.MustCompile(`.*`)
    myTradeCaptureReportAck_RegexValidator.MessageEncoding = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.LastMsgSeqNumProcessed = regexp.MustCompile(`\d+`)
    myTradeCaptureReportAck_RegexValidator.ApplVerID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.CstmApplVerID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.TradeReportID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.TradeReportTransType = regexp.MustCompile(`-?\d+`)
    myTradeCaptureReportAck_RegexValidator.TradeReportType = regexp.MustCompile(`-?\d+`)
    myTradeCaptureReportAck_RegexValidator.TrdType = regexp.MustCompile(`-?\d+`)
    myTradeCaptureReportAck_RegexValidator.TrdSubType = regexp.MustCompile(`-?\d+`)
    myTradeCaptureReportAck_RegexValidator.SecondaryTrdType = regexp.MustCompile(`-?\d+`)
    myTradeCaptureReportAck_RegexValidator.TransferReason = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.ExecType = regexp.MustCompile(`[^|]`)
    myTradeCaptureReportAck_RegexValidator.TradeReportRefID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.SecondaryTradeReportRefID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.TrdRptStatus = regexp.MustCompile(`-?\d+`)
    myTradeCaptureReportAck_RegexValidator.TradeReportRejectReason = regexp.MustCompile(`-?\d+`)
    myTradeCaptureReportAck_RegexValidator.SecondaryTradeReportID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.SubscriptionRequestType = regexp.MustCompile(`[^|]`)
    myTradeCaptureReportAck_RegexValidator.TradeLinkID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.TrdMatchID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.ExecID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.SecondaryExecID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.Symbol = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.SymbolSfx = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.SecurityID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.SecurityIDSource = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.Product = regexp.MustCompile(`-?\d+`)
    myTradeCaptureReportAck_RegexValidator.CFICode = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.SecurityType = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.SecuritySubType = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.MaturityMonthYear = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myTradeCaptureReportAck_RegexValidator.MaturityDate = regexp.MustCompile(`[0-9]{8}`)
    myTradeCaptureReportAck_RegexValidator.CouponPaymentDate = regexp.MustCompile(`[0-9]{8}`)
    myTradeCaptureReportAck_RegexValidator.IssueDate = regexp.MustCompile(`[0-9]{8}`)
    myTradeCaptureReportAck_RegexValidator.RepoCollateralSecurityType = regexp.MustCompile(`-?\d+`)
    myTradeCaptureReportAck_RegexValidator.RepurchaseTerm = regexp.MustCompile(`-?\d+`)
    myTradeCaptureReportAck_RegexValidator.RepurchaseRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradeCaptureReportAck_RegexValidator.Factor = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradeCaptureReportAck_RegexValidator.CreditRating = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.InstrRegistry = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.CountryOfIssue = regexp.MustCompile(`[A-Z]{2}`)
    myTradeCaptureReportAck_RegexValidator.StateOrProvinceOfIssue = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.LocaleOfIssue = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.RedemptionDate = regexp.MustCompile(`[0-9]{8}`)
    myTradeCaptureReportAck_RegexValidator.StrikePrice = regexp.MustCompile(``)
    myTradeCaptureReportAck_RegexValidator.StrikeCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myTradeCaptureReportAck_RegexValidator.OptAttribute = regexp.MustCompile(`[^|]`)
    myTradeCaptureReportAck_RegexValidator.ContractMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradeCaptureReportAck_RegexValidator.CouponRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradeCaptureReportAck_RegexValidator.SecurityExchange = regexp.MustCompile(`[A-Z0-9]{4}`)
    myTradeCaptureReportAck_RegexValidator.Issuer = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.EncodedIssuerLen = regexp.MustCompile(`\d+`)
    myTradeCaptureReportAck_RegexValidator.EncodedIssuer = regexp.MustCompile(`.*`)
    myTradeCaptureReportAck_RegexValidator.SecurityDesc = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.EncodedSecurityDescLen = regexp.MustCompile(`\d+`)
    myTradeCaptureReportAck_RegexValidator.EncodedSecurityDesc = regexp.MustCompile(`.*`)
    myTradeCaptureReportAck_RegexValidator.Pool = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.ContractSettlMonth = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myTradeCaptureReportAck_RegexValidator.CPProgram = regexp.MustCompile(`-?\d+`)
    myTradeCaptureReportAck_RegexValidator.CPRegType = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.DatedDate = regexp.MustCompile(`[0-9]{8}`)
    myTradeCaptureReportAck_RegexValidator.InterestAccrualDate = regexp.MustCompile(`[0-9]{8}`)
    myTradeCaptureReportAck_RegexValidator.SecurityStatus = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.SettleOnOpenFlag = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.InstrmtAssignmentMethod = regexp.MustCompile(`[^|]`)
    myTradeCaptureReportAck_RegexValidator.StrikeMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradeCaptureReportAck_RegexValidator.StrikeValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradeCaptureReportAck_RegexValidator.MinPriceIncrement = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradeCaptureReportAck_RegexValidator.PositionLimit = regexp.MustCompile(`-?\d+`)
    myTradeCaptureReportAck_RegexValidator.NTPositionLimit = regexp.MustCompile(`-?\d+`)
    myTradeCaptureReportAck_RegexValidator.UnitofMeasure = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.TimeUnit = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.MaturityTime = regexp.MustCompile(`\d{2}:\d{2}(:\d{2})?(Z|[\+-]\d{2}(:\d{2})?)`)
    myTradeCaptureReportAck_RegexValidator.TransactTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myTradeCaptureReportAck_RegexValidator.ResponseTransportType = regexp.MustCompile(`-?\d+`)
    myTradeCaptureReportAck_RegexValidator.ResponseDestination = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.Text = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.EncodedTextLen = regexp.MustCompile(`\d+`)
    myTradeCaptureReportAck_RegexValidator.EncodedText = regexp.MustCompile(`.*`)
    myTradeCaptureReportAck_RegexValidator.ClearingFeeIndicator = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.OrdStatus = regexp.MustCompile(`[^|]`)
    myTradeCaptureReportAck_RegexValidator.ExecRestatementReason = regexp.MustCompile(`-?\d+`)
    myTradeCaptureReportAck_RegexValidator.PreviouslyReported = regexp.MustCompile(`[YN]`)
    myTradeCaptureReportAck_RegexValidator.PriceType = regexp.MustCompile(`-?\d+`)
    myTradeCaptureReportAck_RegexValidator.UnderlyingTradingSessionID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.QtyType = regexp.MustCompile(`-?\d+`)
    myTradeCaptureReportAck_RegexValidator.UnderlyingTradingSessionSubID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.LastQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradeCaptureReportAck_RegexValidator.LastPx = regexp.MustCompile(``)
    myTradeCaptureReportAck_RegexValidator.LastParPx = regexp.MustCompile(``)
    myTradeCaptureReportAck_RegexValidator.LastSpotRate = regexp.MustCompile(``)
    myTradeCaptureReportAck_RegexValidator.LastForwardPoints = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradeCaptureReportAck_RegexValidator.LastMkt = regexp.MustCompile(`[A-Z0-9]{4}`)
    myTradeCaptureReportAck_RegexValidator.TradeDate = regexp.MustCompile(`[0-9]{8}`)
    myTradeCaptureReportAck_RegexValidator.ClearingBusinessDate = regexp.MustCompile(`[0-9]{8}`)
    myTradeCaptureReportAck_RegexValidator.AvgPx = regexp.MustCompile(``)
    myTradeCaptureReportAck_RegexValidator.AvgPxIndicator = regexp.MustCompile(`-?\d+`)
    myTradeCaptureReportAck_RegexValidator.MultiLegReportingType = regexp.MustCompile(`[^|]`)
    myTradeCaptureReportAck_RegexValidator.TradeLegRefID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.SettlType = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.MatchStatus = regexp.MustCompile(`[^|]`)
    myTradeCaptureReportAck_RegexValidator.MatchType = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.CopyMsgIndicator = regexp.MustCompile(`[YN]`)
    myTradeCaptureReportAck_RegexValidator.PublishTrdIndicator = regexp.MustCompile(`[YN]`)
    myTradeCaptureReportAck_RegexValidator.ShortSaleReason = regexp.MustCompile(`-?\d+`)
    myTradeCaptureReportAck_RegexValidator.SettlDate = regexp.MustCompile(`[0-9]{8}`)
    myTradeCaptureReportAck_RegexValidator.SettlSessID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.SettlSessSubID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.TierCode = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.MessageEventSource = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.LastUpdateTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myTradeCaptureReportAck_RegexValidator.RndPx = regexp.MustCompile(``)
    myTradeCaptureReportAck_RegexValidator.AsOfIndicator = regexp.MustCompile(`[^|]`)
    myTradeCaptureReportAck_RegexValidator.TradeID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.SecondaryTradeID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.FirmTradeID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.SecondaryFirmTradeID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.CalculatedCcyLastQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradeCaptureReportAck_RegexValidator.LastSwapPoints = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradeCaptureReportAck_RegexValidator.GrossTradeAmt = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradeCaptureReportAck_RegexValidator.TradeHandlingInstr = regexp.MustCompile(`[^|]`)
    myTradeCaptureReportAck_RegexValidator.OrigTradeHandlingInstr = regexp.MustCompile(`[^|]`)
    myTradeCaptureReportAck_RegexValidator.OrigTradeDate = regexp.MustCompile(`[0-9]{8}`)
    myTradeCaptureReportAck_RegexValidator.OrigTradeID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.OrigSecondaryTradeID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.RptSys = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportAck_RegexValidator.SignatureLength = regexp.MustCompile(`\d+`)
    myTradeCaptureReportAck_RegexValidator.Signature = regexp.MustCompile(`.*`)
    myTradeCaptureReportAck_RegexValidator.CheckSum = regexp.MustCompile(`[^|]*`)
}



func (m *TradeCaptureReportAck) HasRequiredFields() bool {
    valid := true
    
    isSetTrdCapRptAckSideGrp := false
    isValidTrdCapRptAckSideGrp := true

    
    
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
    
    
    
    for _, g := range(m.TrdRegTimestamps) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Legs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.PosAmt) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    isSetTrdCapRptAckSideGrp = isSetTrdCapRptAckSideGrp || len(m.Sides) > 0
    for _, g := range(m.Sides) {
        isValidTrdCapRptAckSideGrp = isValidTrdCapRptAckSideGrp && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.RootPartyIDs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Underlyings) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.CheckSum.HasValue()
    
    
    
    valid = valid && (isSetTrdCapRptAckSideGrp == false || isValidTrdCapRptAckSideGrp)
    
    return valid
}



