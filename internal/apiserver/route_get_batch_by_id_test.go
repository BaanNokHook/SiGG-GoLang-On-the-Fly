package apiserver

import (
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetBatchByID(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/mynamespace/batches/abcd12345", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	o.On("GetBatchByID", mock.Anything, "abcd12345").
		Return(&core.BatchPersisted{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
