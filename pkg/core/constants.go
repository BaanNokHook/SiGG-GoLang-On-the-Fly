// SiGG-GoLang-On-the-Fly //
package core

const (

	// LegacySystemNamespace is the system reserved namespace name (deprecated)
	LegacySystemNamespace = "ff_system"
)

const (
	// SystemTopicDefinitions is the FireFly event topic for events that are confirmations of definition of pre-defined datatypes
	SystemTopicDefinitions = "ff_definition"
	// SystemBatchPinTopic is the FireFly event topic for events from the FireFly batch pin listener
	SystemBatchPinTopic = "ff_batch_pin"
)

const (

	// SystemTagDefineDatatype is the tag for messages that broadcast data definitions
	SystemTagDefineDatatype = "ff_define_datatype"

	// DeprecatedSystemTagDefineOrganization is the tag for messages that broadcast organization definitions
	DeprecatedSystemTagDefineOrganization = "ff_define_organization"

	// DeprecatedSystemTagDefineNode is the tag for messages that broadcast node definitions
	DeprecatedSystemTagDefineNode = "ff_define_node"

	// SystemTagDefineGroup is the tag for messages that send the definition of a group, to all parties in that group
	SystemTagDefineGroup = "ff_define_group"

	// SystemTagDefinePool is the tag for messages that broadcast data definitions
	SystemTagDefinePool = "ff_define_pool"

	// SystemTagDefineFFI is the tag for messages that broadcast contract FFIs
	SystemTagDefineFFI = "ff_define_ffi"

	// SystemTagDefineContractAPI is the tag for messages that broadcast contract APIs
	SystemTagDefineContractAPI = "ff_define_contract_api"

	// SystemTagIdentityClaim is the tag for messages that broadcast an identity claim
	SystemTagIdentityClaim = "ff_identity_claim"

	// SystemTagIdentityVerification is the tag for messages that broadcast an identity verification
	SystemTagIdentityVerification = "ff_identity_verification"

	// SystemTagIdentityUpdate is the tag for messages that broadcast an identity update
	SystemTagIdentityUpdate = "ff_identity_update"
)
