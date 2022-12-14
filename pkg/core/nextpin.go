// SiGG-GoLang-On-the-Fly //
package core

import "github.com/hyperledger/firefly-common/pkg/fftypes"

type NextPin struct {
	Namespace string           `json:"namespace"`
	Context   *fftypes.Bytes32 `json:"context"`
	Identity  string           `json:"identity"`
	Hash      *fftypes.Bytes32 `json:"hash"`
	Nonce     int64            `json:"nonce"`
	Sequence  int64            `json:"_"` // Local database sequence used internally for update efficiency
}
