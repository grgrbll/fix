package QuotReqLegsGrp_Legs

import "bufio"
import "grgrbll/Fix/FixDef/Types"

// Autogenerated at , do not edit



import LegSecAltIDGrp_LegSecurityAltID "grgrbll/fix/Shared/LegSecAltIDGrp_LegSecurityAltID"

import LegStipulations_LegStipulations "grgrbll/fix/Shared/LegStipulations_LegStipulations"

import NestedParties_NestedPartyIDs "grgrbll/fix/Shared/NestedParties_NestedPartyIDs"


type QuotReqLegsGrp_Legs struct {
    LegSymbol Types.String `id:"600", required:"N" `
    LegSymbolSfx Types.String `id:"601", required:"N" `
    LegSecurityID Types.String `id:"602", required:"N" `
    LegSecurityIDSource Types.String `id:"603", required:"N" `
    LegSecurityAltID []LegSecAltIDGrp_LegSecurityAltID.LegSecAltIDGrp_LegSecurityAltID `counter_name:"NoLegSecurityAltID", counter_id:"604"`
    LegProduct Types.Int `id:"607", required:"N" `
    LegCFICode Types.String `id:"608", required:"N" `
    LegSecurityType Types.String `id:"609", required:"N" `
    LegSecuritySubType Types.String `id:"764", required:"N" `
    LegMaturityMonthYear Types.MonthYear `id:"610", required:"N" `
    LegMaturityDate Types.LocalMktDate `id:"611", required:"N" `
    LegCouponPaymentDate Types.LocalMktDate `id:"248", required:"N" `
    LegIssueDate Types.LocalMktDate `id:"249", required:"N" `
    LegRepoCollateralSecurityType Types.Int `id:"250", required:"N" `
    LegRepurchaseTerm Types.Int `id:"251", required:"N" `
    LegRepurchaseRate Types.Percentage `id:"252", required:"N" `
    LegFactor Types.Float `id:"253", required:"N" `
    LegCreditRating Types.String `id:"257", required:"N" `
    LegInstrRegistry Types.String `id:"599", required:"N" `
    LegCountryOfIssue Types.Country `id:"596", required:"N" `
    LegStateOrProvinceOfIssue Types.String `id:"597", required:"N" `
    LegLocaleOfIssue Types.String `id:"598", required:"N" `
    LegRedemptionDate Types.LocalMktDate `id:"254", required:"N" `
    LegStrikePrice Types.Price `id:"612", required:"N" `
    LegStrikeCurrency Types.Currency `id:"942", required:"N" `
    LegOptAttribute Types.Char `id:"613", required:"N" `
    LegContractMultiplier Types.Float `id:"614", required:"N" `
    LegCouponRate Types.Percentage `id:"615", required:"N" `
    LegSecurityExchange Types.Exchange `id:"616", required:"N" `
    LegIssuer Types.String `id:"617", required:"N" `
    EncodedLegIssuerLen Types.Length `id:"618", required:"N" `
    EncodedLegIssuer Types.Data `id:"619", required:"N" `
    LegSecurityDesc Types.String `id:"620", required:"N" `
    EncodedLegSecurityDescLen Types.Length `id:"621", required:"N" `
    EncodedLegSecurityDesc Types.Data `id:"622", required:"N" `
    LegRatioQty Types.Float `id:"623", required:"N" `
    LegSide Types.Char `id:"624", required:"N" `
    LegCurrency Types.Currency `id:"556", required:"N" `
    LegPool Types.String `id:"740", required:"N" `
    LegDatedDate Types.LocalMktDate `id:"739", required:"N" `
    LegContractSettlMonth Types.MonthYear `id:"955", required:"N" `
    LegInterestAccrualDate Types.LocalMktDate `id:"956", required:"N" `
    LegUnitofMeasure Types.String `id:"999", required:"N" `
    LegTimeUnit Types.String `id:"1001", required:"N" `
    LegOptionRatio Types.Float `id:"1017", required:"N" `
    LegPrice Types.Price `id:"566", required:"N" `
    LegQty Types.Qty `id:"687", required:"N" `
    LegSwapType Types.Int `id:"690", required:"N" `
    LegSettlType Types.Char `id:"587", required:"N" `
    LegSettlDate Types.LocalMktDate `id:"588", required:"N" `
    LegStipulations []LegStipulations_LegStipulations.LegStipulations_LegStipulations `counter_name:"NoLegStipulations", counter_id:"683"`
    NestedPartyIDs []NestedParties_NestedPartyIDs.NestedParties_NestedPartyIDs `counter_name:"NoNestedPartyIDs", counter_id:"539"`
    LegBenchmarkCurveCurrency Types.Currency `id:"676", required:"N" `
    LegBenchmarkCurveName Types.String `id:"677", required:"N" `
    LegBenchmarkCurvePoint Types.String `id:"678", required:"N" `
    LegBenchmarkPrice Types.Price `id:"679", required:"N" `
    LegBenchmarkPriceType Types.Int `id:"680", required:"N" `
    LegOrderQty Types.Qty `id:"685", required:"N" `
    LegRefID Types.String `id:"654", required:"N" `
    _controlBlock fixMessageControlBlock
}
