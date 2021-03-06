package TrdCapRptSideGrp_Sides

import "bufio"
import "grgrbll/Fix/FixDef/Types"

// Autogenerated at , do not edit



import Parties_PartyIDs "grgrbll/fix/Shared/Parties_PartyIDs"

import ClrInstGrp_ClearingInstructions "grgrbll/fix/Shared/ClrInstGrp_ClearingInstructions"

import ContAmtGrp_ContAmts "grgrbll/fix/Shared/ContAmtGrp_ContAmts"

import Stipulations_Stipulations "grgrbll/fix/Shared/Stipulations_Stipulations"

import MiscFeesGrp_MiscFees "grgrbll/fix/Shared/MiscFeesGrp_MiscFees"

import TrdAllocGrp_Allocs "grgrbll/fix/Shared/TrdAllocGrp_Allocs"

import SideTrdRegTS_SideTrdRegTS "grgrbll/fix/Shared/SideTrdRegTS_SideTrdRegTS"


type TrdCapRptSideGrp_Sides struct {
    Side Types.Char `id:"54", required:"Y" `
    OrderID Types.String `id:"37", required:"N" `
    SecondaryOrderID Types.String `id:"198", required:"N" `
    ClOrdID Types.String `id:"11", required:"N" `
    SecondaryClOrdID Types.String `id:"526", required:"N" `
    ListID Types.String `id:"66", required:"N" `
    PartyIDs []Parties_PartyIDs.Parties_PartyIDs `counter_name:"NoPartyIDs", counter_id:"453"`
    Account Types.String `id:"1", required:"N" `
    AcctIDSource Types.Int `id:"660", required:"N" `
    AccountType Types.Int `id:"581", required:"N" `
    ProcessCode Types.Char `id:"81", required:"N" `
    OddLot Types.Boolean `id:"575", required:"N" `
    ClearingInstructions []ClrInstGrp_ClearingInstructions.ClrInstGrp_ClearingInstructions `counter_name:"NoClearingInstructions", counter_id:"576"`
    TradeInputSource Types.String `id:"578", required:"N" `
    TradeInputDevice Types.String `id:"579", required:"N" `
    OrderInputDevice Types.String `id:"821", required:"N" `
    Currency Types.Currency `id:"15", required:"N" `
    ComplianceID Types.String `id:"376", required:"N" `
    SolicitedFlag Types.Boolean `id:"377", required:"N" `
    OrderCapacity Types.Char `id:"528", required:"N" `
    OrderRestrictions Types.MultipleCharValue `id:"529", required:"N" `
    CustOrderCapacity Types.Int `id:"582", required:"N" `
    OrdType Types.Char `id:"40", required:"N" `
    ExecInst Types.MultipleCharValue `id:"18", required:"N" `
    TransBkdTime Types.UTCTimestamp `id:"483", required:"N" `
    TradingSessionID Types.String `id:"336", required:"N" `
    TradingSessionSubID Types.String `id:"625", required:"N" `
    TimeBracket Types.String `id:"943", required:"N" `
    Commission Types.Amt `id:"12", required:"N" `
    CommType Types.Char `id:"13", required:"N" `
    CommCurrency Types.Currency `id:"479", required:"N" `
    FundRenewWaiv Types.Char `id:"497", required:"N" `
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
    NetMoney Types.Amt `id:"118", required:"N" `
    SettlCurrAmt Types.Amt `id:"119", required:"N" `
    SettlCurrency Types.Currency `id:"120", required:"N" `
    SettlCurrFxRate Types.Float `id:"155", required:"N" `
    SettlCurrFxRateCalc Types.Char `id:"156", required:"N" `
    PositionEffect Types.Char `id:"77", required:"N" `
    Text Types.String `id:"58", required:"N" `
    EncodedTextLen Types.Length `id:"354", required:"N" `
    EncodedText Types.Data `id:"355", required:"N" `
    SideMultiLegReportingType Types.Int `id:"752", required:"N" `
    ContAmts []ContAmtGrp_ContAmts.ContAmtGrp_ContAmts `counter_name:"NoContAmts", counter_id:"518"`
    Stipulations []Stipulations_Stipulations.Stipulations_Stipulations `counter_name:"NoStipulations", counter_id:"232"`
    MiscFees []MiscFeesGrp_MiscFees.MiscFeesGrp_MiscFees `counter_name:"NoMiscFees", counter_id:"136"`
    ExchangeRule Types.String `id:"825", required:"N" `
    TradeAllocIndicator Types.Int `id:"826", required:"N" `
    PreallocMethod Types.Char `id:"591", required:"N" `
    AllocID Types.String `id:"70", required:"N" `
    Allocs []TrdAllocGrp_Allocs.TrdAllocGrp_Allocs `counter_name:"NoAllocs", counter_id:"78"`
    SideQty Types.Int `id:"1009", required:"N" `
    SideTradeReportID Types.String `id:"1005", required:"N" `
    SideFillStationCd Types.String `id:"1006", required:"N" `
    SideReasonCd Types.String `id:"1007", required:"N" `
    RptSeq Types.Int `id:"83", required:"N" `
    SideTrdSubTyp Types.Int `id:"1008", required:"N" `
    SideTrdRegTS []SideTrdRegTS_SideTrdRegTS.SideTrdRegTS_SideTrdRegTS `counter_name:"NoSideTrdRegTS", counter_id:"1016"`
    ExecRefID Types.String `id:"19", required:"N" `
    LotType Types.Char `id:"1093", required:"N" `
    SideGrossTradeAmt Types.Amt `id:"1072", required:"N" `
    AggressorIndicator Types.Boolean `id:"1057", required:"N" `
    ExchangeSpecialInstructions Types.String `id:"1139", required:"N" `
    _controlBlock fixMessageControlBlock
}
