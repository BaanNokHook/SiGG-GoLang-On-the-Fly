// SiGG-GoLang-On-the-Fly //

package tifactory

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/hyperledger/firefly/internal/coreconfig"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/tokens/fftokens"
	"github.com/hyperledger/firefly/pkg/tokens"
)

var pluginsByName = map[string]func() tokens.Plugin{
	(*fftokens.FFTokens)(nil).Name(): func() tokens.Plugin { return &fftokens.FFTokens{} },
}

func InitConfig(config config.ArraySection) {
	config.AddKnownKey(coreconfig.PluginConfigName)
	config.AddKnownKey(coreconfig.PluginConfigType)
	config.AddKnownKey(coreconfig.PluginBroadcastName)
	for name, plugin := range pluginsByName {
		plugin().InitConfig(config.SubSection(name))
	}
}

func InitConfigDeprecated(config config.ArraySection) {
	config.AddKnownKey(tokens.TokensConfigPlugin)
	config.AddKnownKey(coreconfig.PluginConfigName)
	for _, plugin := range pluginsByName {
		// Accept a superset of configs allowed by all plugins
		plugin().InitConfig(config)
	}
}

func GetPlugin(ctx context.Context, connectorName string) (tokens.Plugin, error) {
	plugin, ok := pluginsByName[connectorName]
	if !ok {
		return nil, i18n.NewError(ctx, coremsgs.MsgUnknownTokensPlugin, connectorName)
	}
	return plugin(), nil
}
