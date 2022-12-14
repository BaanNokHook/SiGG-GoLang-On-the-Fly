// SiGG-GoLang-On-the-Fly //

package core

import (
	"context"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/stretchr/testify/assert"
)

func TestDatatypeValidation(t *testing.T) {

	dt := &Datatype{
		Validator: ValidatorType("wrong"),
	}
	assert.Regexp(t, "FF00111.*wrong", dt.Validate(context.Background(), false))

	dt = &Datatype{
		Validator: ValidatorTypeJSON,
		Namespace: "ok",
		Name:      "!wrong",
	}
	assert.Regexp(t, "FF00140.*name", dt.Validate(context.Background(), false))

	dt = &Datatype{
		Validator: ValidatorTypeJSON,
		Namespace: "ok",
		Name:      "ok",
		Version:   "!wrong",
	}
	assert.Regexp(t, "FF00140.*version", dt.Validate(context.Background(), false))

	dt = &Datatype{
		Validator: ValidatorTypeJSON,
		Namespace: "ok",
		Name:      "ok",
		Version:   "ok",
	}
	assert.Regexp(t, "FF00112.*value", dt.Validate(context.Background(), false))

	dt = &Datatype{
		Validator: ValidatorTypeJSON,
		Namespace: "ok",
		Name:      "ok",
		Version:   "ok",
		Value:     fftypes.JSONAnyPtr(`{}`),
	}
	assert.NoError(t, dt.Validate(context.Background(), false))

	assert.Regexp(t, "FF00114", dt.Validate(context.Background(), true))

	dt.ID = fftypes.NewUUID()
	dt.Hash = fftypes.NewRandB32()
	assert.Regexp(t, "FF00113", dt.Validate(context.Background(), true))

	var def Definition = dt
	assert.Equal(t, "8e23c0a7fa9ec15c68a662e0e502933facb3d249409efa2b4f89d479b9f990cb", def.Topic())
	def.SetBroadcastMessage(fftypes.NewUUID())
	assert.NotNil(t, dt.Message)
}
