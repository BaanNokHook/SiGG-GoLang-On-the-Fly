package apiserver

import (
	"bytes"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/mocks/datamocks"
	"github.com/hyperledger/firefly/mocks/multipartymocks"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetDataBlob(t *testing.T) {
	o, r := newTestAPIServer()
	o.On("Authorize", mock.Anything, mock.Anything).Return(nil)
	mdm := &datamocks.Manager{}
	mdm.On("BlobsEnabled").Return(true)
	o.On("Data").Return(mdm)
	o.On("MultiParty").Return(&multipartymocks.Manager{})
	req := httptest.NewRequest("GET", "/api/v1/namespaces/mynamespace/data/abcd1234/blob", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	blobHash := fftypes.NewRandB32()
	mdm.On("DownloadBlob", mock.Anything, "abcd1234").
		Return(&core.Blob{
			Hash: blobHash,
			Size: 12345,
		}, ioutil.NopCloser(bytes.NewReader([]byte("hello"))), nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
	b, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, "hello", string(b))
	assert.Equal(t, "12345", res.Result().Header.Get(core.HTTPHeadersBlobSize))
	assert.Equal(t, blobHash.String(), res.Result().Header.Get(core.HTTPHeadersBlobHashSHA256))
}
