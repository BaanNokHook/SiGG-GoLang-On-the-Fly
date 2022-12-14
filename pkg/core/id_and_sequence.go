// SiGG-GoLang-On-the-Fly //
package core

import "github.com/hyperledger/firefly-common/pkg/fftypes"

// IDAndSequence is a combination of a UUID and a stored sequence
type IDAndSequence struct {
	ID       fftypes.UUID
	Sequence int64
}
