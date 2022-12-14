// SiGG-GoLang-On-the-Fly //

package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPint(t *testing.T) {

	p := &Pin{Sequence: 12345}
	var ls LocallySequenced = p
	assert.Equal(t, int64(12345), ls.LocalSequence())

}
