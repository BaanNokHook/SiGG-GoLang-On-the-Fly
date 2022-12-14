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

func TestPutContractAPI(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	mds := &definitionsmocks.Sender{}
	o.On("DefinitionSender").Return(mds)
	o.On("Contracts").Return(&contractmocks.Manager{})
	input := core.Datatype{}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(&input)
	req := httptest.NewRequest("PUT", "/api/v1/namespaces/ns1/apis/99EEE458-037C-4C78-B66B-31E52F93D2E9", &buf)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mds.On("DefineContractAPI", mock.Anything, mock.Anything, mock.AnythingOfType("*core.ContractAPI"), false).Return(nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 202, res.Result().StatusCode)
}

func TestPutContractAPISync(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	mds := &definitionsmocks.Sender{}
	o.On("DefinitionSender").Return(mds)
	o.On("Contracts").Return(&contractmocks.Manager{})
	input := core.Datatype{}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(&input)
	req := httptest.NewRequest("PUT", "/api/v1/namespaces/ns1/apis/99EEE458-037C-4C78-B66B-31E52F93D2E9?confirm", &buf)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mds.On("DefineContractAPI", mock.Anything, mock.Anything, mock.AnythingOfType("*core.ContractAPI"), true).Return(nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
