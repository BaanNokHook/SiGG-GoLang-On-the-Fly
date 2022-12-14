// SiGG-GoLang-On-the-Fly //

package definitions

import (
	"context"
	"fmt"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/mocks/broadcastmocks"
	"github.com/hyperledger/firefly/mocks/databasemocks"
	"github.com/hyperledger/firefly/mocks/datamocks"
	"github.com/hyperledger/firefly/mocks/identitymanagermocks"
	"github.com/hyperledger/firefly/mocks/syncasyncmocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBroadcastTokenPoolInvalid(t *testing.T) {
	ds, cancel := newTestDefinitionSender(t)
	defer cancel()
	ds.multiparty = true

	mdm := ds.data.(*datamocks.Manager)

	pool := &core.TokenPoolAnnouncement{
		Pool: &core.TokenPool{
			ID:        fftypes.NewUUID(),
			Namespace: "",
			Name:      "",
			Type:      core.TokenTypeNonFungible,
			Locator:   "N1",
			Symbol:    "COIN",
		},
	}

	err := ds.DefineTokenPool(context.Background(), pool, false)
	assert.Regexp(t, "FF10420", err)

	mdm.AssertExpectations(t)
}

func TestBroadcastTokenPoolInvalidNonMultiparty(t *testing.T) {
	ds, cancel := newTestDefinitionSender(t)
	defer cancel()
	ds.multiparty = false

	mdm := ds.data.(*datamocks.Manager)

	pool := &core.TokenPoolAnnouncement{
		Pool: &core.TokenPool{
			ID:        fftypes.NewUUID(),
			Namespace: "",
			Name:      "",
			Type:      core.TokenTypeNonFungible,
			Locator:   "N1",
			Symbol:    "COIN",
		},
	}

	err := ds.DefineTokenPool(context.Background(), pool, false)
	assert.Regexp(t, "FF10420", err)

	mdm.AssertExpectations(t)
}

func TestDefineTokenPoolOk(t *testing.T) {
	ds, cancel := newTestDefinitionSender(t)
	defer cancel()
	ds.multiparty = true

	mdm := ds.data.(*datamocks.Manager)
	mim := ds.identity.(*identitymanagermocks.Manager)
	mbm := ds.broadcast.(*broadcastmocks.Manager)
	mms := &syncasyncmocks.Sender{}

	pool := &core.TokenPoolAnnouncement{
		Pool: &core.TokenPool{
			ID:        fftypes.NewUUID(),
			Namespace: "ns1",
			Name:      "mypool",
			Type:      core.TokenTypeNonFungible,
			Locator:   "N1",
			Symbol:    "COIN",
			Connector: "connector1",
		},
	}

	mim.On("ResolveInputSigningIdentity", mock.Anything, mock.Anything).Return(nil)
	mbm.On("NewBroadcast", mock.Anything).Return(mms)
	mms.On("Send", context.Background()).Return(nil)

	err := ds.DefineTokenPool(context.Background(), pool, false)
	assert.NoError(t, err)

	mdm.AssertExpectations(t)
	mim.AssertExpectations(t)
	mbm.AssertExpectations(t)
	mms.AssertExpectations(t)
}

func TestDefineTokenPoolNonMultipartyTokenPoolFail(t *testing.T) {
	ds, cancel := newTestDefinitionSender(t)
	defer cancel()

	mdm := ds.data.(*datamocks.Manager)
	mbm := ds.broadcast.(*broadcastmocks.Manager)
	mdi := ds.database.(*databasemocks.Plugin)

	pool := &core.TokenPoolAnnouncement{
		Pool: &core.TokenPool{
			ID:        fftypes.NewUUID(),
			Namespace: "ns1",
			Name:      "mypool",
			Type:      core.TokenTypeNonFungible,
			Locator:   "N1",
			Symbol:    "COIN",
			Connector: "connector1",
		},
	}

	mdi.On("GetTokenPoolByID", context.Background(), "ns1", pool.Pool.ID).Return(nil, fmt.Errorf("pop"))

	err := ds.DefineTokenPool(context.Background(), pool, false)
	assert.Regexp(t, "pop", err)

	mdm.AssertExpectations(t)
	mbm.AssertExpectations(t)
}

func TestDefineTokenPoolBadName(t *testing.T) {
	ds, cancel := newTestDefinitionSender(t)
	defer cancel()
	ds.multiparty = true

	mim := ds.identity.(*identitymanagermocks.Manager)
	mbm := ds.broadcast.(*broadcastmocks.Manager)
	mms := &syncasyncmocks.Sender{}

	pool := &core.TokenPoolAnnouncement{
		Pool: &core.TokenPool{
			ID:        fftypes.NewUUID(),
			Namespace: "ns1",
			Name:      "///bad/////",
			Type:      core.TokenTypeNonFungible,
			Locator:   "N1",
			Symbol:    "COIN",
			Connector: "connector1",
		},
	}

	mim.On("ResolveInputSigningIdentity", mock.Anything, mock.Anything).Return(nil)
	mbm.On("NewBroadcast", mock.Anything).Return(mms)
	mms.On("Send", context.Background()).Return(nil)

	err := ds.DefineTokenPool(context.Background(), pool, false)
	assert.Regexp(t, "FF00140", err)
}
