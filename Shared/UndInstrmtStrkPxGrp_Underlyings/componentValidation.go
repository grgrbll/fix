package UndInstrmtStrkPxGrp_Underlyings

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import UndSecAltIDGrp_UnderlyingSecurityAltID "grgrbll/fix/Shared/UndSecAltIDGrp_UnderlyingSecurityAltID"

import UnderlyingStipulations_UnderlyingStips "grgrbll/fix/Shared/UnderlyingStipulations_UnderlyingStips"

import UndlyInstrumentParties_UndlyInstrumentParties "grgrbll/fix/Shared/UndlyInstrumentParties_UndlyInstrumentParties"


type undInstrmtStrkPxGrp_Underlyings_RegexValidator struct {
    UnderlyingSymbol *(regexp.Regexp)
    UnderlyingSymbolSfx *(regexp.Regexp)
    UnderlyingSecurityID *(regexp.Regexp)
    UnderlyingSecurityIDSource *(regexp.Regexp)
    UnderlyingProduct *(regexp.Regexp)
    UnderlyingCFICode *(regexp.Regexp)
    UnderlyingSecurityType *(regexp.Regexp)
    UnderlyingSecuritySubType *(regexp.Regexp)
    UnderlyingMaturityMonthYear *(regexp.Regexp)
    UnderlyingMaturityDate *(regexp.Regexp)
    UnderlyingCouponPaymentDate *(regexp.Regexp)
    UnderlyingIssueDate *(regexp.Regexp)
    UnderlyingRepoCollateralSecurityType *(regexp.Regexp)
    UnderlyingRepurchaseTerm *(regexp.Regexp)
    UnderlyingRepurchaseRate *(regexp.Regexp)
    UnderlyingFactor *(regexp.Regexp)
    UnderlyingCreditRating *(regexp.Regexp)
    UnderlyingInstrRegistry *(regexp.Regexp)
    UnderlyingCountryOfIssue *(regexp.Regexp)
    UnderlyingStateOrProvinceOfIssue *(regexp.Regexp)
    UnderlyingLocaleOfIssue *(regexp.Regexp)
    UnderlyingRedemptionDate *(regexp.Regexp)
    UnderlyingStrikePrice *(regexp.Regexp)
    UnderlyingStrikeCurrency *(regexp.Regexp)
    UnderlyingOptAttribute *(regexp.Regexp)
    UnderlyingContractMultiplier *(regexp.Regexp)
    UnderlyingCouponRate *(regexp.Regexp)
    UnderlyingSecurityExchange *(regexp.Regexp)
    UnderlyingIssuer *(regexp.Regexp)
    EncodedUnderlyingIssuerLen *(regexp.Regexp)
    EncodedUnderlyingIssuer *(regexp.Regexp)
    UnderlyingSecurityDesc *(regexp.Regexp)
    EncodedUnderlyingSecurityDescLen *(regexp.Regexp)
    EncodedUnderlyingSecurityDesc *(regexp.Regexp)
    UnderlyingCPProgram *(regexp.Regexp)
    UnderlyingCPRegType *(regexp.Regexp)
    UnderlyingCurrency *(regexp.Regexp)
    UnderlyingQty *(regexp.Regexp)
    UnderlyingPx *(regexp.Regexp)
    UnderlyingDirtyPrice *(regexp.Regexp)
    UnderlyingEndPrice *(regexp.Regexp)
    UnderlyingStartValue *(regexp.Regexp)
    UnderlyingCurrentValue *(regexp.Regexp)
    UnderlyingEndValue *(regexp.Regexp)
    UnderlyingAllocationPercent *(regexp.Regexp)
    UnderlyingSettlementType *(regexp.Regexp)
    UnderlyingCashAmount *(regexp.Regexp)
    UnderlyingCashType *(regexp.Regexp)
    UnderlyingUnitofMeasure *(regexp.Regexp)
    UnderlyingTimeUnit *(regexp.Regexp)
    UnderlyingCapValue *(regexp.Regexp)
    UnderlyingSettlMethod *(regexp.Regexp)
    UnderlyingAdjustedQuantity *(regexp.Regexp)
    UnderlyingFXRate *(regexp.Regexp)
    UnderlyingFXRateCalc *(regexp.Regexp)
    PrevClosePx *(regexp.Regexp)
    ClOrdID *(regexp.Regexp)
    SecondaryClOrdID *(regexp.Regexp)
    Side *(regexp.Regexp)
    Price *(regexp.Regexp)
    Currency *(regexp.Regexp)
    Text *(regexp.Regexp)
    EncodedTextLen *(regexp.Regexp)
    EncodedText *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myUndInstrmtStrkPxGrp_Underlyings_RegexValidator undInstrmtStrkPxGrp_Underlyings_RegexValidator

func init() {
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingSymbol = regexp.MustCompile(`[^|]*`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingSymbolSfx = regexp.MustCompile(`[^|]*`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingSecurityID = regexp.MustCompile(`[^|]*`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingSecurityIDSource = regexp.MustCompile(`[^|]*`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingProduct = regexp.MustCompile(`-?\d+`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingCFICode = regexp.MustCompile(`[^|]*`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingSecurityType = regexp.MustCompile(`[^|]*`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingSecuritySubType = regexp.MustCompile(`[^|]*`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingMaturityMonthYear = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingMaturityDate = regexp.MustCompile(`[0-9]{8}`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingCouponPaymentDate = regexp.MustCompile(`[0-9]{8}`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingIssueDate = regexp.MustCompile(`[0-9]{8}`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingRepoCollateralSecurityType = regexp.MustCompile(`-?\d+`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingRepurchaseTerm = regexp.MustCompile(`-?\d+`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingRepurchaseRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingFactor = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingCreditRating = regexp.MustCompile(`[^|]*`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingInstrRegistry = regexp.MustCompile(`[^|]*`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingCountryOfIssue = regexp.MustCompile(`[A-Z]{2}`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingStateOrProvinceOfIssue = regexp.MustCompile(`[^|]*`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingLocaleOfIssue = regexp.MustCompile(`[^|]*`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingRedemptionDate = regexp.MustCompile(`[0-9]{8}`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingStrikePrice = regexp.MustCompile(``)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingStrikeCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingOptAttribute = regexp.MustCompile(`[^|]`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingContractMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingCouponRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingSecurityExchange = regexp.MustCompile(`[A-Z0-9]{4}`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingIssuer = regexp.MustCompile(`[^|]*`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.EncodedUnderlyingIssuerLen = regexp.MustCompile(`\d+`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.EncodedUnderlyingIssuer = regexp.MustCompile(`.*`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingSecurityDesc = regexp.MustCompile(`[^|]*`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.EncodedUnderlyingSecurityDescLen = regexp.MustCompile(`\d+`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.EncodedUnderlyingSecurityDesc = regexp.MustCompile(`.*`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingCPProgram = regexp.MustCompile(`[^|]*`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingCPRegType = regexp.MustCompile(`[^|]*`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingPx = regexp.MustCompile(``)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingDirtyPrice = regexp.MustCompile(``)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingEndPrice = regexp.MustCompile(``)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingStartValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingCurrentValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingEndValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingAllocationPercent = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingSettlementType = regexp.MustCompile(`-?\d+`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingCashAmount = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingCashType = regexp.MustCompile(`[^|]*`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingUnitofMeasure = regexp.MustCompile(`[^|]*`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingTimeUnit = regexp.MustCompile(`[^|]*`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingCapValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingSettlMethod = regexp.MustCompile(`[^|]*`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingAdjustedQuantity = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingFXRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.UnderlyingFXRateCalc = regexp.MustCompile(`[^|]`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.PrevClosePx = regexp.MustCompile(``)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.ClOrdID = regexp.MustCompile(`[^|]*`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.SecondaryClOrdID = regexp.MustCompile(`[^|]*`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.Side = regexp.MustCompile(`[^|]`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.Price = regexp.MustCompile(``)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.Currency = regexp.MustCompile(`[A-Z]{3}`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.Text = regexp.MustCompile(`[^|]*`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.EncodedTextLen = regexp.MustCompile(`\d+`)
    myUndInstrmtStrkPxGrp_Underlyings_RegexValidator.EncodedText = regexp.MustCompile(`.*`)
}



func (m *UndInstrmtStrkPxGrp_Underlyings) HasRequiredFields() bool {
    valid := true
    
    

    
    
    for _, g := range(m.UnderlyingSecurityAltID) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.UnderlyingStips) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.UndlyInstrumentParties) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.Price.HasValue()
    
    
    
    
    
    return valid
}




