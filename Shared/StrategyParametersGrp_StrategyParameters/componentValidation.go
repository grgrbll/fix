package StrategyParametersGrp_StrategyParameters

import "bufio"
import "regexp"

// Autogenerated at , do not edit




type strategyParametersGrp_StrategyParameters_RegexValidator struct {
    StrategyParameterName *(regexp.Regexp)
    StrategyParameterType *(regexp.Regexp)
    StrategyParameterValue *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myStrategyParametersGrp_StrategyParameters_RegexValidator strategyParametersGrp_StrategyParameters_RegexValidator

func init() {
    myStrategyParametersGrp_StrategyParameters_RegexValidator
    myStrategyParametersGrp_StrategyParameters_RegexValidator.StrategyParameterName = regexp.MustCompile(`[^|]*`)
    myStrategyParametersGrp_StrategyParameters_RegexValidator.StrategyParameterType = regexp.MustCompile(`-?\d+`)
    myStrategyParametersGrp_StrategyParameters_RegexValidator.StrategyParameterValue = regexp.MustCompile(`[^|]*`)
}



func (m *StrategyParametersGrp_StrategyParameters) HasRequiredFields() bool {
    valid := true
    
    

    
    
    
    
    return valid
}




