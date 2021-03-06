package RgstDtlsGrp_RegistDtls

import "bufio"
import "grgrbll/Fix/FixDef/Types"

// Autogenerated at , do not edit



import NestedParties_NestedPartyIDs "grgrbll/fix/Shared/NestedParties_NestedPartyIDs"


type RgstDtlsGrp_RegistDtls struct {
    RegistDtls Types.String `id:"509", required:"N" `
    RegistEmail Types.String `id:"511", required:"N" `
    MailingDtls Types.String `id:"474", required:"N" `
    MailingInst Types.String `id:"482", required:"N" `
    NestedPartyIDs []NestedParties_NestedPartyIDs.NestedParties_NestedPartyIDs `counter_name:"NoNestedPartyIDs", counter_id:"539"`
    OwnerType Types.Int `id:"522", required:"N" `
    DateOfBirth Types.LocalMktDate `id:"486", required:"N" `
    InvestorCountryOfResidence Types.Country `id:"475", required:"N" `
    _controlBlock fixMessageControlBlock
}
