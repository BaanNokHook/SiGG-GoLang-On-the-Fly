// SiGG-GoLang-On-the-Fly //

package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly/mocks/networkmapmocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewIdentity(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	mnm := &networkmapmocks.Manager{}
	o.On("NetworkMap").Return(mnm)
	input := core.Identity{}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(&input)
	req := httptest.NewRequest("POST", "/api/v1/namespaces/ns1/identities", &buf)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mnm.On("RegisterIdentity", mock.Anything, mock.AnythingOfType("*core.IdentityCreateDTO"), false).
		Return(&core.Identity{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 202, res.Result().StatusCode)
}
