// SiGG-GoLang-On-the-Fly //
package networkmap

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

func (nm *networkMap) UpdateIdentity(ctx context.Context, uuidStr string, dto *core.IdentityUpdateDTO, waitConfirm bool) (identity *core.Identity, err error) {
	id, err := fftypes.ParseUUID(ctx, uuidStr)
	if err != nil {
		return nil, err
	}
	return nm.updateIdentityID(ctx, id, dto, waitConfirm)
}

func (nm *networkMap) updateIdentityID(ctx context.Context, id *fftypes.UUID, dto *core.IdentityUpdateDTO, waitConfirm bool) (identity *core.Identity, err error) {

	// Get the original identity
	identity, err = nm.identity.CachedIdentityLookupByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if identity == nil || identity.Namespace != nm.namespace {
		return nil, i18n.NewError(ctx, coremsgs.Msg404NoResult)
	}

	var updateSigner *core.SignerRef

	if nm.multiparty != nil {
		// Resolve the signer of the original claim
		updateSigner, err = nm.identity.ResolveIdentitySigner(ctx, identity)
		if err != nil {
			return nil, err
		}
	}

	identity.IdentityProfile = dto.IdentityProfile
	if err := identity.Validate(ctx); err != nil {
		return nil, err
	}

	// Send the update
	err = nm.defsender.UpdateIdentity(ctx, identity, &core.IdentityUpdate{
		Identity: identity.IdentityBase,
		Updates:  dto.IdentityProfile,
	}, updateSigner, waitConfirm)
	return identity, err
}
