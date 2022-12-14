// SiGG-GoLang-On-the-Fly //

package networkmap

import (
	"fmt"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/mocks/definitionsmocks"
	"github.com/hyperledger/firefly/mocks/identitymanagermocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateIdentityProfileOk(t *testing.T) {

	nm, cancel := newTestNetworkmap(t)
	defer cancel()

	identity := testOrg("org1")

	mim := nm.identity.(*identitymanagermocks.Manager)
	mim.On("CachedIdentityLookupByID", nm.ctx, identity.ID).Return(identity, nil)
	signerRef := &core.SignerRef{Key: "0x12345"}
	mim.On("ResolveIdentitySigner", nm.ctx, identity).Return(signerRef, nil)

	mds := nm.defsender.(*definitionsmocks.Sender)

	mds.On("UpdateIdentity", nm.ctx,
		mock.AnythingOfType("*core.Identity"),
		mock.AnythingOfType("*core.IdentityUpdate"),
		mock.MatchedBy(func(sr *core.SignerRef) bool {
			return sr.Key == "0x12345"
		}),
		true).Return(nil)

	org, err := nm.UpdateIdentity(nm.ctx, identity.ID.String(), &core.IdentityUpdateDTO{
		IdentityProfile: core.IdentityProfile{
			Description: "new desc",
			Profile:     fftypes.JSONObject{"new": "profile"},
		},
	}, true)
	assert.NoError(t, err)
	assert.NotNil(t, org)

	mim.AssertExpectations(t)
	mds.AssertExpectations(t)
}

func TestUpdateIdentityProfileBroadcastFail(t *testing.T) {

	nm, cancel := newTestNetworkmap(t)
	defer cancel()

	identity := testOrg("org1")

	mim := nm.identity.(*identitymanagermocks.Manager)
	mim.On("CachedIdentityLookupByID", nm.ctx, identity.ID).Return(identity, nil)
	signerRef := &core.SignerRef{Key: "0x12345"}
	mim.On("ResolveIdentitySigner", nm.ctx, identity).Return(signerRef, nil)

	mds := nm.defsender.(*definitionsmocks.Sender)
	mds.On("UpdateIdentity", nm.ctx,
		mock.AnythingOfType("*core.Identity"),
		mock.AnythingOfType("*core.IdentityUpdate"),
		mock.MatchedBy(func(sr *core.SignerRef) bool {
			return sr.Key == "0x12345"
		}),
		true).Return(fmt.Errorf("pop"))

	_, err := nm.UpdateIdentity(nm.ctx, identity.ID.String(), &core.IdentityUpdateDTO{
		IdentityProfile: core.IdentityProfile{
			Description: "new desc",
			Profile:     fftypes.JSONObject{"new": "profile"},
		},
	}, true)
	assert.Regexp(t, "pop", err)

	mim.AssertExpectations(t)
	mds.AssertExpectations(t)
}

func TestUpdateIdentityProfileBadProfile(t *testing.T) {

	nm, cancel := newTestNetworkmap(t)
	defer cancel()

	identity := testOrg("org1")

	mim := nm.identity.(*identitymanagermocks.Manager)
	mim.On("CachedIdentityLookupByID", nm.ctx, identity.ID).Return(identity, nil)
	signerRef := &core.SignerRef{Key: "0x12345"}
	mim.On("ResolveIdentitySigner", nm.ctx, identity).Return(signerRef, nil)

	_, err := nm.UpdateIdentity(nm.ctx, identity.ID.String(), &core.IdentityUpdateDTO{
		IdentityProfile: core.IdentityProfile{
			Description: string(make([]byte, 4097)),
			Profile:     fftypes.JSONObject{"new": "profile"},
		},
	}, true)
	assert.Regexp(t, "FF00135", err)

	mim.AssertExpectations(t)
}

func TestUpdateIdentityProfileNotFound(t *testing.T) {

	nm, cancel := newTestNetworkmap(t)
	defer cancel()

	identity := testOrg("org1")

	mim := nm.identity.(*identitymanagermocks.Manager)
	mim.On("CachedIdentityLookupByID", nm.ctx, identity.ID).Return(nil, nil)

	_, err := nm.UpdateIdentity(nm.ctx, identity.ID.String(), &core.IdentityUpdateDTO{
		IdentityProfile: core.IdentityProfile{
			Description: string(make([]byte, 4097)),
			Profile:     fftypes.JSONObject{"new": "profile"},
		},
	}, true)
	assert.Regexp(t, "FF10143", err)

	mim.AssertExpectations(t)
}

func TestUpdateIdentityProfileLookupFail(t *testing.T) {

	nm, cancel := newTestNetworkmap(t)
	defer cancel()

	identity := testOrg("org1")

	mim := nm.identity.(*identitymanagermocks.Manager)
	mim.On("CachedIdentityLookupByID", nm.ctx, identity.ID).Return(nil, fmt.Errorf("pop"))

	_, err := nm.UpdateIdentity(nm.ctx, identity.ID.String(), &core.IdentityUpdateDTO{
		IdentityProfile: core.IdentityProfile{
			Description: string(make([]byte, 4097)),
			Profile:     fftypes.JSONObject{"new": "profile"},
		},
	}, true)
	assert.Regexp(t, "pop", err)

	mim.AssertExpectations(t)
}

func TestUpdateIdentityProfileClaimLookupFail(t *testing.T) {

	nm, cancel := newTestNetworkmap(t)
	defer cancel()

	identity := testOrg("org1")

	mim := nm.identity.(*identitymanagermocks.Manager)
	mim.On("CachedIdentityLookupByID", nm.ctx, identity.ID).Return(identity, nil)
	signerRef := &core.SignerRef{Key: "0x12345"}
	mim.On("ResolveIdentitySigner", nm.ctx, identity).Return(signerRef, fmt.Errorf("pop"))

	_, err := nm.UpdateIdentity(nm.ctx, identity.ID.String(), &core.IdentityUpdateDTO{
		IdentityProfile: core.IdentityProfile{
			Description: "Desc1",
			Profile:     fftypes.JSONObject{"new": "profile"},
		},
	}, true)
	assert.Regexp(t, "pop", err)

	mim.AssertExpectations(t)
}

func TestUpdateIdentityProfileBadID(t *testing.T) {

	nm, cancel := newTestNetworkmap(t)
	defer cancel()

	_, err := nm.UpdateIdentity(nm.ctx, "badness", &core.IdentityUpdateDTO{}, true)
	assert.Regexp(t, "FF00138", err)
}
