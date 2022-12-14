// SiGG-GoLang-On-the-Fly //
package core

import (
	"context"
	"fmt"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/stretchr/testify/assert"
)

func TestBatchStateFinalizers(t *testing.T) {
	bs := BatchState{}

	run1 := false
	bs.AddPreFinalize(func(ctx context.Context) error { run1 = true; return nil })
	run2 := false
	bs.AddFinalize(func(ctx context.Context) error { run2 = true; return nil })

	bs.RunPreFinalize(context.Background())
	bs.RunFinalize(context.Background())
	assert.True(t, run1)
	assert.True(t, run2)
}

func TestBatchStateFinalizerErrors(t *testing.T) {
	bs := BatchState{}

	bs.AddPreFinalize(func(ctx context.Context) error { return fmt.Errorf("pop") })
	bs.AddFinalize(func(ctx context.Context) error { return fmt.Errorf("pop") })

	err := bs.RunPreFinalize(context.Background())
	assert.EqualError(t, err, "pop")
	err = bs.RunFinalize(context.Background())
	assert.EqualError(t, err, "pop")
}

func TestBatchStateIdentities(t *testing.T) {
	bs := BatchState{
		PendingConfirms: make(map[fftypes.UUID]*Message),
	}

	id := fftypes.NewUUID()
	msg := &Message{}
	bs.AddPendingConfirm(id, msg)

	did := "did:firefly:id1"
	bs.AddConfirmedDIDClaim(did)

	assert.Equal(t, msg, bs.PendingConfirms[*id])
	assert.Equal(t, bs.ConfirmedDIDClaims[0], did)
}
