package BidCompReqGrp_BidComponents

import "bufio"
import "strconv"
import "fmt"

// Autogenerated at , do not edit




func (m *BidCompReqGrp_BidComponents) MarshalFIX(res *bufio.Writer) error {
    var err error
    
    if m.ListID.HasValue() {
        (*res).WriteString("66=")
        val, err := m.ListID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.Side.HasValue() {
        (*res).WriteString("54=")
        val, err := m.Side.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.TradingSessionID.HasValue() {
        (*res).WriteString("336=")
        val, err := m.TradingSessionID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.TradingSessionSubID.HasValue() {
        (*res).WriteString("625=")
        val, err := m.TradingSessionSubID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.NetGrossInd.HasValue() {
        (*res).WriteString("430=")
        val, err := m.NetGrossInd.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SettlType.HasValue() {
        (*res).WriteString("63=")
        val, err := m.SettlType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.SettlDate.HasValue() {
        (*res).WriteString("64=")
        val, err := m.SettlDate.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.Account.HasValue() {
        (*res).WriteString("1=")
        val, err := m.Account.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.AcctIDSource.HasValue() {
        (*res).WriteString("660=")
        val, err := m.AcctIDSource.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    return err
}

func (m *BidCompReqGrp_BidComponents) UnmarshalFIX(input io.Reader) error {
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
                err = errors.New(fmt.Sprintf("Field unused by message BidCompReqGrp_BidComponents (%s)", string(field)))
            }
        }
        if err != nil {
            break
        }
    }
    return err
}


func (m *BidCompReqGrp_BidComponents) UnmarshalFieldFIX(id int, value []byte) (bool, error) {
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
        
        case 66:
            if !m.ListID.HasValue() {
                used = true
                err = m.ListID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 54:
            if !m.Side.HasValue() {
                used = true
                err = m.Side.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 336:
            if !m.TradingSessionID.HasValue() {
                used = true
                err = m.TradingSessionID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 625:
            if !m.TradingSessionSubID.HasValue() {
                used = true
                err = m.TradingSessionSubID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 430:
            if !m.NetGrossInd.HasValue() {
                used = true
                err = m.NetGrossInd.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 63:
            if !m.SettlType.HasValue() {
                used = true
                err = m.SettlType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 64:
            if !m.SettlDate.HasValue() {
                used = true
                err = m.SettlDate.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 1:
            if !m.Account.HasValue() {
                used = true
                err = m.Account.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 660:
            if !m.AcctIDSource.HasValue() {
                used = true
                err = m.AcctIDSource.UnmarshalFIX(value)
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


