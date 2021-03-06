package TrdRegTimestamps_TrdRegTimestamps

import "bufio"
import "strconv"
import "fmt"

// Autogenerated at , do not edit




func (m *TrdRegTimestamps_TrdRegTimestamps) MarshalFIX(res *bufio.Writer) error {
    var err error
    
    if m.TrdRegTimestamp.HasValue() {
        (*res).WriteString("769=")
        val, err := m.TrdRegTimestamp.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.TrdRegTimestampType.HasValue() {
        (*res).WriteString("770=")
        val, err := m.TrdRegTimestampType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.TrdRegTimestampOrigin.HasValue() {
        (*res).WriteString("771=")
        val, err := m.TrdRegTimestampOrigin.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.DeskType.HasValue() {
        (*res).WriteString("1033=")
        val, err := m.DeskType.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.DeskTypeSource.HasValue() {
        (*res).WriteString("1034=")
        val, err := m.DeskTypeSource.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    if m.DeskOrderHandlingInst.HasValue() {
        (*res).WriteString("1035=")
        val, err := m.DeskOrderHandlingInst.MarshalFIX()
        if err != nil {
            return err
        }
        (*res).Write(val)
        (*res).WriteString("\x01")
    }
    return err
}

func (m *TrdRegTimestamps_TrdRegTimestamps) UnmarshalFIX(input io.Reader) error {
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
                err = errors.New(fmt.Sprintf("Field unused by message TrdRegTimestamps_TrdRegTimestamps (%s)", string(field)))
            }
        }
        if err != nil {
            break
        }
    }
    return err
}


func (m *TrdRegTimestamps_TrdRegTimestamps) UnmarshalFieldFIX(id int, value []byte) (bool, error) {
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
        
        case 769:
            if !m.TrdRegTimestamp.HasValue() {
                used = true
                err = m.TrdRegTimestamp.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 770:
            if !m.TrdRegTimestampType.HasValue() {
                used = true
                err = m.TrdRegTimestampType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 771:
            if !m.TrdRegTimestampOrigin.HasValue() {
                used = true
                err = m.TrdRegTimestampOrigin.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 1033:
            if !m.DeskType.HasValue() {
                used = true
                err = m.DeskType.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 1034:
            if !m.DeskTypeSource.HasValue() {
                used = true
                err = m.DeskTypeSource.UnmarshalFIX(value)
                if err != nil {
                    return used, err
                }
            }
        
        
        
        case 1035:
            if !m.DeskOrderHandlingInst.HasValue() {
                used = true
                err = m.DeskOrderHandlingInst.UnmarshalFIX(value)
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


