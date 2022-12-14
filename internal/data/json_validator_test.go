// SiGG-GoLang-On-the-Fly //

package data

import (
	"context"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
)

func TestJSONValidator(t *testing.T) {

	schemaBinary := []byte(`{
		"properties": {
			"prop1": {
				"type": "string"
			}
		},
		"required": ["prop1"]
	}`)

	dt := &core.Datatype{
		Validator: core.ValidatorTypeJSON,
		Name:      "customer",
		Version:   "0.0.1",
		Value:     fftypes.JSONAnyPtrBytes(schemaBinary),
	}

	jv, err := newJSONValidator(context.Background(), "ns1", dt)
	assert.NoError(t, err)

	err = jv.validateJSONString(context.Background(), `{}`)
	assert.Regexp(t, "FF10198.*prop1", err)

	err = jv.validateJSONString(context.Background(), `{"prop1": "a value"}`)
	assert.NoError(t, err)

	err = jv.validateJSONString(context.Background(), `{!bad json`)
	assert.Regexp(t, "FF00127", err)

	assert.Equal(t, int64(len(schemaBinary)), jv.Size())

}

func TestJSONValidatorParseSchemaFail(t *testing.T) {

	dt := &core.Datatype{
		Validator: core.ValidatorTypeJSON,
		Name:      "customer",
		Version:   "0.0.1",
		Value:     fftypes.JSONAnyPtr(`{!json`),
	}

	_, err := newJSONValidator(context.Background(), "ns1", dt)
	assert.Regexp(t, "FF10196", err)

}

func TestJSONValidatorNilData(t *testing.T) {

	v := &jsonValidator{}
	err := v.Validate(context.Background(), &core.Data{})
	assert.Regexp(t, "FF10199", err)

}
