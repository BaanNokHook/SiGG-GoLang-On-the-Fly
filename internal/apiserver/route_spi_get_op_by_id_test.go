// SiGG-GoLang-On-the-Fly //

package apiserver

import (
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSPIGetOperationByID(t *testing.T) {
	mgr, _, as := newTestServer()
	r := as.createAdminMuxRouter(mgr)
	req := httptest.NewRequest("GET", "/spi/v1/operations/ns1:0df3d864-2646-4e5d-8585-51eb154a8d23", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mgr.On("GetOperationByNamespacedID", mock.Anything, "ns1:0df3d864-2646-4e5d-8585-51eb154a8d23").
		Return(&core.Operation{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
