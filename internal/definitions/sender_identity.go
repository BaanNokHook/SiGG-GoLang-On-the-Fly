// SiGG-GoLang-On-the-Fly //
package definitions

import (
	"context"

	"github.com/hyperledger/firefly/internal/identity"
	"github.com/hyperledger/firefly/pkg/core"
)

// ClaimIdentity is a special form of CreateDefinition where the signing identity does not need to have been pre-registered
// The blockchain "key" will be normalized, but the "author" will pass through unchecked
func (bm *definitionSender) ClaimIdentity(ctx context.Context, claim *core.IdentityClaim, signingIdentity *core.SignerRef, parentSigner *core.SignerRef, waitConfirm bool) error {
	if bm.multiparty {
		var err error
		signingIdentity.Key, err = bm.identity.NormalizeSigningKey(ctx, signingIdentity.Key, identity.KeyNormalizationBlockchainPlugin)
		if err != nil {
			return err
		}

		claim.Identity.Namespace = ""
		claimMsg, err := bm.sendDefinitionCommon(ctx, claim, signingIdentity, core.SystemTagIdentityClaim, waitConfirm)
		if err != nil {
			return err
		}
		claim.Identity.Messages.Claim = claimMsg.Header.ID

		// Send the verification if one is required.
		if parentSigner != nil {
			verifyMsg, err := bm.sendDefinition(ctx, &core.IdentityVerification{
				Claim: core.MessageRef{
					ID:   claimMsg.Header.ID,
					Hash: claimMsg.Hash,
				},
				Identity: claim.Identity.IdentityBase,
			}, parentSigner, core.SystemTagIdentityVerification, false)
			if err != nil {
				return err
			}
			claim.Identity.Messages.Verification = verifyMsg.Header.ID
		}

		return nil
	}

	claim.Identity.Namespace = bm.namespace
	return fakeBatch(ctx, func(ctx context.Context, state *core.BatchState) (HandlerResult, error) {
		return bm.handler.handleIdentityClaim(ctx, state, &identityMsgInfo{SignerRef: *signingIdentity}, claim)
	})
}

func (bm *definitionSender) UpdateIdentity(ctx context.Context, identity *core.Identity, def *core.IdentityUpdate, signingIdentity *core.SignerRef, waitConfirm bool) error {
	if bm.multiparty {
		updateMsg, err := bm.sendDefinition(ctx, def, signingIdentity, core.SystemTagIdentityUpdate, waitConfirm)
		identity.Messages.Update = updateMsg.Header.ID
		return err
	}

	return fakeBatch(ctx, func(ctx context.Context, state *core.BatchState) (HandlerResult, error) {
		return bm.handler.handleIdentityUpdate(ctx, state, &identityUpdateMsgInfo{}, def)
	})
}
