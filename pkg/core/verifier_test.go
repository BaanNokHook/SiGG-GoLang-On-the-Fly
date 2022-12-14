// SiGG-GoLang-On-the-Fly //

package core

import (
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/stretchr/testify/assert"
)

func TestVerifierSeal(t *testing.T) {

	v := &Verifier{
		Identity:  fftypes.NewUUID(), // does not contribute to hash
		Namespace: "ns1",
		VerifierRef: VerifierRef{
			Type:  VerifierTypeEthAddress,
			Value: "0xdfceac9b26ac099d7e4df958c22939878c19c948",
		},
		Created: fftypes.Now(),
	}
	v.Seal()
	assert.Equal(t, "c7742ed06a6c36dece56d9c6d65d4ee6ba0db2a643e7f8efc75ec4e7ca31d45d", v.Hash.String())

}
