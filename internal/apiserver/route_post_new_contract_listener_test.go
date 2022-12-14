// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly/mocks/contractmocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostNewContractListener(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	mcm := &contractmocks.Manager{}
	o.On("Contracts").Return(mcm)
	input := core.ContractListenerInput{}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(&input)
	req := httptest.NewRequest("POST", "/api/v1/namespaces/mynamespace/contracts/listeners", &buf)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mcm.On("AddContractListener", mock.Anything, mock.AnythingOfType("*core.ContractListenerInput")).
		Return(&core.ContractListener{}, nil, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
