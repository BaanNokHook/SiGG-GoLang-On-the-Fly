// SiGG-GoLang-On-the-Fly //
package data

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/hyperledger/firefly-common/pkg/log"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

type jsonValidator struct {
	id       *fftypes.UUID
	size     int64
	ns       string
	datatype *core.DatatypeRef
	schema   *jsonschema.Schema
}

func newJSONValidator(ctx context.Context, ns string, datatype *core.Datatype) (*jsonValidator, error) {
	jv := &jsonValidator{
		id: datatype.ID,
		ns: ns,
		datatype: &core.DatatypeRef{
			Name:    datatype.Name,
			Version: datatype.Version,
		},
	}

	var schemaBytes []byte
	if datatype.Value != nil {
		schemaBytes = []byte(*datatype.Value)
	}
	c := jsonschema.NewCompiler()
	c.Draft = jsonschema.Draft2020
	err := c.AddResource(datatype.Name, strings.NewReader(datatype.Value.String()))
	if err != nil {
		return nil, i18n.WrapError(ctx, err, coremsgs.MsgSchemaLoadFailed, jv.datatype)
	}
	schema, err := c.Compile(datatype.Name)
	if err != nil {
		return nil, i18n.WrapError(ctx, err, coremsgs.MsgSchemaLoadFailed, jv.datatype)
	}
	jv.schema = schema
	jv.size = int64(len(schemaBytes))

	log.L(ctx).Debugf("Found JSON schema validator for json:%s:%s: %v", jv.ns, datatype, jv.id)
	return jv, nil
}

func (jv *jsonValidator) Validate(ctx context.Context, data *core.Data) error {
	return jv.ValidateValue(ctx, data.Value, data.Hash)
}

func (jv *jsonValidator) ValidateValue(ctx context.Context, value *fftypes.JSONAny, expectedHash *fftypes.Bytes32) error {
	if value == nil {
		return i18n.NewError(ctx, coremsgs.MsgDataValueIsNull)
	}

	if expectedHash != nil {
		hash := value.Hash()
		if *hash != *expectedHash {
			return i18n.NewError(ctx, coremsgs.MsgDataInvalidHash, hash, expectedHash)
		}
	}

	return jv.validateJSONString(ctx, value.String())
}

func (jv *jsonValidator) validateJSONString(ctx context.Context, input string) error {
	inputValue, err := jsonDecode(input)
	if err != nil {
		return i18n.NewError(ctx, i18n.MsgJSONObjectParseFailed, input, err)
	}
	if err := jv.schema.Validate(inputValue); err != nil {
		log.L(ctx).Warnf("JSON schema %s [%v] validation failed: %s", jv.datatype, jv.id, err)
		return i18n.NewError(ctx, coremsgs.MsgJSONDataInvalidPerSchema, jv.datatype, err)
	}
	return nil
}

func (jv *jsonValidator) Size() int64 {
	return jv.size
}

func jsonDecode(input string) (interface{}, error) {
	var output interface{}
	if err := json.Unmarshal([]byte(input), &output); err != nil {
		return nil, err
	}
	return output, nil
}
