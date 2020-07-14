package NestedParties_NestedPartyIDs

import "bufio"
import "strconv"
import "fmt"

// Autogenerated at , do not edit



import NstdPtysSubGrp_NestedPartySubIDs "grgrbll/fix/Shared/NstdPtysSubGrp_NestedPartySubIDs"


func (m *NestedParties_NestedPartyIDs) MarshalFIX(res *bufio.Writer) error {
    var err error
    
    if m.NestedPartyID.HasValue() {
        (*res).WriteString("524=")
        val, err := m.NestedPartyID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.NestedPartyIDSource.HasValue() {
        (*res).WriteString("525=")
        val, err := m.NestedPartyIDSource.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.NestedPartyRole.HasValue() {
        (*res).WriteString("538=")
        val, err := m.NestedPartyRole.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if len(m.NestedPartySubIDs) > 0 {
    
    (*res).WriteString("804=")
    (*res).WriteString(strconv.Itoa(len(m.NestedPartySubIDs)))
    (*res).WriteString("\x01")
    for _, g := range m.NestedPartySubIDs {
        err = g.MarshalFIX(res)
    }
    }
    
    return err
}

func (m *NestedParties_NestedPartyIDs) UnmarshalFIX(input io.Reader) error {
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
                err = errors.New(fmt.Sprintf("Field unused by message NestedParties_NestedPartyIDs (%s)", string(field)))
            }
        }
        if err != nil {
            break
        }
    }
    return err
}


func (m *NestedParties_NestedPartyIDs) UnmarshalFieldFIX(id int, value []byte) (bool, error) {
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
        
        case 524:
            if !m.NestedPartyID.HasValue() {
                used = true
                err = m.NestedPartyID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 525:
            if !m.NestedPartyIDSource.HasValue() {
                used = true
                err = m.NestedPartyIDSource.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 538:
            if !m.NestedPartyRole.HasValue() {
                used = true
                err = m.NestedPartyRole.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 804:
            // This counter (NoNestedPartySubIDs) indicates the start of repeated group block 
            count, _ := strconv.Atoi(value)
            m.NestedPartySubIDs = make([]NstdPtysSubGrp_NestedPartySubIDs, count)
            m._controlBlock.mostRecentRepeatingGroup = make([]FixMessage, count)
            for i := 0; i < count; i++ {
                m._controlBlock.mostRecentRepeatingGroup[i] = &m.NestedPartySubIDs[i]
            }
            m._controlBlock.mostRecentRepeatingGroupCounter = 0
            used = true
        
        
        default:
            used = false
        }
    }
   
    return used, err
}


