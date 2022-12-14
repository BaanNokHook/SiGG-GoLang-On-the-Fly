// SiGG-GoLang-On-the-Fly //
package ipfs

import (
	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/hyperledger/firefly-common/pkg/ffresty"
)

const (
	// IPFSConfAPISubconf is the http configuration to connect to the API endpoint of IPFS
	IPFSConfAPISubconf = "api"
	// IPFSConfGatewaySubconf is the http configuration to connect to the Gateway endpoint of IPFS
	IPFSConfGatewaySubconf = "gateway"
)

func (i *IPFS) InitConfig(config config.Section) {
	ffresty.InitConfig(config.SubSection(IPFSConfAPISubconf))
	ffresty.InitConfig(config.SubSection(IPFSConfGatewaySubconf))
}
