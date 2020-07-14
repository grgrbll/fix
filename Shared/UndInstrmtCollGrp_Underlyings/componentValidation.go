package UndInstrmtCollGrp_Underlyings

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import UndSecAltIDGrp_UnderlyingSecurityAltID "grgrbll/fix/Shared/UndSecAltIDGrp_UnderlyingSecurityAltID"

import UnderlyingStipulations_UnderlyingStips "grgrbll/fix/Shared/UnderlyingStipulations_UnderlyingStips"

import UndlyInstrumentParties_UndlyInstrumentParties "grgrbll/fix/Shared/UndlyInstrumentParties_UndlyInstrumentParties"


type undInstrmtCollGrp_Underlyings_RegexValidator struct {
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
    CollAction *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myUndInstrmtCollGrp_Underlyings_RegexValidator undInstrmtCollGrp_Underlyings_RegexValidator

func init() {
    myUndInstrmtCollGrp_Underlyings_RegexValidator
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingSymbol = regexp.MustCompile(`[^|]*`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingSymbolSfx = regexp.MustCompile(`[^|]*`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingSecurityID = regexp.MustCompile(`[^|]*`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingSecurityIDSource = regexp.MustCompile(`[^|]*`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingProduct = regexp.MustCompile(`-?\d+`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingCFICode = regexp.MustCompile(`[^|]*`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingSecurityType = regexp.MustCompile(`[^|]*`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingSecuritySubType = regexp.MustCompile(`[^|]*`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingMaturityMonthYear = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingMaturityDate = regexp.MustCompile(`[0-9]{8}`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingCouponPaymentDate = regexp.MustCompile(`[0-9]{8}`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingIssueDate = regexp.MustCompile(`[0-9]{8}`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingRepoCollateralSecurityType = regexp.MustCompile(`-?\d+`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingRepurchaseTerm = regexp.MustCompile(`-?\d+`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingRepurchaseRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingFactor = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingCreditRating = regexp.MustCompile(`[^|]*`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingInstrRegistry = regexp.MustCompile(`[^|]*`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingCountryOfIssue = regexp.MustCompile(`[A-Z]{2}`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingStateOrProvinceOfIssue = regexp.MustCompile(`[^|]*`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingLocaleOfIssue = regexp.MustCompile(`[^|]*`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingRedemptionDate = regexp.MustCompile(`[0-9]{8}`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingStrikePrice = regexp.MustCompile(``)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingStrikeCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingOptAttribute = regexp.MustCompile(`[^|]`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingContractMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingCouponRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingSecurityExchange = regexp.MustCompile(`[A-Z0-9]{4}`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingIssuer = regexp.MustCompile(`[^|]*`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.EncodedUnderlyingIssuerLen = regexp.MustCompile(`\d+`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.EncodedUnderlyingIssuer = regexp.MustCompile(`.*`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingSecurityDesc = regexp.MustCompile(`[^|]*`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.EncodedUnderlyingSecurityDescLen = regexp.MustCompile(`\d+`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.EncodedUnderlyingSecurityDesc = regexp.MustCompile(`.*`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingCPProgram = regexp.MustCompile(`[^|]*`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingCPRegType = regexp.MustCompile(`[^|]*`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingPx = regexp.MustCompile(``)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingDirtyPrice = regexp.MustCompile(``)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingEndPrice = regexp.MustCompile(``)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingStartValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingCurrentValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingEndValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingAllocationPercent = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingSettlementType = regexp.MustCompile(`-?\d+`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingCashAmount = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingCashType = regexp.MustCompile(`[^|]*`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingUnitofMeasure = regexp.MustCompile(`[^|]*`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingTimeUnit = regexp.MustCompile(`[^|]*`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingCapValue = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingSettlMethod = regexp.MustCompile(`[^|]*`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingAdjustedQuantity = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingFXRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.UnderlyingFXRateCalc = regexp.MustCompile(`[^|]`)
    myUndInstrmtCollGrp_Underlyings_RegexValidator.CollAction = regexp.MustCompile(`-?\d+`)
}



func (m *UndInstrmtCollGrp_Underlyings) HasRequiredFields() bool {
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
    
    
    
    
    
    return valid
}




