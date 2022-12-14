// SiGG-GoLang-On-the-Fly //

package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly/mocks/broadcastmocks"
	"github.com/hyperledger/firefly/mocks/multipartymocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostNewMessageBroadcast(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	mmp := &multipartymocks.Manager{}
	o.On("MultiParty").Return(mmp)
	mbm := &broadcastmocks.Manager{}
	o.On("Broadcast").Return(mbm)
	input := core.MessageInOut{}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(&input)
	req := httptest.NewRequest("POST", "/api/v1/namespaces/ns1/messages/broadcast", &buf)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mbm.On("BroadcastMessage", mock.Anything, mock.AnythingOfType("*core.MessageInOut"), false).
		Return(&core.Message{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 202, res.Result().StatusCode)
}

func TestPostNewMessageBroadcastSync(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	mmp := &multipartymocks.Manager{}
	o.On("MultiParty").Return(mmp)
	mbm := &broadcastmocks.Manager{}
	o.On("Broadcast").Return(mbm)
	input := core.MessageInOut{}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(&input)
	req := httptest.NewRequest("POST", "/api/v1/namespaces/ns1/messages/broadcast?confirm", &buf)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mbm.On("BroadcastMessage", mock.Anything, mock.AnythingOfType("*core.MessageInOut"), true).
		Return(&core.Message{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
