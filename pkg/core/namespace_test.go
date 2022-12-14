// SiGG-GoLang-On-the-Fly //

package core

import (
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/stretchr/testify/assert"
)

func TestMultipartyContractsDatabaseSerialization(t *testing.T) {
	contracts1 := &MultipartyContracts{
		Active: &MultipartyContract{
			Index:      1,
			FirstEvent: "oldest",
			Location: fftypes.JSONAnyPtr(fftypes.JSONObject{
				"address": "0x123",
			}.String()),
			Info: MultipartyContractInfo{
				Subscription: "1234",
			},
		},
		Terminated: []*MultipartyContract{
			{
				Index:      0,
				FirstEvent: "oldest",
				Location: fftypes.JSONAnyPtr(fftypes.JSONObject{
					"address": "0x1234",
				}.String()),
				Info: MultipartyContractInfo{
					Subscription: "12345",
					FinalEvent:   "50",
				},
			},
		},
	}

	// Verify it serializes as bytes to the database
	val1, err := contracts1.Value()
	assert.NoError(t, err)
	assert.Equal(t, `{"active":{"index":1,"location":{"address":"0x123"},"firstEvent":"oldest","info":{"subscription":"1234"}},"terminated":[{"index":0,"location":{"address":"0x1234"},"firstEvent":"oldest","info":{"subscription":"12345","finalEvent":"50"}}]}`, string(val1.([]byte)))

	// Verify it restores ok
	contracts2 := &MultipartyContracts{}
	err = contracts2.Scan(val1)
	assert.NoError(t, err)
	assert.Equal(t, 1, contracts2.Active.Index)
	assert.Equal(t, *fftypes.JSONAnyPtr(fftypes.JSONObject{
		"address": "0x123",
	}.String()), *contracts2.Active.Location)
	assert.Len(t, contracts2.Terminated, 1)

	// Verify it ignores a blank string
	err = contracts2.Scan("")
	assert.NoError(t, err)
	assert.Equal(t, 1, contracts2.Active.Index)

	// Out of luck with anything else
	err = contracts2.Scan(false)
	assert.Regexp(t, "FF00105", err)
}
