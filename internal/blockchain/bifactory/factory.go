// SiGG-GoLang-On-the-Fly //
package bifactory

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/hyperledger/firefly/internal/blockchain/ethereum"
	"github.com/hyperledger/firefly/internal/blockchain/fabric"
	"github.com/hyperledger/firefly/internal/coreconfig"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/blockchain"
)

var pluginsByType = map[string]func() blockchain.Plugin{
	(*ethereum.Ethereum)(nil).Name(): func() blockchain.Plugin { return &ethereum.Ethereum{} },
	(*fabric.Fabric)(nil).Name():     func() blockchain.Plugin { return &fabric.Fabric{} },
}

func InitConfig(config config.ArraySection) {
	config.AddKnownKey(coreconfig.PluginConfigName)
	config.AddKnownKey(coreconfig.PluginConfigType)
	for name, plugin := range pluginsByType {
		plugin().InitConfig(config.SubSection(name))
	}
}

func InitConfigDeprecated(config config.Section) {
	config.AddKnownKey(coreconfig.PluginConfigType)
	for name, plugin := range pluginsByType {
		plugin().InitConfig(config.SubSection(name))
	}
}

func GetPlugin(ctx context.Context, pluginType string) (blockchain.Plugin, error) {
	plugin, ok := pluginsByType[pluginType]
	if !ok {
		return nil, i18n.NewError(ctx, coremsgs.MsgUnknownBlockchainPlugin, pluginType)
	}
	return plugin(), nil
}