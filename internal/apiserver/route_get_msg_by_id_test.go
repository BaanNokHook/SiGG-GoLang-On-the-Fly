// SiGG-GoLang-On-the-Fly //

package apiserver

import (
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetMessageByID(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/mynamespace/messages/abcd12345", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	o.On("GetMessageByID", mock.Anything, "abcd12345").
		Return(&core.Message{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}

func TestGetMessageByIDWithData(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/mynamespace/messages/abcd12345?fetchdata", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	o.On("GetMessageByIDWithData", mock.Anything, "abcd12345").
		Return(&core.MessageInOut{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
