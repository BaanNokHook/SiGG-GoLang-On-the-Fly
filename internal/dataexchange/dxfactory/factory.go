// SiGG-GoLang-On-the-Fly //
package dxfactory

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/hyperledger/firefly/internal/coreconfig"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/dataexchange/ffdx"
	"github.com/hyperledger/firefly/pkg/dataexchange"
)

var (
	NewFFDXPluginName = (*ffdx.FFDX)(nil).Name()
)

var pluginsByName = map[string]func() dataexchange.Plugin{
	NewFFDXPluginName: func() dataexchange.Plugin { return &ffdx.FFDX{} },
}

func InitConfig(config config.ArraySection) {
	config.AddKnownKey(coreconfig.PluginConfigName)
	config.AddKnownKey(coreconfig.PluginConfigType)
	for name, plugin := range pluginsByName {
		plugin().InitConfig(config.SubSection(name))
	}
}

func InitConfigDeprecated(config config.Section) {
	config.AddKnownKey(coreconfig.PluginConfigType)
	for name, plugin := range pluginsByName {
		plugin().InitConfig(config.SubSection(name))
	}
}

func GetPlugin(ctx context.Context, pluginType string) (dataexchange.Plugin, error) {
	plugin, ok := pluginsByName[pluginType]
	if !ok {
		return nil, i18n.NewError(ctx, coremsgs.MsgUnknownDataExchangePlugin, pluginType)
	}
	return plugin(), nil
}
