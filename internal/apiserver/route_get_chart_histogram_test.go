
package apiserver

import (
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetChartHistogramBadStartTime(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/mynamespace/charts/histogram/test?startTime=abc&endTime=456&buckets=30", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, 400, res.Result().StatusCode)
}

func TestGetChartHistogramBadEndTime(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/mynamespace/charts/histogram/test?startTime=123&endTime=abc&buckets=30", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, 400, res.Result().StatusCode)
}

func TestGetChartHistogramBadBuckets(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/mynamespace/charts/histogram/test?startTime=123&endTime=456&buckets=abc", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, 400, res.Result().StatusCode)
}

func TestGetChartHistogramSuccess(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/mynamespace/charts/histogram/test?startTime=1234567890&endTime=1234567891&buckets=30", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	startTime, _ := fftypes.ParseTimeString("1234567890")
	endtime, _ := fftypes.ParseTimeString("1234567891")

	o.On("GetChartHistogram", mock.Anything, startTime.UnixNano(), endtime.UnixNano(), int64(30), database.CollectionName("test")).
		Return([]*core.ChartHistogram{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
