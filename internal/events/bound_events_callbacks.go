// SiGG-GoLang-On-the-Fly //
package events

import (
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/events"
)

type boundCallbacks struct {
	sm *subscriptionManager
	ei events.Plugin
}

func (bc *boundCallbacks) RegisterConnection(connID string, matcher events.SubscriptionMatcher) error {
	return bc.sm.registerConnection(bc.ei, connID, matcher)
}

func (bc *boundCallbacks) EphemeralSubscription(connID, namespace string, filter *core.SubscriptionFilter, options *core.SubscriptionOptions) error {
	return bc.sm.ephemeralSubscription(bc.ei, connID, namespace, filter, options)
}

func (bc *boundCallbacks) DeliveryResponse(connID string, inflight *core.EventDeliveryResponse) {
	bc.sm.deliveryResponse(bc.ei, connID, inflight)
}

func (bc *boundCallbacks) ConnectionClosed(connID string) {
	bc.sm.connectionClosed(bc.ei, connID)
}
