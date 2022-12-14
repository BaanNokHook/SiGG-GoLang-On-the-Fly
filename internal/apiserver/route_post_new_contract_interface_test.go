// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly/mocks/contractmocks"
	"github.com/hyperledger/firefly/mocks/definitionsmocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostNewContractInterface(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	mds := &definitionsmocks.Sender{}
	o.On("Contracts").Return(&contractmocks.Manager{})
	o.On("DefinitionSender").Return(mds)
	input := core.Datatype{}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(&input)
	req := httptest.NewRequest("POST", "/api/v1/namespaces/ns1/contracts/interfaces", &buf)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mds.On("DefineFFI", mock.Anything, mock.AnythingOfType("*fftypes.FFI"), false).Return(nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 202, res.Result().StatusCode)
}

func TestPostNewContractInterfaceSync(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	mds := &definitionsmocks.Sender{}
	o.On("Contracts").Return(&contractmocks.Manager{})
	o.On("DefinitionSender").Return(mds)
	input := core.Datatype{}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(&input)
	req := httptest.NewRequest("POST", "/api/v1/namespaces/ns1/contracts/interfaces?confirm", &buf)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mds.On("DefineFFI", mock.Anything, mock.AnythingOfType("*fftypes.FFI"), true).Return(nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
