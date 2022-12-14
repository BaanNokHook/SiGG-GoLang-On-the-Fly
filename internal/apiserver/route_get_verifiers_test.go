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

func TestGetVerifiers(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	mnm := &networkmapmocks.Manager{}
	o.On("NetworkMap").Return(mnm)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/ns1/verifiers", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mnm.On("GetVerifiers", mock.Anything, mock.Anything).Return([]*core.Verifier{}, nil, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
