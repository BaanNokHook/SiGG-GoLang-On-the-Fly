// SiGG-GoLang-On-the-Fly //

package events

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPersistBatch(t *testing.T) {

	em := newTestEventManager(t)
	defer em.cleanup(t)

	em.mdi.On("UpsertBatch", em.ctx, mock.Anything).Return(fmt.Errorf(("pop")))

	org := newTestOrg("org1")
	orgBytes, err := json.Marshal(&org)
	assert.NoError(t, err)
	data := &core.Data{
		ID:        fftypes.NewUUID(),
		Value:     fftypes.JSONAnyPtrBytes(orgBytes),
		Validator: core.MessageTypeDefinition,
	}

	batch := &core.Batch{
		BatchHeader: core.BatchHeader{
			ID: fftypes.NewUUID(),
			SignerRef: core.SignerRef{
				Author: "did:firefly:org/12345",
				Key:    "0x12345",
			},
		},
		Payload: core.BatchPayload{
			TX: core.TransactionRef{
				ID:   fftypes.NewUUID(),
				Type: core.TransactionTypeBatchPin,
			},
			Messages: []*core.Message{
				{
					Header: core.MessageHeader{
						ID:   fftypes.NewUUID(),
						Type: core.MessageTypeDefinition,
						SignerRef: core.SignerRef{
							Author: "did:firefly:org/12345",
							Key:    "0x12345",
						},
					},
					Data: core.DataRefs{
						{
							ID:   data.ID,
							Hash: data.Hash,
						},
					},
				},
			},
			Data: core.DataArray{
				data,
			},
		},
	}
	bp, _ := batch.Confirmed()
	batch.Hash = fftypes.HashString(bp.Manifest.String())

	_, _, err = em.persistBatch(em.ctx, batch)
	assert.EqualError(t, err, "pop") // Confirms we got to upserting the batch

}

func TestPersistBatchNoCacheDataNotInBatch(t *testing.T) {

	em := newTestEventManager(t)
	defer em.cleanup(t)

	em.mdi.On("UpsertBatch", em.ctx, mock.Anything).Return(nil)

	data := &core.Data{ID: fftypes.NewUUID(), Value: fftypes.JSONAnyPtr(`"test"`)}
	batch := sampleBatch(t, core.BatchTypeBroadcast, core.TransactionTypeBatchPin, core.DataArray{data})
	data.ID = fftypes.NewUUID()
	_ = data.Seal(em.ctx, nil)
	bp, _ := batch.Confirmed()
	batch.Hash = fftypes.HashString(bp.Manifest.String())

	_, valid, err := em.persistBatch(em.ctx, batch)
	assert.False(t, valid)
	assert.NoError(t, err)

}

func TestPersistBatchExtraDataInBatch(t *testing.T) {

	em := newTestEventManager(t)
	defer em.cleanup(t)

	em.mdi.On("UpsertBatch", em.ctx, mock.Anything).Return(nil)

	data := &core.Data{ID: fftypes.NewUUID(), Value: fftypes.JSONAnyPtr(`"test"`)}
	batch := sampleBatch(t, core.BatchTypeBroadcast, core.TransactionTypeBatchPin, core.DataArray{data})
	data2 := &core.Data{ID: fftypes.NewUUID(), Value: fftypes.JSONAnyPtr(`"test2"`)}
	_ = data2.Seal(em.ctx, nil)
	batch.Payload.Data = append(batch.Payload.Data, data2)
	bp, _ := batch.Confirmed()
	batch.Hash = fftypes.HashString(bp.Manifest.String())

	_, valid, err := em.persistBatch(em.ctx, batch)
	assert.False(t, valid)
	assert.NoError(t, err)

}

func TestPersistBatchNilMessageEntryop(t *testing.T) {

	em := newTestEventManager(t)
	defer em.cleanup(t)

	valid := em.validateBatchMessage(em.ctx, &core.Batch{}, 0, nil)
	assert.False(t, valid)

}

func TestPersistBatchContentSendByUsOK(t *testing.T) {

	em := newTestEventManager(t)
	defer em.cleanup(t)

	data := &core.Data{ID: fftypes.NewUUID(), Value: fftypes.JSONAnyPtr(`"test"`)}
	batch := sampleBatch(t, core.BatchTypeBroadcast, core.TransactionTypeBatchPin, core.DataArray{data})
	batch.Node = testNodeID

	em.mdm.On("GetMessageWithDataCached", em.ctx, batch.Payload.Messages[0].Header.ID).Return(batch.Payload.Messages[0], batch.Payload.Data, true, nil)

	em.mim.On("GetLocalNode", mock.Anything).Return(testNode, nil)

	ok, err := em.persistBatchContent(em.ctx, batch, []*messageAndData{})
	assert.NoError(t, err)
	assert.True(t, ok)

}

func TestPersistBatchContentSentByNil(t *testing.T) {

	em := newTestEventManager(t)
	defer em.cleanup(t)

	data := &core.Data{ID: fftypes.NewUUID(), Value: fftypes.JSONAnyPtr(`"test"`)}
	batch := sampleBatch(t, core.BatchTypeBroadcast, core.TransactionTypeBatchPin, core.DataArray{data})
	batch.Node = nil

	em.mdi.On("InsertDataArray", mock.Anything, mock.Anything).Return(nil)
	em.mdi.On("InsertMessages", em.ctx, mock.Anything, mock.AnythingOfType("database.PostCompletionHook")).Return(nil, nil).Run(func(args mock.Arguments) {
		args[2].(database.PostCompletionHook)()
	})

	em.mim.On("GetLocalNode", mock.Anything).Return(nil, nil)

	ok, err := em.persistBatchContent(em.ctx, batch, []*messageAndData{})
	assert.NoError(t, err)
	assert.True(t, ok)

}

func TestPersistBatchContentSentByUsNotFoundFallback(t *testing.T) {

	em := newTestEventManager(t)
	defer em.cleanup(t)

	data := &core.Data{ID: fftypes.NewUUID(), Value: fftypes.JSONAnyPtr(`"test"`)}
	batch := sampleBatch(t, core.BatchTypeBroadcast, core.TransactionTypeBatchPin, core.DataArray{data})
	batch.Node = testNodeID

	em.mdm.On("GetMessageWithDataCached", em.ctx, batch.Payload.Messages[0].Header.ID).Return(nil, nil, false, nil)

	em.mdi.On("InsertDataArray", mock.Anything, mock.Anything).Return(nil)
	em.mdi.On("InsertMessages", em.ctx, mock.Anything, mock.AnythingOfType("database.PostCompletionHook")).Return(nil, nil).Run(func(args mock.Arguments) {
		args[2].(database.PostCompletionHook)()
	})

	em.mim.On("GetLocalNode", mock.Anything).Return(testNode, nil)

	ok, err := em.persistBatchContent(em.ctx, batch, []*messageAndData{})
	assert.NoError(t, err)
	assert.True(t, ok)

}

func TestPersistBatchContentSentByUsFoundMismatch(t *testing.T) {

	em := newTestEventManager(t)
	defer em.cleanup(t)

	data := &core.Data{ID: fftypes.NewUUID(), Value: fftypes.JSONAnyPtr(`"test"`)}
	batch := sampleBatch(t, core.BatchTypeBroadcast, core.TransactionTypeBatchPin, core.DataArray{data})
	batch.Node = testNodeID

	em.mdm.On("GetMessageWithDataCached", em.ctx, batch.Payload.Messages[0].Header.ID).Return(&core.Message{
		Header: core.MessageHeader{
			ID: fftypes.NewUUID(),
		},
	}, nil, true, nil)

	em.mdi.On("InsertDataArray", mock.Anything, mock.Anything).Return(nil)
	em.mdi.On("InsertMessages", mock.Anything, mock.Anything, mock.AnythingOfType("database.PostCompletionHook")).Return(fmt.Errorf("optimization miss"))
	em.mdi.On("UpsertMessage", mock.Anything, mock.Anything, database.UpsertOptimizationExisting, mock.AnythingOfType("database.PostCompletionHook")).Return(database.HashMismatch)

	em.mim.On("GetLocalNode", mock.Anything).Return(testNode, nil)

	ok, err := em.persistBatchContent(em.ctx, batch, []*messageAndData{})
	assert.NoError(t, err)
	assert.False(t, ok)

}

func TestPersistBatchContentInsertMessagesFail(t *testing.T) {

	em := newTestEventManager(t)
	defer em.cleanup(t)

	data := &core.Data{ID: fftypes.NewUUID(), Value: fftypes.JSONAnyPtr(`"test"`)}
	batch := sampleBatch(t, core.BatchTypeBroadcast, core.TransactionTypeBatchPin, core.DataArray{data})

	em.mdi.On("InsertDataArray", mock.Anything, mock.Anything).Return(nil)
	em.mdi.On("InsertMessages", mock.Anything, mock.Anything, mock.AnythingOfType("database.PostCompletionHook")).Return(fmt.Errorf("optimization miss"))
	em.mdi.On("UpsertMessage", mock.Anything, mock.Anything, database.UpsertOptimizationExisting, mock.AnythingOfType("database.PostCompletionHook")).Return(nil).Run(func(args mock.Arguments) {
		args[3].(database.PostCompletionHook)()
	})

	msgData := &messageAndData{
		message: batch.Payload.Messages[0],
		data:    batch.Payload.Data,
	}
	em.mdm.On("UpdateMessageCache", msgData.message, msgData.data).Return()

	em.mim.On("GetLocalNode", mock.Anything).Return(testNode, nil)

	ok, err := em.persistBatchContent(em.ctx, batch, []*messageAndData{msgData})
	assert.NoError(t, err)
	assert.True(t, ok)

}

func TestPersistBatchContentSentByUsFoundError(t *testing.T) {

	em := newTestEventManager(t)
	defer em.cleanup(t)

	data := &core.Data{ID: fftypes.NewUUID(), Value: fftypes.JSONAnyPtr(`"test"`)}
	batch := sampleBatch(t, core.BatchTypeBroadcast, core.TransactionTypeBatchPin, core.DataArray{data})
	batch.Node = testNodeID

	em.mdm.On("GetMessageWithDataCached", em.ctx, batch.Payload.Messages[0].Header.ID).Return(nil, nil, false, fmt.Errorf("pop"))

	em.mim.On("GetLocalNode", mock.Anything).Return(testNode, nil)

	ok, err := em.persistBatchContent(em.ctx, batch, []*messageAndData{})
	assert.Regexp(t, "pop", err)
	assert.False(t, ok)

}

func TestPersistBatchContentDataHashMismatch(t *testing.T) {

	em := newTestEventManager(t)
	defer em.cleanup(t)

	data := &core.Data{ID: fftypes.NewUUID(), Value: fftypes.JSONAnyPtr(`"test"`)}
	batch := sampleBatch(t, core.BatchTypeBroadcast, core.TransactionTypeBatchPin, core.DataArray{data})

	em.mdi.On("InsertDataArray", mock.Anything, mock.Anything).Return(fmt.Errorf("optimization miss"))
	em.mdi.On("UpsertData", mock.Anything, mock.Anything, database.UpsertOptimizationExisting).Return(database.HashMismatch)

	em.mim.On("GetLocalNode", mock.Anything).Return(testNode, nil)

	ok, err := em.persistBatchContent(em.ctx, batch, []*messageAndData{})
	assert.NoError(t, err)
	assert.False(t, ok)

}

func TestPersistBatchContentDataMissingBlobRef(t *testing.T) {

	em := newTestEventManager(t)
	defer em.cleanup(t)

	blob := &core.Blob{
		Hash: fftypes.NewRandB32(),
	}
	data := &core.Data{ID: fftypes.NewUUID(), Value: fftypes.JSONAnyPtr(`"test"`), Blob: &core.BlobRef{
		Hash: blob.Hash,
	}}
	batch := sampleBatch(t, core.BatchTypeBroadcast, core.TransactionTypeBatchPin, core.DataArray{data}, blob)

	em.mdi.On("GetBlobMatchingHash", mock.Anything, mock.Anything).Return(nil, nil)

	ok, err := em.validateAndPersistBatchContent(em.ctx, batch)
	assert.NoError(t, err)
	assert.False(t, ok)

}

func TestPersistBatchInvalidTXType(t *testing.T) {

	em := newTestEventManager(t)
	defer em.cleanup(t)

	batch := &core.Batch{
		BatchHeader: core.BatchHeader{
			ID: fftypes.NewUUID(),
		},
		Hash: fftypes.NewRandB32(),
		Payload: core.BatchPayload{
			TX: core.TransactionRef{
				ID:   fftypes.NewUUID(),
				Type: core.TransactionTypeContractInvoke,
			},
			Messages: []*core.Message{{}},
			Data:     core.DataArray{{}},
		},
	}

	_, ok, err := em.persistBatch(em.ctx, batch)
	assert.NoError(t, err)
	assert.False(t, ok)

}
