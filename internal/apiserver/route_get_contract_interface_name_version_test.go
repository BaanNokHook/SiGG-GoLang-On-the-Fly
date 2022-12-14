package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/mocks/contractmocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetContractInterfaceNameVersion(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	mcm := &contractmocks.Manager{}
	o.On("Contracts").Return(mcm)
	input := core.Datatype{}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(&input)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/ns1/contracts/interfaces/banana/v1.0.0", &buf)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mcm.On("GetFFI", mock.Anything, "banana", "v1.0.0").
		Return(&fftypes.FFI{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}

func TestGetContractInterfaceNameVersionWithChildren(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	mcm := &contractmocks.Manager{}
	o.On("Contracts").Return(mcm)
	input := core.Datatype{}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(&input)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/ns1/contracts/interfaces/banana/v1.0.0?fetchchildren", &buf)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mcm.On("GetFFIWithChildren", mock.Anything, "banana", "v1.0.0").
		Return(&fftypes.FFI{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
