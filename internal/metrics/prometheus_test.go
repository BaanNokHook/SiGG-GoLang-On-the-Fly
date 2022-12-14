// SiGG-GoLang-On-the-Fly //
package metrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrometheusMiddleware(t *testing.T) {
	Registry()
	adminInstrumentation = nil
	restInstrumentation = nil
	assert.NotNil(t, GetAdminServerInstrumentation())
	assert.NotNil(t, GetRestServerInstrumentation())
}
