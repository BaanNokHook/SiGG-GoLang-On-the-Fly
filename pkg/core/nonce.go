// SiGG-GoLang-On-the-Fly //
package core

import "github.com/hyperledger/firefly-common/pkg/fftypes"

// Nonce is this local node's state record for the context of a group+topic+author combination.
// The Hash is the state of the hash before the nonce is added on to make it unique to the message.
type Nonce struct {
	Hash  *fftypes.Bytes32 `json:"hash"`
	Nonce int64            `json:"nonce"`
}
