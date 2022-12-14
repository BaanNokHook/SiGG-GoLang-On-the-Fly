// SiGG-GoLang-On-the-Fly //
package fftokens

import (
	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/hyperledger/firefly-common/pkg/wsclient"
)

func (ft *FFTokens) InitConfig(config config.KeySet) {
	wsclient.InitConfig(config)
}
