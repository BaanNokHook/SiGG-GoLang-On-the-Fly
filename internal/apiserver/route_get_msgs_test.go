package apiserver

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetMessages(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/mynamespace/messages", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	o.On("GetMessages", mock.Anything, mock.Anything).
		Return([]*core.Message{}, nil, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}

func TestGetMessagesWithCount(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/mynamespace/messages?count", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	var ten int64 = 10
	o.On("GetMessages", mock.Anything, mock.Anything).
		Return([]*core.Message{}, &ffapi.FilterResult{
			TotalCount: &ten,
		}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
	var resWithCount ffapi.FilterResultsWithCount
	err := json.NewDecoder(res.Body).Decode(&resWithCount)
	assert.NoError(t, err)
	assert.NotNil(t, resWithCount.Items)
	assert.Equal(t, int64(0), resWithCount.Count)
	assert.Equal(t, int64(10), resWithCount.Total)
}

func TestGetMessagesWithCountAndData(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/mynamespace/messages?count&fetchdata", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	var ten int64 = 10
	o.On("GetMessagesWithData", mock.Anything, mock.Anything).
		Return([]*core.MessageInOut{}, &ffapi.FilterResult{
			TotalCount: &ten,
		}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
	var resWithCount ffapi.FilterResultsWithCount
	err := json.NewDecoder(res.Body).Decode(&resWithCount)
	assert.NoError(t, err)
	assert.NotNil(t, resWithCount.Items)
	assert.Equal(t, int64(0), resWithCount.Count)
	assert.Equal(t, int64(10), resWithCount.Total)
}
