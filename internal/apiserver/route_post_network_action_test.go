// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly/mocks/multipartymocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostNetworkAction(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	o.On("MultiParty").Return(&multipartymocks.Manager{})
	input := core.NetworkAction{}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(&input)
	req := httptest.NewRequest("POST", "/api/v1/network/action", &buf)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	o.On("SubmitNetworkAction", mock.Anything, mock.AnythingOfType("*core.NetworkAction")).Return(nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 202, res.Result().StatusCode)
}
