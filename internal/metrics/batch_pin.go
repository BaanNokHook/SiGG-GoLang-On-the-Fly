// SiGG-GoLang-On-the-Fly //
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var BatchPinCounter prometheus.Counter

// MetricsBatchPin is the prometheus metric for total number of batch pins submitted
var MetricsBatchPin = "ff_batchpin_total"

func InitBatchPinMetrics() {
	BatchPinCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: MetricsBatchPin,
		Help: "Number of batch pins submitted",
	})
}

func RegisterBatchPinMetrics() {
	registry.MustRegister(BatchPinCounter)
}
