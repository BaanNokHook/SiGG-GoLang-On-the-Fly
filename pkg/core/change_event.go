// SiGG-GoLang-On-the-Fly //
package core

import "github.com/hyperledger/firefly-common/pkg/fftypes"

// ChangeEventType
type ChangeEventType string

const (
	ChangeEventTypeCreated ChangeEventType = "created"
	ChangeEventTypeUpdated ChangeEventType = "updated" // note bulk updates might not results in change events.
	ChangeEventTypeDeleted ChangeEventType = "deleted"
	ChangeEventTypeDropped ChangeEventType = "dropped" // See ChangeEventDropped structure, sent to client instead of ChangeEvent when dropping notifications
)

type WSChangeEventCommandType = fftypes.FFEnum

var (
	// WSChangeEventCommandTypeStart is the command to start listening
	WSChangeEventCommandTypeStart = fftypes.FFEnumValue("changeevent_cmd_type", "start")
)

// WSChangeEventCommand is the WebSocket command to send to start listening for change events.
// Replaces any previous start requests.
type WSChangeEventCommand struct {
	Type        WSChangeEventCommandType `json:"type" ffenum:"changeevent_cmd_type"`
	Collections []string                 `json:"collections"`
	Filter      ChangeEventFilter        `json:"filter"`
}

type ChangeEventFilter struct {
	Types      []ChangeEventType `json:"types,omitempty"`
	Namespaces []string          `json:"namespaces,omitempty"`
}

// ChangeEvent is a change to the local FireFly core node.
type ChangeEvent struct {
	// The resource collection where the changed resource exists
	Collection string `json:"collection"`
	// The type of event
	Type ChangeEventType `json:"type"`
	// Namespace is set if there is a namespace associated with the changed resource
	Namespace string `json:"namespace,omitempty"`
	// UUID is set if the resource is identified by ID
	ID *fftypes.UUID `json:"id,omitempty"`
	// Hash is set if the resource is identified primarily by hash (groups is currently the only example)
	Hash *fftypes.Bytes32 `json:"hash,omitempty"`
	// Sequence is set if there is a local ordered sequence associated with the changed resource
	Sequence *int64 `json:"sequence,omitempty"`
	// DroppedSince only for ChangeEventTypeDropped. When the first miss happened
	DroppedSince *fftypes.FFTime `json:"droppedSince,omitempty"`
	// DroppedCount only for ChangeEventTypeDropped. How many events dropped
	DroppedCount int64 `json:"droppedCount,omitempty"`
}
