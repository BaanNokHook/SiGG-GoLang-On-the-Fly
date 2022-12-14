// SiGG-GoLang-On-the-Fly //
package definitions

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/hyperledger/firefly-common/pkg/log"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

func (bm *definitionSender) DefineTokenPool(ctx context.Context, pool *core.TokenPoolAnnouncement, waitConfirm bool) error {
	// Map token connector name -> broadcast name
	if broadcastName, exists := bm.tokenBroadcastNames[pool.Pool.Connector]; exists {
		pool.Pool.Connector = broadcastName
	} else {
		log.L(ctx).Infof("Could not find broadcast name for token connector: %s", pool.Pool.Connector)
		return i18n.NewError(ctx, coremsgs.MsgInvalidConnectorName, broadcastName, "token")
	}

	if bm.multiparty {
		if err := pool.Pool.Validate(ctx); err != nil {
			return err
		}

		pool.Pool.Namespace = ""
		msg, err := bm.sendDefinitionDefault(ctx, pool, core.SystemTagDefinePool, waitConfirm)
		if msg != nil {
			pool.Pool.Message = msg.Header.ID
		}
		pool.Pool.Namespace = bm.namespace
		return err
	}

	return fakeBatch(ctx, func(ctx context.Context, state *core.BatchState) (HandlerResult, error) {
		return bm.handler.handleTokenPoolDefinition(ctx, state, pool.Pool)
	})
}
