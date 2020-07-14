package PreAllocMlegGrp_Allocs

import "bufio"
import "strconv"
import "fmt"

// Autogenerated at , do not edit



import NestedParties3_Nested3PartyIDs "grgrbll/fix/Shared/NestedParties3_Nested3PartyIDs"


func (m *PreAllocMlegGrp_Allocs) MarshalFIX(res *bufio.Writer) error {
    var err error
    
    if m.AllocAccount.HasValue() {
        (*res).WriteString("79=")
        val, err := m.AllocAccount.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.AllocAcctIDSource.HasValue() {
        (*res).WriteString("661=")
        val, err := m.AllocAcctIDSource.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.AllocSettlCurrency.HasValue() {
        (*res).WriteString("736=")
        val, err := m.AllocSettlCurrency.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.IndividualAllocID.HasValue() {
        (*res).WriteString("467=")
        val, err := m.IndividualAllocID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if len(m.Nested3PartyIDs) > 0 {
    
    (*res).WriteString("948=")
    (*res).WriteString(strconv.Itoa(len(m.Nested3PartyIDs)))
    (*res).WriteString("\x01")
    for _, g := range m.Nested3PartyIDs {
        err = g.MarshalFIX(res)
    }
    }
    
    if m.AllocQty.HasValue() {
        (*res).WriteString("80=")
        val, err := m.AllocQty.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    return err
}

func (m *PreAllocMlegGrp_Allocs) UnmarshalFIX(input io.Reader) error {
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
                err = errors.New(fmt.Sprintf("Field unused by message PreAllocMlegGrp_Allocs (%s)", string(field)))
            }
        }
        if err != nil {
            break
        }
    }
    return err
}


func (m *PreAllocMlegGrp_Allocs) UnmarshalFieldFIX(id int, value []byte) (bool, error) {
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
        
        case 79:
            if !m.AllocAccount.HasValue() {
                used = true
                err = m.AllocAccount.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 661:
            if !m.AllocAcctIDSource.HasValue() {
                used = true
                err = m.AllocAcctIDSource.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 736:
            if !m.AllocSettlCurrency.HasValue() {
                used = true
                err = m.AllocSettlCurrency.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 467:
            if !m.IndividualAllocID.HasValue() {
                used = true
                err = m.IndividualAllocID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 948:
            // This counter (NoNested3PartyIDs) indicates the start of repeated group block 
            count, _ := strconv.Atoi(value)
            m.Nested3PartyIDs = make([]NestedParties3_Nested3PartyIDs, count)
            m._controlBlock.mostRecentRepeatingGroup = make([]FixMessage, count)
            for i := 0; i < count; i++ {
                m._controlBlock.mostRecentRepeatingGroup[i] = &m.Nested3PartyIDs[i]
            }
            m._controlBlock.mostRecentRepeatingGroupCounter = 0
            used = true
        
        
        
        case 80:
            if !m.AllocQty.HasValue() {
                used = true
                err = m.AllocQty.UnmarshalFIX(value)
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

