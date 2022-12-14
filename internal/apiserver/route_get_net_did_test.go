// SiGG-GoLang-On-the-Fly //

package apiserver

import (
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly/mocks/networkmapmocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetNetIdentityByDID(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	nmn := &networkmapmocks.Manager{}
	o.On("NetworkMap").Return(nmn)
	req := httptest.NewRequest("GET", "/api/v1/network/identities/did:firefly:org/org_1", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	nmn.On("GetIdentityByDID", mock.Anything, "did:firefly:org/org_1").
		Return(&core.Identity{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}

func TestGetNetIdentityByDIDWithVerifiers(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	nmn := &networkmapmocks.Manager{}
	o.On("NetworkMap").Return(nmn)
	req := httptest.NewRequest("GET", "/api/v1/network/identities/did:firefly:org/org_1?fetchverifiers", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	nmn.On("GetIdentityByDIDWithVerifiers", mock.Anything, "did:firefly:org/org_1").
		Return(&core.IdentityWithVerifiers{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
