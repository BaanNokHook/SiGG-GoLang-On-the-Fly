// SiGG-GoLang-On-the-Fly //

package definitions

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/mocks/databasemocks"
	"github.com/hyperledger/firefly/mocks/datamocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleDefinitionBroadcastDatatypeOk(t *testing.T) {
	dh, bs := newTestDefinitionHandler(t)

	dt := &core.Datatype{
		ID:        fftypes.NewUUID(),
		Validator: core.ValidatorTypeJSON,
		Namespace: "ns1",
		Name:      "name1",
		Version:   "ver1",
		Value:     fftypes.JSONAnyPtr(`{}`),
	}
	dt.Hash = dt.Value.Hash()
	b, err := json.Marshal(&dt)
	assert.NoError(t, err)
	data := &core.Data{
		Value: fftypes.JSONAnyPtrBytes(b),
	}

	mdm := dh.data.(*datamocks.Manager)
	mdm.On("CheckDatatype", mock.Anything, mock.Anything).Return(nil)
	mbi := dh.database.(*databasemocks.Plugin)
	mbi.On("GetDatatypeByName", mock.Anything, "ns1", "name1", "ver1").Return(nil, nil)
	mbi.On("UpsertDatatype", mock.Anything, mock.Anything, false).Return(nil)
	mbi.On("InsertEvent", mock.Anything, mock.Anything).Return(nil)
	action, err := dh.HandleDefinitionBroadcast(context.Background(), &bs.BatchState, &core.Message{
		Header: core.MessageHeader{
			Tag: core.SystemTagDefineDatatype,
		},
	}, core.DataArray{data}, fftypes.NewUUID())
	assert.Equal(t, HandlerResult{Action: ActionConfirm}, action)
	assert.NoError(t, err)
	err = bs.RunFinalize(context.Background())
	assert.NoError(t, err)

	mdm.AssertExpectations(t)
	mbi.AssertExpectations(t)
}

func TestHandleDefinitionBroadcastDatatypeEventFail(t *testing.T) {
	dh, bs := newTestDefinitionHandler(t)

	dt := &core.Datatype{
		ID:        fftypes.NewUUID(),
		Validator: core.ValidatorTypeJSON,
		Namespace: "ns1",
		Name:      "name1",
		Version:   "ver1",
		Value:     fftypes.JSONAnyPtr(`{}`),
	}
	dt.Hash = dt.Value.Hash()
	b, err := json.Marshal(&dt)
	assert.NoError(t, err)
	data := &core.Data{
		Value: fftypes.JSONAnyPtrBytes(b),
	}

	mdm := dh.data.(*datamocks.Manager)
	mdm.On("CheckDatatype", mock.Anything, mock.Anything).Return(nil)
	mbi := dh.database.(*databasemocks.Plugin)
	mbi.On("GetDatatypeByName", mock.Anything, "ns1", "name1", "ver1").Return(nil, nil)
	mbi.On("UpsertDatatype", mock.Anything, mock.Anything, false).Return(nil)
	mbi.On("InsertEvent", mock.Anything, mock.Anything).Return(fmt.Errorf("pop"))
	action, err := dh.HandleDefinitionBroadcast(context.Background(), &bs.BatchState, &core.Message{
		Header: core.MessageHeader{
			Tag: core.SystemTagDefineDatatype,
		},
	}, core.DataArray{data}, fftypes.NewUUID())
	assert.Equal(t, HandlerResult{Action: ActionConfirm}, action)
	assert.NoError(t, err)
	err = bs.RunFinalize(context.Background())
	assert.EqualError(t, err, "pop")

	mdm.AssertExpectations(t)
	mbi.AssertExpectations(t)
}

func TestHandleDefinitionBroadcastDatatypeMissingID(t *testing.T) {
	dh, bs := newTestDefinitionHandler(t)

	dt := &core.Datatype{
		Validator: core.ValidatorTypeJSON,
		Namespace: "ns1",
		Name:      "name1",
		Version:   "ver1",
		Value:     fftypes.JSONAnyPtr(`{}`),
	}
	dt.Hash = dt.Value.Hash()
	b, err := json.Marshal(&dt)
	assert.NoError(t, err)
	data := &core.Data{
		Value: fftypes.JSONAnyPtrBytes(b),
	}

	action, err := dh.HandleDefinitionBroadcast(context.Background(), &bs.BatchState, &core.Message{
		Header: core.MessageHeader{
			Tag: core.SystemTagDefineDatatype,
		},
	}, core.DataArray{data}, fftypes.NewUUID())
	assert.Equal(t, HandlerResult{Action: ActionReject}, action)
	assert.Error(t, err)
	bs.assertNoFinalizers()
}

func TestHandleDefinitionBroadcastBadSchema(t *testing.T) {
	dh, bs := newTestDefinitionHandler(t)

	dt := &core.Datatype{
		ID:        fftypes.NewUUID(),
		Validator: core.ValidatorTypeJSON,
		Namespace: "ns1",
		Name:      "name1",
		Version:   "ver1",
		Value:     fftypes.JSONAnyPtr(`{}`),
	}
	dt.Hash = dt.Value.Hash()
	b, err := json.Marshal(&dt)
	assert.NoError(t, err)
	data := &core.Data{
		Value: fftypes.JSONAnyPtrBytes(b),
	}

	mdm := dh.data.(*datamocks.Manager)
	mdm.On("CheckDatatype", mock.Anything, mock.Anything).Return(fmt.Errorf("pop"))
	action, err := dh.HandleDefinitionBroadcast(context.Background(), &bs.BatchState, &core.Message{
		Header: core.MessageHeader{
			Tag: core.SystemTagDefineDatatype,
		},
	}, core.DataArray{data}, fftypes.NewUUID())
	assert.Equal(t, HandlerResult{Action: ActionReject}, action)
	assert.Error(t, err)

	mdm.AssertExpectations(t)
	bs.assertNoFinalizers()
}

func TestHandleDefinitionBroadcastMissingData(t *testing.T) {
	dh, bs := newTestDefinitionHandler(t)

	dt := &core.Datatype{
		ID:        fftypes.NewUUID(),
		Validator: core.ValidatorTypeJSON,
		Namespace: "ns1",
		Name:      "name1",
		Version:   "ver1",
		Value:     fftypes.JSONAnyPtr(`{}`),
	}
	dt.Hash = dt.Value.Hash()

	action, err := dh.HandleDefinitionBroadcast(context.Background(), &bs.BatchState, &core.Message{
		Header: core.MessageHeader{
			Tag: core.SystemTagDefineDatatype,
		},
	}, core.DataArray{}, fftypes.NewUUID())
	assert.Equal(t, HandlerResult{Action: ActionReject}, action)
	assert.Error(t, err)
	bs.assertNoFinalizers()
}

func TestHandleDefinitionBroadcastDatatypeLookupFail(t *testing.T) {
	dh, bs := newTestDefinitionHandler(t)

	dt := &core.Datatype{
		ID:        fftypes.NewUUID(),
		Validator: core.ValidatorTypeJSON,
		Namespace: "ns1",
		Name:      "name1",
		Version:   "ver1",
		Value:     fftypes.JSONAnyPtr(`{}`),
	}
	dt.Hash = dt.Value.Hash()
	b, err := json.Marshal(&dt)
	assert.NoError(t, err)
	data := &core.Data{
		Value: fftypes.JSONAnyPtrBytes(b),
	}

	mdm := dh.data.(*datamocks.Manager)
	mdm.On("CheckDatatype", mock.Anything, mock.Anything).Return(nil)
	mbi := dh.database.(*databasemocks.Plugin)
	mbi.On("GetDatatypeByName", mock.Anything, "ns1", "name1", "ver1").Return(nil, fmt.Errorf("pop"))
	action, err := dh.HandleDefinitionBroadcast(context.Background(), &bs.BatchState, &core.Message{
		Header: core.MessageHeader{
			Namespace: "ns1",
			Tag:       core.SystemTagDefineDatatype,
		},
	}, core.DataArray{data}, fftypes.NewUUID())
	assert.Equal(t, HandlerResult{Action: ActionRetry}, action)
	assert.EqualError(t, err, "pop")

	mdm.AssertExpectations(t)
	mbi.AssertExpectations(t)
	bs.assertNoFinalizers()
}

func TestHandleDefinitionBroadcastUpsertFail(t *testing.T) {
	dh, bs := newTestDefinitionHandler(t)

	dt := &core.Datatype{
		ID:        fftypes.NewUUID(),
		Validator: core.ValidatorTypeJSON,
		Namespace: "ns1",
		Name:      "name1",
		Version:   "ver1",
		Value:     fftypes.JSONAnyPtr(`{}`),
	}
	dt.Hash = dt.Value.Hash()
	b, err := json.Marshal(&dt)
	assert.NoError(t, err)
	data := &core.Data{
		Value: fftypes.JSONAnyPtrBytes(b),
	}

	mdm := dh.data.(*datamocks.Manager)
	mdm.On("CheckDatatype", mock.Anything, mock.Anything).Return(nil)
	mbi := dh.database.(*databasemocks.Plugin)
	mbi.On("GetDatatypeByName", mock.Anything, "ns1", "name1", "ver1").Return(nil, nil)
	mbi.On("UpsertDatatype", mock.Anything, mock.Anything, false).Return(fmt.Errorf("pop"))
	action, err := dh.HandleDefinitionBroadcast(context.Background(), &bs.BatchState, &core.Message{
		Header: core.MessageHeader{
			Tag: core.SystemTagDefineDatatype,
		},
	}, core.DataArray{data}, fftypes.NewUUID())
	assert.Equal(t, HandlerResult{Action: ActionRetry}, action)
	assert.EqualError(t, err, "pop")

	mdm.AssertExpectations(t)
	mbi.AssertExpectations(t)
	bs.assertNoFinalizers()
}

func TestHandleDefinitionBroadcastDatatypeDuplicate(t *testing.T) {
	dh, bs := newTestDefinitionHandler(t)

	dt := &core.Datatype{
		ID:        fftypes.NewUUID(),
		Validator: core.ValidatorTypeJSON,
		Namespace: "ns1",
		Name:      "name1",
		Version:   "ver1",
		Value:     fftypes.JSONAnyPtr(`{}`),
	}
	dt.Hash = dt.Value.Hash()
	b, err := json.Marshal(&dt)
	assert.NoError(t, err)
	data := &core.Data{
		Value: fftypes.JSONAnyPtrBytes(b),
	}

	mdm := dh.data.(*datamocks.Manager)
	mdm.On("CheckDatatype", mock.Anything, mock.Anything).Return(nil)
	mbi := dh.database.(*databasemocks.Plugin)
	mbi.On("GetDatatypeByName", mock.Anything, "ns1", "name1", "ver1").Return(dt, nil)
	action, err := dh.HandleDefinitionBroadcast(context.Background(), &bs.BatchState, &core.Message{
		Header: core.MessageHeader{
			Tag: core.SystemTagDefineDatatype,
		},
	}, core.DataArray{data}, fftypes.NewUUID())
	assert.Equal(t, HandlerResult{Action: ActionReject}, action)
	assert.Error(t, err)

	mdm.AssertExpectations(t)
	mbi.AssertExpectations(t)
	bs.assertNoFinalizers()
}
