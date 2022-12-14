// SiGG-GoLang-On-the-Fly //
package eifactory

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/events/system"
	"github.com/hyperledger/firefly/internal/events/webhooks"
	"github.com/hyperledger/firefly/internal/events/websockets"
	"github.com/hyperledger/firefly/pkg/events"
)

var plugins = []events.Plugin{
	&websockets.WebSockets{},
	&webhooks.WebHooks{},
	&system.Events{},
}

var pluginsByName = make(map[string]events.Plugin)

func init() {
	for _, p := range plugins {
		pluginsByName[p.Name()] = p
	}
}

func InitConfig(config config.Section) {
	for name, plugin := range pluginsByName {
		plugin.InitConfig(config.SubSection(name))
	}
}

func GetPlugin(ctx context.Context, pluginType string) (events.Plugin, error) {
	plugin, ok := pluginsByName[pluginType]
	if !ok {
		return nil, i18n.NewError(ctx, coremsgs.MsgUnknownEventTransportPlugin, pluginType)
	}
	return plugin, nil
}
