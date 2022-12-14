// SiGG-GoLang-On-the-Fly //

package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly/mocks/broadcastmocks"
	"github.com/hyperledger/firefly/mocks/datamocks"
	"github.com/hyperledger/firefly/mocks/multipartymocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostDataBlobPublish(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	mdm := &datamocks.Manager{}
	o.On("Data").Return(mdm)
	mbm := &broadcastmocks.Manager{}
	o.On("MultiParty").Return(&multipartymocks.Manager{})
	o.On("Broadcast").Return(mbm)
	input := core.Data{}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(&input)
	req := httptest.NewRequest("POST", "/api/v1/namespaces/ns1/data/id1/blob/publish", &buf)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mbm.On("PublishDataBlob", mock.Anything, "id1", core.IdempotencyKey("")).
		Return(&core.Data{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
