package apiserver

import (
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/mocks/contractmocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetContractListenerByNameOrID(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	mcm := &contractmocks.Manager{}
	o.On("Contracts").Return(mcm)
	id := fftypes.NewUUID()
	req := httptest.NewRequest("GET", "/api/v1/namespaces/mynamespace/contracts/listeners/"+id.String(), nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mcm.On("GetContractListenerByNameOrID", mock.Anything, id.String()).
		Return(&core.ContractListener{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}

func TestGetContractListenerByNameOrIDWithStatus(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	mcm := &contractmocks.Manager{}
	o.On("Contracts").Return(mcm)
	id := fftypes.NewUUID()
	req := httptest.NewRequest("GET", "/api/v1/namespaces/mynamespace/contracts/listeners/"+id.String()+"?fetchstatus", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mcm.On("GetContractListenerByNameOrIDWithStatus", mock.Anything, id.String()).
		Return(&core.ContractListenerWithStatus{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
