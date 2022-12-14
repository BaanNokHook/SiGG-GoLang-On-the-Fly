// SiGG-GoLang-On-the-Fly //

package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTXSizeEstimate(t *testing.T) {
	assert.Equal(t, transactionBaseSizeEstimate, (&Transaction{}).Size())
}
