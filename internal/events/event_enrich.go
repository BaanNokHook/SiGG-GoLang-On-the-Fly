// SiGG-GoLang-On-the-Fly //
package events

import (
	"context"

	"github.com/hyperledger/firefly/internal/data"
	"github.com/hyperledger/firefly/internal/operations"
	"github.com/hyperledger/firefly/internal/txcommon"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

type eventEnricher struct {
	namespace  string
	data       data.Manager
	database   database.Plugin
	operations operations.Manager
	txHelper   txcommon.Helper
}

func newEventEnricher(ns string, di database.Plugin, dm data.Manager, om operations.Manager, txHelper txcommon.Helper) *eventEnricher {
	return &eventEnricher{
		namespace:  ns,
		data:       dm,
		database:   di,
		operations: om,
		txHelper:   txHelper,
	}
}

func (em *eventEnricher) enrichEvent(ctx context.Context, event *core.Event) (*core.EnrichedEvent, error) {
	e := &core.EnrichedEvent{
		Event: *event,
	}

	switch event.Type {
	case core.EventTypeTransactionSubmitted:
		tx, err := em.txHelper.GetTransactionByIDCached(ctx, event.Reference)
		if err != nil {
			return nil, err
		}
		e.Transaction = tx
	case core.EventTypeMessageConfirmed, core.EventTypeMessageRejected:
		msg, _, _, err := em.data.GetMessageWithDataCached(ctx, event.Reference)
		if err != nil {
			return nil, err
		}
		e.Message = msg
	case core.EventTypeBlockchainEventReceived:
		be, err := em.txHelper.GetBlockchainEventByIDCached(ctx, event.Reference)
		if err != nil {
			return nil, err
		}
		e.BlockchainEvent = be
	case core.EventTypeContractAPIConfirmed:
		contractAPI, err := em.database.GetContractAPIByID(ctx, em.namespace, event.Reference)
		if err != nil {
			return nil, err
		}
		e.ContractAPI = contractAPI
	case core.EventTypeContractInterfaceConfirmed:
		contractInterface, err := em.database.GetFFIByID(ctx, em.namespace, event.Reference)
		if err != nil {
			return nil, err
		}
		e.ContractInterface = contractInterface
	case core.EventTypeDatatypeConfirmed:
		dt, err := em.database.GetDatatypeByID(ctx, em.namespace, event.Reference)
		if err != nil {
			return nil, err
		}
		e.Datatype = dt
	case core.EventTypeIdentityConfirmed, core.EventTypeIdentityUpdated:
		identity, err := em.database.GetIdentityByID(ctx, em.namespace, event.Reference)
		if err != nil {
			return nil, err
		}
		e.Identity = identity
	case core.EventTypePoolConfirmed:
		tokenPool, err := em.database.GetTokenPoolByID(ctx, em.namespace, event.Reference)
		if err != nil {
			return nil, err
		}
		e.TokenPool = tokenPool
	case core.EventTypeApprovalConfirmed:
		approval, err := em.database.GetTokenApprovalByID(ctx, em.namespace, event.Reference)
		if err != nil {
			return nil, err
		}
		e.TokenApproval = approval
	case core.EventTypeTransferConfirmed:
		transfer, err := em.database.GetTokenTransferByID(ctx, em.namespace, event.Reference)
		if err != nil {
			return nil, err
		}
		e.TokenTransfer = transfer
	case core.EventTypeApprovalOpFailed, core.EventTypeTransferOpFailed, core.EventTypeBlockchainInvokeOpFailed, core.EventTypePoolOpFailed, core.EventTypeBlockchainInvokeOpSucceeded:
		operation, err := em.operations.GetOperationByIDCached(ctx, event.Reference)
		if err != nil {
			return nil, err
		}
		e.Operation = operation
	}
	return e, nil
}
