// SiGG-GoLang-On-the-Fly //

package networkmap

import (
	"context"
	"testing"

	"github.com/hyperledger/firefly/internal/coreconfig"
	"github.com/hyperledger/firefly/mocks/databasemocks"
	"github.com/hyperledger/firefly/mocks/dataexchangemocks"
	"github.com/hyperledger/firefly/mocks/definitionsmocks"
	"github.com/hyperledger/firefly/mocks/identitymanagermocks"
	"github.com/hyperledger/firefly/mocks/multipartymocks"
	"github.com/hyperledger/firefly/mocks/syncasyncmocks"
	"github.com/stretchr/testify/assert"
)

func newTestNetworkmap(t *testing.T) (*networkMap, func()) {
	coreconfig.Reset()
	ctx, cancel := context.WithCancel(context.Background())
	mdi := &databasemocks.Plugin{}
	mds := &definitionsmocks.Sender{}
	mdx := &dataexchangemocks.Plugin{}
	mim := &identitymanagermocks.Manager{}
	msa := &syncasyncmocks.Bridge{}
	mmp := &multipartymocks.Manager{}
	nm, err := NewNetworkMap(ctx, "ns1", mdi, mdx, mds, mim, msa, mmp)
	assert.NoError(t, err)
	return nm.(*networkMap), cancel

}

func TestNewNetworkMapMissingDep(t *testing.T) {
	_, err := NewNetworkMap(context.Background(), "", nil, nil, nil, nil, nil, nil)
	assert.Regexp(t, "FF10128", err)
}
