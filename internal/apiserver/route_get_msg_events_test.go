package apiserver

import (
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetMessageEvents(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/mynamespace/messages/uuid1/events", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	o.On("GetMessageEvents", mock.Anything, "uuid1", mock.Anything).
		Return([]*core.Event{}, nil, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
