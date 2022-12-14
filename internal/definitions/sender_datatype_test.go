// SiGG-GoLang-On-the-Fly //
package definitions

import (
	"context"
	"fmt"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/mocks/broadcastmocks"
	"github.com/hyperledger/firefly/mocks/datamocks"
	"github.com/hyperledger/firefly/mocks/identitymanagermocks"
	"github.com/hyperledger/firefly/mocks/syncasyncmocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDefineDatatypeBadType(t *testing.T) {
	ds, cancel := newTestDefinitionSender(t)
	defer cancel()
	ds.multiparty = true
	err := ds.DefineDatatype(context.Background(), &core.Datatype{
		Validator: core.ValidatorType("wrong"),
	}, false)
	assert.Regexp(t, "FF00111.*validator", err)
}

func TestBroadcastDatatypeBadValue(t *testing.T) {
	ds, cancel := newTestDefinitionSender(t)
	defer cancel()
	ds.multiparty = true
	mdm := ds.data.(*datamocks.Manager)
	mdm.On("CheckDatatype", mock.Anything, mock.Anything).Return(nil)
	mim := ds.identity.(*identitymanagermocks.Manager)
	mim.On("ResolveInputSigningIdentity", mock.Anything, mock.Anything).Return(nil)
	err := ds.DefineDatatype(context.Background(), &core.Datatype{
		Namespace: "ns1",
		Name:      "ent1",
		Version:   "0.0.1",
		Value:     fftypes.JSONAnyPtr(`!unparsable`),
	}, false)
	assert.Regexp(t, "FF10137.*value", err)

	mdm.AssertExpectations(t)
	mim.AssertExpectations(t)
}

func TestDefineDatatypeInvalid(t *testing.T) {
	ds, cancel := newTestDefinitionSender(t)
	defer cancel()
	ds.multiparty = true
	mdm := ds.data.(*datamocks.Manager)
	mim := ds.identity.(*identitymanagermocks.Manager)

	mim.On("ResolveInputIdentity", mock.Anything, mock.Anything).Return(nil)
	mdm.On("CheckDatatype", mock.Anything, mock.Anything).Return(fmt.Errorf("pop"))

	err := ds.DefineDatatype(context.Background(), &core.Datatype{
		Namespace: "ns1",
		Name:      "ent1",
		Version:   "0.0.1",
		Value:     fftypes.JSONAnyPtr(`{"some": "data"}`),
	}, false)
	assert.EqualError(t, err, "pop")

	mdm.AssertExpectations(t)
}

func TestBroadcastOk(t *testing.T) {
	ds, cancel := newTestDefinitionSender(t)
	defer cancel()
	ds.multiparty = true
	mdm := ds.data.(*datamocks.Manager)
	mim := ds.identity.(*identitymanagermocks.Manager)
	mbm := ds.broadcast.(*broadcastmocks.Manager)
	mms := &syncasyncmocks.Sender{}

	mim.On("ResolveInputSigningIdentity", mock.Anything, mock.Anything).Return(nil)
	mdm.On("CheckDatatype", mock.Anything, mock.Anything).Return(nil)
	mbm.On("NewBroadcast", mock.Anything).Return(mms)
	mms.On("Send", context.Background()).Return(nil)

	err := ds.DefineDatatype(context.Background(), &core.Datatype{
		Namespace: "ns1",
		Name:      "ent1",
		Version:   "0.0.1",
		Value:     fftypes.JSONAnyPtr(`{"some": "data"}`),
	}, false)
	assert.NoError(t, err)

	mdm.AssertExpectations(t)
	mim.AssertExpectations(t)
	mbm.AssertExpectations(t)
	mms.AssertExpectations(t)
}

func TestDefineDatatypeNonMultiparty(t *testing.T) {
	ds, cancel := newTestDefinitionSender(t)
	defer cancel()
	ds.multiparty = false

	err := ds.DefineDatatype(context.Background(), &core.Datatype{
		Namespace: "ns1",
		Name:      "ent1",
		Version:   "0.0.1",
		Value:     fftypes.JSONAnyPtr(`{"some": "data"}`),
	}, false)
	assert.Regexp(t, "FF10414", err)
}
