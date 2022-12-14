// SiGG-GoLang-On-the-Fly //

package core

import (
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/stretchr/testify/assert"
)

func TestFFISerializedEventScan(t *testing.T) {
	params := &FFISerializedEvent{}
	err := params.Scan([]byte(`{"name":"event1","description":"a super event","params":[{"name":"details","type":"integer","details":{"type":"uint256"}}]}`))
	assert.NoError(t, err)
}

func TestFFISerializedEventScanNil(t *testing.T) {
	params := &FFISerializedEvent{}
	err := params.Scan(nil)
	assert.Nil(t, err)
}

func TestFFISerializedEventScanString(t *testing.T) {
	params := &FFISerializedEvent{}
	err := params.Scan(`{"name":"event1","description":"a super event","params":[{"name":"details","type":"integer","details":{"type":"uint256"}}]}`)
	assert.NoError(t, err)
}

func TestFFISerializedEventScanError(t *testing.T) {
	params := &FFISerializedEvent{}
	err := params.Scan(map[string]interface{}{"this is": "not a supported serialization of a FFISerializedEvent"})
	assert.Regexp(t, "FF00105", err)
}

func TestFFISerializedEventValue(t *testing.T) {
	params := &FFISerializedEvent{
		FFIEventDefinition: fftypes.FFIEventDefinition{
			Name:        "event1",
			Description: "a super event",
			Params: fftypes.FFIParams{
				&fftypes.FFIParam{Name: "details", Schema: fftypes.JSONAnyPtr(`{"type": "integer", "details": {"type": "uint256"}}`)},
			},
		},
	}

	val, err := params.Value()
	assert.NoError(t, err)
	assert.Equal(t, `{"name":"event1","description":"a super event","params":[{"name":"details","schema":{"type":"integer","details":{"type":"uint256"}}}]}`, string(val.([]byte)))
}

func TestContractListenerOptionsScan(t *testing.T) {
	options := &ContractListenerOptions{}
	err := options.Scan([]byte(`{"firstBlock":"newest"}`))
	assert.NoError(t, err)
}

func TestContractListenerOptionsScanNil(t *testing.T) {
	options := &ContractListenerOptions{}
	err := options.Scan(nil)
	assert.Nil(t, err)
}

func TestContractListenerOptionsScanString(t *testing.T) {
	options := &ContractListenerOptions{}
	err := options.Scan(`{"firstBlock":"newest"}`)
	assert.NoError(t, err)
}

func TestContractListenerOptionsScanError(t *testing.T) {
	options := &ContractListenerOptions{}
	err := options.Scan(false)
	assert.Regexp(t, "FF00105", err)
}

func TestContractListenerOptionsValue(t *testing.T) {
	options := &ContractListenerOptions{
		FirstEvent: "newest",
	}

	val, err := options.Value()
	assert.NoError(t, err)
	assert.Equal(t, `{"firstEvent":"newest"}`, string(val.([]byte)))
}
