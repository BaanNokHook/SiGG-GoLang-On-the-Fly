// SiGG-GoLang-On-the-Fly //
package core

import "github.com/hyperledger/firefly-common/pkg/fftypes"

// Definition is implemented by all objects that can be broadcast as system definitions to the network
type Definition interface {
	// Topic returns the topic on which the object should be broadcast
	Topic() string
	// SetBroadcastMessage sets the message that broadcast the definition
	SetBroadcastMessage(msgID *fftypes.UUID)
}
