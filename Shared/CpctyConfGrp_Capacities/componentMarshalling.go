package CpctyConfGrp_Capacities

import "bufio"
import "strconv"
import "fmt"

// Autogenerated at , do not edit




func (m *CpctyConfGrp_Capacities) MarshalFIX(res *bufio.Writer) error {
    var err error
    
    if m.OrderCapacity.HasValue() {
        (*res).WriteString("528=")
        val, err := m.OrderCapacity.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.OrderRestrictions.HasValue() {
        (*res).WriteString("529=")
        val, err := m.OrderRestrictions.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.OrderCapacityQty.HasValue() {
        (*res).WriteString("863=")
        val, err := m.OrderCapacityQty.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    return err
}

func (m *CpctyConfGrp_Capacities) UnmarshalFIX(input io.Reader) error {
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
                err = errors.New(fmt.Sprintf("Field unused by message CpctyConfGrp_Capacities (%s)", string(field)))
            }
        }
        if err != nil {
            break
        }
    }
    return err
}


func (m *CpctyConfGrp_Capacities) UnmarshalFieldFIX(id int, value []byte) (bool, error) {
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
        
        case 528:
            if !m.OrderCapacity.HasValue() {
                used = true
                err = m.OrderCapacity.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 529:
            if !m.OrderRestrictions.HasValue() {
                used = true
                err = m.OrderRestrictions.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 863:
            if !m.OrderCapacityQty.HasValue() {
                used = true
                err = m.OrderCapacityQty.UnmarshalFIX(value)
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


