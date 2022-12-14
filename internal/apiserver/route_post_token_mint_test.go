// SiGG-GoLang-On-the-Fly //

package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly/mocks/assetmocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostTokenMint(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	mam := &assetmocks.Manager{}
	o.On("Assets").Return(mam)
	input := core.TokenTransferInput{}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(&input)
	req := httptest.NewRequest("POST", "/api/v1/namespaces/ns1/tokens/mint", &buf)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mam.On("MintTokens", mock.Anything, mock.AnythingOfType("*core.TokenTransferInput"), false).
		Return(&core.TokenTransfer{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 202, res.Result().StatusCode)
}
