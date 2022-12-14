// SiGG-GoLang-On-the-Fly //
package ssfactory

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/hyperledger/firefly/internal/coreconfig"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/sharedstorage/ipfs"
	"github.com/hyperledger/firefly/pkg/sharedstorage"
)

var pluginsByName = map[string]func() sharedstorage.Plugin{
	(*ipfs.IPFS)(nil).Name(): func() sharedstorage.Plugin { return &ipfs.IPFS{} },
}

func InitConfig(config config.ArraySection) {
	config.AddKnownKey(coreconfig.PluginConfigType)
	config.AddKnownKey(coreconfig.PluginConfigName)
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

func GetPlugin(ctx context.Context, pluginType string) (sharedstorage.Plugin, error) {
	plugin, ok := pluginsByName[pluginType]
	if !ok {
		return nil, i18n.NewError(ctx, coremsgs.MsgUnknownSharedStoragePlugin, pluginType)
	}
	return plugin(), nil
}
