// SiGG-GoLang-On-the-Fly //
package events

import (
	"context"
	"sync"

	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/hyperledger/firefly-common/pkg/log"
	"github.com/hyperledger/firefly/internal/coremsgs"
)

type eventNotifier struct {
	ctx            context.Context
	desc           string
	newEvents      chan int64
	latestSequence int64
	cond           *sync.Cond
	closed         bool
}

func newEventNotifier(ctx context.Context, desc string) *eventNotifier {
	mux := &sync.Mutex{}
	en := &eventNotifier{
		ctx:            ctx,
		newEvents:      make(chan int64),
		latestSequence: -1,
		cond:           sync.NewCond(mux),
		desc:           desc,
	}
	go en.newEventLoop()
	return en
}

func (en *eventNotifier) waitNext(lastSequence int64) error {
	var seq int64
	en.cond.L.Lock()
	closed := en.closed
	for en.latestSequence <= lastSequence && !en.closed {
		en.cond.Wait()
	}
	seq = en.latestSequence
	en.cond.L.Unlock()
	if closed {
		return i18n.NewError(en.ctx, coremsgs.MsgEventListenerClosing)
	}
	log.L(en.ctx).Tracef("Detected new %s (%d)", en.desc, seq)
	return nil
}

func (en *eventNotifier) close() {
	en.cond.L.Lock()
	en.closed = true
	en.cond.Broadcast()
	en.cond.L.Unlock()
}

func (en *eventNotifier) newEventLoop() {
	l := log.L(en.ctx)
	defer en.close()
	for {
		select {
		case <-en.ctx.Done():
			l.Debugf("New event notifier loop ending (context cancelled)")
			return
		case seq, ok := <-en.newEvents:
			if !ok {
				l.Debugf("New event notifier loop ending (closed channel)")
				return
			}
			log.L(en.ctx).Tracef("Notifying new %s %d", en.desc, seq)
			en.cond.L.Lock()
			en.latestSequence = seq
			en.cond.Broadcast()
			en.cond.L.Unlock()
		}
	}
}
