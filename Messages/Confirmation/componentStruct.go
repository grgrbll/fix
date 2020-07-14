package Confirmation

import "bufio"
import "grgrbll/Fix/FixDef/Types"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import Parties_PartyIDs "grgrbll/fix/Shared/Parties_PartyIDs"

import OrdAllocGrp_Orders "grgrbll/fix/Shared/OrdAllocGrp_Orders"

import TrdRegTimestamps_TrdRegTimestamps "grgrbll/fix/Shared/TrdRegTimestamps_TrdRegTimestamps"

import SecAltIDGrp_SecurityAltID "grgrbll/fix/Shared/SecAltIDGrp_SecurityAltID"

import EvntGrp_Events "grgrbll/fix/Shared/EvntGrp_Events"

import InstrumentParties_InstrumentParties "grgrbll/fix/Shared/InstrumentParties_InstrumentParties"

import AttrbGrp_InstrAttrib "grgrbll/fix/Shared/AttrbGrp_InstrAttrib"

import UndInstrmtGrp_Underlyings "grgrbll/fix/Shared/UndInstrmtGrp_Underlyings"

import InstrmtLegGrp_Legs "grgrbll/fix/Shared/InstrmtLegGrp_Legs"

import CpctyConfGrp_Capacities "grgrbll/fix/Shared/CpctyConfGrp_Capacities"

import DlvyInstGrp_DlvyInst "grgrbll/fix/Shared/DlvyInstGrp_DlvyInst"

import Stipulations_Stipulations "grgrbll/fix/Shared/Stipulations_Stipulations"

import MiscFeesGrp_MiscFees "grgrbll/fix/Shared/MiscFeesGrp_MiscFees"


type Confirmation struct {
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
    ConfirmID Types.String `id:"664", required:"Y" `
    ConfirmRefID Types.String `id:"772", required:"N" `
    ConfirmReqID Types.String `id:"859", required:"N" `
    ConfirmTransType Types.Int `id:"666", required:"Y" `
    ConfirmType Types.Int `id:"773", required:"Y" `
    CopyMsgIndicator Types.Boolean `id:"797", required:"N" `
    LegalConfirm Types.Boolean `id:"650", required:"N" `
    ConfirmStatus Types.Int `id:"665", required:"Y" `
    PartyIDs []Parties_PartyIDs.Parties_PartyIDs `counter_name:"NoPartyIDs", counter_id:"453"`
    Orders []OrdAllocGrp_Orders.OrdAllocGrp_Orders `counter_name:"NoOrders", counter_id:"73"`
    AllocID Types.String `id:"70", required:"N" `
    SecondaryAllocID Types.String `id:"793", required:"N" `
    IndividualAllocID Types.String `id:"467", required:"N" `
    TransactTime Types.UTCTimestamp `id:"60", required:"Y" `
    TradeDate Types.LocalMktDate `id:"75", required:"Y" `
    TrdRegTimestamps []TrdRegTimestamps_TrdRegTimestamps.TrdRegTimestamps_TrdRegTimestamps `counter_name:"NoTrdRegTimestamps", counter_id:"768"`
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
    DeliveryForm Types.Int `id:"668", required:"N" `
    PctAtRisk Types.Percentage `id:"869", required:"N" `
    InstrAttrib []AttrbGrp_InstrAttrib.AttrbGrp_InstrAttrib `counter_name:"NoInstrAttrib", counter_id:"870"`
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
    Legs []InstrmtLegGrp_Legs.InstrmtLegGrp_Legs `counter_name:"NoLegs", counter_id:"555"`
    YieldType Types.String `id:"235", required:"N" `
    Yield Types.Percentage `id:"236", required:"N" `
    YieldCalcDate Types.LocalMktDate `id:"701", required:"N" `
    YieldRedemptionDate Types.LocalMktDate `id:"696", required:"N" `
    YieldRedemptionPrice Types.Price `id:"697", required:"N" `
    YieldRedemptionPriceType Types.Int `id:"698", required:"N" `
    AllocQty Types.Qty `id:"80", required:"Y" `
    QtyType Types.Int `id:"854", required:"N" `
    Side Types.Char `id:"54", required:"Y" `
    Currency Types.Currency `id:"15", required:"N" `
    LastMkt Types.Exchange `id:"30", required:"N" `
    Capacities []CpctyConfGrp_Capacities.CpctyConfGrp_Capacities `counter_name:"NoCapacities", counter_id:"862"`
    AllocAccount Types.String `id:"79", required:"Y" `
    AllocAcctIDSource Types.Int `id:"661", required:"N" `
    AllocAccountType Types.Int `id:"798", required:"N" `
    AvgPx Types.Price `id:"6", required:"Y" `
    AvgPxPrecision Types.Int `id:"74", required:"N" `
    PriceType Types.Int `id:"423", required:"N" `
    AvgParPx Types.Price `id:"860", required:"N" `
    Spread Types.PriceOffset `id:"218", required:"N" `
    BenchmarkCurveCurrency Types.Currency `id:"220", required:"N" `
    BenchmarkCurveName Types.String `id:"221", required:"N" `
    BenchmarkCurvePoint Types.String `id:"222", required:"N" `
    BenchmarkPrice Types.Price `id:"662", required:"N" `
    BenchmarkPriceType Types.Int `id:"663", required:"N" `
    BenchmarkSecurityID Types.String `id:"699", required:"N" `
    BenchmarkSecurityIDSource Types.String `id:"761", required:"N" `
    ReportedPx Types.Price `id:"861", required:"N" `
    Text Types.String `id:"58", required:"N" `
    EncodedTextLen Types.Length `id:"354", required:"N" `
    EncodedText Types.Data `id:"355", required:"N" `
    ProcessCode Types.Char `id:"81", required:"N" `
    GrossTradeAmt Types.Amt `id:"381", required:"Y" `
    NumDaysInterest Types.Int `id:"157", required:"N" `
    ExDate Types.LocalMktDate `id:"230", required:"N" `
    AccruedInterestRate Types.Percentage `id:"158", required:"N" `
    AccruedInterestAmt Types.Amt `id:"159", required:"N" `
    InterestAtMaturity Types.Amt `id:"738", required:"N" `
    EndAccruedInterestAmt Types.Amt `id:"920", required:"N" `
    StartCash Types.Amt `id:"921", required:"N" `
    EndCash Types.Amt `id:"922", required:"N" `
    Concession Types.Amt `id:"238", required:"N" `
    TotalTakedown Types.Amt `id:"237", required:"N" `
    NetMoney Types.Amt `id:"118", required:"Y" `
    MaturityNetMoney Types.Amt `id:"890", required:"N" `
    SettlCurrAmt Types.Amt `id:"119", required:"N" `
    SettlCurrency Types.Currency `id:"120", required:"N" `
    SettlCurrFxRate Types.Float `id:"155", required:"N" `
    SettlCurrFxRateCalc Types.Char `id:"156", required:"N" `
    SettlType Types.String `id:"63", required:"N" `
    SettlDate Types.LocalMktDate `id:"64", required:"N" `
    SettlDeliveryType Types.Int `id:"172", required:"N" `
    StandInstDbType Types.Int `id:"169", required:"N" `
    StandInstDbName Types.String `id:"170", required:"N" `
    StandInstDbID Types.String `id:"171", required:"N" `
    DlvyInst []DlvyInstGrp_DlvyInst.DlvyInstGrp_DlvyInst `counter_name:"NoDlvyInst", counter_id:"85"`
    Commission Types.Amt `id:"12", required:"N" `
    CommType Types.Char `id:"13", required:"N" `
    CommCurrency Types.Currency `id:"479", required:"N" `
    FundRenewWaiv Types.Char `id:"497", required:"N" `
    SharedCommission Types.Amt `id:"858", required:"N" `
    Stipulations []Stipulations_Stipulations.Stipulations_Stipulations `counter_name:"NoStipulations", counter_id:"232"`
    MiscFees []MiscFeesGrp_MiscFees.MiscFeesGrp_MiscFees `counter_name:"NoMiscFees", counter_id:"136"`
    SignatureLength Types.Length `id:"93", required:"N" `
    Signature Types.Data `id:"89", required:"N" `
    CheckSum Types.String `id:"10", required:"Y" `
    _controlBlock fixMessageControlBlock
}