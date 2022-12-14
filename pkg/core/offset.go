// SiGG-GoLang-On-the-Fly //
package core

import "github.com/hyperledger/firefly-common/pkg/fftypes"

type OffsetType = fftypes.FFEnum

var (
	// OffsetTypeBatch is an offset stored by the batch manager on the messages table
	OffsetTypeBatch = fftypes.FFEnumValue("offsettype", "batch")
	// OffsetTypeAggregator is an offset stored by the aggregator on the events table
	OffsetTypeAggregator = fftypes.FFEnumValue("offsettype", "aggregator")
	// OffsetTypeSubscription is an offeset stored by a dispatcher on the events table
	OffsetTypeSubscription = fftypes.FFEnumValue("offsettype", "subscription")
)

// Offset is a simple stored data structure that records a sequence position within another collection
type Offset struct {
	Type    OffsetType `json:"type" ffenum:"offsettype"`
	Name    string     `json:"name"`
	Current int64      `json:"current,omitempty"`

	RowID int64 `json:"-"`
}
