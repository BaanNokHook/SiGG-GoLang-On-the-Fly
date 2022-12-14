// SiGG-GoLang-On-the-Fly //

package core

import (
	"context"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/stretchr/testify/assert"
)

func TestTokenPoolValidation(t *testing.T) {
	pool := &TokenPool{
		Namespace: "ok",
		Name:      "!wrong",
	}
	err := pool.Validate(context.Background())
	assert.Regexp(t, "FF00140.*'name'", err)

	pool = &TokenPool{
		Namespace: "ok",
		Name:      "ok",
	}
	err = pool.Validate(context.Background())
	assert.NoError(t, err)
}

func TestTokenPoolDefinition(t *testing.T) {
	pool := &TokenPool{
		Namespace: "ok",
		Name:      "ok",
	}
	var def Definition = &TokenPoolAnnouncement{Pool: pool}
	assert.Equal(t, "73008386c5579b7015385528eb892f7773e13a20015c692f6b90b26e413fe8a4", def.Topic())

	id := fftypes.NewUUID()
	def.SetBroadcastMessage(id)
	assert.Equal(t, id, pool.Message)
}
