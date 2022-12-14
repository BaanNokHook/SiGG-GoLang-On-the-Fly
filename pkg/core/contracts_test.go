// SiGG-GoLang-On-the-Fly //

package core

import (
	"context"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/stretchr/testify/assert"
)

func TestValidateContractAPI(t *testing.T) {
	api := &ContractAPI{
		Namespace: "ns1",
		Name:      "banana",
	}
	err := api.Validate(context.Background(), false)
	assert.NoError(t, err)
}

func TestValidateInvalidContractAPI(t *testing.T) {
	api := &ContractAPI{
		Namespace: "&%&^#()#",
		Name:      "banana",
	}
	err := api.Validate(context.Background(), false)
	assert.Regexp(t, "FF00140", err)

	api = &ContractAPI{
		Namespace: "ns1",
		Name:      "(%&@!^%^)",
	}
	err = api.Validate(context.Background(), false)
	assert.Regexp(t, "FF00140", err)
}

func TestContractAPITopic(t *testing.T) {
	api := &ContractAPI{
		Namespace: "ns1",
	}
	assert.Equal(t, "4cccc66c1f0eebcf578f1e63b73a2047d4eb4c84c0a00c69b0e00c7490403d20", api.Topic())
}

func TestContractAPISetBroadCastMessage(t *testing.T) {
	msgID := fftypes.NewUUID()
	api := &ContractAPI{}
	api.SetBroadcastMessage(msgID)
	assert.Equal(t, api.Message, msgID)
}

func TestLocationAndLedgerEquals(t *testing.T) {
	var c1 *ContractAPI = nil
	var c2 *ContractAPI = nil
	assert.False(t, c1.LocationAndLedgerEquals(c2))

	c1 = &ContractAPI{
		ID:       fftypes.NewUUID(),
		Location: fftypes.JSONAnyPtr("abc"),
	}
	c2 = &ContractAPI{
		ID:       fftypes.NewUUID(),
		Location: fftypes.JSONAnyPtr("abc"),
	}
	assert.True(t, c1.LocationAndLedgerEquals(c2))

	c1 = &ContractAPI{
		ID:       fftypes.NewUUID(),
		Location: fftypes.JSONAnyPtr("fff"),
	}
	c2 = &ContractAPI{
		ID:       fftypes.NewUUID(),
		Location: fftypes.JSONAnyPtr("abc"),
	}
	assert.False(t, c1.LocationAndLedgerEquals(c2))

	c1 = &ContractAPI{
		ID:       fftypes.NewUUID(),
		Location: nil,
	}
	c2 = &ContractAPI{
		ID:       fftypes.NewUUID(),
		Location: nil,
	}
	assert.True(t, c1.LocationAndLedgerEquals(c2))
}
