// SiGG-GoLang-On-the-Fly //
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var PrivateMsgSubmittedCounter prometheus.Counter
var PrivateMsgConfirmedCounter prometheus.Counter
var PrivateMsgRejectedCounter prometheus.Counter
var PrivateMsgHistogram prometheus.Histogram

// PrivateMsgSubmittedCounterName is the prometheus metric for tracking the total number of private messages submitted
var PrivateMsgSubmittedCounterName = "ff_private_msg_submitted_total"

// PrivateMsgConfirmedCounterName is the prometheus metric for tracking the total number of private messages confirmed
var PrivateMsgConfirmedCounterName = "ff_private_msg_confirmed_total"

// PrivateMsgRejectedCounterName is the prometheus metric for tracking the total number of private messages rejected
var PrivateMsgRejectedCounterName = "ff_private_msg_rejected_total"

// PrivateMsgHistogramName is the prometheus metric for tracking the total number of private messages - histogram
var PrivateMsgHistogramName = "ff_private_msg_histogram"

func InitPrivateMsgMetrics() {
	PrivateMsgSubmittedCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: PrivateMsgSubmittedCounterName,
		Help: "Number of submitted private messages",
	})
	PrivateMsgConfirmedCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: PrivateMsgConfirmedCounterName,
		Help: "Number of confirmed private messages",
	})
	PrivateMsgRejectedCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: PrivateMsgRejectedCounterName,
		Help: "Number of rejected private messages",
	})
	PrivateMsgHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: PrivateMsgHistogramName,
		Help: "Histogram of private messages, bucketed by time to finished",
	})
}

func RegisterPrivateMsgMetrics() {
	registry.MustRegister(PrivateMsgSubmittedCounter)
	registry.MustRegister(PrivateMsgConfirmedCounter)
	registry.MustRegister(PrivateMsgRejectedCounter)
	registry.MustRegister(PrivateMsgHistogram)
}
