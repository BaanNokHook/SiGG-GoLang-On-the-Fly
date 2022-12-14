// SiGG-GoLang-On-the-Fly //

package orchestrator

import (
	"fmt"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/mocks/eventmocks"
	"github.com/hyperledger/firefly/mocks/operationmocks"
	"github.com/hyperledger/firefly/mocks/sharedstoragemocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
)

func TestBoundCallbacks(t *testing.T) {
	mei := &eventmocks.EventManager{}
	mss := &sharedstoragemocks.Plugin{}
	mom := &operationmocks.Manager{}
	bc := boundCallbacks{ei: mei, ss: mss, om: mom}

	info := fftypes.JSONObject{"hello": "world"}
	hash := fftypes.NewRandB32()
	opID := fftypes.NewUUID()
	nsOpID := "ns1:" + opID.String()

	update := &core.OperationUpdate{
		NamespacedOpID: nsOpID,
		Status:         core.OpStatusFailed,
		BlockchainTXID: "0xffffeeee",
		ErrorMessage:   "error info",
		Output:         info,
	}
	mom.On("SubmitOperationUpdate", update).Return().Once()
	bc.OperationUpdate(update)

	mei.On("SharedStorageBatchDownloaded", mss, "payload1", []byte(`{}`)).Return(nil, fmt.Errorf("pop"))
	_, err := bc.SharedStorageBatchDownloaded("payload1", []byte(`{}`))
	assert.EqualError(t, err, "pop")

	mei.On("SharedStorageBlobDownloaded", mss, *hash, int64(12345), "payload1").Return()
	bc.SharedStorageBlobDownloaded(*hash, 12345, "payload1")

	mei.AssertExpectations(t)
	mss.AssertExpectations(t)
	mom.AssertExpectations(t)
}
