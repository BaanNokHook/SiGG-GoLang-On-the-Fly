// SiGG-GoLang-On-the-Fly //

package apiserver

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSPIPatchOperationByID(t *testing.T) {
	mgr, _, as := newTestServer()
	r := as.createAdminMuxRouter(mgr)
	req := httptest.NewRequest("PATCH", "/spi/v1/operations/ns1:0df3d864-2646-4e5d-8585-51eb154a8d23", bytes.NewReader([]byte("{}")))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mgr.On("ResolveOperationByNamespacedID", mock.Anything, "ns1:0df3d864-2646-4e5d-8585-51eb154a8d23", mock.AnythingOfType("*core.OperationUpdateDTO")).Return(nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
