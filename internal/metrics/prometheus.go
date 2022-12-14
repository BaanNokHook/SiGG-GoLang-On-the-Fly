// SiGG-GoLang-On-the-Fly //
package metrics

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	muxprom "gitlab.com/hfuss/mux-prometheus/pkg/middleware"
)

var regMux sync.Mutex
var registry *prometheus.Registry
var adminInstrumentation *muxprom.Instrumentation
var restInstrumentation *muxprom.Instrumentation

// Registry returns FireFly's customized Prometheus registry
func Registry() *prometheus.Registry {
	if registry == nil {
		initMetricsCollectors()
		registry = prometheus.NewRegistry()
		registerMetricsCollectors()
	}

	return registry
}

// GetAdminServerInstrumentation returns the admin server's Prometheus middleware, ensuring its metrics are never
// registered twice
func GetAdminServerInstrumentation() *muxprom.Instrumentation {
	regMux.Lock()
	defer regMux.Unlock()
	if adminInstrumentation == nil {
		adminInstrumentation = NewInstrumentation("admin")
	}
	return adminInstrumentation
}

// GetRestServerInstrumentation returns the REST server's Prometheus middleware, ensuring its metrics are never
// registered twice
func GetRestServerInstrumentation() *muxprom.Instrumentation {
	regMux.Lock()
	defer regMux.Unlock()
	if restInstrumentation == nil {
		restInstrumentation = NewInstrumentation("rest")
	}
	return restInstrumentation
}

func NewInstrumentation(subsystem string) *muxprom.Instrumentation {
	return muxprom.NewCustomInstrumentation(
		true,
		"ff_apiserver",
		subsystem,
		prometheus.DefBuckets,
		map[string]string{},
		Registry(),
	)
}

// Clear will reset the Prometheus metrics registry and instrumentations, useful for testing
func Clear() {
	registry = nil
	adminInstrumentation = nil
	restInstrumentation = nil
}

func initMetricsCollectors() {
	InitBroadcastMetrics()
	InitPrivateMsgMetrics()
	InitTokenMintMetrics()
	InitTokenTransferMetrics()
	InitTokenBurnMetrics()
	InitBatchPinMetrics()
	InitBlockchainMetrics()
}

func registerMetricsCollectors() {
	registry.MustRegister(collectors.NewGoCollector())
	registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	RegisterBatchPinMetrics()
	RegisterBroadcastMetrics()
	RegisterPrivateMsgMetrics()
	RegisterTokenMintMetrics()
	RegisterTokenTransferMetrics()
	RegisterTokenBurnMetrics()
	RegisterBlockchainMetrics()
}
