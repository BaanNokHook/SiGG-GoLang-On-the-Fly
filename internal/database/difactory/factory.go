// SiGG-GoLang-On-the-Fly //
package difactory

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/hyperledger/firefly/internal/coreconfig"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/database"
)

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

func GetPlugin(ctx context.Context, pluginType string) (database.Plugin, error) {
	plugin, ok := pluginsByName[pluginType]
	if !ok {
		return nil, i18n.NewError(ctx, coremsgs.MsgUnknownDatabasePlugin, pluginType)
	}
	return plugin(), nil
}
