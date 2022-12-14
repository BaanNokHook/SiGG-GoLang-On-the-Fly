// SiGG-GoLang-On-the-Fly //
package coreconfig

import (
	"testing"

	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/stretchr/testify/assert"
)

const configDir = "../../test/data/config"

func TestInitConfigOK(t *testing.T) {
	Reset()

	assert.Equal(t, 25, config.GetInt(APIDefaultFilterLimit))
	assert.Equal(t, "localhost", config.GetString(DebugAddress))
	assert.Equal(t, -1, config.GetInt(DebugPort))
}
