// SiGG-GoLang-On-the-Fly //
package networkmap

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

func (nm *networkMap) RegisterNode(ctx context.Context, waitConfirm bool) (identity *core.Identity, err error) {

	nodeOwningOrg, err := nm.identity.GetMultipartyRootOrg(ctx)
	if err != nil {
		return nil, err
	}

	localNodeName := nm.multiparty.LocalNode().Name
	if localNodeName == "" {
		return nil, i18n.NewError(ctx, coremsgs.MsgNodeAndOrgIDMustBeSet)
	}
	nodeRequest := &core.IdentityCreateDTO{
		Parent: nodeOwningOrg.ID.String(),
		Name:   localNodeName,
		Type:   core.IdentityTypeNode,
		IdentityProfile: core.IdentityProfile{
			Description: nm.multiparty.LocalNode().Description,
		},
	}

	nodeRequest.Profile, err = nm.exchange.GetEndpointInfo(ctx, localNodeName)
	if err != nil {
		return nil, err
	}

	return nm.RegisterIdentity(ctx, nodeRequest, waitConfirm)
}
