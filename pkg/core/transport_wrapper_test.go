// SiGG-GoLang-On-the-Fly //
package core

import (
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/stretchr/testify/assert"
)

func TestBatchManifest(t *testing.T) {

	tw := TransportWrapper{
		Batch: &Batch{
			Payload: BatchPayload{
				Messages: []*Message{
					{Header: MessageHeader{ID: fftypes.NewUUID()}, Hash: fftypes.NewRandB32()},
					{Header: MessageHeader{ID: fftypes.NewUUID()}, Hash: fftypes.NewRandB32()},
				},
				Data: []*Data{
					{ID: fftypes.NewUUID(), Hash: fftypes.NewRandB32()},
					{ID: fftypes.NewUUID(), Hash: fftypes.NewRandB32()},
				},
			},
		},
	}
	bp, _ := tw.Batch.Confirmed()
	tm := bp.GenManifest(tw.Batch.Payload.Messages, tw.Batch.Payload.Data)
	assert.Equal(t, 2, len(tm.Messages))
	assert.Equal(t, tw.Batch.Payload.Messages[0].Header.ID.String(), tm.Messages[0].ID.String())
	assert.Equal(t, tw.Batch.Payload.Messages[1].Header.ID.String(), tm.Messages[1].ID.String())
	assert.Equal(t, tw.Batch.Payload.Messages[0].Hash.String(), tm.Messages[0].Hash.String())
	assert.Equal(t, tw.Batch.Payload.Messages[1].Hash.String(), tm.Messages[1].Hash.String())
	assert.Equal(t, 2, len(tm.Data))
	assert.Equal(t, tw.Batch.Payload.Data[0].ID.String(), tm.Data[0].ID.String())
	assert.Equal(t, tw.Batch.Payload.Data[1].ID.String(), tm.Data[1].ID.String())
	assert.Equal(t, tw.Batch.Payload.Data[0].Hash.String(), tm.Data[0].Hash.String())
	assert.Equal(t, tw.Batch.Payload.Data[1].Hash.String(), tm.Data[1].Hash.String())

}
