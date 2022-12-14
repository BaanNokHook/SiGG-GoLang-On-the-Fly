// SiGG-GoLang-On-the-Fly //
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var BroadcastSubmittedCounter prometheus.Counter
var BroadcastConfirmedCounter prometheus.Counter
var BroadcastRejectedCounter prometheus.Counter
var BroadcastHistogram prometheus.Histogram

// BroadcastSubmittedCounterName is the prometheus metric for tracking the total number of broadcasts submitted
var BroadcastSubmittedCounterName = "ff_broadcast_submitted_total"

// BroadcastConfirmedCounterName is the prometheus metric for tracking the total number of broadcasts confirmed
var BroadcastConfirmedCounterName = "ff_broadcast_confirmed_total"

// BroadcastRejectedCounterName is the prometheus metric for tracking the total number of broadcasts rejected
var BroadcastRejectedCounterName = "ff_broadcast_rejected_total"

// BroadcastHistogramName is the prometheus metric for tracking the total number of broadcast messages - histogram
var BroadcastHistogramName = "ff_broadcast_histogram"

func InitBroadcastMetrics() {
	BroadcastSubmittedCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: BroadcastSubmittedCounterName,
		Help: "Number of submitted broadcasts",
	})
	BroadcastConfirmedCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: BroadcastConfirmedCounterName,
		Help: "Number of confirmed broadcasts",
	})
	BroadcastRejectedCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: BroadcastRejectedCounterName,
		Help: "Number of rejected broadcasts",
	})
	BroadcastHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: BroadcastHistogramName,
		Help: "Histogram of broadcasts, bucketed by time to finished",
	})
}

func RegisterBroadcastMetrics() {
	registry.MustRegister(BroadcastSubmittedCounter)
	registry.MustRegister(BroadcastConfirmedCounter)
	registry.MustRegister(BroadcastRejectedCounter)
	registry.MustRegister(BroadcastHistogram)
}
