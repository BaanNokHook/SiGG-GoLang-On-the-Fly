// SiGG-GoLang-On-the-Fly //

package core

import (
	"context"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/stretchr/testify/assert"
)

type fakePlugin struct{}

func (f *fakePlugin) Name() string { return "fake" }

func TestNewPendingMessageOp(t *testing.T) {

	txID := fftypes.NewUUID()
	op := NewOperation(&fakePlugin{}, "ns1", txID, OpTypeSharedStorageUploadBatch)
	assert.Equal(t, Operation{
		ID:          op.ID,
		Namespace:   "ns1",
		Transaction: txID,
		Plugin:      "fake",
		Type:        OpTypeSharedStorageUploadBatch,
		Status:      OpStatusPending,
		Created:     op.Created,
		Updated:     op.Created,
	}, *op)
}

func TestParseNamespacedOpID(t *testing.T) {

	ctx := context.Background()
	u := fftypes.NewUUID()

	_, _, err := ParseNamespacedOpID(ctx, "")
	assert.Regexp(t, "FF10411", err)

	_, _, err = ParseNamespacedOpID(ctx, "a::"+u.String())
	assert.Regexp(t, "FF10411", err)

	_, _, err = ParseNamespacedOpID(ctx, "bad%namespace:"+u.String())
	assert.Regexp(t, "FF00140", err)

	_, _, err = ParseNamespacedOpID(ctx, "ns1:Bad UUID")
	assert.Regexp(t, "FF00138", err)

	po := &PreparedOperation{
		ID:        u,
		Namespace: "ns1",
	}
	ns, u1, err := ParseNamespacedOpID(ctx, po.NamespacedIDString())
	assert.NoError(t, err)
	assert.Equal(t, u, u1)
	assert.Equal(t, "ns1", ns)

}
