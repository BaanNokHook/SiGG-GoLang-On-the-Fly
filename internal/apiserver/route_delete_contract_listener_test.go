package apiserver

import (
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/mocks/contractmocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteContractListenerByID(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	mcm := &contractmocks.Manager{}
	o.On("Contracts").Return(mcm)
	id := fftypes.NewUUID()
	req := httptest.NewRequest("DELETE", "/api/v1/namespaces/mynamespace/contracts/listeners/"+id.String(), nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mcm.On("DeleteContractListenerByNameOrID", mock.Anything, id.String()).
		Return(nil, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 204, res.Result().StatusCode)
}
