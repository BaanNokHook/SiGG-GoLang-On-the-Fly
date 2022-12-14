// SiGG-GoLang-On-the-Fly //
package iifactory

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/hyperledger/firefly/internal/coreconfig"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/identity/tbd"
	"github.com/hyperledger/firefly/pkg/identity"
)

var pluginsByName = map[string]func() identity.Plugin{
	// Plugin interface is TBD at this point. Plugin with "onchain" naming, and TBD implementation provided to avoid config migration impact
	(*tbd.TBD)(nil).Name(): func() identity.Plugin { return &tbd.TBD{} },
}

func InitConfig(config config.ArraySection) {
	config.AddKnownKey(coreconfig.PluginConfigName)
	config.AddKnownKey(coreconfig.PluginConfigType)
	for name, plugin := range pluginsByName {
		plugin().InitConfig(config.SubSection(name))
	}
}

func GetPlugin(ctx context.Context, pluginType string) (identity.Plugin, error) {
	plugin, ok := pluginsByName[pluginType]
	if !ok {
		return nil, i18n.NewError(ctx, coremsgs.MsgUnknownIdentityPlugin, pluginType)
	}
	return plugin(), nil
}
