package HopGrp_Hops

import "bufio"
import "strconv"
import "fmt"

// Autogenerated at , do not edit




func (m *HopGrp_Hops) MarshalFIX(res *bufio.Writer) error {
    var err error
    
    if m.HopCompID.HasValue() {
        (*res).WriteString("628=")
        val, err := m.HopCompID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.HopSendingTime.HasValue() {
        (*res).WriteString("629=")
        val, err := m.HopSendingTime.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.HopRefID.HasValue() {
        (*res).WriteString("630=")
        val, err := m.HopRefID.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    return err
}

func (m *HopGrp_Hops) UnmarshalFIX(input io.Reader) error {
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
                err = errors.New(fmt.Sprintf("Field unused by message HopGrp_Hops (%s)", string(field)))
            }
        }
        if err != nil {
            break
        }
    }
    return err
}


func (m *HopGrp_Hops) UnmarshalFieldFIX(id int, value []byte) (bool, error) {
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
        
        case 628:
            if !m.HopCompID.HasValue() {
                used = true
                err = m.HopCompID.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 629:
            if !m.HopSendingTime.HasValue() {
                used = true
                err = m.HopSendingTime.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 630:
            if !m.HopRefID.HasValue() {
                used = true
                err = m.HopRefID.UnmarshalFIX(value)
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

