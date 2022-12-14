// SiGG-GoLang-On-the-Fly //

package apiserver

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/mocks/assetmocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetTokenTransfers(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	mam := &assetmocks.Manager{}
	o.On("Assets").Return(mam)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/ns1/tokens/transfers", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mam.On("GetTokenTransfers", mock.Anything, mock.Anything).
		Return([]*core.TokenTransfer{}, nil, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}

func TestGetTokenTransfersFromOrTo(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	mam := &assetmocks.Manager{}
	o.On("Assets").Return(mam)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/ns1/tokens/transfers?fromOrTo=0x1", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mam.On("GetTokenTransfers", mock.Anything, mock.MatchedBy(func(filter ffapi.AndFilter) bool {
		f, _ := filter.Finalize()
		filterStr := f.String()
		return strings.Contains(filterStr, "( ( from == '0x1' ) || ( to == '0x1' ) )")
	})).Return([]*core.TokenTransfer{}, nil, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
