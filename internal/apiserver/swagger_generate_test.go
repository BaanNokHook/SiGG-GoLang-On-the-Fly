// SiGG-GoLang-On-the-Fly //
//go:build reference
// +build reference

package apiserver

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/hyperledger/firefly/internal/coreconfig"
	"github.com/stretchr/testify/assert"
)

func TestDownloadSwaggerYAML(t *testing.T) {
	coreconfig.Reset()
	config.Set(coreconfig.APIOASPanicOnMissingDescription, true)
	as := &apiServer{}
	hf := as.handlerFactory()
	handler := hf.APIWrapper(as.swaggerHandler(as.swaggerGenerator(routes, "http://localhost:5000")))
	s := httptest.NewServer(http.HandlerFunc(handler))
	defer s.Close()

	res, err := http.Get(fmt.Sprintf("http://%s/api/swagger.yaml", s.Listener.Addr()))
	assert.Nil(t, err)
	b, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, 200, res.StatusCode, string(b))
	doc, err := openapi3.NewLoader().LoadFromData(b)
	assert.NoError(t, err)
	err = doc.Validate(context.Background())
	assert.NoError(t, err)
	err = os.WriteFile(filepath.Join("..", "..", "docs", "swagger", "swagger.yaml"), b, 0644)
	assert.NoError(t, err)
}
