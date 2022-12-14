// SiGG-GoLang-On-the-Fly //
package ffdx

import (
	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/hyperledger/firefly-common/pkg/wsclient"
)

const (
	// DataExchangeManifestEnabled determines whether to require+validate a manifest from other DX instances in the network. Must be supported by the connector
	DataExchangeManifestEnabled = "manifestEnabled"
	// DataExchangeInitEnabled instructs FireFly to always post all current nodes to the /init API before connecting or reconnecting to the connector
	DataExchangeInitEnabled = "initEnabled"
)

func (h *FFDX) InitConfig(config config.Section) {
	wsclient.InitConfig(config)
	config.AddKnownKey(DataExchangeManifestEnabled, false)
	config.AddKnownKey(DataExchangeInitEnabled, false)
}
