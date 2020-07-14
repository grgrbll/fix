package UndInstrmtCollGrp_Underlyings

import "bufio"
import "strconv"
import "fmt"

// Autogenerated at , do not edit



import UndSecAltIDGrp_UnderlyingSecurityAltID "grgrbll/fix/Shared/UndSecAltIDGrp_UnderlyingSecurityAltID"

import UnderlyingStipulations_UnderlyingStips "grgrbll/fix/Shared/UnderlyingStipulations_UnderlyingStips"

import UndlyInstrumentParties_UndlyInstrumentParties "grgrbll/fix/Shared/UndlyInstrumentParties_UndlyInstrumentParties"


func (m *UndInstrmtCollGrp_Underlyings) MarshalFIX(res *bufio.Writer) error {
    var err error
    
    if m.UnderlyingSymbol.HasValue() {
        (*res).WriteString("311=")
        val, err := m.UnderlyingSymbol.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingSymbolSfx.HasValue() {
        (*res).WriteString("312=")
        val, err := m.UnderlyingSymbolSfx.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingSecurityID.HasValue() {
        (*res).WriteString("309=")
        val, err := m.UnderlyingSecurityID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingSecurityIDSource.HasValue() {
        (*res).WriteString("305=")
        val, err := m.UnderlyingSecurityIDSource.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if len(m.UnderlyingSecurityAltID) > 0 {
    
    (*res).WriteString("457=")
    (*res).WriteString(strconv.Itoa(len(m.UnderlyingSecurityAltID)))
    (*res).WriteString("\x01")
    for _, g := range m.UnderlyingSecurityAltID {
        err = g.MarshalFIX(res)
    }
    }
    
    if m.UnderlyingProduct.HasValue() {
        (*res).WriteString("462=")
        val, err := m.UnderlyingProduct.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingCFICode.HasValue() {
        (*res).WriteString("463=")
        val, err := m.UnderlyingCFICode.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingSecurityType.HasValue() {
        (*res).WriteString("310=")
        val, err := m.UnderlyingSecurityType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingSecuritySubType.HasValue() {
        (*res).WriteString("763=")
        val, err := m.UnderlyingSecuritySubType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingMaturityMonthYear.HasValue() {
        (*res).WriteString("313=")
        val, err := m.UnderlyingMaturityMonthYear.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingMaturityDate.HasValue() {
        (*res).WriteString("542=")
        val, err := m.UnderlyingMaturityDate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingCouponPaymentDate.HasValue() {
        (*res).WriteString("241=")
        val, err := m.UnderlyingCouponPaymentDate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingIssueDate.HasValue() {
        (*res).WriteString("242=")
        val, err := m.UnderlyingIssueDate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingRepoCollateralSecurityType.HasValue() {
        (*res).WriteString("243=")
        val, err := m.UnderlyingRepoCollateralSecurityType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingRepurchaseTerm.HasValue() {
        (*res).WriteString("244=")
        val, err := m.UnderlyingRepurchaseTerm.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingRepurchaseRate.HasValue() {
        (*res).WriteString("245=")
        val, err := m.UnderlyingRepurchaseRate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingFactor.HasValue() {
        (*res).WriteString("246=")
        val, err := m.UnderlyingFactor.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingCreditRating.HasValue() {
        (*res).WriteString("256=")
        val, err := m.UnderlyingCreditRating.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingInstrRegistry.HasValue() {
        (*res).WriteString("595=")
        val, err := m.UnderlyingInstrRegistry.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingCountryOfIssue.HasValue() {
        (*res).WriteString("592=")
        val, err := m.UnderlyingCountryOfIssue.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingStateOrProvinceOfIssue.HasValue() {
        (*res).WriteString("593=")
        val, err := m.UnderlyingStateOrProvinceOfIssue.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingLocaleOfIssue.HasValue() {
        (*res).WriteString("594=")
        val, err := m.UnderlyingLocaleOfIssue.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingRedemptionDate.HasValue() {
        (*res).WriteString("247=")
        val, err := m.UnderlyingRedemptionDate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingStrikePrice.HasValue() {
        (*res).WriteString("316=")
        val, err := m.UnderlyingStrikePrice.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingStrikeCurrency.HasValue() {
        (*res).WriteString("941=")
        val, err := m.UnderlyingStrikeCurrency.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingOptAttribute.HasValue() {
        (*res).WriteString("317=")
        val, err := m.UnderlyingOptAttribute.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingContractMultiplier.HasValue() {
        (*res).WriteString("436=")
        val, err := m.UnderlyingContractMultiplier.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingCouponRate.HasValue() {
        (*res).WriteString("435=")
        val, err := m.UnderlyingCouponRate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingSecurityExchange.HasValue() {
        (*res).WriteString("308=")
        val, err := m.UnderlyingSecurityExchange.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingIssuer.HasValue() {
        (*res).WriteString("306=")
        val, err := m.UnderlyingIssuer.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.EncodedUnderlyingIssuerLen.HasValue() {
        (*res).WriteString("362=")
        val, err := m.EncodedUnderlyingIssuerLen.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.EncodedUnderlyingIssuer.HasValue() {
        (*res).WriteString("363=")
        val, err := m.EncodedUnderlyingIssuer.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingSecurityDesc.HasValue() {
        (*res).WriteString("307=")
        val, err := m.UnderlyingSecurityDesc.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.EncodedUnderlyingSecurityDescLen.HasValue() {
        (*res).WriteString("364=")
        val, err := m.EncodedUnderlyingSecurityDescLen.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.EncodedUnderlyingSecurityDesc.HasValue() {
        (*res).WriteString("365=")
        val, err := m.EncodedUnderlyingSecurityDesc.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingCPProgram.HasValue() {
        (*res).WriteString("877=")
        val, err := m.UnderlyingCPProgram.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingCPRegType.HasValue() {
        (*res).WriteString("878=")
        val, err := m.UnderlyingCPRegType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingCurrency.HasValue() {
        (*res).WriteString("318=")
        val, err := m.UnderlyingCurrency.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingQty.HasValue() {
        (*res).WriteString("879=")
        val, err := m.UnderlyingQty.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingPx.HasValue() {
        (*res).WriteString("810=")
        val, err := m.UnderlyingPx.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingDirtyPrice.HasValue() {
        (*res).WriteString("882=")
        val, err := m.UnderlyingDirtyPrice.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingEndPrice.HasValue() {
        (*res).WriteString("883=")
        val, err := m.UnderlyingEndPrice.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingStartValue.HasValue() {
        (*res).WriteString("884=")
        val, err := m.UnderlyingStartValue.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingCurrentValue.HasValue() {
        (*res).WriteString("885=")
        val, err := m.UnderlyingCurrentValue.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingEndValue.HasValue() {
        (*res).WriteString("886=")
        val, err := m.UnderlyingEndValue.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if len(m.UnderlyingStips) > 0 {
    
    (*res).WriteString("887=")
    (*res).WriteString(strconv.Itoa(len(m.UnderlyingStips)))
    (*res).WriteString("\x01")
    for _, g := range m.UnderlyingStips {
        err = g.MarshalFIX(res)
    }
    }
    
    if m.UnderlyingAllocationPercent.HasValue() {
        (*res).WriteString("972=")
        val, err := m.UnderlyingAllocationPercent.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingSettlementType.HasValue() {
        (*res).WriteString("975=")
        val, err := m.UnderlyingSettlementType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingCashAmount.HasValue() {
        (*res).WriteString("973=")
        val, err := m.UnderlyingCashAmount.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingCashType.HasValue() {
        (*res).WriteString("974=")
        val, err := m.UnderlyingCashType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingUnitofMeasure.HasValue() {
        (*res).WriteString("998=")
        val, err := m.UnderlyingUnitofMeasure.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingTimeUnit.HasValue() {
        (*res).WriteString("1000=")
        val, err := m.UnderlyingTimeUnit.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingCapValue.HasValue() {
        (*res).WriteString("1038=")
        val, err := m.UnderlyingCapValue.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if len(m.UndlyInstrumentParties) > 0 {
    
    (*res).WriteString("1058=")
    (*res).WriteString(strconv.Itoa(len(m.UndlyInstrumentParties)))
    (*res).WriteString("\x01")
    for _, g := range m.UndlyInstrumentParties {
        err = g.MarshalFIX(res)
    }
    }
    
    if m.UnderlyingSettlMethod.HasValue() {
        (*res).WriteString("1039=")
        val, err := m.UnderlyingSettlMethod.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingAdjustedQuantity.HasValue() {
        (*res).WriteString("1044=")
        val, err := m.UnderlyingAdjustedQuantity.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingFXRate.HasValue() {
        (*res).WriteString("1045=")
        val, err := m.UnderlyingFXRate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.UnderlyingFXRateCalc.HasValue() {
        (*res).WriteString("1046=")
        val, err := m.UnderlyingFXRateCalc.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.CollAction.HasValue() {
        (*res).WriteString("944=")
        val, err := m.CollAction.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    return err
}

func (m *UndInstrmtCollGrp_Underlyings) UnmarshalFIX(input io.Reader) error {
    var err error
    var field []byte
    for field, err = input.ReadSlice([]byte("\x01")); err == nil {
        slices = bytes.Split(field, []byte("="))
        if len(slices) != 2 {
            // TODO handle data fields
            err = errors.New(fmt.Sprintf("Found field without seperator '=' (%s)", string(field))
        } else {
            var used bool
            used, err m.UnmarshalFieldFIX(strconv.Atoi(string(slices[0])), slices[1])
            if !used {
                err = errors.New(fmt.Sprintf("Field unused by message UndInstrmtCollGrp_Underlyings (%s)", string(field)))
            }
        }
        if err != nil {
            break
        }
    }
    return err
}


func (m *UndInstrmtCollGrp_Underlyings) UnmarshalFieldFIX(id int, value []byte) (bool, error) {
    // Check if we are currently populating a group
    used := false
    var err error = nil
    for !used && m._controlBlock.mostRecentRepeatingGroup != nil {
        used = m._controlBlock.mostRecentRepeatingGroup[m._controlBlock.mostRecentRepeatingGroupCounter].PopulateNextFieldById(id, value)
        if !used {
            // This group did not need use the latest field, and has all its mandatory fields, so consider it complete.
            m._controlBlock.mostRecentRepeatingGroupCounter++
            if m._controlBlock.mostRecentRepeatingGroupCounter >= len(m._controlBlock.mostRecentRepeatingGroup)  {
                // we have all the repeated groups we expected.
                m._controlBlock.mostRecentRepeatingGroup = nil
            }
        }
    }
    
    // If one of the groups did not consume the KV, try our own fields 
    if !used {
        switch id {
        
        case 311:
            if !m.UnderlyingSymbol.HasValue() {
                used = true
                err = m.UnderlyingSymbol.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 312:
            if !m.UnderlyingSymbolSfx.HasValue() {
                used = true
                err = m.UnderlyingSymbolSfx.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 309:
            if !m.UnderlyingSecurityID.HasValue() {
                used = true
                err = m.UnderlyingSecurityID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 305:
            if !m.UnderlyingSecurityIDSource.HasValue() {
                used = true
                err = m.UnderlyingSecurityIDSource.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 457:
            // This counter (NoUnderlyingSecurityAltID) indicates the start of repeated group block 
            count, _ := strconv.Atoi(value)
            m.UnderlyingSecurityAltID = make([]UndSecAltIDGrp_UnderlyingSecurityAltID, count)
            m._controlBlock.mostRecentRepeatingGroup = make([]FixMessage, count)
            for i := 0; i < count; i++ {
                m._controlBlock.mostRecentRepeatingGroup[i] = &m.UnderlyingSecurityAltID[i]
            }
            m._controlBlock.mostRecentRepeatingGroupCounter = 0
            used = true
        
        
        
        case 462:
            if !m.UnderlyingProduct.HasValue() {
                used = true
                err = m.UnderlyingProduct.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 463:
            if !m.UnderlyingCFICode.HasValue() {
                used = true
                err = m.UnderlyingCFICode.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 310:
            if !m.UnderlyingSecurityType.HasValue() {
                used = true
                err = m.UnderlyingSecurityType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 763:
            if !m.UnderlyingSecuritySubType.HasValue() {
                used = true
                err = m.UnderlyingSecuritySubType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 313:
            if !m.UnderlyingMaturityMonthYear.HasValue() {
                used = true
                err = m.UnderlyingMaturityMonthYear.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 542:
            if !m.UnderlyingMaturityDate.HasValue() {
                used = true
                err = m.UnderlyingMaturityDate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 241:
            if !m.UnderlyingCouponPaymentDate.HasValue() {
                used = true
                err = m.UnderlyingCouponPaymentDate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 242:
            if !m.UnderlyingIssueDate.HasValue() {
                used = true
                err = m.UnderlyingIssueDate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 243:
            if !m.UnderlyingRepoCollateralSecurityType.HasValue() {
                used = true
                err = m.UnderlyingRepoCollateralSecurityType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 244:
            if !m.UnderlyingRepurchaseTerm.HasValue() {
                used = true
                err = m.UnderlyingRepurchaseTerm.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 245:
            if !m.UnderlyingRepurchaseRate.HasValue() {
                used = true
                err = m.UnderlyingRepurchaseRate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 246:
            if !m.UnderlyingFactor.HasValue() {
                used = true
                err = m.UnderlyingFactor.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 256:
            if !m.UnderlyingCreditRating.HasValue() {
                used = true
                err = m.UnderlyingCreditRating.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 595:
            if !m.UnderlyingInstrRegistry.HasValue() {
                used = true
                err = m.UnderlyingInstrRegistry.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 592:
            if !m.UnderlyingCountryOfIssue.HasValue() {
                used = true
                err = m.UnderlyingCountryOfIssue.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 593:
            if !m.UnderlyingStateOrProvinceOfIssue.HasValue() {
                used = true
                err = m.UnderlyingStateOrProvinceOfIssue.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 594:
            if !m.UnderlyingLocaleOfIssue.HasValue() {
                used = true
                err = m.UnderlyingLocaleOfIssue.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 247:
            if !m.UnderlyingRedemptionDate.HasValue() {
                used = true
                err = m.UnderlyingRedemptionDate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 316:
            if !m.UnderlyingStrikePrice.HasValue() {
                used = true
                err = m.UnderlyingStrikePrice.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 941:
            if !m.UnderlyingStrikeCurrency.HasValue() {
                used = true
                err = m.UnderlyingStrikeCurrency.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 317:
            if !m.UnderlyingOptAttribute.HasValue() {
                used = true
                err = m.UnderlyingOptAttribute.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 436:
            if !m.UnderlyingContractMultiplier.HasValue() {
                used = true
                err = m.UnderlyingContractMultiplier.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 435:
            if !m.UnderlyingCouponRate.HasValue() {
                used = true
                err = m.UnderlyingCouponRate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 308:
            if !m.UnderlyingSecurityExchange.HasValue() {
                used = true
                err = m.UnderlyingSecurityExchange.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 306:
            if !m.UnderlyingIssuer.HasValue() {
                used = true
                err = m.UnderlyingIssuer.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 362:
            if !m.EncodedUnderlyingIssuerLen.HasValue() {
                used = true
                err = m.EncodedUnderlyingIssuerLen.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 363:
            if !m.EncodedUnderlyingIssuer.HasValue() {
                used = true
                err = m.EncodedUnderlyingIssuer.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 307:
            if !m.UnderlyingSecurityDesc.HasValue() {
                used = true
                err = m.UnderlyingSecurityDesc.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 364:
            if !m.EncodedUnderlyingSecurityDescLen.HasValue() {
                used = true
                err = m.EncodedUnderlyingSecurityDescLen.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 365:
            if !m.EncodedUnderlyingSecurityDesc.HasValue() {
                used = true
                err = m.EncodedUnderlyingSecurityDesc.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 877:
            if !m.UnderlyingCPProgram.HasValue() {
                used = true
                err = m.UnderlyingCPProgram.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 878:
            if !m.UnderlyingCPRegType.HasValue() {
                used = true
                err = m.UnderlyingCPRegType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 318:
            if !m.UnderlyingCurrency.HasValue() {
                used = true
                err = m.UnderlyingCurrency.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 879:
            if !m.UnderlyingQty.HasValue() {
                used = true
                err = m.UnderlyingQty.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 810:
            if !m.UnderlyingPx.HasValue() {
                used = true
                err = m.UnderlyingPx.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 882:
            if !m.UnderlyingDirtyPrice.HasValue() {
                used = true
                err = m.UnderlyingDirtyPrice.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 883:
            if !m.UnderlyingEndPrice.HasValue() {
                used = true
                err = m.UnderlyingEndPrice.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 884:
            if !m.UnderlyingStartValue.HasValue() {
                used = true
                err = m.UnderlyingStartValue.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 885:
            if !m.UnderlyingCurrentValue.HasValue() {
                used = true
                err = m.UnderlyingCurrentValue.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 886:
            if !m.UnderlyingEndValue.HasValue() {
                used = true
                err = m.UnderlyingEndValue.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 887:
            // This counter (NoUnderlyingStips) indicates the start of repeated group block 
            count, _ := strconv.Atoi(value)
            m.UnderlyingStips = make([]UnderlyingStipulations_UnderlyingStips, count)
            m._controlBlock.mostRecentRepeatingGroup = make([]FixMessage, count)
            for i := 0; i < count; i++ {
                m._controlBlock.mostRecentRepeatingGroup[i] = &m.UnderlyingStips[i]
            }
            m._controlBlock.mostRecentRepeatingGroupCounter = 0
            used = true
        
        
        
        case 972:
            if !m.UnderlyingAllocationPercent.HasValue() {
                used = true
                err = m.UnderlyingAllocationPercent.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 975:
            if !m.UnderlyingSettlementType.HasValue() {
                used = true
                err = m.UnderlyingSettlementType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 973:
            if !m.UnderlyingCashAmount.HasValue() {
                used = true
                err = m.UnderlyingCashAmount.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 974:
            if !m.UnderlyingCashType.HasValue() {
                used = true
                err = m.UnderlyingCashType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 998:
            if !m.UnderlyingUnitofMeasure.HasValue() {
                used = true
                err = m.UnderlyingUnitofMeasure.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 1000:
            if !m.UnderlyingTimeUnit.HasValue() {
                used = true
                err = m.UnderlyingTimeUnit.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 1038:
            if !m.UnderlyingCapValue.HasValue() {
                used = true
                err = m.UnderlyingCapValue.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 1058:
            // This counter (NoUndlyInstrumentParties) indicates the start of repeated group block 
            count, _ := strconv.Atoi(value)
            m.UndlyInstrumentParties = make([]UndlyInstrumentParties_UndlyInstrumentParties, count)
            m._controlBlock.mostRecentRepeatingGroup = make([]FixMessage, count)
            for i := 0; i < count; i++ {
                m._controlBlock.mostRecentRepeatingGroup[i] = &m.UndlyInstrumentParties[i]
            }
            m._controlBlock.mostRecentRepeatingGroupCounter = 0
            used = true
        
        
        
        case 1039:
            if !m.UnderlyingSettlMethod.HasValue() {
                used = true
                err = m.UnderlyingSettlMethod.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 1044:
            if !m.UnderlyingAdjustedQuantity.HasValue() {
                used = true
                err = m.UnderlyingAdjustedQuantity.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 1045:
            if !m.UnderlyingFXRate.HasValue() {
                used = true
                err = m.UnderlyingFXRate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 1046:
            if !m.UnderlyingFXRateCalc.HasValue() {
                used = true
                err = m.UnderlyingFXRateCalc.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 944:
            if !m.CollAction.HasValue() {
                used = true
                err = m.CollAction.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        default:
            used = false
        }
    }
   
    return used, err
}


