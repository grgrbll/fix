package TrdCapRptSideGrp_Sides

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import Parties_PartyIDs "grgrbll/fix/Shared/Parties_PartyIDs"

import ClrInstGrp_ClearingInstructions "grgrbll/fix/Shared/ClrInstGrp_ClearingInstructions"

import ContAmtGrp_ContAmts "grgrbll/fix/Shared/ContAmtGrp_ContAmts"

import Stipulations_Stipulations "grgrbll/fix/Shared/Stipulations_Stipulations"

import MiscFeesGrp_MiscFees "grgrbll/fix/Shared/MiscFeesGrp_MiscFees"

import TrdAllocGrp_Allocs "grgrbll/fix/Shared/TrdAllocGrp_Allocs"

import SideTrdRegTS_SideTrdRegTS "grgrbll/fix/Shared/SideTrdRegTS_SideTrdRegTS"


type trdCapRptSideGrp_Sides_RegexValidator struct {
    Side *(regexp.Regexp)
    OrderID *(regexp.Regexp)
    SecondaryOrderID *(regexp.Regexp)
    ClOrdID *(regexp.Regexp)
    SecondaryClOrdID *(regexp.Regexp)
    ListID *(regexp.Regexp)
    Account *(regexp.Regexp)
    AcctIDSource *(regexp.Regexp)
    AccountType *(regexp.Regexp)
    ProcessCode *(regexp.Regexp)
    OddLot *(regexp.Regexp)
    TradeInputSource *(regexp.Regexp)
    TradeInputDevice *(regexp.Regexp)
    OrderInputDevice *(regexp.Regexp)
    Currency *(regexp.Regexp)
    ComplianceID *(regexp.Regexp)
    SolicitedFlag *(regexp.Regexp)
    OrderCapacity *(regexp.Regexp)
    OrderRestrictions *(regexp.Regexp)
    CustOrderCapacity *(regexp.Regexp)
    OrdType *(regexp.Regexp)
    ExecInst *(regexp.Regexp)
    TransBkdTime *(regexp.Regexp)
    TradingSessionID *(regexp.Regexp)
    TradingSessionSubID *(regexp.Regexp)
    TimeBracket *(regexp.Regexp)
    Commission *(regexp.Regexp)
    CommType *(regexp.Regexp)
    CommCurrency *(regexp.Regexp)
    FundRenewWaiv *(regexp.Regexp)
    NumDaysInterest *(regexp.Regexp)
    ExDate *(regexp.Regexp)
    AccruedInterestRate *(regexp.Regexp)
    AccruedInterestAmt *(regexp.Regexp)
    InterestAtMaturity *(regexp.Regexp)
    EndAccruedInterestAmt *(regexp.Regexp)
    StartCash *(regexp.Regexp)
    EndCash *(regexp.Regexp)
    Concession *(regexp.Regexp)
    TotalTakedown *(regexp.Regexp)
    NetMoney *(regexp.Regexp)
    SettlCurrAmt *(regexp.Regexp)
    SettlCurrency *(regexp.Regexp)
    SettlCurrFxRate *(regexp.Regexp)
    SettlCurrFxRateCalc *(regexp.Regexp)
    PositionEffect *(regexp.Regexp)
    Text *(regexp.Regexp)
    EncodedTextLen *(regexp.Regexp)
    EncodedText *(regexp.Regexp)
    SideMultiLegReportingType *(regexp.Regexp)
    ExchangeRule *(regexp.Regexp)
    TradeAllocIndicator *(regexp.Regexp)
    PreallocMethod *(regexp.Regexp)
    AllocID *(regexp.Regexp)
    SideQty *(regexp.Regexp)
    SideTradeReportID *(regexp.Regexp)
    SideFillStationCd *(regexp.Regexp)
    SideReasonCd *(regexp.Regexp)
    RptSeq *(regexp.Regexp)
    SideTrdSubTyp *(regexp.Regexp)
    ExecRefID *(regexp.Regexp)
    LotType *(regexp.Regexp)
    SideGrossTradeAmt *(regexp.Regexp)
    AggressorIndicator *(regexp.Regexp)
    ExchangeSpecialInstructions *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myTrdCapRptSideGrp_Sides_RegexValidator trdCapRptSideGrp_Sides_RegexValidator

func init() {
    myTrdCapRptSideGrp_Sides_RegexValidator
    myTrdCapRptSideGrp_Sides_RegexValidator.Side = regexp.MustCompile(`[^|]`)
    myTrdCapRptSideGrp_Sides_RegexValidator.OrderID = regexp.MustCompile(`[^|]*`)
    myTrdCapRptSideGrp_Sides_RegexValidator.SecondaryOrderID = regexp.MustCompile(`[^|]*`)
    myTrdCapRptSideGrp_Sides_RegexValidator.ClOrdID = regexp.MustCompile(`[^|]*`)
    myTrdCapRptSideGrp_Sides_RegexValidator.SecondaryClOrdID = regexp.MustCompile(`[^|]*`)
    myTrdCapRptSideGrp_Sides_RegexValidator.ListID = regexp.MustCompile(`[^|]*`)
    myTrdCapRptSideGrp_Sides_RegexValidator.Account = regexp.MustCompile(`[^|]*`)
    myTrdCapRptSideGrp_Sides_RegexValidator.AcctIDSource = regexp.MustCompile(`-?\d+`)
    myTrdCapRptSideGrp_Sides_RegexValidator.AccountType = regexp.MustCompile(`-?\d+`)
    myTrdCapRptSideGrp_Sides_RegexValidator.ProcessCode = regexp.MustCompile(`[^|]`)
    myTrdCapRptSideGrp_Sides_RegexValidator.OddLot = regexp.MustCompile(`[YN]`)
    myTrdCapRptSideGrp_Sides_RegexValidator.TradeInputSource = regexp.MustCompile(`[^|]*`)
    myTrdCapRptSideGrp_Sides_RegexValidator.TradeInputDevice = regexp.MustCompile(`[^|]*`)
    myTrdCapRptSideGrp_Sides_RegexValidator.OrderInputDevice = regexp.MustCompile(`[^|]*`)
    myTrdCapRptSideGrp_Sides_RegexValidator.Currency = regexp.MustCompile(`[A-Z]{3}`)
    myTrdCapRptSideGrp_Sides_RegexValidator.ComplianceID = regexp.MustCompile(`[^|]*`)
    myTrdCapRptSideGrp_Sides_RegexValidator.SolicitedFlag = regexp.MustCompile(`[YN]`)
    myTrdCapRptSideGrp_Sides_RegexValidator.OrderCapacity = regexp.MustCompile(`[^|]`)
    myTrdCapRptSideGrp_Sides_RegexValidator.OrderRestrictions = regexp.MustCompile(`[ ]?(\w[ ])+\w?`)
    myTrdCapRptSideGrp_Sides_RegexValidator.CustOrderCapacity = regexp.MustCompile(`-?\d+`)
    myTrdCapRptSideGrp_Sides_RegexValidator.OrdType = regexp.MustCompile(`[^|]`)
    myTrdCapRptSideGrp_Sides_RegexValidator.ExecInst = regexp.MustCompile(`[ ]?(\w[ ])+\w?`)
    myTrdCapRptSideGrp_Sides_RegexValidator.TransBkdTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myTrdCapRptSideGrp_Sides_RegexValidator.TradingSessionID = regexp.MustCompile(`[^|]*`)
    myTrdCapRptSideGrp_Sides_RegexValidator.TradingSessionSubID = regexp.MustCompile(`[^|]*`)
    myTrdCapRptSideGrp_Sides_RegexValidator.TimeBracket = regexp.MustCompile(`[^|]*`)
    myTrdCapRptSideGrp_Sides_RegexValidator.Commission = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTrdCapRptSideGrp_Sides_RegexValidator.CommType = regexp.MustCompile(`[^|]`)
    myTrdCapRptSideGrp_Sides_RegexValidator.CommCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myTrdCapRptSideGrp_Sides_RegexValidator.FundRenewWaiv = regexp.MustCompile(`[^|]`)
    myTrdCapRptSideGrp_Sides_RegexValidator.NumDaysInterest = regexp.MustCompile(`-?\d+`)
    myTrdCapRptSideGrp_Sides_RegexValidator.ExDate = regexp.MustCompile(`[0-9]{8}`)
    myTrdCapRptSideGrp_Sides_RegexValidator.AccruedInterestRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTrdCapRptSideGrp_Sides_RegexValidator.AccruedInterestAmt = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTrdCapRptSideGrp_Sides_RegexValidator.InterestAtMaturity = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTrdCapRptSideGrp_Sides_RegexValidator.EndAccruedInterestAmt = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTrdCapRptSideGrp_Sides_RegexValidator.StartCash = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTrdCapRptSideGrp_Sides_RegexValidator.EndCash = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTrdCapRptSideGrp_Sides_RegexValidator.Concession = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTrdCapRptSideGrp_Sides_RegexValidator.TotalTakedown = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTrdCapRptSideGrp_Sides_RegexValidator.NetMoney = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTrdCapRptSideGrp_Sides_RegexValidator.SettlCurrAmt = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTrdCapRptSideGrp_Sides_RegexValidator.SettlCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myTrdCapRptSideGrp_Sides_RegexValidator.SettlCurrFxRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTrdCapRptSideGrp_Sides_RegexValidator.SettlCurrFxRateCalc = regexp.MustCompile(`[^|]`)
    myTrdCapRptSideGrp_Sides_RegexValidator.PositionEffect = regexp.MustCompile(`[^|]`)
    myTrdCapRptSideGrp_Sides_RegexValidator.Text = regexp.MustCompile(`[^|]*`)
    myTrdCapRptSideGrp_Sides_RegexValidator.EncodedTextLen = regexp.MustCompile(`\d+`)
    myTrdCapRptSideGrp_Sides_RegexValidator.EncodedText = regexp.MustCompile(`.*`)
    myTrdCapRptSideGrp_Sides_RegexValidator.SideMultiLegReportingType = regexp.MustCompile(`-?\d+`)
    myTrdCapRptSideGrp_Sides_RegexValidator.ExchangeRule = regexp.MustCompile(`[^|]*`)
    myTrdCapRptSideGrp_Sides_RegexValidator.TradeAllocIndicator = regexp.MustCompile(`-?\d+`)
    myTrdCapRptSideGrp_Sides_RegexValidator.PreallocMethod = regexp.MustCompile(`[^|]`)
    myTrdCapRptSideGrp_Sides_RegexValidator.AllocID = regexp.MustCompile(`[^|]*`)
    myTrdCapRptSideGrp_Sides_RegexValidator.SideQty = regexp.MustCompile(`-?\d+`)
    myTrdCapRptSideGrp_Sides_RegexValidator.SideTradeReportID = regexp.MustCompile(`[^|]*`)
    myTrdCapRptSideGrp_Sides_RegexValidator.SideFillStationCd = regexp.MustCompile(`[^|]*`)
    myTrdCapRptSideGrp_Sides_RegexValidator.SideReasonCd = regexp.MustCompile(`[^|]*`)
    myTrdCapRptSideGrp_Sides_RegexValidator.RptSeq = regexp.MustCompile(`-?\d+`)
    myTrdCapRptSideGrp_Sides_RegexValidator.SideTrdSubTyp = regexp.MustCompile(`-?\d+`)
    myTrdCapRptSideGrp_Sides_RegexValidator.ExecRefID = regexp.MustCompile(`[^|]*`)
    myTrdCapRptSideGrp_Sides_RegexValidator.LotType = regexp.MustCompile(`[^|]`)
    myTrdCapRptSideGrp_Sides_RegexValidator.SideGrossTradeAmt = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTrdCapRptSideGrp_Sides_RegexValidator.AggressorIndicator = regexp.MustCompile(`[YN]`)
    myTrdCapRptSideGrp_Sides_RegexValidator.ExchangeSpecialInstructions = regexp.MustCompile(`[^|]*`)
}



func (m *TrdCapRptSideGrp_Sides) HasRequiredFields() bool {
    valid := true
    
    

    
    
    valid = valid && m.Side.HasValue()
    
    
    
    for _, g := range(m.PartyIDs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.ClearingInstructions) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.ContAmts) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Stipulations) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.MiscFees) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.Allocs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.SideTrdRegTS) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    
    
    return valid
}



