package UndInstrmtGrp_Underlyings

import "bufio"
import "grgrbll/Fix/FixDef/Types"

// Autogenerated at , do not edit



import UndSecAltIDGrp_UnderlyingSecurityAltID "grgrbll/fix/Shared/UndSecAltIDGrp_UnderlyingSecurityAltID"

import UnderlyingStipulations_UnderlyingStips "grgrbll/fix/Shared/UnderlyingStipulations_UnderlyingStips"

import UndlyInstrumentParties_UndlyInstrumentParties "grgrbll/fix/Shared/UndlyInstrumentParties_UndlyInstrumentParties"


type UndInstrmtGrp_Underlyings struct {
    UnderlyingSymbol Types.String `id:"311", required:"N" `
    UnderlyingSymbolSfx Types.String `id:"312", required:"N" `
    UnderlyingSecurityID Types.String `id:"309", required:"N" `
    UnderlyingSecurityIDSource Types.String `id:"305", required:"N" `
    UnderlyingSecurityAltID []UndSecAltIDGrp_UnderlyingSecurityAltID.UndSecAltIDGrp_UnderlyingSecurityAltID `counter_name:"NoUnderlyingSecurityAltID", counter_id:"457"`
    UnderlyingProduct Types.Int `id:"462", required:"N" `
    UnderlyingCFICode Types.String `id:"463", required:"N" `
    UnderlyingSecurityType Types.String `id:"310", required:"N" `
    UnderlyingSecuritySubType Types.String `id:"763", required:"N" `
    UnderlyingMaturityMonthYear Types.MonthYear `id:"313", required:"N" `
    UnderlyingMaturityDate Types.LocalMktDate `id:"542", required:"N" `
    UnderlyingCouponPaymentDate Types.LocalMktDate `id:"241", required:"N" `
    UnderlyingIssueDate Types.LocalMktDate `id:"242", required:"N" `
    UnderlyingRepoCollateralSecurityType Types.Int `id:"243", required:"N" `
    UnderlyingRepurchaseTerm Types.Int `id:"244", required:"N" `
    UnderlyingRepurchaseRate Types.Percentage `id:"245", required:"N" `
    UnderlyingFactor Types.Float `id:"246", required:"N" `
    UnderlyingCreditRating Types.String `id:"256", required:"N" `
    UnderlyingInstrRegistry Types.String `id:"595", required:"N" `
    UnderlyingCountryOfIssue Types.Country `id:"592", required:"N" `
    UnderlyingStateOrProvinceOfIssue Types.String `id:"593", required:"N" `
    UnderlyingLocaleOfIssue Types.String `id:"594", required:"N" `
    UnderlyingRedemptionDate Types.LocalMktDate `id:"247", required:"N" `
    UnderlyingStrikePrice Types.Price `id:"316", required:"N" `
    UnderlyingStrikeCurrency Types.Currency `id:"941", required:"N" `
    UnderlyingOptAttribute Types.Char `id:"317", required:"N" `
    UnderlyingContractMultiplier Types.Float `id:"436", required:"N" `
    UnderlyingCouponRate Types.Percentage `id:"435", required:"N" `
    UnderlyingSecurityExchange Types.Exchange `id:"308", required:"N" `
    UnderlyingIssuer Types.String `id:"306", required:"N" `
    EncodedUnderlyingIssuerLen Types.Length `id:"362", required:"N" `
    EncodedUnderlyingIssuer Types.Data `id:"363", required:"N" `
    UnderlyingSecurityDesc Types.String `id:"307", required:"N" `
    EncodedUnderlyingSecurityDescLen Types.Length `id:"364", required:"N" `
    EncodedUnderlyingSecurityDesc Types.Data `id:"365", required:"N" `
    UnderlyingCPProgram Types.String `id:"877", required:"N" `
    UnderlyingCPRegType Types.String `id:"878", required:"N" `
    UnderlyingCurrency Types.Currency `id:"318", required:"N" `
    UnderlyingQty Types.Qty `id:"879", required:"N" `
    UnderlyingPx Types.Price `id:"810", required:"N" `
    UnderlyingDirtyPrice Types.Price `id:"882", required:"N" `
    UnderlyingEndPrice Types.Price `id:"883", required:"N" `
    UnderlyingStartValue Types.Amt `id:"884", required:"N" `
    UnderlyingCurrentValue Types.Amt `id:"885", required:"N" `
    UnderlyingEndValue Types.Amt `id:"886", required:"N" `
    UnderlyingStips []UnderlyingStipulations_UnderlyingStips.UnderlyingStipulations_UnderlyingStips `counter_name:"NoUnderlyingStips", counter_id:"887"`
    UnderlyingAllocationPercent Types.Percentage `id:"972", required:"N" `
    UnderlyingSettlementType Types.Int `id:"975", required:"N" `
    UnderlyingCashAmount Types.Amt `id:"973", required:"N" `
    UnderlyingCashType Types.String `id:"974", required:"N" `
    UnderlyingUnitofMeasure Types.String `id:"998", required:"N" `
    UnderlyingTimeUnit Types.String `id:"1000", required:"N" `
    UnderlyingCapValue Types.Amt `id:"1038", required:"N" `
    UndlyInstrumentParties []UndlyInstrumentParties_UndlyInstrumentParties.UndlyInstrumentParties_UndlyInstrumentParties `counter_name:"NoUndlyInstrumentParties", counter_id:"1058"`
    UnderlyingSettlMethod Types.String `id:"1039", required:"N" `
    UnderlyingAdjustedQuantity Types.Qty `id:"1044", required:"N" `
    UnderlyingFXRate Types.Float `id:"1045", required:"N" `
    UnderlyingFXRateCalc Types.Char `id:"1046", required:"N" `
    _controlBlock fixMessageControlBlock
}
