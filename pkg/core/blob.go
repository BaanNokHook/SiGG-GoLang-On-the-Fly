// SiGG-GoLang-On-the-Fly //
package core

import "github.com/hyperledger/firefly-common/pkg/fftypes"

type Blob struct {
	Hash       *fftypes.Bytes32 `json:"hash"`
	Size       int64            `json:"size"`
	PayloadRef string           `json:"payloadRef,omitempty"`
	Peer       string           `json:"peer,omitempty"`
	Created    *fftypes.FFTime  `json:"created,omitempty"`
	Sequence   int64            `json:"-"`
}
