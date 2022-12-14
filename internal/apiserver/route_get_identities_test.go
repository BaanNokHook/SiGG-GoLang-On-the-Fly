package apiserver

import (
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly/mocks/networkmapmocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetIdentities(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	mnm := &networkmapmocks.Manager{}
	o.On("NetworkMap").Return(mnm)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/ns1/identities", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mnm.On("GetIdentities", mock.Anything, mock.Anything).Return([]*core.Identity{}, nil, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}

func TestGetIdentitiesWithVerifiers(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	mnm := &networkmapmocks.Manager{}
	o.On("NetworkMap").Return(mnm)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/ns1/identities?fetchverifiers", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mnm.On("GetIdentitiesWithVerifiers", mock.Anything, mock.Anything).Return([]*core.IdentityWithVerifiers{}, nil, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
