// SiGG-GoLang-On-the-Fly //

package networkmap

import (
	"fmt"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/internal/multiparty"
	"github.com/hyperledger/firefly/mocks/dataexchangemocks"
	"github.com/hyperledger/firefly/mocks/definitionsmocks"
	"github.com/hyperledger/firefly/mocks/identitymanagermocks"
	"github.com/hyperledger/firefly/mocks/multipartymocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterNodeOk(t *testing.T) {

	nm, cancel := newTestNetworkmap(t)
	defer cancel()

	parentOrg := testOrg("org1")
	signerRef := &core.SignerRef{Key: "0x23456"}

	mim := nm.identity.(*identitymanagermocks.Manager)
	mim.On("GetMultipartyRootOrg", nm.ctx).Return(parentOrg, nil)
	mim.On("VerifyIdentityChain", nm.ctx, mock.AnythingOfType("*core.Identity")).Return(parentOrg, false, nil)
	mim.On("ResolveIdentitySigner", nm.ctx, parentOrg).Return(signerRef, nil)

	mdx := nm.exchange.(*dataexchangemocks.Plugin)
	mdx.On("GetEndpointInfo", nm.ctx, "node1").Return(fftypes.JSONObject{
		"id":       "peer1",
		"endpoint": "details",
	}, nil)

	mds := nm.defsender.(*definitionsmocks.Sender)
	mds.On("ClaimIdentity", nm.ctx,
		mock.AnythingOfType("*core.IdentityClaim"),
		signerRef,
		(*core.SignerRef)(nil),
		false).Return(nil)

	mmp := nm.multiparty.(*multipartymocks.Manager)
	mmp.On("LocalNode").Return(multiparty.LocalNode{Name: "node1"})

	node, err := nm.RegisterNode(nm.ctx, false)
	assert.NoError(t, err)
	assert.NotNil(t, node)

	mim.AssertExpectations(t)
	mdx.AssertExpectations(t)
	mds.AssertExpectations(t)
	mmp.AssertExpectations(t)
}

func TestRegisterNodeMissingName(t *testing.T) {

	nm, cancel := newTestNetworkmap(t)
	defer cancel()

	parentOrg := testOrg("org1")

	mim := nm.identity.(*identitymanagermocks.Manager)
	mim.On("GetMultipartyRootOrg", nm.ctx).Return(parentOrg, nil)

	mmp := nm.multiparty.(*multipartymocks.Manager)
	mmp.On("LocalNode").Return(multiparty.LocalNode{})

	_, err := nm.RegisterNode(nm.ctx, false)
	assert.Regexp(t, "FF10216", err)

	mim.AssertExpectations(t)
	mmp.AssertExpectations(t)
}

func TestRegisterNodePeerInfoFail(t *testing.T) {

	nm, cancel := newTestNetworkmap(t)
	defer cancel()

	parentOrg := testOrg("org1")

	mim := nm.identity.(*identitymanagermocks.Manager)
	mim.On("GetMultipartyRootOrg", nm.ctx).Return(parentOrg, nil)

	mdx := nm.exchange.(*dataexchangemocks.Plugin)
	mdx.On("GetEndpointInfo", nm.ctx, "node1").Return(fftypes.JSONObject{}, fmt.Errorf("pop"))

	mmp := nm.multiparty.(*multipartymocks.Manager)
	mmp.On("LocalNode").Return(multiparty.LocalNode{Name: "node1"})

	_, err := nm.RegisterNode(nm.ctx, false)
	assert.Regexp(t, "pop", err)

	mim.AssertExpectations(t)
	mdx.AssertExpectations(t)
	mmp.AssertExpectations(t)
}

func TestRegisterNodeGetOwnerFail(t *testing.T) {

	nm, cancel := newTestNetworkmap(t)
	defer cancel()

	mim := nm.identity.(*identitymanagermocks.Manager)
	mim.On("GetMultipartyRootOrg", nm.ctx).Return(nil, fmt.Errorf("pop"))

	_, err := nm.RegisterNode(nm.ctx, false)
	assert.Regexp(t, "pop", err)

}
