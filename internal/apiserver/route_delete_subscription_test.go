package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteSubscription(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	input := core.Subscription{}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(&input)
	u := fftypes.NewUUID()
	req := httptest.NewRequest("DELETE", fmt.Sprintf("/api/v1/namespaces/ns1/subscriptions/%s", u), &buf)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	o.On("DeleteSubscription", mock.Anything, u.String()).
		Return(nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 204, res.Result().StatusCode)
}
