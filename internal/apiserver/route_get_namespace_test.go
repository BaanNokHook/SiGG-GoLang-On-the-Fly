// SiGG-GoLang-On-the-Fly //

package apiserver

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetNamespace(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/ns1", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	o.On("GetNamespace", mock.Anything).
		Return(&core.Namespace{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}

func TestGetNamespaceInvalid(t *testing.T) {
	mgr, o, as := newTestServer()
	r := as.createMuxRouter(context.Background(), mgr)
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/BAD", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mgr.On("Orchestrator", "BAD").Return(nil, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 404, res.Result().StatusCode)
}
