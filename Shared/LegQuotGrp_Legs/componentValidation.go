package LegQuotGrp_Legs

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import LegSecAltIDGrp_LegSecurityAltID "grgrbll/fix/Shared/LegSecAltIDGrp_LegSecurityAltID"

import LegStipulations_LegStipulations "grgrbll/fix/Shared/LegStipulations_LegStipulations"

import NestedParties_NestedPartyIDs "grgrbll/fix/Shared/NestedParties_NestedPartyIDs"


type legQuotGrp_Legs_RegexValidator struct {
    LegSymbol *(regexp.Regexp)
    LegSymbolSfx *(regexp.Regexp)
    LegSecurityID *(regexp.Regexp)
    LegSecurityIDSource *(regexp.Regexp)
    LegProduct *(regexp.Regexp)
    LegCFICode *(regexp.Regexp)
    LegSecurityType *(regexp.Regexp)
    LegSecuritySubType *(regexp.Regexp)
    LegMaturityMonthYear *(regexp.Regexp)
    LegMaturityDate *(regexp.Regexp)
    LegCouponPaymentDate *(regexp.Regexp)
    LegIssueDate *(regexp.Regexp)
    LegRepoCollateralSecurityType *(regexp.Regexp)
    LegRepurchaseTerm *(regexp.Regexp)
    LegRepurchaseRate *(regexp.Regexp)
    LegFactor *(regexp.Regexp)
    LegCreditRating *(regexp.Regexp)
    LegInstrRegistry *(regexp.Regexp)
    LegCountryOfIssue *(regexp.Regexp)
    LegStateOrProvinceOfIssue *(regexp.Regexp)
    LegLocaleOfIssue *(regexp.Regexp)
    LegRedemptionDate *(regexp.Regexp)
    LegStrikePrice *(regexp.Regexp)
    LegStrikeCurrency *(regexp.Regexp)
    LegOptAttribute *(regexp.Regexp)
    LegContractMultiplier *(regexp.Regexp)
    LegCouponRate *(regexp.Regexp)
    LegSecurityExchange *(regexp.Regexp)
    LegIssuer *(regexp.Regexp)
    EncodedLegIssuerLen *(regexp.Regexp)
    EncodedLegIssuer *(regexp.Regexp)
    LegSecurityDesc *(regexp.Regexp)
    EncodedLegSecurityDescLen *(regexp.Regexp)
    EncodedLegSecurityDesc *(regexp.Regexp)
    LegRatioQty *(regexp.Regexp)
    LegSide *(regexp.Regexp)
    LegCurrency *(regexp.Regexp)
    LegPool *(regexp.Regexp)
    LegDatedDate *(regexp.Regexp)
    LegContractSettlMonth *(regexp.Regexp)
    LegInterestAccrualDate *(regexp.Regexp)
    LegUnitofMeasure *(regexp.Regexp)
    LegTimeUnit *(regexp.Regexp)
    LegOptionRatio *(regexp.Regexp)
    LegPrice *(regexp.Regexp)
    LegQty *(regexp.Regexp)
    LegSwapType *(regexp.Regexp)
    LegSettlType *(regexp.Regexp)
    LegSettlDate *(regexp.Regexp)
    LegPriceType *(regexp.Regexp)
    LegBidPx *(regexp.Regexp)
    LegOfferPx *(regexp.Regexp)
    LegBenchmarkCurveCurrency *(regexp.Regexp)
    LegBenchmarkCurveName *(regexp.Regexp)
    LegBenchmarkCurvePoint *(regexp.Regexp)
    LegBenchmarkPrice *(regexp.Regexp)
    LegBenchmarkPriceType *(regexp.Regexp)
    LegOrderQty *(regexp.Regexp)
    LegRefID *(regexp.Regexp)
    LegBidForwardPoints *(regexp.Regexp)
    LegOfferForwardPoints *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myLegQuotGrp_Legs_RegexValidator legQuotGrp_Legs_RegexValidator

func init() {
    myLegQuotGrp_Legs_RegexValidator
    myLegQuotGrp_Legs_RegexValidator.LegSymbol = regexp.MustCompile(`[^|]*`)
    myLegQuotGrp_Legs_RegexValidator.LegSymbolSfx = regexp.MustCompile(`[^|]*`)
    myLegQuotGrp_Legs_RegexValidator.LegSecurityID = regexp.MustCompile(`[^|]*`)
    myLegQuotGrp_Legs_RegexValidator.LegSecurityIDSource = regexp.MustCompile(`[^|]*`)
    myLegQuotGrp_Legs_RegexValidator.LegProduct = regexp.MustCompile(`-?\d+`)
    myLegQuotGrp_Legs_RegexValidator.LegCFICode = regexp.MustCompile(`[^|]*`)
    myLegQuotGrp_Legs_RegexValidator.LegSecurityType = regexp.MustCompile(`[^|]*`)
    myLegQuotGrp_Legs_RegexValidator.LegSecuritySubType = regexp.MustCompile(`[^|]*`)
    myLegQuotGrp_Legs_RegexValidator.LegMaturityMonthYear = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myLegQuotGrp_Legs_RegexValidator.LegMaturityDate = regexp.MustCompile(`[0-9]{8}`)
    myLegQuotGrp_Legs_RegexValidator.LegCouponPaymentDate = regexp.MustCompile(`[0-9]{8}`)
    myLegQuotGrp_Legs_RegexValidator.LegIssueDate = regexp.MustCompile(`[0-9]{8}`)
    myLegQuotGrp_Legs_RegexValidator.LegRepoCollateralSecurityType = regexp.MustCompile(`-?\d+`)
    myLegQuotGrp_Legs_RegexValidator.LegRepurchaseTerm = regexp.MustCompile(`-?\d+`)
    myLegQuotGrp_Legs_RegexValidator.LegRepurchaseRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myLegQuotGrp_Legs_RegexValidator.LegFactor = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myLegQuotGrp_Legs_RegexValidator.LegCreditRating = regexp.MustCompile(`[^|]*`)
    myLegQuotGrp_Legs_RegexValidator.LegInstrRegistry = regexp.MustCompile(`[^|]*`)
    myLegQuotGrp_Legs_RegexValidator.LegCountryOfIssue = regexp.MustCompile(`[A-Z]{2}`)
    myLegQuotGrp_Legs_RegexValidator.LegStateOrProvinceOfIssue = regexp.MustCompile(`[^|]*`)
    myLegQuotGrp_Legs_RegexValidator.LegLocaleOfIssue = regexp.MustCompile(`[^|]*`)
    myLegQuotGrp_Legs_RegexValidator.LegRedemptionDate = regexp.MustCompile(`[0-9]{8}`)
    myLegQuotGrp_Legs_RegexValidator.LegStrikePrice = regexp.MustCompile(``)
    myLegQuotGrp_Legs_RegexValidator.LegStrikeCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myLegQuotGrp_Legs_RegexValidator.LegOptAttribute = regexp.MustCompile(`[^|]`)
    myLegQuotGrp_Legs_RegexValidator.LegContractMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myLegQuotGrp_Legs_RegexValidator.LegCouponRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myLegQuotGrp_Legs_RegexValidator.LegSecurityExchange = regexp.MustCompile(`[A-Z0-9]{4}`)
    myLegQuotGrp_Legs_RegexValidator.LegIssuer = regexp.MustCompile(`[^|]*`)
    myLegQuotGrp_Legs_RegexValidator.EncodedLegIssuerLen = regexp.MustCompile(`\d+`)
    myLegQuotGrp_Legs_RegexValidator.EncodedLegIssuer = regexp.MustCompile(`.*`)
    myLegQuotGrp_Legs_RegexValidator.LegSecurityDesc = regexp.MustCompile(`[^|]*`)
    myLegQuotGrp_Legs_RegexValidator.EncodedLegSecurityDescLen = regexp.MustCompile(`\d+`)
    myLegQuotGrp_Legs_RegexValidator.EncodedLegSecurityDesc = regexp.MustCompile(`.*`)
    myLegQuotGrp_Legs_RegexValidator.LegRatioQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myLegQuotGrp_Legs_RegexValidator.LegSide = regexp.MustCompile(`[^|]`)
    myLegQuotGrp_Legs_RegexValidator.LegCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myLegQuotGrp_Legs_RegexValidator.LegPool = regexp.MustCompile(`[^|]*`)
    myLegQuotGrp_Legs_RegexValidator.LegDatedDate = regexp.MustCompile(`[0-9]{8}`)
    myLegQuotGrp_Legs_RegexValidator.LegContractSettlMonth = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myLegQuotGrp_Legs_RegexValidator.LegInterestAccrualDate = regexp.MustCompile(`[0-9]{8}`)
    myLegQuotGrp_Legs_RegexValidator.LegUnitofMeasure = regexp.MustCompile(`[^|]*`)
    myLegQuotGrp_Legs_RegexValidator.LegTimeUnit = regexp.MustCompile(`[^|]*`)
    myLegQuotGrp_Legs_RegexValidator.LegOptionRatio = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myLegQuotGrp_Legs_RegexValidator.LegPrice = regexp.MustCompile(``)
    myLegQuotGrp_Legs_RegexValidator.LegQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myLegQuotGrp_Legs_RegexValidator.LegSwapType = regexp.MustCompile(`-?\d+`)
    myLegQuotGrp_Legs_RegexValidator.LegSettlType = regexp.MustCompile(`[^|]`)
    myLegQuotGrp_Legs_RegexValidator.LegSettlDate = regexp.MustCompile(`[0-9]{8}`)
    myLegQuotGrp_Legs_RegexValidator.LegPriceType = regexp.MustCompile(`-?\d+`)
    myLegQuotGrp_Legs_RegexValidator.LegBidPx = regexp.MustCompile(``)
    myLegQuotGrp_Legs_RegexValidator.LegOfferPx = regexp.MustCompile(``)
    myLegQuotGrp_Legs_RegexValidator.LegBenchmarkCurveCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myLegQuotGrp_Legs_RegexValidator.LegBenchmarkCurveName = regexp.MustCompile(`[^|]*`)
    myLegQuotGrp_Legs_RegexValidator.LegBenchmarkCurvePoint = regexp.MustCompile(`[^|]*`)
    myLegQuotGrp_Legs_RegexValidator.LegBenchmarkPrice = regexp.MustCompile(``)
    myLegQuotGrp_Legs_RegexValidator.LegBenchmarkPriceType = regexp.MustCompile(`-?\d+`)
    myLegQuotGrp_Legs_RegexValidator.LegOrderQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myLegQuotGrp_Legs_RegexValidator.LegRefID = regexp.MustCompile(`[^|]*`)
    myLegQuotGrp_Legs_RegexValidator.LegBidForwardPoints = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myLegQuotGrp_Legs_RegexValidator.LegOfferForwardPoints = regexp.MustCompile(`-?\d*(\.\d*)?`)
}



func (m *LegQuotGrp_Legs) HasRequiredFields() bool {
    valid := true
    
    

    
    
    for _, g := range(m.LegSecurityAltID) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.LegStipulations) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    for _, g := range(m.NestedPartyIDs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    
    
    return valid
}



