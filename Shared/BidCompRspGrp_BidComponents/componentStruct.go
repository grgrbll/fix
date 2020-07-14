package BidCompRspGrp_BidComponents

import "bufio"
import "grgrbll/Fix/FixDef/Types"

// Autogenerated at , do not edit




type BidCompRspGrp_BidComponents struct {
    Commission Types.Amt `id:"12", required:"N" `
    CommType Types.Char `id:"13", required:"N" `
    CommCurrency Types.Currency `id:"479", required:"N" `
    FundRenewWaiv Types.Char `id:"497", required:"N" `
    ListID Types.String `id:"66", required:"N" `
    Country Types.Country `id:"421", required:"N" `
    Side Types.Char `id:"54", required:"N" `
    Price Types.Price `id:"44", required:"N" `
    PriceType Types.Int `id:"423", required:"N" `
    FairValue Types.Amt `id:"406", required:"N" `
    NetGrossInd Types.Int `id:"430", required:"N" `
    SettlType Types.String `id:"63", required:"N" `
    SettlDate Types.LocalMktDate `id:"64", required:"N" `
    TradingSessionID Types.String `id:"336", required:"N" `
    TradingSessionSubID Types.String `id:"625", required:"N" `
    Text Types.String `id:"58", required:"N" `
    EncodedTextLen Types.Length `id:"354", required:"N" `
    EncodedText Types.Data `id:"355", required:"N" `
    _controlBlock fixMessageControlBlock
}