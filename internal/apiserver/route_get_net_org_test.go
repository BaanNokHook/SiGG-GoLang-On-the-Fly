package apiserver

import (
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly/mocks/networkmapmocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetOrg(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	nmn := &networkmapmocks.Manager{}
	o.On("NetworkMap").Return(nmn)
	req := httptest.NewRequest("GET", "/api/v1/network/organizations/org12345", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	nmn.On("GetOrganizationByNameOrID", mock.Anything, "org12345").
		Return(&core.Identity{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
