// SiGG-GoLang-On-the-Fly //
package core

import "github.com/hyperledger/firefly-common/pkg/fftypes"

// Pin represents a ledger-pinning event that has been
// detected from the blockchain, in the sequence that it was detected.
//
// A batch contains many messages, and each of those messages can be on a different
// topic (or topics)
// All messages on the same topic must be processed in the order that
// the batch pinning events arrive from the blockchain.
//
// As we need to correlate the on-chain events, with off-chain data that might
// arrive at a different time (or never), we "park" all pinned sequences first,
// then only complete them (and generate the associated events) once all the data
// has been assembled for all messages on that sequence, within that batch.
//
// We might park the pin first (from the blockchain), or park the batch first
// (if it arrived first off-chain).
// There's a third part as well that can block a message, which is large blob data
// moving separately to the batch. If we get the private message, then the batch,
// before receiving the blob data - we have to upgrade a batch-park, to a pin-park.
// This is because the sequence must be in the order the pins arrive.
type Pin struct {
	Sequence   int64            `ffstruct:"Pin" json:"sequence"`
	Namespace  string           `ffstruct:"Pin" json:"namespace"`
	Masked     bool             `ffstruct:"Pin" json:"masked,omitempty"`
	Hash       *fftypes.Bytes32 `ffstruct:"Pin" json:"hash,omitempty"`
	Batch      *fftypes.UUID    `ffstruct:"Pin" json:"batch,omitempty"`
	BatchHash  *fftypes.Bytes32 `ffstruct:"Pin" json:"batchHash,omitempty"`
	Index      int64            `ffstruct:"Pin" json:"index"`
	Dispatched bool             `ffstruct:"Pin" json:"dispatched,omitempty"`
	Signer     string           `ffstruct:"Pin" json:"signer,omitempty"`
	Created    *fftypes.FFTime  `ffstruct:"Pin" json:"created,omitempty"`
}

func (p *Pin) LocalSequence() int64 {
	return p.Sequence
}

type PinRewind struct {
	Sequence int64         `ffstruct:"PinRewind" json:"sequence"`
	Batch    *fftypes.UUID `ffstruct:"PinRewind" json:"batch"`
}
