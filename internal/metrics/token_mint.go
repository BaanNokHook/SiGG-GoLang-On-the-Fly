// SiGG-GoLang-On-the-Fly //
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var MintSubmittedCounter prometheus.Counter
var MintConfirmedCounter prometheus.Counter
var MintRejectedCounter prometheus.Counter
var MintHistogram prometheus.Histogram

// MintSubmittedCounterName is the prometheus metric for tracking the total number of mints submitted
var MintSubmittedCounterName = "ff_mint_submitted_total"

// MintConfirmedCounterName is the prometheus metric for tracking the total number of mints confirmed
var MintConfirmedCounterName = "ff_mint_confirmed_total"

// MintRejectedCounterName is the prometheus metric for tracking the total number of mints rejected
var MintRejectedCounterName = "ff_mint_rejected_total"

// MintHistogramName is the prometheus metric for tracking the total number of mints - histogram
var MintHistogramName = "ff_mint_histogram"

func InitTokenMintMetrics() {
	MintSubmittedCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: MintSubmittedCounterName,
		Help: "Number of submitted mints",
	})
	MintConfirmedCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: MintConfirmedCounterName,
		Help: "Number of confirmed mints",
	})
	MintRejectedCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: MintRejectedCounterName,
		Help: "Number of rejected mints",
	})
	MintHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: MintHistogramName,
		Help: "Histogram of mints, bucketed by time to finished",
	})
}

func RegisterTokenMintMetrics() {
	registry.MustRegister(MintSubmittedCounter)
	registry.MustRegister(MintConfirmedCounter)
	registry.MustRegister(MintRejectedCounter)
	registry.MustRegister(MintHistogram)
}
