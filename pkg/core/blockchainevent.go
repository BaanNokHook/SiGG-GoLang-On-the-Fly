// SiGG-GoLang-On-the-Fly //
package core

import "github.com/hyperledger/firefly-common/pkg/fftypes"

type BlockchainEvent struct {
	ID         *fftypes.UUID            `ffstruct:"BlockchainEvent" json:"id,omitempty"`
	Source     string                   `ffstruct:"BlockchainEvent" json:"source,omitempty"`
	Namespace  string                   `ffstruct:"BlockchainEvent" json:"namespace,omitempty"`
	Name       string                   `ffstruct:"BlockchainEvent" json:"name,omitempty"`
	Listener   *fftypes.UUID            `ffstruct:"BlockchainEvent" json:"listener,omitempty"`
	ProtocolID string                   `ffstruct:"BlockchainEvent" json:"protocolId,omitempty"`
	Output     fftypes.JSONObject       `ffstruct:"BlockchainEvent" json:"output,omitempty"`
	Info       fftypes.JSONObject       `ffstruct:"BlockchainEvent" json:"info,omitempty"`
	Timestamp  *fftypes.FFTime          `ffstruct:"BlockchainEvent" json:"timestamp,omitempty"`
	TX         BlockchainTransactionRef `ffstruct:"BlockchainEvent" json:"tx"`
}
