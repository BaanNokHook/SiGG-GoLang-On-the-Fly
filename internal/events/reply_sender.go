// SiGG-GoLang-On-the-Fly //
package events

import (
	"context"
	"fmt"

	"github.com/hyperledger/firefly-common/pkg/log"
	"github.com/hyperledger/firefly/pkg/core"
)

func (ed *eventDispatcher) sendReply(ctx context.Context, event *core.Event, reply *core.MessageInOut) {
	var err error
	if reply.Header.Group != nil {
		if ed.messaging == nil {
			err = fmt.Errorf("private messaging manager not initialized")
		} else {
			err = ed.messaging.NewMessage(reply).Send(ctx)
		}
	} else {
		if ed.broadcast == nil {
			err = fmt.Errorf("broadcast manager not initialized")
		} else {
			err = ed.broadcast.NewBroadcast(reply).Send(ctx)
		}
	}
	if err != nil {
		log.L(ctx).Errorf("Failed to send reply: %s", err)
	} else {
		log.L(ctx).Infof("Sent reply %s (%s) cid=%s to event '%s'", reply.Header.ID, reply.Header.Type, reply.Header.CID, event.ID)
	}
}
