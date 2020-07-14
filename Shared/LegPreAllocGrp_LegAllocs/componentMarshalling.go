package LegPreAllocGrp_LegAllocs

import "bufio"
import "strconv"
import "fmt"

// Autogenerated at , do not edit



import NestedParties2_Nested2PartyIDs "grgrbll/fix/Shared/NestedParties2_Nested2PartyIDs"


func (m *LegPreAllocGrp_LegAllocs) MarshalFIX(res *bufio.Writer) error {
    var err error
    
    if m.LegAllocAccount.HasValue() {
        (*res).WriteString("671=")
        val, err := m.LegAllocAccount.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegIndividualAllocID.HasValue() {
        (*res).WriteString("672=")
        val, err := m.LegIndividualAllocID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if len(m.Nested2PartyIDs) > 0 {
    
    (*res).WriteString("756=")
    (*res).WriteString(strconv.Itoa(len(m.Nested2PartyIDs)))
    (*res).WriteString("\x01")
    for _, g := range m.Nested2PartyIDs {
        err = g.MarshalFIX(res)
    }
    }
    
    if m.LegAllocQty.HasValue() {
        (*res).WriteString("673=")
        val, err := m.LegAllocQty.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegAllocAcctIDSource.HasValue() {
        (*res).WriteString("674=")
        val, err := m.LegAllocAcctIDSource.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.LegSettlCurrency.HasValue() {
        (*res).WriteString("675=")
        val, err := m.LegSettlCurrency.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    return err
}

func (m *LegPreAllocGrp_LegAllocs) UnmarshalFIX(input io.Reader) error {
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
                err = errors.New(fmt.Sprintf("Field unused by message LegPreAllocGrp_LegAllocs (%s)", string(field)))
            }
        }
        if err != nil {
            break
        }
    }
    return err
}


func (m *LegPreAllocGrp_LegAllocs) UnmarshalFieldFIX(id int, value []byte) (bool, error) {
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
        
        case 671:
            if !m.LegAllocAccount.HasValue() {
                used = true
                err = m.LegAllocAccount.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 672:
            if !m.LegIndividualAllocID.HasValue() {
                used = true
                err = m.LegIndividualAllocID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 756:
            // This counter (NoNested2PartyIDs) indicates the start of repeated group block 
            count, _ := strconv.Atoi(value)
            m.Nested2PartyIDs = make([]NestedParties2_Nested2PartyIDs, count)
            m._controlBlock.mostRecentRepeatingGroup = make([]FixMessage, count)
            for i := 0; i < count; i++ {
                m._controlBlock.mostRecentRepeatingGroup[i] = &m.Nested2PartyIDs[i]
            }
            m._controlBlock.mostRecentRepeatingGroupCounter = 0
            used = true
        
        
        
        case 673:
            if !m.LegAllocQty.HasValue() {
                used = true
                err = m.LegAllocQty.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 674:
            if !m.LegAllocAcctIDSource.HasValue() {
                used = true
                err = m.LegAllocAcctIDSource.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 675:
            if !m.LegSettlCurrency.HasValue() {
                used = true
                err = m.LegSettlCurrency.UnmarshalFIX(value)
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


