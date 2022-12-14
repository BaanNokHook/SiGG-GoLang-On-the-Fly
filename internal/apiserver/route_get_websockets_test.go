// SiGG-GoLang-On-the-Fly //

package apiserver

import (
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly/mocks/eventmocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetWebSockets(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	mem := &eventmocks.EventManager{}
	o.On("Events").Return(mem)
	req := httptest.NewRequest("GET", "/api/v1/websockets", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mem.On("GetWebSocketStatus").Return(&core.WebSocketStatus{})
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
