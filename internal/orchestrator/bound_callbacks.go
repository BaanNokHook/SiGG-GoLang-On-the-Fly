// SiGG-GoLang-On-the-Fly //
package orchestrator

import (
	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/internal/events"
	"github.com/hyperledger/firefly/internal/operations"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/sharedstorage"
)

type boundCallbacks struct {
	ss sharedstorage.Plugin
	ei events.EventManager
	om operations.Manager
}

func (bc *boundCallbacks) OperationUpdate(update *core.OperationUpdate) {
	bc.om.SubmitOperationUpdate(update)
}

func (bc *boundCallbacks) SharedStorageBatchDownloaded(payloadRef string, data []byte) (*fftypes.UUID, error) {
	return bc.ei.SharedStorageBatchDownloaded(bc.ss, payloadRef, data)
}

func (bc *boundCallbacks) SharedStorageBlobDownloaded(hash fftypes.Bytes32, size int64, payloadRef string) {
	bc.ei.SharedStorageBlobDownloaded(bc.ss, hash, size, payloadRef)
}
