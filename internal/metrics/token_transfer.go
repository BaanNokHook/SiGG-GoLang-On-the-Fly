// SiGG-GoLang-On-the-Fly //
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var TransferSubmittedCounter prometheus.Counter
var TransferConfirmedCounter prometheus.Counter
var TransferRejectedCounter prometheus.Counter
var TransferHistogram prometheus.Histogram

// TransferSubmittedCounterName is the prometheus metric for tracking the total number of transfers submitted
var TransferSubmittedCounterName = "ff_transfer_submitted_total"

// TransferConfirmedCounterName is the prometheus metric for tracking the total number of transfers confirmed
var TransferConfirmedCounterName = "ff_transfer_confirmed_total"

// TransferRejectedCounterName is the prometheus metric for tracking the total number of transfers rejected
var TransferRejectedCounterName = "ff_transfer_rejected_total"

// TransferHistogramName is the prometheus metric for tracking the total number of transfers - histogram
var TransferHistogramName = "ff_transfer_histogram"

func InitTokenTransferMetrics() {
	TransferSubmittedCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: TransferSubmittedCounterName,
		Help: "Number of submitted transfers",
	})
	TransferConfirmedCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: TransferConfirmedCounterName,
		Help: "Number of confirmed transfers",
	})
	TransferRejectedCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: TransferRejectedCounterName,
		Help: "Number of rejected transfers",
	})
	TransferHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: TransferHistogramName,
		Help: "Histogram of transfers, bucketed by time to finished",
	})
}

func RegisterTokenTransferMetrics() {
	registry.MustRegister(TransferSubmittedCounter)
	registry.MustRegister(TransferConfirmedCounter)
	registry.MustRegister(TransferRejectedCounter)
	registry.MustRegister(TransferHistogram)
}
