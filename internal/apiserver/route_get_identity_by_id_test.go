package apiserver

import (
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly/mocks/networkmapmocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetIdentityByID(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	mnm := &networkmapmocks.Manager{}
	o.On("NetworkMap").Return(mnm)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/ns1/identities/id1", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mnm.On("GetIdentityByID", mock.Anything, "id1").Return(&core.Identity{}, nil, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}

func TestGetIdentityByIDWithVerifiers(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	mnm := &networkmapmocks.Manager{}
	o.On("NetworkMap").Return(mnm)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/ns1/identities/id1?fetchverifiers", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mnm.On("GetIdentityByIDWithVerifiers", mock.Anything, "id1").Return(&core.IdentityWithVerifiers{}, nil, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
