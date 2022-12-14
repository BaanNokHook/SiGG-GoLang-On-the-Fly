// SiGG-GoLang-On-the-Fly //
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var BlockchainTransactionsCounter *prometheus.CounterVec
var BlockchainQueriesCounter *prometheus.CounterVec
var BlockchainEventsCounter *prometheus.CounterVec

// BlockchainTransactionsCounterName is the prometheus metric for tracking the total number of blockchain transactions
var BlockchainTransactionsCounterName = "ff_blockchain_transactions_total"

// BlockchainQueriesCounterName is the prometheus metric for tracking the total number of blockchain queries
var BlockchainQueriesCounterName = "ff_blockchain_queries_total"

// BlockchainEventsCounterName is the prometheus metric for tracking the total number of blockchain events
var BlockchainEventsCounterName = "ff_blockchain_events_total"

var LocationLabelName = "location"
var MethodNameLabelName = "methodName"
var SignatureLabelName = "signature"

func InitBlockchainMetrics() {
	BlockchainTransactionsCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: BlockchainTransactionsCounterName,
		Help: "Number of blockchain transactions",
	}, []string{LocationLabelName, MethodNameLabelName})
	BlockchainQueriesCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: BlockchainQueriesCounterName,
		Help: "Number of blockchain queries",
	}, []string{LocationLabelName, MethodNameLabelName})
	BlockchainEventsCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: BlockchainEventsCounterName,
		Help: "Number of blockchain events",
	}, []string{LocationLabelName, SignatureLabelName})
}

func RegisterBlockchainMetrics() {
	registry.MustRegister(BlockchainTransactionsCounter)
	registry.MustRegister(BlockchainQueriesCounter)
	registry.MustRegister(BlockchainEventsCounter)
}
