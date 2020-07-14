package ListExecute

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"


type listExecute_RegexValidator struct {
    BeginString *(regexp.Regexp)
    BodyLength *(regexp.Regexp)
    MsgType *(regexp.Regexp)
    SenderCompID *(regexp.Regexp)
    TargetCompID *(regexp.Regexp)
    OnBehalfOfCompID *(regexp.Regexp)
    DeliverToCompID *(regexp.Regexp)
    SecureDataLen *(regexp.Regexp)
    SecureData *(regexp.Regexp)
    MsgSeqNum *(regexp.Regexp)
    SenderSubID *(regexp.Regexp)
    SenderLocationID *(regexp.Regexp)
    TargetSubID *(regexp.Regexp)
    TargetLocationID *(regexp.Regexp)
    OnBehalfOfSubID *(regexp.Regexp)
    OnBehalfOfLocationID *(regexp.Regexp)
    DeliverToSubID *(regexp.Regexp)
    DeliverToLocationID *(regexp.Regexp)
    PossDupFlag *(regexp.Regexp)
    PossResend *(regexp.Regexp)
    SendingTime *(regexp.Regexp)
    OrigSendingTime *(regexp.Regexp)
    XmlDataLen *(regexp.Regexp)
    XmlData *(regexp.Regexp)
    MessageEncoding *(regexp.Regexp)
    LastMsgSeqNumProcessed *(regexp.Regexp)
    ApplVerID *(regexp.Regexp)
    CstmApplVerID *(regexp.Regexp)
    ListID *(regexp.Regexp)
    ClientBidID *(regexp.Regexp)
    BidID *(regexp.Regexp)
    TransactTime *(regexp.Regexp)
    Text *(regexp.Regexp)
    EncodedTextLen *(regexp.Regexp)
    EncodedText *(regexp.Regexp)
    SignatureLength *(regexp.Regexp)
    Signature *(regexp.Regexp)
    CheckSum *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var myListExecute_RegexValidator listExecute_RegexValidator

func init() {
    myListExecute_RegexValidator
    myListExecute_RegexValidator.BeginString = regexp.MustCompile(`[^|]*`)
    myListExecute_RegexValidator.BodyLength = regexp.MustCompile(`\d+`)
    myListExecute_RegexValidator.MsgType = regexp.MustCompile(`[^|]*`)
    myListExecute_RegexValidator.SenderCompID = regexp.MustCompile(`[^|]*`)
    myListExecute_RegexValidator.TargetCompID = regexp.MustCompile(`[^|]*`)
    myListExecute_RegexValidator.OnBehalfOfCompID = regexp.MustCompile(`[^|]*`)
    myListExecute_RegexValidator.DeliverToCompID = regexp.MustCompile(`[^|]*`)
    myListExecute_RegexValidator.SecureDataLen = regexp.MustCompile(`\d+`)
    myListExecute_RegexValidator.SecureData = regexp.MustCompile(`.*`)
    myListExecute_RegexValidator.MsgSeqNum = regexp.MustCompile(`\d+`)
    myListExecute_RegexValidator.SenderSubID = regexp.MustCompile(`[^|]*`)
    myListExecute_RegexValidator.SenderLocationID = regexp.MustCompile(`[^|]*`)
    myListExecute_RegexValidator.TargetSubID = regexp.MustCompile(`[^|]*`)
    myListExecute_RegexValidator.TargetLocationID = regexp.MustCompile(`[^|]*`)
    myListExecute_RegexValidator.OnBehalfOfSubID = regexp.MustCompile(`[^|]*`)
    myListExecute_RegexValidator.OnBehalfOfLocationID = regexp.MustCompile(`[^|]*`)
    myListExecute_RegexValidator.DeliverToSubID = regexp.MustCompile(`[^|]*`)
    myListExecute_RegexValidator.DeliverToLocationID = regexp.MustCompile(`[^|]*`)
    myListExecute_RegexValidator.PossDupFlag = regexp.MustCompile(`[YN]`)
    myListExecute_RegexValidator.PossResend = regexp.MustCompile(`[YN]`)
    myListExecute_RegexValidator.SendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myListExecute_RegexValidator.OrigSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myListExecute_RegexValidator.XmlDataLen = regexp.MustCompile(`\d+`)
    myListExecute_RegexValidator.XmlData = regexp.MustCompile(`.*`)
    myListExecute_RegexValidator.MessageEncoding = regexp.MustCompile(`[^|]*`)
    myListExecute_RegexValidator.LastMsgSeqNumProcessed = regexp.MustCompile(`\d+`)
    myListExecute_RegexValidator.ApplVerID = regexp.MustCompile(`[^|]*`)
    myListExecute_RegexValidator.CstmApplVerID = regexp.MustCompile(`[^|]*`)
    myListExecute_RegexValidator.ListID = regexp.MustCompile(`[^|]*`)
    myListExecute_RegexValidator.ClientBidID = regexp.MustCompile(`[^|]*`)
    myListExecute_RegexValidator.BidID = regexp.MustCompile(`[^|]*`)
    myListExecute_RegexValidator.TransactTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    myListExecute_RegexValidator.Text = regexp.MustCompile(`[^|]*`)
    myListExecute_RegexValidator.EncodedTextLen = regexp.MustCompile(`\d+`)
    myListExecute_RegexValidator.EncodedText = regexp.MustCompile(`.*`)
    myListExecute_RegexValidator.SignatureLength = regexp.MustCompile(`\d+`)
    myListExecute_RegexValidator.Signature = regexp.MustCompile(`.*`)
    myListExecute_RegexValidator.CheckSum = regexp.MustCompile(`[^|]*`)
}



func (m *ListExecute) HasRequiredFields() bool {
    valid := true
    
    

    
    
    valid = valid && m.BeginString.HasValue()
    
    
    
    valid = valid && m.BodyLength.HasValue()
    
    
    
    valid = valid && m.MsgType.HasValue()
    
    
    
    valid = valid && m.SenderCompID.HasValue()
    
    
    
    valid = valid && m.TargetCompID.HasValue()
    
    
    
    valid = valid && m.MsgSeqNum.HasValue()
    
    
    
    valid = valid && m.SendingTime.HasValue()
    
    
    
    for _, g := range(m.Hops) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.ListID.HasValue()
    
    
    
    valid = valid && m.TransactTime.HasValue()
    
    
    
    valid = valid && m.CheckSum.HasValue()
    
    
    
    
    
    return valid
}



