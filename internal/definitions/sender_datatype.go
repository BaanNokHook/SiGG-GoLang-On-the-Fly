// SiGG-GoLang-On-the-Fly //
package definitions

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

func (bm *definitionSender) DefineDatatype(ctx context.Context, datatype *core.Datatype, waitConfirm bool) error {
	// Validate the input data definition data
	datatype.ID = fftypes.NewUUID()
	datatype.Created = fftypes.Now()
	if datatype.Validator == "" {
		datatype.Validator = core.ValidatorTypeJSON
	}
	datatype.Hash = datatype.Value.Hash()

	if bm.multiparty {
		if err := datatype.Validate(ctx, false); err != nil {
			return err
		}
		// Verify the data type is now all valid, before we broadcast it
		if err := bm.data.CheckDatatype(ctx, datatype); err != nil {
			return err
		}

		datatype.Namespace = ""
		msg, err := bm.sendDefinitionDefault(ctx, datatype, core.SystemTagDefineDatatype, waitConfirm)
		if msg != nil {
			datatype.Message = msg.Header.ID
		}
		datatype.Namespace = bm.namespace
		return err
	}

	return i18n.NewError(ctx, coremsgs.MsgActionNotSupported)
}
