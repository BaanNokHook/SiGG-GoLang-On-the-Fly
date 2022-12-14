// SiGG-GoLang-On-the-Fly //

package core

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/stretchr/testify/assert"
)

func TestSQLSerializedManifest(t *testing.T) {

	msgID1 := fftypes.NewUUID()
	msgID2 := fftypes.NewUUID()
	batch := &Batch{
		BatchHeader: BatchHeader{
			ID: fftypes.NewUUID(),
		},
		Payload: BatchPayload{
			TX: TransactionRef{
				ID: fftypes.NewUUID(),
			},
			Messages: []*Message{
				{Header: MessageHeader{ID: msgID1}},
				{Header: MessageHeader{ID: msgID2}},
			},
		},
	}

	bp, manifest := batch.Confirmed()
	mfString := manifest.String()
	assert.Equal(t, batch.BatchHeader, bp.BatchHeader)
	assert.Equal(t, batch.Payload.TX, bp.TX)
	assert.Equal(t, mfString, bp.Manifest.String())
	assert.NotNil(t, bp.Confirmed)

	var mf *BatchManifest
	err := json.Unmarshal([]byte(mfString), &mf)
	assert.NoError(t, err)
	assert.Equal(t, msgID1, mf.Messages[0].ID)
	assert.Equal(t, msgID2, mf.Messages[1].ID)
	mfHash := sha256.Sum256([]byte(mfString))
	assert.Equal(t, fftypes.HashString(bp.GenManifest(batch.Payload.Messages, batch.Payload.Data).String()).String(), hex.EncodeToString(mfHash[:]))

	assert.Equal(t, batch, bp.GenInflight(batch.Payload.Messages, batch.Payload.Data))

	assert.NotEqual(t, batch.Payload.Hash().String(), hex.EncodeToString(mfHash[:]))

}
