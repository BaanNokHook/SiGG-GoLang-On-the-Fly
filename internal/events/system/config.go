// SiGG-GoLang-On-the-Fly //
package system

import (
	"github.com/hyperledger/firefly-common/pkg/config"
)

const (
	readAhead = 50
)

const (
	// SystemEventsConfReadAhead is the readahead used for system events
	SystemEventsConfReadAhead = "readAhead"
)

func (se *Events) InitConfig(config config.Section) {
	config.AddKnownKey(SystemEventsConfReadAhead, readAhead)
}
