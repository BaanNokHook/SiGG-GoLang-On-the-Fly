// SiGG-GoLang-On-the-Fly //
package orchestrator

import (
	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly-common/pkg/log"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

func (or *orchestrator) OrderedUUIDCollectionNSEvent(resType database.OrderedUUIDCollectionNS, eventType core.ChangeEventType, ns string, id *fftypes.UUID, sequence int64) {
	if ns != or.namespace.Name {
		log.L(or.ctx).Debugf("Ignoring database event from different namespace '%s'", ns)
		return
	}
	switch {
	case eventType == core.ChangeEventTypeCreated && resType == database.CollectionMessages:
		or.batch.NewMessages() <- sequence
	case eventType == core.ChangeEventTypeCreated && resType == database.CollectionEvents:
		or.events.NewEvents() <- sequence
	}
}

func (or *orchestrator) OrderedCollectionNSEvent(resType database.OrderedCollectionNS, eventType core.ChangeEventType, ns string, sequence int64) {
	if ns != or.namespace.Name {
		log.L(or.ctx).Debugf("Ignoring database event from different namespace '%s'", ns)
		return
	}
	if eventType == core.ChangeEventTypeCreated && resType == database.CollectionPins {
		or.events.NewPins() <- sequence
	}
}

func (or *orchestrator) UUIDCollectionNSEvent(resType database.UUIDCollectionNS, eventType core.ChangeEventType, ns string, id *fftypes.UUID) {
	if ns != or.namespace.Name {
		log.L(or.ctx).Debugf("Ignoring database event from different namespace '%s'", ns)
		return
	}
	switch {
	case eventType == core.ChangeEventTypeCreated && resType == database.CollectionSubscriptions:
		or.events.NewSubscriptions() <- id
	case eventType == core.ChangeEventTypeDeleted && resType == database.CollectionSubscriptions:
		or.events.DeletedSubscriptions() <- id
	case eventType == core.ChangeEventTypeUpdated && resType == database.CollectionSubscriptions:
		or.events.SubscriptionUpdates() <- id
	}
}

func (or *orchestrator) HashCollectionNSEvent(resType database.HashCollectionNS, eventType core.ChangeEventType, ns string, hash *fftypes.Bytes32) {
	// do nothing
}
