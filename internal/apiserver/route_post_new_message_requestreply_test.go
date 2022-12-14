// SiGG-GoLang-On-the-Fly //

package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly/mocks/multipartymocks"
	"github.com/hyperledger/firefly/mocks/privatemessagingmocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostNewMessageRequestReply(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	o.On("PrivateMessaging").Return(&privatemessagingmocks.Manager{})
	o.On("RequestReply", mock.Anything, mock.Anything).Return(&core.MessageInOut{}, nil)
	mmp := &multipartymocks.Manager{}
	o.On("MultiParty").Return(mmp)
	input := &core.MessageInOut{}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(&input)
	req := httptest.NewRequest("POST", "/api/v1/namespaces/ns1/messages/requestreply", &buf)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
