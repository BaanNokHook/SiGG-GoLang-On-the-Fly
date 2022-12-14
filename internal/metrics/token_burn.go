// SiGG-GoLang-On-the-Fly //
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var BurnSubmittedCounter prometheus.Counter
var BurnConfirmedCounter prometheus.Counter
var BurnRejectedCounter prometheus.Counter
var BurnHistogram prometheus.Histogram

// BurnSubmittedCounterName is the prometheus metric for tracking the total number of burns submitted
var BurnSubmittedCounterName = "ff_burn_submitted_total"

// BurnConfirmedCounterName is the prometheus metric for tracking the total number of burns confirmed
var BurnConfirmedCounterName = "ff_burn_confirmed_total"

// BurnRejectedCounterName is the prometheus metric for tracking the total number of burns rejected
var BurnRejectedCounterName = "ff_burn_rejected_total"

// BurnHistogramName is the prometheus metric for tracking the total number of burns - histogram
var BurnHistogramName = "ff_burn_histogram"

func InitTokenBurnMetrics() {
	BurnSubmittedCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: BurnSubmittedCounterName,
		Help: "Number of submitted burns",
	})
	BurnConfirmedCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: BurnConfirmedCounterName,
		Help: "Number of confirmed burns",
	})
	BurnRejectedCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: BurnRejectedCounterName,
		Help: "Number of rejected burns",
	})
	BurnHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: BurnHistogramName,
		Help: "Histogram of burns, bucketed by time to finished",
	})
}

func RegisterTokenBurnMetrics() {
	registry.MustRegister(BurnSubmittedCounter)
	registry.MustRegister(BurnConfirmedCounter)
	registry.MustRegister(BurnRejectedCounter)
	registry.MustRegister(BurnHistogram)
}
