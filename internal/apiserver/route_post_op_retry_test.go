// SiGG-GoLang-On-the-Fly //

package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/mocks/operationmocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostOpRetry(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	mom := &operationmocks.Manager{}
	o.On("Operations").Return(mom)
	input := core.EmptyInput{}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(&input)
	opID := fftypes.NewUUID()
	req := httptest.NewRequest("POST", "/api/v1/namespaces/ns1/operations/"+opID.String()+"/retry", &buf)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mom.On("RetryOperation", mock.Anything, opID).
		Return(&core.Operation{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 202, res.Result().StatusCode)
}

func TestPostOpRetryBadID(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	input := core.EmptyInput{}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(&input)
	req := httptest.NewRequest("POST", "/api/v1/namespaces/ns1/operations/bad/retry", &buf)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, 400, res.Result().StatusCode)
}
