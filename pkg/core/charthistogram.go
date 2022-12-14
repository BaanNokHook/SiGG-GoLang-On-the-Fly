// SiGG-GoLang-On-the-Fly //
package core

import "github.com/hyperledger/firefly-common/pkg/fftypes"

const (
	// ChartHistogramMaxBuckets max buckets that can be requested
	ChartHistogramMaxBuckets = 100
	// ChartHistogramMinBuckets min buckets that can be requested
	ChartHistogramMinBuckets = 1
)

// ChartHistogram is a list of buckets with types
type ChartHistogram struct {
	Count     string                `ffstruct:"ChartHistogram" json:"count"`
	Timestamp *fftypes.FFTime       `ffstruct:"ChartHistogram" json:"timestamp"`
	Types     []*ChartHistogramType `ffstruct:"ChartHistogram" json:"types"`
	IsCapped  bool                  `ffstruct:"ChartHistogram" json:"isCapped"`
}

// ChartHistogramType is a type and count
type ChartHistogramType struct {
	Count string `ffstruct:"ChartHistogramType" json:"count"`
	Type  string `ffstruct:"ChartHistogramType" json:"type"`
}

// ChartHistogramInterval specifies lower and upper timestamps for histogram bucket
type ChartHistogramInterval struct {
	// StartTime start time of histogram interval
	StartTime *fftypes.FFTime `json:"startTime"`
	// EndTime end time of histogram interval
	EndTime *fftypes.FFTime `json:"endTime"`
}
