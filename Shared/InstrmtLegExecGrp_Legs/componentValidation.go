package InstrmtLegExecGrp_Legs

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import LegSecAltIDGrp_LegSecurityAltID "grgrbll/fix/Shared/LegSecAltIDGrp_LegSecurityAltID"

import LegStipulations_LegStipulations "grgrbll/fix/Shared/LegStipulations_LegStipulations"

import NestedParties_NestedPartyIDs "grgrbll/fix/Shared/NestedParties_NestedPartyIDs"


type instrmtLegExecGrp_Legs_RegexValidator struct {
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
    LegPositionEffect *(regexp.Regexp)
    LegCoveredOrUncovered *(regexp.Regexp)
    LegRefID *(regexp.Regexp)
    LegSettlType *(regexp.Regexp)
    LegSettlDate *(regexp.Regexp)
    LegLastPx *(regexp.Regexp)
    LegOrderQty *(regexp.Regexp)
    LegSettlCurrency *(regexp.Regexp)
    LegLastForwardPoints *(regexp.Regexp)
    LegCalculatedCcyLastQty *(regexp.Regexp)
    LegGrossTradeAmt *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myInstrmtLegExecGrp_Legs_RegexValidator instrmtLegExecGrp_Legs_RegexValidator

func init() {
    myInstrmtLegExecGrp_Legs_RegexValidator
    myInstrmtLegExecGrp_Legs_RegexValidator.LegSymbol = regexp.MustCompile(`[^|]*`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegSymbolSfx = regexp.MustCompile(`[^|]*`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegSecurityID = regexp.MustCompile(`[^|]*`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegSecurityIDSource = regexp.MustCompile(`[^|]*`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegProduct = regexp.MustCompile(`-?\d+`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegCFICode = regexp.MustCompile(`[^|]*`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegSecurityType = regexp.MustCompile(`[^|]*`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegSecuritySubType = regexp.MustCompile(`[^|]*`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegMaturityMonthYear = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegMaturityDate = regexp.MustCompile(`[0-9]{8}`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegCouponPaymentDate = regexp.MustCompile(`[0-9]{8}`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegIssueDate = regexp.MustCompile(`[0-9]{8}`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegRepoCollateralSecurityType = regexp.MustCompile(`-?\d+`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegRepurchaseTerm = regexp.MustCompile(`-?\d+`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegRepurchaseRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegFactor = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegCreditRating = regexp.MustCompile(`[^|]*`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegInstrRegistry = regexp.MustCompile(`[^|]*`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegCountryOfIssue = regexp.MustCompile(`[A-Z]{2}`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegStateOrProvinceOfIssue = regexp.MustCompile(`[^|]*`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegLocaleOfIssue = regexp.MustCompile(`[^|]*`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegRedemptionDate = regexp.MustCompile(`[0-9]{8}`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegStrikePrice = regexp.MustCompile(``)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegStrikeCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegOptAttribute = regexp.MustCompile(`[^|]`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegContractMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegCouponRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegSecurityExchange = regexp.MustCompile(`[A-Z0-9]{4}`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegIssuer = regexp.MustCompile(`[^|]*`)
    myInstrmtLegExecGrp_Legs_RegexValidator.EncodedLegIssuerLen = regexp.MustCompile(`\d+`)
    myInstrmtLegExecGrp_Legs_RegexValidator.EncodedLegIssuer = regexp.MustCompile(`.*`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegSecurityDesc = regexp.MustCompile(`[^|]*`)
    myInstrmtLegExecGrp_Legs_RegexValidator.EncodedLegSecurityDescLen = regexp.MustCompile(`\d+`)
    myInstrmtLegExecGrp_Legs_RegexValidator.EncodedLegSecurityDesc = regexp.MustCompile(`.*`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegRatioQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegSide = regexp.MustCompile(`[^|]`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegPool = regexp.MustCompile(`[^|]*`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegDatedDate = regexp.MustCompile(`[0-9]{8}`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegContractSettlMonth = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegInterestAccrualDate = regexp.MustCompile(`[0-9]{8}`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegUnitofMeasure = regexp.MustCompile(`[^|]*`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegTimeUnit = regexp.MustCompile(`[^|]*`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegOptionRatio = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegPrice = regexp.MustCompile(``)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegSwapType = regexp.MustCompile(`-?\d+`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegPositionEffect = regexp.MustCompile(`[^|]`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegCoveredOrUncovered = regexp.MustCompile(`-?\d+`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegRefID = regexp.MustCompile(`[^|]*`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegSettlType = regexp.MustCompile(`[^|]`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegSettlDate = regexp.MustCompile(`[0-9]{8}`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegLastPx = regexp.MustCompile(``)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegOrderQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegSettlCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegLastForwardPoints = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegCalculatedCcyLastQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myInstrmtLegExecGrp_Legs_RegexValidator.LegGrossTradeAmt = regexp.MustCompile(`-?\d*(\.\d*)?`)
}



func (m *InstrmtLegExecGrp_Legs) HasRequiredFields() bool {
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




