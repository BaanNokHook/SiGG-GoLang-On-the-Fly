// SiGG-GoLang-On-the-Fly //
package definitions

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

func (dh *definitionHandler) handleDeprecatedNodeBroadcast(ctx context.Context, state *core.BatchState, msg *core.Message, data core.DataArray) (HandlerResult, error) {
	var nodeOld core.DeprecatedNode
	if valid := dh.getSystemBroadcastPayload(ctx, msg, data, &nodeOld); !valid {
		return HandlerResult{Action: ActionReject}, i18n.NewError(ctx, coremsgs.MsgDefRejectedBadPayload, "node", msg.Header.ID)
	}

	owner, err := dh.identity.FindIdentityForVerifier(ctx, []core.IdentityType{core.IdentityTypeOrg}, &core.VerifierRef{
		Type:  dh.blockchain.VerifierType(),
		Value: nodeOld.Owner,
	})
	if err != nil {
		return HandlerResult{Action: ActionRetry}, err // We only return database errors
	}
	if owner == nil {
		return HandlerResult{Action: ActionReject}, i18n.NewError(ctx, coremsgs.MsgDefRejectedIdentityNotFound, "node", nodeOld.ID, nodeOld.Owner)
	}

	return dh.handleIdentityClaim(ctx, state, buildIdentityMsgInfo(msg, nil), nodeOld.AddMigratedParent(owner.ID))

}
