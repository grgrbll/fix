package NestedParties3_Nested3PartyIDs

import "bufio"
import "strconv"
import "fmt"

// Autogenerated at , do not edit



import NstdPtys3SubGrp_Nested3PartySubIDs "grgrbll/fix/Shared/NstdPtys3SubGrp_Nested3PartySubIDs"


func (m *NestedParties3_Nested3PartyIDs) MarshalFIX(res *bufio.Writer) error {
    var err error
    
    if m.Nested3PartyID.HasValue() {
        (*res).WriteString("949=")
        val, err := m.Nested3PartyID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.Nested3PartyIDSource.HasValue() {
        (*res).WriteString("950=")
        val, err := m.Nested3PartyIDSource.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.Nested3PartyRole.HasValue() {
        (*res).WriteString("951=")
        val, err := m.Nested3PartyRole.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if len(m.Nested3PartySubIDs) > 0 {
    
    (*res).WriteString("952=")
    (*res).WriteString(strconv.Itoa(len(m.Nested3PartySubIDs)))
    (*res).WriteString("\x01")
    for _, g := range m.Nested3PartySubIDs {
        err = g.MarshalFIX(res)
    }
    }
    
    return err
}

func (m *NestedParties3_Nested3PartyIDs) UnmarshalFIX(input io.Reader) error {
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
                err = errors.New(fmt.Sprintf("Field unused by message NestedParties3_Nested3PartyIDs (%s)", string(field)))
            }
        }
        if err != nil {
            break
        }
    }
    return err
}


func (m *NestedParties3_Nested3PartyIDs) UnmarshalFieldFIX(id int, value []byte) (bool, error) {
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
        
        case 949:
            if !m.Nested3PartyID.HasValue() {
                used = true
                err = m.Nested3PartyID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 950:
            if !m.Nested3PartyIDSource.HasValue() {
                used = true
                err = m.Nested3PartyIDSource.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 951:
            if !m.Nested3PartyRole.HasValue() {
                used = true
                err = m.Nested3PartyRole.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 952:
            // This counter (NoNested3PartySubIDs) indicates the start of repeated group block 
            count, _ := strconv.Atoi(value)
            m.Nested3PartySubIDs = make([]NstdPtys3SubGrp_Nested3PartySubIDs, count)
            m._controlBlock.mostRecentRepeatingGroup = make([]FixMessage, count)
            for i := 0; i < count; i++ {
                m._controlBlock.mostRecentRepeatingGroup[i] = &m.Nested3PartySubIDs[i]
            }
            m._controlBlock.mostRecentRepeatingGroupCounter = 0
            used = true
        
        
        default:
            used = false
        }
    }
   
    return used, err
}


