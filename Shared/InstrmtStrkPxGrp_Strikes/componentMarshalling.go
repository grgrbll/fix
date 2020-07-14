package InstrmtStrkPxGrp_Strikes

import "bufio"
import "strconv"
import "fmt"

// Autogenerated at , do not edit



import SecAltIDGrp_SecurityAltID "grgrbll/fix/Shared/SecAltIDGrp_SecurityAltID"

import EvntGrp_Events "grgrbll/fix/Shared/EvntGrp_Events"

import InstrumentParties_InstrumentParties "grgrbll/fix/Shared/InstrumentParties_InstrumentParties"


func (m *InstrmtStrkPxGrp_Strikes) MarshalFIX(res *bufio.Writer) error {
    var err error
    
    if m.Symbol.HasValue() {
        (*res).WriteString("55=")
        val, err := m.Symbol.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SymbolSfx.HasValue() {
        (*res).WriteString("65=")
        val, err := m.SymbolSfx.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SecurityID.HasValue() {
        (*res).WriteString("48=")
        val, err := m.SecurityID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SecurityIDSource.HasValue() {
        (*res).WriteString("22=")
        val, err := m.SecurityIDSource.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if len(m.SecurityAltID) > 0 {
    
    (*res).WriteString("454=")
    (*res).WriteString(strconv.Itoa(len(m.SecurityAltID)))
    (*res).WriteString("\x01")
    for _, g := range m.SecurityAltID {
        err = g.MarshalFIX(res)
    }
    }
    
    if m.Product.HasValue() {
        (*res).WriteString("460=")
        val, err := m.Product.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.CFICode.HasValue() {
        (*res).WriteString("461=")
        val, err := m.CFICode.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SecurityType.HasValue() {
        (*res).WriteString("167=")
        val, err := m.SecurityType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SecuritySubType.HasValue() {
        (*res).WriteString("762=")
        val, err := m.SecuritySubType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.MaturityMonthYear.HasValue() {
        (*res).WriteString("200=")
        val, err := m.MaturityMonthYear.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.MaturityDate.HasValue() {
        (*res).WriteString("541=")
        val, err := m.MaturityDate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.CouponPaymentDate.HasValue() {
        (*res).WriteString("224=")
        val, err := m.CouponPaymentDate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.IssueDate.HasValue() {
        (*res).WriteString("225=")
        val, err := m.IssueDate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.RepoCollateralSecurityType.HasValue() {
        (*res).WriteString("239=")
        val, err := m.RepoCollateralSecurityType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.RepurchaseTerm.HasValue() {
        (*res).WriteString("226=")
        val, err := m.RepurchaseTerm.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.RepurchaseRate.HasValue() {
        (*res).WriteString("227=")
        val, err := m.RepurchaseRate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.Factor.HasValue() {
        (*res).WriteString("228=")
        val, err := m.Factor.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.CreditRating.HasValue() {
        (*res).WriteString("255=")
        val, err := m.CreditRating.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.InstrRegistry.HasValue() {
        (*res).WriteString("543=")
        val, err := m.InstrRegistry.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.CountryOfIssue.HasValue() {
        (*res).WriteString("470=")
        val, err := m.CountryOfIssue.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.StateOrProvinceOfIssue.HasValue() {
        (*res).WriteString("471=")
        val, err := m.StateOrProvinceOfIssue.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LocaleOfIssue.HasValue() {
        (*res).WriteString("472=")
        val, err := m.LocaleOfIssue.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.RedemptionDate.HasValue() {
        (*res).WriteString("240=")
        val, err := m.RedemptionDate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.StrikePrice.HasValue() {
        (*res).WriteString("202=")
        val, err := m.StrikePrice.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.StrikeCurrency.HasValue() {
        (*res).WriteString("947=")
        val, err := m.StrikeCurrency.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.OptAttribute.HasValue() {
        (*res).WriteString("206=")
        val, err := m.OptAttribute.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.ContractMultiplier.HasValue() {
        (*res).WriteString("231=")
        val, err := m.ContractMultiplier.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.CouponRate.HasValue() {
        (*res).WriteString("223=")
        val, err := m.CouponRate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SecurityExchange.HasValue() {
        (*res).WriteString("207=")
        val, err := m.SecurityExchange.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.Issuer.HasValue() {
        (*res).WriteString("106=")
        val, err := m.Issuer.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.EncodedIssuerLen.HasValue() {
        (*res).WriteString("348=")
        val, err := m.EncodedIssuerLen.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.EncodedIssuer.HasValue() {
        (*res).WriteString("349=")
        val, err := m.EncodedIssuer.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SecurityDesc.HasValue() {
        (*res).WriteString("107=")
        val, err := m.SecurityDesc.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.EncodedSecurityDescLen.HasValue() {
        (*res).WriteString("350=")
        val, err := m.EncodedSecurityDescLen.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.EncodedSecurityDesc.HasValue() {
        (*res).WriteString("351=")
        val, err := m.EncodedSecurityDesc.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.Pool.HasValue() {
        (*res).WriteString("691=")
        val, err := m.Pool.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.ContractSettlMonth.HasValue() {
        (*res).WriteString("667=")
        val, err := m.ContractSettlMonth.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.CPProgram.HasValue() {
        (*res).WriteString("875=")
        val, err := m.CPProgram.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.CPRegType.HasValue() {
        (*res).WriteString("876=")
        val, err := m.CPRegType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if len(m.Events) > 0 {
    
    (*res).WriteString("864=")
    (*res).WriteString(strconv.Itoa(len(m.Events)))
    (*res).WriteString("\x01")
    for _, g := range m.Events {
        err = g.MarshalFIX(res)
    }
    }
    
    if m.DatedDate.HasValue() {
        (*res).WriteString("873=")
        val, err := m.DatedDate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.InterestAccrualDate.HasValue() {
        (*res).WriteString("874=")
        val, err := m.InterestAccrualDate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SecurityStatus.HasValue() {
        (*res).WriteString("965=")
        val, err := m.SecurityStatus.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SettleOnOpenFlag.HasValue() {
        (*res).WriteString("966=")
        val, err := m.SettleOnOpenFlag.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.InstrmtAssignmentMethod.HasValue() {
        (*res).WriteString("1049=")
        val, err := m.InstrmtAssignmentMethod.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.StrikeMultiplier.HasValue() {
        (*res).WriteString("967=")
        val, err := m.StrikeMultiplier.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.StrikeValue.HasValue() {
        (*res).WriteString("968=")
        val, err := m.StrikeValue.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.MinPriceIncrement.HasValue() {
        (*res).WriteString("969=")
        val, err := m.MinPriceIncrement.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.PositionLimit.HasValue() {
        (*res).WriteString("970=")
        val, err := m.PositionLimit.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.NTPositionLimit.HasValue() {
        (*res).WriteString("971=")
        val, err := m.NTPositionLimit.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if len(m.InstrumentParties) > 0 {
    
    (*res).WriteString("1018=")
    (*res).WriteString(strconv.Itoa(len(m.InstrumentParties)))
    (*res).WriteString("\x01")
    for _, g := range m.InstrumentParties {
        err = g.MarshalFIX(res)
    }
    }
    
    if m.UnitofMeasure.HasValue() {
        (*res).WriteString("996=")
        val, err := m.UnitofMeasure.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.TimeUnit.HasValue() {
        (*res).WriteString("997=")
        val, err := m.TimeUnit.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.MaturityTime.HasValue() {
        (*res).WriteString("1079=")
        val, err := m.MaturityTime.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    return err
}

func (m *InstrmtStrkPxGrp_Strikes) UnmarshalFIX(input io.Reader) error {
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
                err = errors.New(fmt.Sprintf("Field unused by message InstrmtStrkPxGrp_Strikes (%s)", string(field)))
            }
        }
        if err != nil {
            break
        }
    }
    return err
}


func (m *InstrmtStrkPxGrp_Strikes) UnmarshalFieldFIX(id int, value []byte) (bool, error) {
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
        
        case 55:
            if !m.Symbol.HasValue() {
                used = true
                err = m.Symbol.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 65:
            if !m.SymbolSfx.HasValue() {
                used = true
                err = m.SymbolSfx.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 48:
            if !m.SecurityID.HasValue() {
                used = true
                err = m.SecurityID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 22:
            if !m.SecurityIDSource.HasValue() {
                used = true
                err = m.SecurityIDSource.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 454:
            // This counter (NoSecurityAltID) indicates the start of repeated group block 
            count, _ := strconv.Atoi(value)
            m.SecurityAltID = make([]SecAltIDGrp_SecurityAltID, count)
            m._controlBlock.mostRecentRepeatingGroup = make([]FixMessage, count)
            for i := 0; i < count; i++ {
                m._controlBlock.mostRecentRepeatingGroup[i] = &m.SecurityAltID[i]
            }
            m._controlBlock.mostRecentRepeatingGroupCounter = 0
            used = true
        
        
        
        case 460:
            if !m.Product.HasValue() {
                used = true
                err = m.Product.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 461:
            if !m.CFICode.HasValue() {
                used = true
                err = m.CFICode.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 167:
            if !m.SecurityType.HasValue() {
                used = true
                err = m.SecurityType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 762:
            if !m.SecuritySubType.HasValue() {
                used = true
                err = m.SecuritySubType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 200:
            if !m.MaturityMonthYear.HasValue() {
                used = true
                err = m.MaturityMonthYear.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 541:
            if !m.MaturityDate.HasValue() {
                used = true
                err = m.MaturityDate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 224:
            if !m.CouponPaymentDate.HasValue() {
                used = true
                err = m.CouponPaymentDate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 225:
            if !m.IssueDate.HasValue() {
                used = true
                err = m.IssueDate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 239:
            if !m.RepoCollateralSecurityType.HasValue() {
                used = true
                err = m.RepoCollateralSecurityType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 226:
            if !m.RepurchaseTerm.HasValue() {
                used = true
                err = m.RepurchaseTerm.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 227:
            if !m.RepurchaseRate.HasValue() {
                used = true
                err = m.RepurchaseRate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 228:
            if !m.Factor.HasValue() {
                used = true
                err = m.Factor.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 255:
            if !m.CreditRating.HasValue() {
                used = true
                err = m.CreditRating.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 543:
            if !m.InstrRegistry.HasValue() {
                used = true
                err = m.InstrRegistry.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 470:
            if !m.CountryOfIssue.HasValue() {
                used = true
                err = m.CountryOfIssue.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 471:
            if !m.StateOrProvinceOfIssue.HasValue() {
                used = true
                err = m.StateOrProvinceOfIssue.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 472:
            if !m.LocaleOfIssue.HasValue() {
                used = true
                err = m.LocaleOfIssue.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 240:
            if !m.RedemptionDate.HasValue() {
                used = true
                err = m.RedemptionDate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 202:
            if !m.StrikePrice.HasValue() {
                used = true
                err = m.StrikePrice.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 947:
            if !m.StrikeCurrency.HasValue() {
                used = true
                err = m.StrikeCurrency.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 206:
            if !m.OptAttribute.HasValue() {
                used = true
                err = m.OptAttribute.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 231:
            if !m.ContractMultiplier.HasValue() {
                used = true
                err = m.ContractMultiplier.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 223:
            if !m.CouponRate.HasValue() {
                used = true
                err = m.CouponRate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 207:
            if !m.SecurityExchange.HasValue() {
                used = true
                err = m.SecurityExchange.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 106:
            if !m.Issuer.HasValue() {
                used = true
                err = m.Issuer.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 348:
            if !m.EncodedIssuerLen.HasValue() {
                used = true
                err = m.EncodedIssuerLen.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 349:
            if !m.EncodedIssuer.HasValue() {
                used = true
                err = m.EncodedIssuer.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 107:
            if !m.SecurityDesc.HasValue() {
                used = true
                err = m.SecurityDesc.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 350:
            if !m.EncodedSecurityDescLen.HasValue() {
                used = true
                err = m.EncodedSecurityDescLen.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 351:
            if !m.EncodedSecurityDesc.HasValue() {
                used = true
                err = m.EncodedSecurityDesc.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 691:
            if !m.Pool.HasValue() {
                used = true
                err = m.Pool.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 667:
            if !m.ContractSettlMonth.HasValue() {
                used = true
                err = m.ContractSettlMonth.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 875:
            if !m.CPProgram.HasValue() {
                used = true
                err = m.CPProgram.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 876:
            if !m.CPRegType.HasValue() {
                used = true
                err = m.CPRegType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 864:
            // This counter (NoEvents) indicates the start of repeated group block 
            count, _ := strconv.Atoi(value)
            m.Events = make([]EvntGrp_Events, count)
            m._controlBlock.mostRecentRepeatingGroup = make([]FixMessage, count)
            for i := 0; i < count; i++ {
                m._controlBlock.mostRecentRepeatingGroup[i] = &m.Events[i]
            }
            m._controlBlock.mostRecentRepeatingGroupCounter = 0
            used = true
        
        
        
        case 873:
            if !m.DatedDate.HasValue() {
                used = true
                err = m.DatedDate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 874:
            if !m.InterestAccrualDate.HasValue() {
                used = true
                err = m.InterestAccrualDate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 965:
            if !m.SecurityStatus.HasValue() {
                used = true
                err = m.SecurityStatus.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 966:
            if !m.SettleOnOpenFlag.HasValue() {
                used = true
                err = m.SettleOnOpenFlag.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 1049:
            if !m.InstrmtAssignmentMethod.HasValue() {
                used = true
                err = m.InstrmtAssignmentMethod.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 967:
            if !m.StrikeMultiplier.HasValue() {
                used = true
                err = m.StrikeMultiplier.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 968:
            if !m.StrikeValue.HasValue() {
                used = true
                err = m.StrikeValue.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 969:
            if !m.MinPriceIncrement.HasValue() {
                used = true
                err = m.MinPriceIncrement.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 970:
            if !m.PositionLimit.HasValue() {
                used = true
                err = m.PositionLimit.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 971:
            if !m.NTPositionLimit.HasValue() {
                used = true
                err = m.NTPositionLimit.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 1018:
            // This counter (NoInstrumentParties) indicates the start of repeated group block 
            count, _ := strconv.Atoi(value)
            m.InstrumentParties = make([]InstrumentParties_InstrumentParties, count)
            m._controlBlock.mostRecentRepeatingGroup = make([]FixMessage, count)
            for i := 0; i < count; i++ {
                m._controlBlock.mostRecentRepeatingGroup[i] = &m.InstrumentParties[i]
            }
            m._controlBlock.mostRecentRepeatingGroupCounter = 0
            used = true
        
        
        
        case 996:
            if !m.UnitofMeasure.HasValue() {
                used = true
                err = m.UnitofMeasure.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 997:
            if !m.TimeUnit.HasValue() {
                used = true
                err = m.TimeUnit.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 1079:
            if !m.MaturityTime.HasValue() {
                used = true
                err = m.MaturityTime.UnmarshalFIX(value)
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


