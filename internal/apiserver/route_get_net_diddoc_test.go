// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly/internal/networkmap"
	"github.com/hyperledger/firefly/mocks/networkmapmocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetDIDDocByDID(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	nmn := &networkmapmocks.Manager{}
	o.On("NetworkMap").Return(nmn)
	req := httptest.NewRequest("GET", "/api/v1/network/diddocs/did:firefly:org/org_1", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	nmn.On("GetDIDDocForIndentityByDID", mock.Anything, "did:firefly:org/org_1").
		Return(&networkmap.DIDDocument{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
