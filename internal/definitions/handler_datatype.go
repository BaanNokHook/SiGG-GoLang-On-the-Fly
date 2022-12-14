// SiGG-GoLang-On-the-Fly //
package definitions

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

func (dh *definitionHandler) handleDatatypeBroadcast(ctx context.Context, state *core.BatchState, msg *core.Message, data core.DataArray, tx *fftypes.UUID) (HandlerResult, error) {
	var dt core.Datatype
	valid := dh.getSystemBroadcastPayload(ctx, msg, data, &dt)
	if !valid {
		return HandlerResult{Action: ActionReject}, i18n.NewError(ctx, coremsgs.MsgDefRejectedBadPayload, "datatype", msg.Header.ID)
	}
	dt.Namespace = dh.namespace.Name
	if err := dt.Validate(ctx, true); err != nil {
		return HandlerResult{Action: ActionReject}, i18n.NewError(ctx, coremsgs.MsgDefRejectedValidateFail, "datatype", dt.ID, err)
	}
	if err := dh.data.CheckDatatype(ctx, &dt); err != nil {
		return HandlerResult{Action: ActionReject}, i18n.NewError(ctx, coremsgs.MsgDefRejectedSchemaFail, "datatype", dt.ID, err)
	}

	existing, err := dh.database.GetDatatypeByName(ctx, dt.Namespace, dt.Name, dt.Version)
	if err != nil {
		return HandlerResult{Action: ActionRetry}, err
	} else if existing != nil {
		return HandlerResult{Action: ActionReject}, i18n.NewError(ctx, coremsgs.MsgDefRejectedConflict, "datatype", dt.ID, existing.ID)
	}

	if err = dh.database.UpsertDatatype(ctx, &dt, false); err != nil {
		return HandlerResult{Action: ActionRetry}, err
	}

	state.AddFinalize(func(ctx context.Context) error {
		event := core.NewEvent(core.EventTypeDatatypeConfirmed, dt.Namespace, dt.ID, tx, core.SystemTopicDefinitions)
		return dh.database.InsertEvent(ctx, event)
	})
	return HandlerResult{Action: ActionConfirm}, nil
}
