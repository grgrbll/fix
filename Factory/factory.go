// Autogenerated at , do not edit

import "errors"
import "fmt"

import (
     "grgrbll/fix/Messages/Quote"
     "grgrbll/fix/Messages/QuoteStatusRequest"
     "grgrbll/fix/Messages/ConfirmationAck"
     "grgrbll/fix/Messages/CollateralAssignment"
     "grgrbll/fix/Messages/MarketDataIncrementalRefresh"
     "grgrbll/fix/Messages/TradeCaptureReport"
     "grgrbll/fix/Messages/OrderMassCancelRequest"
     "grgrbll/fix/Messages/SecurityStatus"
     "grgrbll/fix/Messages/MarketDataSnapshotFullRefresh"
     "grgrbll/fix/Messages/CollateralInquiry"
     "grgrbll/fix/Messages/RegistrationInstructions"
     "grgrbll/fix/Messages/QuoteStatusReport"
     "grgrbll/fix/Messages/SecurityListUpdateReport"
     "grgrbll/fix/Messages/DerivativeSecurityListRequest"
     "grgrbll/fix/Messages/TradeCaptureReportRequest"
     "grgrbll/fix/Messages/TradeCaptureReportAck"
     "grgrbll/fix/Messages/ContraryIntentionReport"
     "grgrbll/fix/Messages/SecurityDefinitionUpdateReport"
     "grgrbll/fix/Messages/QuoteRequestReject"
     "grgrbll/fix/Messages/MarketDataRequestReject"
     "grgrbll/fix/Messages/OrderCancelRequest"
     "grgrbll/fix/Messages/CollateralReport"
     "grgrbll/fix/Messages/ResendRequest"
     "grgrbll/fix/Messages/ListStatus"
     "grgrbll/fix/Messages/QuoteRequest"
     "grgrbll/fix/Messages/BusinessMessageReject"
     "grgrbll/fix/Messages/ListCancelRequest"
     "grgrbll/fix/Messages/RequestForPositions"
     "grgrbll/fix/Messages/NewOrderSingle"
     "grgrbll/fix/Messages/SecurityTypeRequest"
     "grgrbll/fix/Messages/Advertisement"
     "grgrbll/fix/Messages/RequestForPositionsAck"
     "grgrbll/fix/Messages/OrderCancelReplaceRequest"
     "grgrbll/fix/Messages/DerivativeSecurityList"
     "grgrbll/fix/Messages/ListStrikePrice"
     "grgrbll/fix/Messages/CollateralRequest"
     "grgrbll/fix/Messages/ExecutionAcknowledgement"
     "grgrbll/fix/Messages/BidRequest"
     "grgrbll/fix/Messages/MassQuote"
     "grgrbll/fix/Messages/SecurityDefinition"
     "grgrbll/fix/Messages/BidResponse"
     "grgrbll/fix/Messages/AllocationInstruction"
     "grgrbll/fix/Messages/SettlementInstructionRequest"
     "grgrbll/fix/Messages/CollateralInquiryAck"
     "grgrbll/fix/Messages/AssignmentReport"
     "grgrbll/fix/Messages/NetworkCounterpartySystemStatusRequest"
     "grgrbll/fix/Messages/PositionMaintenanceRequest"
     "grgrbll/fix/Messages/OrderCancelReject"
     "grgrbll/fix/Messages/PositionMaintenanceReport"
     "grgrbll/fix/Messages/ListExecute"
     "grgrbll/fix/Messages/TradingSessionList"
     "grgrbll/fix/Messages/OrderMassCancelReport"
     "grgrbll/fix/Messages/SecurityList"
     "grgrbll/fix/Messages/UserResponse"
     "grgrbll/fix/Messages/PositionReport"
     "grgrbll/fix/Messages/NewOrderCross"
     "grgrbll/fix/Messages/TradeCaptureReportRequestAck"
     "grgrbll/fix/Messages/DontKnowTradeDK"
     "grgrbll/fix/Messages/SequenceReset"
     "grgrbll/fix/Messages/MarketDataRequest"
     "grgrbll/fix/Messages/SecurityStatusRequest"
     "grgrbll/fix/Messages/ListStatusRequest"
     "grgrbll/fix/Messages/Email"
     "grgrbll/fix/Messages/AllocationReportAck"
     "grgrbll/fix/Messages/TestRequest"
     "grgrbll/fix/Messages/News"
     "grgrbll/fix/Messages/OrderMassStatusRequest"
     "grgrbll/fix/Messages/SecurityTypes"
     "grgrbll/fix/Messages/ConfirmationRequest"
     "grgrbll/fix/Messages/MultilegOrderCancelReplace"
     "grgrbll/fix/Messages/Logout"
     "grgrbll/fix/Messages/Reject"
     "grgrbll/fix/Messages/NewOrderList"
     "grgrbll/fix/Messages/IOI"
     "grgrbll/fix/Messages/TradingSessionStatusRequest"
     "grgrbll/fix/Messages/XMLnonFIX"
     "grgrbll/fix/Messages/RegistrationInstructionsResponse"
     "grgrbll/fix/Messages/TradingSessionStatus"
     "grgrbll/fix/Messages/AllocationInstructionAck"
     "grgrbll/fix/Messages/Heartbeat"
     "grgrbll/fix/Messages/AllocationInstructionAlert"
     "grgrbll/fix/Messages/OrderStatusRequest"
     "grgrbll/fix/Messages/CrossOrderCancelReplaceRequest"
     "grgrbll/fix/Messages/CrossOrderCancelRequest"
     "grgrbll/fix/Messages/UserRequest"
     "grgrbll/fix/Messages/QuoteCancel"
     "grgrbll/fix/Messages/Confirmation"
     "grgrbll/fix/Messages/CollateralResponse"
     "grgrbll/fix/Messages/SecurityListRequest"
     "grgrbll/fix/Messages/SettlementInstructions"
     "grgrbll/fix/Messages/RFQRequest"
     "grgrbll/fix/Messages/QuoteResponse"
     "grgrbll/fix/Messages/SecurityDefinitionRequest"
     "grgrbll/fix/Messages/ExecutionReport"
     "grgrbll/fix/Messages/AdjustedPositionReport"
     "grgrbll/fix/Messages/Logon"
     "grgrbll/fix/Messages/NetworkCounterpartySystemStatusResponse"
     "grgrbll/fix/Messages/AllocationReport"
     "grgrbll/fix/Messages/MassQuoteAcknowledgement"
     "grgrbll/fix/Messages/NewOrderMultileg"
     "grgrbll/fix/Messages/TradingSessionListRequest"
)

const (
    
    Quote = "S"
    MsgType_S = "Quote"
    
    QuoteStatusRequest = "a"
    MsgType_a = "QuoteStatusRequest"
    
    ConfirmationAck = "AU"
    MsgType_AU = "ConfirmationAck"
    
    CollateralAssignment = "AY"
    MsgType_AY = "CollateralAssignment"
    
    MarketDataIncrementalRefresh = "X"
    MsgType_X = "MarketDataIncrementalRefresh"
    
    TradeCaptureReport = "AE"
    MsgType_AE = "TradeCaptureReport"
    
    OrderMassCancelRequest = "q"
    MsgType_q = "OrderMassCancelRequest"
    
    SecurityStatus = "f"
    MsgType_f = "SecurityStatus"
    
    MarketDataSnapshotFullRefresh = "W"
    MsgType_W = "MarketDataSnapshotFullRefresh"
    
    CollateralInquiry = "BB"
    MsgType_BB = "CollateralInquiry"
    
    RegistrationInstructions = "o"
    MsgType_o = "RegistrationInstructions"
    
    QuoteStatusReport = "AI"
    MsgType_AI = "QuoteStatusReport"
    
    SecurityListUpdateReport = "BK"
    MsgType_BK = "SecurityListUpdateReport"
    
    DerivativeSecurityListRequest = "z"
    MsgType_z = "DerivativeSecurityListRequest"
    
    TradeCaptureReportRequest = "AD"
    MsgType_AD = "TradeCaptureReportRequest"
    
    TradeCaptureReportAck = "AR"
    MsgType_AR = "TradeCaptureReportAck"
    
    ContraryIntentionReport = "BO"
    MsgType_BO = "ContraryIntentionReport"
    
    SecurityDefinitionUpdateReport = "BP"
    MsgType_BP = "SecurityDefinitionUpdateReport"
    
    QuoteRequestReject = "AG"
    MsgType_AG = "QuoteRequestReject"
    
    MarketDataRequestReject = "Y"
    MsgType_Y = "MarketDataRequestReject"
    
    OrderCancelRequest = "F"
    MsgType_F = "OrderCancelRequest"
    
    CollateralReport = "BA"
    MsgType_BA = "CollateralReport"
    
    ResendRequest = "2"
    MsgType_2 = "ResendRequest"
    
    ListStatus = "N"
    MsgType_N = "ListStatus"
    
    QuoteRequest = "R"
    MsgType_R = "QuoteRequest"
    
    BusinessMessageReject = "j"
    MsgType_j = "BusinessMessageReject"
    
    ListCancelRequest = "K"
    MsgType_K = "ListCancelRequest"
    
    RequestForPositions = "AN"
    MsgType_AN = "RequestForPositions"
    
    NewOrderSingle = "D"
    MsgType_D = "NewOrderSingle"
    
    SecurityTypeRequest = "v"
    MsgType_v = "SecurityTypeRequest"
    
    Advertisement = "7"
    MsgType_7 = "Advertisement"
    
    RequestForPositionsAck = "AO"
    MsgType_AO = "RequestForPositionsAck"
    
    OrderCancelReplaceRequest = "G"
    MsgType_G = "OrderCancelReplaceRequest"
    
    DerivativeSecurityList = "AA"
    MsgType_AA = "DerivativeSecurityList"
    
    ListStrikePrice = "m"
    MsgType_m = "ListStrikePrice"
    
    CollateralRequest = "AX"
    MsgType_AX = "CollateralRequest"
    
    ExecutionAcknowledgement = "BN"
    MsgType_BN = "ExecutionAcknowledgement"
    
    BidRequest = "k"
    MsgType_k = "BidRequest"
    
    MassQuote = "i"
    MsgType_i = "MassQuote"
    
    SecurityDefinition = "d"
    MsgType_d = "SecurityDefinition"
    
    BidResponse = "l"
    MsgType_l = "BidResponse"
    
    AllocationInstruction = "J"
    MsgType_J = "AllocationInstruction"
    
    SettlementInstructionRequest = "AV"
    MsgType_AV = "SettlementInstructionRequest"
    
    CollateralInquiryAck = "BG"
    MsgType_BG = "CollateralInquiryAck"
    
    AssignmentReport = "AW"
    MsgType_AW = "AssignmentReport"
    
    NetworkCounterpartySystemStatusRequest = "BC"
    MsgType_BC = "NetworkCounterpartySystemStatusRequest"
    
    PositionMaintenanceRequest = "AL"
    MsgType_AL = "PositionMaintenanceRequest"
    
    OrderCancelReject = "9"
    MsgType_9 = "OrderCancelReject"
    
    PositionMaintenanceReport = "AM"
    MsgType_AM = "PositionMaintenanceReport"
    
    ListExecute = "L"
    MsgType_L = "ListExecute"
    
    TradingSessionList = "BJ"
    MsgType_BJ = "TradingSessionList"
    
    OrderMassCancelReport = "r"
    MsgType_r = "OrderMassCancelReport"
    
    SecurityList = "y"
    MsgType_y = "SecurityList"
    
    UserResponse = "BF"
    MsgType_BF = "UserResponse"
    
    PositionReport = "AP"
    MsgType_AP = "PositionReport"
    
    NewOrderCross = "s"
    MsgType_s = "NewOrderCross"
    
    TradeCaptureReportRequestAck = "AQ"
    MsgType_AQ = "TradeCaptureReportRequestAck"
    
    DontKnowTradeDK = "Q"
    MsgType_Q = "DontKnowTradeDK"
    
    SequenceReset = "4"
    MsgType_4 = "SequenceReset"
    
    MarketDataRequest = "V"
    MsgType_V = "MarketDataRequest"
    
    SecurityStatusRequest = "e"
    MsgType_e = "SecurityStatusRequest"
    
    ListStatusRequest = "M"
    MsgType_M = "ListStatusRequest"
    
    Email = "C"
    MsgType_C = "Email"
    
    AllocationReportAck = "AT"
    MsgType_AT = "AllocationReportAck"
    
    TestRequest = "1"
    MsgType_1 = "TestRequest"
    
    News = "B"
    MsgType_B = "News"
    
    OrderMassStatusRequest = "AF"
    MsgType_AF = "OrderMassStatusRequest"
    
    SecurityTypes = "w"
    MsgType_w = "SecurityTypes"
    
    ConfirmationRequest = "BH"
    MsgType_BH = "ConfirmationRequest"
    
    MultilegOrderCancelReplace = "AC"
    MsgType_AC = "MultilegOrderCancelReplace"
    
    Logout = "5"
    MsgType_5 = "Logout"
    
    Reject = "3"
    MsgType_3 = "Reject"
    
    NewOrderList = "E"
    MsgType_E = "NewOrderList"
    
    IOI = "6"
    MsgType_6 = "IOI"
    
    TradingSessionStatusRequest = "g"
    MsgType_g = "TradingSessionStatusRequest"
    
    XMLnonFIX = "n"
    MsgType_n = "XMLnonFIX"
    
    RegistrationInstructionsResponse = "p"
    MsgType_p = "RegistrationInstructionsResponse"
    
    TradingSessionStatus = "h"
    MsgType_h = "TradingSessionStatus"
    
    AllocationInstructionAck = "P"
    MsgType_P = "AllocationInstructionAck"
    
    Heartbeat = "0"
    MsgType_0 = "Heartbeat"
    
    AllocationInstructionAlert = "BM"
    MsgType_BM = "AllocationInstructionAlert"
    
    OrderStatusRequest = "H"
    MsgType_H = "OrderStatusRequest"
    
    CrossOrderCancelReplaceRequest = "t"
    MsgType_t = "CrossOrderCancelReplaceRequest"
    
    CrossOrderCancelRequest = "u"
    MsgType_u = "CrossOrderCancelRequest"
    
    UserRequest = "BE"
    MsgType_BE = "UserRequest"
    
    QuoteCancel = "Z"
    MsgType_Z = "QuoteCancel"
    
    Confirmation = "AK"
    MsgType_AK = "Confirmation"
    
    CollateralResponse = "AZ"
    MsgType_AZ = "CollateralResponse"
    
    SecurityListRequest = "x"
    MsgType_x = "SecurityListRequest"
    
    SettlementInstructions = "T"
    MsgType_T = "SettlementInstructions"
    
    RFQRequest = "AH"
    MsgType_AH = "RFQRequest"
    
    QuoteResponse = "AJ"
    MsgType_AJ = "QuoteResponse"
    
    SecurityDefinitionRequest = "c"
    MsgType_c = "SecurityDefinitionRequest"
    
    ExecutionReport = "8"
    MsgType_8 = "ExecutionReport"
    
    AdjustedPositionReport = "BL"
    MsgType_BL = "AdjustedPositionReport"
    
    Logon = "A"
    MsgType_A = "Logon"
    
    NetworkCounterpartySystemStatusResponse = "BD"
    MsgType_BD = "NetworkCounterpartySystemStatusResponse"
    
    AllocationReport = "AS"
    MsgType_AS = "AllocationReport"
    
    MassQuoteAcknowledgement = "b"
    MsgType_b = "MassQuoteAcknowledgement"
    
    NewOrderMultileg = "AB"
    MsgType_AB = "NewOrderMultileg"
    
    TradingSessionListRequest = "BI"
    MsgType_BI = "TradingSessionListRequest"
    
)

type FixMessage interface{} {
    MarshalFIX(*bufio.Writer) error 
    UnmarshalFieldFIX(int, []byte) (bool, error)
    HasRequiredFields() bool
}

func getMessageByType(msgtype string) (FixMessage, error) {
    var err error = nil
    var message FixMessage = nil
    switch msgtype {
    
    case "S":
        message =  &Quote.Quote{}
    
    case "a":
        message =  &QuoteStatusRequest.QuoteStatusRequest{}
    
    case "AU":
        message =  &ConfirmationAck.ConfirmationAck{}
    
    case "AY":
        message =  &CollateralAssignment.CollateralAssignment{}
    
    case "X":
        message =  &MarketDataIncrementalRefresh.MarketDataIncrementalRefresh{}
    
    case "AE":
        message =  &TradeCaptureReport.TradeCaptureReport{}
    
    case "q":
        message =  &OrderMassCancelRequest.OrderMassCancelRequest{}
    
    case "f":
        message =  &SecurityStatus.SecurityStatus{}
    
    case "W":
        message =  &MarketDataSnapshotFullRefresh.MarketDataSnapshotFullRefresh{}
    
    case "BB":
        message =  &CollateralInquiry.CollateralInquiry{}
    
    case "o":
        message =  &RegistrationInstructions.RegistrationInstructions{}
    
    case "AI":
        message =  &QuoteStatusReport.QuoteStatusReport{}
    
    case "BK":
        message =  &SecurityListUpdateReport.SecurityListUpdateReport{}
    
    case "z":
        message =  &DerivativeSecurityListRequest.DerivativeSecurityListRequest{}
    
    case "AD":
        message =  &TradeCaptureReportRequest.TradeCaptureReportRequest{}
    
    case "AR":
        message =  &TradeCaptureReportAck.TradeCaptureReportAck{}
    
    case "BO":
        message =  &ContraryIntentionReport.ContraryIntentionReport{}
    
    case "BP":
        message =  &SecurityDefinitionUpdateReport.SecurityDefinitionUpdateReport{}
    
    case "AG":
        message =  &QuoteRequestReject.QuoteRequestReject{}
    
    case "Y":
        message =  &MarketDataRequestReject.MarketDataRequestReject{}
    
    case "F":
        message =  &OrderCancelRequest.OrderCancelRequest{}
    
    case "BA":
        message =  &CollateralReport.CollateralReport{}
    
    case "2":
        message =  &ResendRequest.ResendRequest{}
    
    case "N":
        message =  &ListStatus.ListStatus{}
    
    case "R":
        message =  &QuoteRequest.QuoteRequest{}
    
    case "j":
        message =  &BusinessMessageReject.BusinessMessageReject{}
    
    case "K":
        message =  &ListCancelRequest.ListCancelRequest{}
    
    case "AN":
        message =  &RequestForPositions.RequestForPositions{}
    
    case "D":
        message =  &NewOrderSingle.NewOrderSingle{}
    
    case "v":
        message =  &SecurityTypeRequest.SecurityTypeRequest{}
    
    case "7":
        message =  &Advertisement.Advertisement{}
    
    case "AO":
        message =  &RequestForPositionsAck.RequestForPositionsAck{}
    
    case "G":
        message =  &OrderCancelReplaceRequest.OrderCancelReplaceRequest{}
    
    case "AA":
        message =  &DerivativeSecurityList.DerivativeSecurityList{}
    
    case "m":
        message =  &ListStrikePrice.ListStrikePrice{}
    
    case "AX":
        message =  &CollateralRequest.CollateralRequest{}
    
    case "BN":
        message =  &ExecutionAcknowledgement.ExecutionAcknowledgement{}
    
    case "k":
        message =  &BidRequest.BidRequest{}
    
    case "i":
        message =  &MassQuote.MassQuote{}
    
    case "d":
        message =  &SecurityDefinition.SecurityDefinition{}
    
    case "l":
        message =  &BidResponse.BidResponse{}
    
    case "J":
        message =  &AllocationInstruction.AllocationInstruction{}
    
    case "AV":
        message =  &SettlementInstructionRequest.SettlementInstructionRequest{}
    
    case "BG":
        message =  &CollateralInquiryAck.CollateralInquiryAck{}
    
    case "AW":
        message =  &AssignmentReport.AssignmentReport{}
    
    case "BC":
        message =  &NetworkCounterpartySystemStatusRequest.NetworkCounterpartySystemStatusRequest{}
    
    case "AL":
        message =  &PositionMaintenanceRequest.PositionMaintenanceRequest{}
    
    case "9":
        message =  &OrderCancelReject.OrderCancelReject{}
    
    case "AM":
        message =  &PositionMaintenanceReport.PositionMaintenanceReport{}
    
    case "L":
        message =  &ListExecute.ListExecute{}
    
    case "BJ":
        message =  &TradingSessionList.TradingSessionList{}
    
    case "r":
        message =  &OrderMassCancelReport.OrderMassCancelReport{}
    
    case "y":
        message =  &SecurityList.SecurityList{}
    
    case "BF":
        message =  &UserResponse.UserResponse{}
    
    case "AP":
        message =  &PositionReport.PositionReport{}
    
    case "s":
        message =  &NewOrderCross.NewOrderCross{}
    
    case "AQ":
        message =  &TradeCaptureReportRequestAck.TradeCaptureReportRequestAck{}
    
    case "Q":
        message =  &DontKnowTradeDK.DontKnowTradeDK{}
    
    case "4":
        message =  &SequenceReset.SequenceReset{}
    
    case "V":
        message =  &MarketDataRequest.MarketDataRequest{}
    
    case "e":
        message =  &SecurityStatusRequest.SecurityStatusRequest{}
    
    case "M":
        message =  &ListStatusRequest.ListStatusRequest{}
    
    case "C":
        message =  &Email.Email{}
    
    case "AT":
        message =  &AllocationReportAck.AllocationReportAck{}
    
    case "1":
        message =  &TestRequest.TestRequest{}
    
    case "B":
        message =  &News.News{}
    
    case "AF":
        message =  &OrderMassStatusRequest.OrderMassStatusRequest{}
    
    case "w":
        message =  &SecurityTypes.SecurityTypes{}
    
    case "BH":
        message =  &ConfirmationRequest.ConfirmationRequest{}
    
    case "AC":
        message =  &MultilegOrderCancelReplace.MultilegOrderCancelReplace{}
    
    case "5":
        message =  &Logout.Logout{}
    
    case "3":
        message =  &Reject.Reject{}
    
    case "E":
        message =  &NewOrderList.NewOrderList{}
    
    case "6":
        message =  &IOI.IOI{}
    
    case "g":
        message =  &TradingSessionStatusRequest.TradingSessionStatusRequest{}
    
    case "n":
        message =  &XMLnonFIX.XMLnonFIX{}
    
    case "p":
        message =  &RegistrationInstructionsResponse.RegistrationInstructionsResponse{}
    
    case "h":
        message =  &TradingSessionStatus.TradingSessionStatus{}
    
    case "P":
        message =  &AllocationInstructionAck.AllocationInstructionAck{}
    
    case "0":
        message =  &Heartbeat.Heartbeat{}
    
    case "BM":
        message =  &AllocationInstructionAlert.AllocationInstructionAlert{}
    
    case "H":
        message =  &OrderStatusRequest.OrderStatusRequest{}
    
    case "t":
        message =  &CrossOrderCancelReplaceRequest.CrossOrderCancelReplaceRequest{}
    
    case "u":
        message =  &CrossOrderCancelRequest.CrossOrderCancelRequest{}
    
    case "BE":
        message =  &UserRequest.UserRequest{}
    
    case "Z":
        message =  &QuoteCancel.QuoteCancel{}
    
    case "AK":
        message =  &Confirmation.Confirmation{}
    
    case "AZ":
        message =  &CollateralResponse.CollateralResponse{}
    
    case "x":
        message =  &SecurityListRequest.SecurityListRequest{}
    
    case "T":
        message =  &SettlementInstructions.SettlementInstructions{}
    
    case "AH":
        message =  &RFQRequest.RFQRequest{}
    
    case "AJ":
        message =  &QuoteResponse.QuoteResponse{}
    
    case "c":
        message =  &SecurityDefinitionRequest.SecurityDefinitionRequest{}
    
    case "8":
        message =  &ExecutionReport.ExecutionReport{}
    
    case "BL":
        message =  &AdjustedPositionReport.AdjustedPositionReport{}
    
    case "A":
        message =  &Logon.Logon{}
    
    case "BD":
        message =  &NetworkCounterpartySystemStatusResponse.NetworkCounterpartySystemStatusResponse{}
    
    case "AS":
        message =  &AllocationReport.AllocationReport{}
    
    case "b":
        message =  &MassQuoteAcknowledgement.MassQuoteAcknowledgement{}
    
    case "AB":
        message =  &NewOrderMultileg.NewOrderMultileg{}
    
    case "BI":
        message =  &TradingSessionListRequest.TradingSessionListRequest{}
    
    default:
        err = errors.New(fmt.Sprintf("Message type (%s) not found", msgtype))
    }
    return message, err
}

const {
    WAIT_FOR_BEGIN = 1
    WAIT_FOR_TYPE = 2
    PROCESS_HISTORIC = 3
    PROCESS_NEW = 4
    ERROR_RECOVERY = 5
}

type tagValue struct {
    Tag int
    Value []byte
}

type MessageBuilder struct {
    init bool
    mode int
    currentMessage FixMessage
    historic []tagValue
}

func (builder *MessageBuilder) reset() {
    init = true
    mode = WAIT_FOR_BEGIN
    currentMessage = nil
    historic = nil
}

func (builder *MessageBuilder) UnmarshalFieldFIX(tag int, value []byte) (FixMessage, error) {
    
    if builder.init == false {
        builder.reset()
    }
    
    var err error = nil
    var message FixMessage = nil
    
    getNext := false
    for !getNext {
        switch builder.mode {
            case WAIT_FOR_BEGIN:
                getNext = true
                if tag == 8 {
                   builder.mode = WAIT_FOR_TYPE 
                }
                
            case WAIT_FOR_TYPE:
                getNext = true
                builder.historic = append(builder.historic, tagValue{tag,value})
                
                if tag == 35 {
                    builder.mode = PROCESS_HISTORIC
                    builder.currentMessage, err = getMessageByType(string(value))
                    if err != nil {
                        builder.mode = ERROR_RECOVERY
                        getNext = false
                    }
                }
                
            case PROCESS_HISTORIC:
                getNext = true
                builder.mode = PROCESS_NEW
                for _, kv := range(builder.historic) {
                    err = currentMessage.UnmarshalFieldFIX(kv.Tag, kv.Value)
                    if err != nil {
                        builder.mode = ERROR_RECOVERY
                        getNext = false
                        break
                    }
                }
                
            case PROCESS_NEW:
                getNext = true
                err = currentMessage.UnmarshalFieldFIX(tag, value)
                if err != nil {
                    builder.mode = ERROR_RECOVERY
                    getNext = false
                }
                if tag == 10 {
                    message = builder.currentMessage
                }
                
            
            case ERROR_RECOVERY:
                getNext = true
                builder.mode = WAIT_FOR_BEGIN
                reset()

        }
    }
    
    return message, err
}