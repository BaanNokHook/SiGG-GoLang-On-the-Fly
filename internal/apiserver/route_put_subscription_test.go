// SiGG-GoLang-On-the-Fly //

package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPutSubscription(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	input := core.Subscription{}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(&input)
	req := httptest.NewRequest("PUT", "/api/v1/namespaces/ns1/subscriptions", &buf)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	o.On("CreateUpdateSubscription", mock.Anything, mock.AnythingOfType("*core.Subscription")).
		Return(&core.Subscription{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
