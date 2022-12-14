// SiGG-GoLang-On-the-Fly //

package events

import (
	"context"
	"fmt"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/mocks/broadcastmocks"
	"github.com/hyperledger/firefly/mocks/privatemessagingmocks"
	"github.com/hyperledger/firefly/mocks/syncasyncmocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/mock"
)

func TestSendReplyBroadcastFail(t *testing.T) {
	ed, cancel := newTestEventDispatcher(&subscription{
		definition: &core.Subscription{},
	})
	defer cancel()

	mms := &syncasyncmocks.Sender{}
	mbm := ed.broadcast.(*broadcastmocks.Manager)
	mbm.On("NewBroadcast", mock.Anything).Return(mms)
	mms.On("Send", context.Background()).Return(fmt.Errorf("pop"))

	ed.sendReply(context.Background(), &core.Event{
		ID:        fftypes.NewUUID(),
		Namespace: "ns1",
	}, &core.MessageInOut{})

	mbm.AssertExpectations(t)
	mms.AssertExpectations(t)
}

func TestSendReplyPrivateFail(t *testing.T) {
	ed, cancel := newTestEventDispatcher(&subscription{
		definition: &core.Subscription{},
	})
	defer cancel()

	mms := &syncasyncmocks.Sender{}
	mpm := ed.messaging.(*privatemessagingmocks.Manager)
	mpm.On("NewMessage", mock.Anything).Return(mms)
	mms.On("Send", context.Background()).Return(fmt.Errorf("pop"))

	ed.sendReply(context.Background(), &core.Event{
		ID:        fftypes.NewUUID(),
		Namespace: "ns1",
	}, &core.MessageInOut{
		Message: core.Message{
			Header: core.MessageHeader{
				Group: fftypes.NewRandB32(),
			},
		},
	})

	mpm.AssertExpectations(t)
	mms.AssertExpectations(t)
}

func TestSendReplyPrivateOk(t *testing.T) {
	ed, cancel := newTestEventDispatcher(&subscription{
		definition: &core.Subscription{},
	})
	defer cancel()

	msg := &core.Message{
		Header: core.MessageHeader{
			Group: fftypes.NewRandB32(),
		},
	}

	mms := &syncasyncmocks.Sender{}
	mpm := ed.messaging.(*privatemessagingmocks.Manager)
	mpm.On("NewMessage", mock.Anything).Return(mms)
	mms.On("Send", context.Background()).Return(nil)

	ed.sendReply(context.Background(), &core.Event{
		ID:        fftypes.NewUUID(),
		Namespace: "ns1",
	}, &core.MessageInOut{
		Message: *msg,
	})

	mpm.AssertExpectations(t)
	mms.AssertExpectations(t)
}

func TestSendReplyBroadcastDisabled(t *testing.T) {
	ed, cancel := newTestEventDispatcher(&subscription{
		definition: &core.Subscription{},
	})
	defer cancel()
	ed.broadcast = nil

	ed.sendReply(context.Background(), &core.Event{
		ID:        fftypes.NewUUID(),
		Namespace: "ns1",
	}, &core.MessageInOut{})
}

func TestSendReplyPrivateDisabled(t *testing.T) {
	ed, cancel := newTestEventDispatcher(&subscription{
		definition: &core.Subscription{},
	})
	defer cancel()
	ed.messaging = nil

	msg := &core.Message{
		Header: core.MessageHeader{
			Group: fftypes.NewRandB32(),
		},
	}

	ed.sendReply(context.Background(), &core.Event{
		ID:        fftypes.NewUUID(),
		Namespace: "ns1",
	}, &core.MessageInOut{
		Message: *msg,
	})
}
