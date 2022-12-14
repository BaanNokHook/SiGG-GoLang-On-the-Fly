package apiserver

import (
	"github.com/hyperledger/firefly-common/pkg/config"
)

const (
	MetricsEnabled = "enabled"
	MetricsPath    = "path"
)

func initMetricsConfig(config config.Section) {
	config.AddKnownKey(MetricsEnabled, true)
	config.AddKnownKey(MetricsPath, "/metrics")
}
