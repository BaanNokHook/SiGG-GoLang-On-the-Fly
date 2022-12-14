// SiGG-GoLang-On-the-Fly //

package core

import (
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/stretchr/testify/assert"
)

func TestTokenBalanceIdentifier(t *testing.T) {
	id := fftypes.NewUUID()
	balance := &TokenBalance{
		Pool:       id,
		TokenIndex: "1",
		Key:        "0x00",
	}
	assert.Equal(t, id.String()+":1:0x00", balance.Identifier())
}
