package apiserver

import (
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetEventByID(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/mynamespace/events/abcd12345", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	o.On("GetEventByID", mock.Anything, "abcd12345").
		Return(&core.Event{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}

func TestGetEventByIDWithReference(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/mynamespace/events/abcd12345?fetchreference", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	o.On("GetEventByIDWithReference", mock.Anything, "abcd12345").
		Return(&core.EnrichedEvent{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
