// SiGG-GoLang-On-the-Fly //

package tbd

import (
	"context"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/hyperledger/firefly/mocks/identitymocks"
	"github.com/hyperledger/firefly/pkg/identity"
	"github.com/stretchr/testify/assert"
)

var utConfig = config.RootSection("onchain_unit_tests")

func TestInit(t *testing.T) {
	var oc identity.Plugin = &TBD{}
	oc.InitConfig(utConfig)
	err := oc.Init(context.Background(), utConfig)
	assert.NoError(t, err)
	assert.Equal(t, "onchain", oc.Name())
	err = oc.Start()
	assert.NoError(t, err)
	capabilities := oc.Capabilities()
	assert.NotNil(t, capabilities)
	cbs := &identitymocks.Callbacks{}
	oc.SetHandler("ns1", cbs) // no-op
}
