package apiserver

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/mocks/datamocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetDataValue(t *testing.T) {
	o, r := newTestAPIServer()
	mdm := &datamocks.Manager{}
	o.On("Data").Return(mdm)
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/mynamespace/data/abcd12345/value", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	o.On("GetDataByID", mock.Anything, "abcd12345").
		Return(&core.Data{
			Value: fftypes.JSONAnyPtr(`{"some":"data"}`),
		}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)

	resData, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"some":"data"}`, string(resData))
}

func TestGetDataValueFail(t *testing.T) {
	o, r := newTestAPIServer()
	mdm := &datamocks.Manager{}
	o.On("Data").Return(mdm)
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/mynamespace/data/abcd12345/value", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	o.On("GetDataByID", mock.Anything, "abcd12345").
		Return(nil, fmt.Errorf("pop"))
	r.ServeHTTP(res, req)

	assert.Equal(t, 500, res.Result().StatusCode)

}
