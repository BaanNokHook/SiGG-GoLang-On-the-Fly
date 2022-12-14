// SiGG-GoLang-On-the-Fly //
package sqlcommon

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/hyperledger/firefly-common/pkg/dbsql"
	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"

	// Import migrate file source
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type SQLCommon struct {
	dbsql.Database
	capabilities *database.Capabilities
	callbacks    callbacks
	features     dbsql.SQLFeatures
}

type callbacks struct {
	handlers map[string]database.Callbacks
}

func (cb *callbacks) OrderedUUIDCollectionNSEvent(resType database.OrderedUUIDCollectionNS, eventType core.ChangeEventType, ns string, id *fftypes.UUID, sequence int64) {
	if cb, ok := cb.handlers[ns]; ok {
		cb.OrderedUUIDCollectionNSEvent(resType, eventType, ns, id, sequence)
	}
	if cb, ok := cb.handlers[database.GlobalHandler]; ok {
		cb.OrderedUUIDCollectionNSEvent(resType, eventType, ns, id, sequence)
	}
}

func (cb *callbacks) OrderedCollectionNSEvent(resType database.OrderedCollectionNS, eventType core.ChangeEventType, ns string, sequence int64) {
	if cb, ok := cb.handlers[ns]; ok {
		cb.OrderedCollectionNSEvent(resType, eventType, ns, sequence)
	}
	if cb, ok := cb.handlers[database.GlobalHandler]; ok {
		cb.OrderedCollectionNSEvent(resType, eventType, ns, sequence)
	}
}

func (cb *callbacks) UUIDCollectionNSEvent(resType database.UUIDCollectionNS, eventType core.ChangeEventType, ns string, id *fftypes.UUID) {
	if cb, ok := cb.handlers[ns]; ok {
		cb.UUIDCollectionNSEvent(resType, eventType, ns, id)
	}
	if cb, ok := cb.handlers[database.GlobalHandler]; ok {
		cb.UUIDCollectionNSEvent(resType, eventType, ns, id)
	}
}

func (cb *callbacks) HashCollectionNSEvent(resType database.HashCollectionNS, eventType core.ChangeEventType, ns string, hash *fftypes.Bytes32) {
	if cb, ok := cb.handlers[ns]; ok {
		cb.HashCollectionNSEvent(resType, eventType, ns, hash)
	}
	if cb, ok := cb.handlers[database.GlobalHandler]; ok {
		cb.HashCollectionNSEvent(resType, eventType, ns, hash)
	}
}

func (s *SQLCommon) Init(ctx context.Context, provider dbsql.Provider, config config.Section, capabilities *database.Capabilities) (err error) {
	s.capabilities = capabilities
	return s.Database.Init(ctx, provider, config)
}

func (s *SQLCommon) SetHandler(namespace string, handler database.Callbacks) {
	if s.callbacks.handlers == nil {
		s.callbacks.handlers = make(map[string]database.Callbacks)
	}
	s.callbacks.handlers[namespace] = handler
}

func (s *SQLCommon) Capabilities() *database.Capabilities { return s.capabilities }
