// SiGG-GoLang-On-the-Fly //

package apiserver

import (
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetSubscriptionByID(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/mynamespace/subscriptions/abcd12345", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	o.On("GetSubscriptionByID", mock.Anything, "abcd12345").
		Return(&core.Subscription{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}

func TestGetSubscriptionByIDWithStatus(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/mynamespace/subscriptions/abcd12345?fetchstatus", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	o.On("GetSubscriptionByID", mock.Anything, "abcd12345").
		Return(&core.Subscription{}, nil)
	o.On("GetSubscriptionByIDWithStatus", mock.Anything, "abcd12345").
		Return(&core.SubscriptionWithStatus{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
