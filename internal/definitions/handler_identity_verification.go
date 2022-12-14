// SiGG-GoLang-On-the-Fly //
package definitions

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

func (dh *definitionHandler) handleIdentityVerificationBroadcast(ctx context.Context, state *core.BatchState, verifyMsg *core.Message, data core.DataArray) (HandlerResult, error) {
	var verification core.IdentityVerification
	valid := dh.getSystemBroadcastPayload(ctx, verifyMsg, data, &verification)
	if !valid {
		return HandlerResult{Action: ActionReject}, i18n.NewError(ctx, coremsgs.MsgDefRejectedBadPayload, "identity verification", verifyMsg.Header.ID)
	}
	verification.Identity.Namespace = dh.namespace.Name
	err := verification.Identity.Validate(ctx)
	if err != nil || verification.Identity.Parent == nil || verification.Claim.ID == nil || verification.Claim.Hash == nil {
		return HandlerResult{Action: ActionReject}, i18n.NewError(ctx, coremsgs.MsgDefRejectedValidateFail, "identity verification", verifyMsg.Header.ID, err)
	}

	// Check the verification is signed by the correct org
	parent, err := dh.identity.CachedIdentityLookupByID(ctx, verification.Identity.Parent)
	if err != nil {
		return HandlerResult{Action: ActionRetry}, err
	}
	if parent == nil {
		return HandlerResult{Action: ActionReject}, i18n.NewError(ctx, coremsgs.MsgDefRejectedIdentityNotFound, "identity verification", verifyMsg.Header.ID, verification.Identity.Parent)
	}
	if parent.DID != verifyMsg.Header.Author {
		return HandlerResult{Action: ActionReject}, i18n.NewError(ctx, coremsgs.MsgDefRejectedWrongAuthor, "identity verification", verifyMsg.Header.ID, verifyMsg.Header.Author)
	}

	// At this point, this is a valid verification, but we don't know if the claim has arrived.
	// See if the message has already arrived, if so we need to queue a rewind to it
	claimMsg, err := dh.database.GetMessageByID(ctx, dh.namespace.Name, verification.Claim.ID)
	if err != nil {
		return HandlerResult{Action: ActionRetry}, err
	}
	// See if the message was processed earlier in this same batch
	if claimMsg == nil || claimMsg.State != core.MessageStateConfirmed {
		claimMsg = state.PendingConfirms[*verification.Claim.ID]
	}

	if claimMsg != nil {
		if !claimMsg.Hash.Equals(verification.Claim.Hash) {
			return HandlerResult{Action: ActionReject}, i18n.NewError(ctx, coremsgs.MsgDefRejectedHashMismatch, "identity verification", verifyMsg.Header.ID, claimMsg.Hash, verification.Claim.Hash)
		}
		data, foundAll, err := dh.data.GetMessageDataCached(ctx, claimMsg)
		if err != nil {
			return HandlerResult{Action: ActionRetry}, err
		}
		if foundAll {
			// The verification came in after the messsage, so we need to call the idempotent
			// handler of the claim logic again
			return dh.handleIdentityClaimBroadcast(ctx, state, claimMsg, data, verifyMsg.Header.ID)
		}
	}

	// Just confirm the verification - when the claim message is processed it will come back and look for
	// this (now confirmed) verification message.
	return HandlerResult{Action: ActionConfirm}, nil

}
