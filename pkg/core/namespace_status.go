// SiGG-GoLang-On-the-Fly //
package core

import "github.com/hyperledger/firefly-common/pkg/fftypes"

// NamespaceStatus is a set of information that represents the configuration and status of a given namespace
type NamespaceStatus struct {
	Namespace  *Namespace                `ffstruct:"NamespaceStatus" json:"namespace"`
	Node       *NamespaceStatusNode      `ffstruct:"NamespaceStatus" json:"node,omitempty"`
	Org        *NamespaceStatusOrg       `ffstruct:"NamespaceStatus" json:"org,omitempty"`
	Plugins    NamespaceStatusPlugins    `ffstruct:"NamespaceStatus" json:"plugins"`
	Multiparty NamespaceStatusMultiparty `ffstruct:"NamespaceStatus" json:"multiparty"`
}

// NamespaceStatusNode is the information about the local node, returned in the namespace status
type NamespaceStatusNode struct {
	Name       string        `ffstruct:"NamespaceStatusNode" json:"name"`
	Registered bool          `ffstruct:"NamespaceStatusNode" json:"registered"`
	ID         *fftypes.UUID `ffstruct:"NamespaceStatusNode" json:"id,omitempty"`
}

// NamespaceStatusOrg is the information about the node owning org, returned in the namespace status
type NamespaceStatusOrg struct {
	Name       string         `ffstruct:"NamespaceStatusOrg" json:"name"`
	Registered bool           `ffstruct:"NamespaceStatusOrg" json:"registered"`
	DID        string         `ffstruct:"NamespaceStatusOrg" json:"did,omitempty"`
	ID         *fftypes.UUID  `ffstruct:"NamespaceStatusOrg" json:"id,omitempty"`
	Verifiers  []*VerifierRef `ffstruct:"NamespaceStatusOrg" json:"verifiers,omitempty"`
}

// NamespaceStatusPlugins is a map of plugins configured in the namespace
type NamespaceStatusPlugins struct {
	Blockchain    []*NamespaceStatusPlugin `ffstruct:"NamespaceStatusPlugins" json:"blockchain"`
	Database      []*NamespaceStatusPlugin `ffstruct:"NamespaceStatusPlugins" json:"database"`
	DataExchange  []*NamespaceStatusPlugin `ffstruct:"NamespaceStatusPlugins" json:"dataExchange"`
	Events        []*NamespaceStatusPlugin `ffstruct:"NamespaceStatusPlugins" json:"events"`
	Identity      []*NamespaceStatusPlugin `ffstruct:"NamespaceStatusPlugins" json:"identity"`
	SharedStorage []*NamespaceStatusPlugin `ffstruct:"NamespaceStatusPlugins" json:"sharedStorage"`
	Tokens        []*NamespaceStatusPlugin `ffstruct:"NamespaceStatusPlugins" json:"tokens"`
}

// NamespaceStatusPlugin is information about a plugin
type NamespaceStatusPlugin struct {
	Name       string `ffstruct:"NamespaceStatusPlugin" json:"name,omitempty"`
	PluginType string `ffstruct:"NamespaceStatusPlugin" json:"pluginType"`
}

// NamespaceStatusMultiparty is information about multiparty mode and any associated multiparty contracts
type NamespaceStatusMultiparty struct {
	Enabled   bool                 `ffstruct:"NamespaceStatusMultiparty" json:"enabled"`
	Contracts *MultipartyContracts `ffstruct:"NamespaceStatusMultiparty" json:"contract,omitempty"`
}
