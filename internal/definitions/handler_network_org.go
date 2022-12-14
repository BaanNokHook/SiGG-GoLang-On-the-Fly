// SiGG-GoLang-On-the-Fly //
package definitions

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

func (dh *definitionHandler) handleDeprecatedOrganizationBroadcast(ctx context.Context, state *core.BatchState, msg *core.Message, data core.DataArray) (HandlerResult, error) {

	var orgOld core.DeprecatedOrganization
	valid := dh.getSystemBroadcastPayload(ctx, msg, data, &orgOld)
	if !valid {
		return HandlerResult{Action: ActionReject}, i18n.NewError(ctx, coremsgs.MsgDefRejectedBadPayload, "org", msg.Header.ID)
	}

	return dh.handleIdentityClaim(ctx, state, buildIdentityMsgInfo(msg, nil), orgOld.Migrated())

}
