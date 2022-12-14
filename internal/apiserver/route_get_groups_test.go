package apiserver

import (
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly/mocks/privatemessagingmocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetGroups(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/mynamespace/groups", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mpm := &privatemessagingmocks.Manager{}
	o.On("PrivateMessaging").Return(mpm)
	mpm.On("GetGroups", mock.Anything, mock.Anything).
		Return([]*core.Group{}, nil, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
