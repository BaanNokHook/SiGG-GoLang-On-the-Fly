// SiGG-GoLang-On-the-Fly //
package definitions

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/hyperledger/firefly-common/pkg/log"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

type identityUpdateMsgInfo struct {
	ID     *fftypes.UUID
	Author string
}

func (dh *definitionHandler) handleIdentityUpdateBroadcast(ctx context.Context, state *core.BatchState, msg *core.Message, data core.DataArray) (HandlerResult, error) {
	var update core.IdentityUpdate
	if valid := dh.getSystemBroadcastPayload(ctx, msg, data, &update); !valid {
		return HandlerResult{Action: ActionReject}, i18n.NewError(ctx, coremsgs.MsgDefRejectedBadPayload, "identity update", msg.Header.ID)
	}
	return dh.handleIdentityUpdate(ctx, state, &identityUpdateMsgInfo{
		ID:     msg.Header.ID,
		Author: msg.Header.Author,
	}, &update)
}

func (dh *definitionHandler) handleIdentityUpdate(ctx context.Context, state *core.BatchState, msg *identityUpdateMsgInfo, update *core.IdentityUpdate) (HandlerResult, error) {
	if err := update.Identity.Validate(ctx); err != nil {
		return HandlerResult{Action: ActionReject}, i18n.NewError(ctx, coremsgs.MsgDefRejectedValidateFail, "identity update", update.Identity.ID, err)
	}

	// Get the existing identity (must be a confirmed identity at the point an update is issued)
	identity, err := dh.identity.CachedIdentityLookupByID(ctx, update.Identity.ID)
	if err != nil {
		return HandlerResult{Action: ActionRetry}, err
	}
	if identity == nil {
		return HandlerResult{Action: ActionReject}, i18n.NewError(ctx, coremsgs.MsgDefRejectedIdentityNotFound, "identity update", update.Identity.ID, update.Identity.ID)
	}

	if dh.multiparty {

		parent, retryable, err := dh.identity.VerifyIdentityChain(ctx, identity)
		if err != nil && retryable {
			return HandlerResult{Action: ActionRetry}, err
		} else if err != nil {
			log.L(ctx).Infof("Unable to process identity update (parked) %s: %s", msg.ID, err)
			return HandlerResult{Action: ActionWait}, nil
		}

		// Check the author matches
		expectedSigner := dh.getExpectedSigner(identity, parent)
		if expectedSigner.DID != msg.Author {
			return HandlerResult{Action: ActionReject}, i18n.NewError(ctx, coremsgs.MsgDefRejectedWrongAuthor, "identity update", update.Identity.ID, msg.Author)
		}

	}

	// Update the profile
	identity.IdentityProfile = update.Updates
	identity.Messages.Update = msg.ID
	err = dh.database.UpsertIdentity(ctx, identity, database.UpsertOptimizationExisting)
	if err != nil {
		return HandlerResult{Action: ActionRetry}, err
	}

	state.AddFinalize(func(ctx context.Context) error {
		event := core.NewEvent(core.EventTypeIdentityUpdated, identity.Namespace, identity.ID, nil, core.SystemTopicDefinitions)
		return dh.database.InsertEvent(ctx, event)
	})
	return HandlerResult{Action: ActionConfirm}, err

}
