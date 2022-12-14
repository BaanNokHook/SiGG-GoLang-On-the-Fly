// SiGG-GoLang-On-the-Fly //

package apiserver

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAdminPostResetConfig(t *testing.T) {
	mgr, _, as := newTestServer()
	r := as.createAdminMuxRouter(mgr)
	req := httptest.NewRequest("POST", "/spi/v1/reset", bytes.NewReader([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mgr.On("Reset", mock.Anything).Return()
	r.ServeHTTP(res, req)

	assert.Equal(t, 204, res.Result().StatusCode)
}
