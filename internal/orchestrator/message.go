// SiGG-GoLang-On-the-Fly //
package orchestrator

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

func (or *orchestrator) RequestReply(ctx context.Context, msg *core.MessageInOut) (reply *core.MessageInOut, err error) {
	if msg.Header.Group == nil && (msg.Group == nil || len(msg.Group.Members) == 0) {
		return nil, i18n.NewError(ctx, coremsgs.MsgRequestMustBePrivate)
	}
	return or.PrivateMessaging().RequestReply(ctx, msg)
}
