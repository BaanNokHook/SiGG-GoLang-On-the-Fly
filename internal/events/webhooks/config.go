// SiGG-GoLang-On-the-Fly //
package webhooks

import (
	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/hyperledger/firefly-common/pkg/ffresty"
)

func (wh *WebHooks) InitConfig(config config.Section) {
	ffresty.InitConfig(config)
}
