// SiGG-GoLang-On-the-Fly //

package apiserver

import (
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSPIGetNamespaces(t *testing.T) {
	mgr, _, as := newTestServer()
	r := as.createAdminMuxRouter(mgr)
	req := httptest.NewRequest("GET", "/spi/v1/namespaces", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mgr.On("GetNamespaces", mock.Anything, mock.Anything).
		Return([]*core.Namespace{}, nil, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
