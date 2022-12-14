// SiGG-GoLang-On-the-Fly //
package networkmap

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

// RegisterNodeOrganization is a convenience helper to register the org configured on the node, without any extra info
func (nm *networkMap) RegisterNodeOrganization(ctx context.Context, waitConfirm bool) (*core.Identity, error) {

	key, err := nm.identity.GetMultipartyRootVerifier(ctx)
	if err != nil {
		return nil, err
	}

	orgName := nm.multiparty.RootOrg().Name
	if orgName == "" {
		return nil, i18n.NewError(ctx, coremsgs.MsgNodeAndOrgIDMustBeSet)
	}
	orgRequest := &core.IdentityCreateDTO{
		Name: orgName,
		IdentityProfile: core.IdentityProfile{
			Description: nm.multiparty.RootOrg().Description,
		},
		Key: key.Value,
	}
	return nm.RegisterOrganization(ctx, orgRequest, waitConfirm)
}

func (nm *networkMap) RegisterOrganization(ctx context.Context, orgRequest *core.IdentityCreateDTO, waitConfirm bool) (*core.Identity, error) {
	orgRequest.Type = core.IdentityTypeOrg
	return nm.RegisterIdentity(ctx, orgRequest, waitConfirm)
}
