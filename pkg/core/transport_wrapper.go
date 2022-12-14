// SiGG-GoLang-On-the-Fly //
package core

import "github.com/hyperledger/firefly-common/pkg/fftypes"

type TransportPayloadType = fftypes.FFEnum

var (
	TransportPayloadTypeMessage = fftypes.FFEnumValue("transportpayload", "message")
	TransportPayloadTypeBatch   = fftypes.FFEnumValue("transportpayload", "batch")
)

// TransportWrapper wraps paylaods over data exchange transfers, for easy deserialization at target
type TransportWrapper struct {
	Group *Group `json:"group,omitempty"`
	Batch *Batch `json:"batch,omitempty"`
}
