// SiGG-GoLang-On-the-Fly //

package namespace

import (
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/mocks/spieventsmocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
	"github.com/stretchr/testify/mock"
)

func TestMessageCreated(t *testing.T) {
	mae := &spieventsmocks.Manager{}
	nm := &namespaceManager{
		adminEvents: mae,
	}
	mae.On("Dispatch", mock.Anything).Return()
	nm.OrderedUUIDCollectionNSEvent(database.CollectionMessages, core.ChangeEventTypeCreated, "ns1", fftypes.NewUUID(), 12345)
	mae.AssertExpectations(t)
}

func TestPinCreated(t *testing.T) {
	mae := &spieventsmocks.Manager{}
	nm := &namespaceManager{
		adminEvents: mae,
	}
	mae.On("Dispatch", mock.Anything).Return()
	nm.OrderedCollectionNSEvent(database.CollectionPins, core.ChangeEventTypeCreated, "ns1", 12345)
	mae.AssertExpectations(t)
}

func TestEventCreated(t *testing.T) {
	mae := &spieventsmocks.Manager{}
	nm := &namespaceManager{
		adminEvents: mae,
	}
	mae.On("Dispatch", mock.Anything).Return()
	nm.OrderedUUIDCollectionNSEvent(database.CollectionEvents, core.ChangeEventTypeCreated, "ns1", fftypes.NewUUID(), 12345)
	mae.AssertExpectations(t)
}

func TestSubscriptionCreated(t *testing.T) {
	mae := &spieventsmocks.Manager{}
	nm := &namespaceManager{
		adminEvents: mae,
	}
	mae.On("Dispatch", mock.Anything).Return()
	nm.UUIDCollectionNSEvent(database.CollectionSubscriptions, core.ChangeEventTypeCreated, "ns1", fftypes.NewUUID())
	mae.AssertExpectations(t)
}

func TestSubscriptionUpdated(t *testing.T) {
	mae := &spieventsmocks.Manager{}
	nm := &namespaceManager{
		adminEvents: mae,
	}
	mae.On("Dispatch", mock.Anything).Return()
	nm.UUIDCollectionNSEvent(database.CollectionSubscriptions, core.ChangeEventTypeUpdated, "ns1", fftypes.NewUUID())
	mae.AssertExpectations(t)
}

func TestSubscriptionDeleted(t *testing.T) {
	mae := &spieventsmocks.Manager{}
	nm := &namespaceManager{
		adminEvents: mae,
	}
	mae.On("Dispatch", mock.Anything).Return()
	nm.UUIDCollectionNSEvent(database.CollectionSubscriptions, core.ChangeEventTypeDeleted, "ns1", fftypes.NewUUID())
	mae.AssertExpectations(t)
}

func TestHashCollectionNSEventOk(t *testing.T) {
	mae := &spieventsmocks.Manager{}
	nm := &namespaceManager{
		adminEvents: mae,
	}
	mae.On("Dispatch", mock.Anything).Return()
	nm.HashCollectionNSEvent(database.CollectionGroups, core.ChangeEventTypeDeleted, "ns1", fftypes.NewRandB32())
	mae.AssertExpectations(t)
}
