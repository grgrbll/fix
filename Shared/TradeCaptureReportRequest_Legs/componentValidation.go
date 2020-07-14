package TradeCaptureReportRequest_Legs

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import LegSecAltIDGrp_LegSecurityAltID "grgrbll/fix/Shared/LegSecAltIDGrp_LegSecurityAltID"


type tradeCaptureReportRequest_Legs_RegexValidator struct {
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
    _controlBlock fixMessageControlBlock
}


var myTradeCaptureReportRequest_Legs_RegexValidator tradeCaptureReportRequest_Legs_RegexValidator

func init() {
    myTradeCaptureReportRequest_Legs_RegexValidator
    myTradeCaptureReportRequest_Legs_RegexValidator.LegSymbol = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegSymbolSfx = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegSecurityID = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegSecurityIDSource = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegProduct = regexp.MustCompile(`-?\d+`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegCFICode = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegSecurityType = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegSecuritySubType = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegMaturityMonthYear = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegMaturityDate = regexp.MustCompile(`[0-9]{8}`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegCouponPaymentDate = regexp.MustCompile(`[0-9]{8}`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegIssueDate = regexp.MustCompile(`[0-9]{8}`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegRepoCollateralSecurityType = regexp.MustCompile(`-?\d+`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegRepurchaseTerm = regexp.MustCompile(`-?\d+`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegRepurchaseRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegFactor = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegCreditRating = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegInstrRegistry = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegCountryOfIssue = regexp.MustCompile(`[A-Z]{2}`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegStateOrProvinceOfIssue = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegLocaleOfIssue = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegRedemptionDate = regexp.MustCompile(`[0-9]{8}`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegStrikePrice = regexp.MustCompile(``)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegStrikeCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegOptAttribute = regexp.MustCompile(`[^|]`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegContractMultiplier = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegCouponRate = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegSecurityExchange = regexp.MustCompile(`[A-Z0-9]{4}`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegIssuer = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportRequest_Legs_RegexValidator.EncodedLegIssuerLen = regexp.MustCompile(`\d+`)
    myTradeCaptureReportRequest_Legs_RegexValidator.EncodedLegIssuer = regexp.MustCompile(`.*`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegSecurityDesc = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportRequest_Legs_RegexValidator.EncodedLegSecurityDescLen = regexp.MustCompile(`\d+`)
    myTradeCaptureReportRequest_Legs_RegexValidator.EncodedLegSecurityDesc = regexp.MustCompile(`.*`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegRatioQty = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegSide = regexp.MustCompile(`[^|]`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegCurrency = regexp.MustCompile(`[A-Z]{3}`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegPool = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegDatedDate = regexp.MustCompile(`[0-9]{8}`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegContractSettlMonth = regexp.MustCompile(`[0-9]{6}([0-9]{2}|w[0-9])?`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegInterestAccrualDate = regexp.MustCompile(`[0-9]{8}`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegUnitofMeasure = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegTimeUnit = regexp.MustCompile(`[^|]*`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegOptionRatio = regexp.MustCompile(`-?\d*(\.\d*)?`)
    myTradeCaptureReportRequest_Legs_RegexValidator.LegPrice = regexp.MustCompile(``)
}



func (m *TradeCaptureReportRequest_Legs) HasRequiredFields() bool {
    valid := true
    
    

    
    
    for _, g := range(m.LegSecurityAltID) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    
    
    return valid
}




