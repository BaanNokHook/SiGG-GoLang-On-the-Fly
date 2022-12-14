// SiGG-GoLang-On-the-Fly //
package shareddownload

import (
	"context"
	"fmt"

	"github.com/hyperledger/firefly-common/pkg/log"
	"github.com/hyperledger/firefly/internal/operations"
)

type downloadWorker struct {
	ctx  context.Context
	done chan struct{}
	dm   *downloadManager
}

func newDownloadWorker(dm *downloadManager, idx int) *downloadWorker {
	dw := &downloadWorker{
		ctx:  log.WithLogField(dm.ctx, "downloadworker", fmt.Sprintf("dw_%.3d", idx)),
		done: make(chan struct{}),
		dm:   dm,
	}
	go dw.downloadWorkerLoop()
	return dw
}

func (dw *downloadWorker) downloadWorkerLoop() {
	defer close(dw.done)

	l := log.L(dw.ctx)
	for {
		select {
		case <-dw.ctx.Done():
			l.Debugf("Download worker shutting down")
			return
		case work := <-dw.dm.work:
			dw.attemptWork(work)
		}
	}
}

func (dw *downloadWorker) attemptWork(work *downloadWork) {

	work.attempts++
	isLastAttempt := work.attempts >= dw.dm.retryMaxAttempts
	options := []operations.RunOperationOption{operations.RemainPendingOnFailure}
	if isLastAttempt {
		options = []operations.RunOperationOption{}
	}

	_, err := dw.dm.operations.RunOperation(dw.ctx, work.preparedOp, options...)
	if err != nil {
		log.L(dw.ctx).Errorf("Download operation %s/%s attempt=%d/%d failed: %s", work.preparedOp.Type, work.preparedOp.ID, work.attempts, dw.dm.retryMaxAttempts, err)
		if !isLastAttempt {
			go dw.dm.waitAndRetryDownload(work)
		}
	}
}
