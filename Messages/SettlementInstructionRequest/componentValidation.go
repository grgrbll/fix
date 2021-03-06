package SettlementInstructionRequest

import "bufio"
import "regexp"

// Autogenerated at , do not edit



import HopGrp_Hops "grgrbll/fix/Shared/HopGrp_Hops"

import Parties_PartyIDs "grgrbll/fix/Shared/Parties_PartyIDs"


type settlementInstructionRequest_RegexValidator struct {
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
    SettlInstReqID *(regexp.Regexp)
    TransactTime *(regexp.Regexp)
    AllocAccount *(regexp.Regexp)
    AllocAcctIDSource *(regexp.Regexp)
    Side *(regexp.Regexp)
    Product *(regexp.Regexp)
    SecurityType *(regexp.Regexp)
    CFICode *(regexp.Regexp)
    EffectiveTime *(regexp.Regexp)
    ExpireTime *(regexp.Regexp)
    LastUpdateTime *(regexp.Regexp)
    StandInstDbType *(regexp.Regexp)
    StandInstDbName *(regexp.Regexp)
    StandInstDbID *(regexp.Regexp)
    SettlCurrency *(regexp.Regexp)
    SignatureLength *(regexp.Regexp)
    Signature *(regexp.Regexp)
    CheckSum *(regexp.Regexp)
    _controlBlock fixMessageControlBlock
}


var mySettlementInstructionRequest_RegexValidator settlementInstructionRequest_RegexValidator

func init() {
    mySettlementInstructionRequest_RegexValidator
    mySettlementInstructionRequest_RegexValidator.BeginString = regexp.MustCompile(`[^|]*`)
    mySettlementInstructionRequest_RegexValidator.BodyLength = regexp.MustCompile(`\d+`)
    mySettlementInstructionRequest_RegexValidator.MsgType = regexp.MustCompile(`[^|]*`)
    mySettlementInstructionRequest_RegexValidator.SenderCompID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructionRequest_RegexValidator.TargetCompID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructionRequest_RegexValidator.OnBehalfOfCompID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructionRequest_RegexValidator.DeliverToCompID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructionRequest_RegexValidator.SecureDataLen = regexp.MustCompile(`\d+`)
    mySettlementInstructionRequest_RegexValidator.SecureData = regexp.MustCompile(`.*`)
    mySettlementInstructionRequest_RegexValidator.MsgSeqNum = regexp.MustCompile(`\d+`)
    mySettlementInstructionRequest_RegexValidator.SenderSubID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructionRequest_RegexValidator.SenderLocationID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructionRequest_RegexValidator.TargetSubID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructionRequest_RegexValidator.TargetLocationID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructionRequest_RegexValidator.OnBehalfOfSubID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructionRequest_RegexValidator.OnBehalfOfLocationID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructionRequest_RegexValidator.DeliverToSubID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructionRequest_RegexValidator.DeliverToLocationID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructionRequest_RegexValidator.PossDupFlag = regexp.MustCompile(`[YN]`)
    mySettlementInstructionRequest_RegexValidator.PossResend = regexp.MustCompile(`[YN]`)
    mySettlementInstructionRequest_RegexValidator.SendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    mySettlementInstructionRequest_RegexValidator.OrigSendingTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    mySettlementInstructionRequest_RegexValidator.XmlDataLen = regexp.MustCompile(`\d+`)
    mySettlementInstructionRequest_RegexValidator.XmlData = regexp.MustCompile(`.*`)
    mySettlementInstructionRequest_RegexValidator.MessageEncoding = regexp.MustCompile(`[^|]*`)
    mySettlementInstructionRequest_RegexValidator.LastMsgSeqNumProcessed = regexp.MustCompile(`\d+`)
    mySettlementInstructionRequest_RegexValidator.ApplVerID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructionRequest_RegexValidator.CstmApplVerID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructionRequest_RegexValidator.SettlInstReqID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructionRequest_RegexValidator.TransactTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    mySettlementInstructionRequest_RegexValidator.AllocAccount = regexp.MustCompile(`[^|]*`)
    mySettlementInstructionRequest_RegexValidator.AllocAcctIDSource = regexp.MustCompile(`-?\d+`)
    mySettlementInstructionRequest_RegexValidator.Side = regexp.MustCompile(`[^|]`)
    mySettlementInstructionRequest_RegexValidator.Product = regexp.MustCompile(`-?\d+`)
    mySettlementInstructionRequest_RegexValidator.SecurityType = regexp.MustCompile(`[^|]*`)
    mySettlementInstructionRequest_RegexValidator.CFICode = regexp.MustCompile(`[^|]*`)
    mySettlementInstructionRequest_RegexValidator.EffectiveTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    mySettlementInstructionRequest_RegexValidator.ExpireTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    mySettlementInstructionRequest_RegexValidator.LastUpdateTime = regexp.MustCompile(`\d{4}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])-([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](.\d{3})?`)
    mySettlementInstructionRequest_RegexValidator.StandInstDbType = regexp.MustCompile(`-?\d+`)
    mySettlementInstructionRequest_RegexValidator.StandInstDbName = regexp.MustCompile(`[^|]*`)
    mySettlementInstructionRequest_RegexValidator.StandInstDbID = regexp.MustCompile(`[^|]*`)
    mySettlementInstructionRequest_RegexValidator.SettlCurrency = regexp.MustCompile(`[A-Z]{3}`)
    mySettlementInstructionRequest_RegexValidator.SignatureLength = regexp.MustCompile(`\d+`)
    mySettlementInstructionRequest_RegexValidator.Signature = regexp.MustCompile(`.*`)
    mySettlementInstructionRequest_RegexValidator.CheckSum = regexp.MustCompile(`[^|]*`)
}



func (m *SettlementInstructionRequest) HasRequiredFields() bool {
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
    
    
    
    valid = valid && m.SettlInstReqID.HasValue()
    
    
    
    valid = valid && m.TransactTime.HasValue()
    
    
    
    for _, g := range(m.PartyIDs) {
        valid = valid && g.HasRequiredFields()
    }
    
    
    
    valid = valid && m.CheckSum.HasValue()
    
    
    
    
    
    return valid
}




