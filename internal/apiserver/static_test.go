// SiGG-GoLang-On-the-Fly //

package apiserver

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStaticHosting(t *testing.T) {
	sc := newStaticHandler(configDir, "index.html", "/config")
	var handler http.HandlerFunc = sc.ServeHTTP
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/config/firefly.core.yaml", nil)
	handler(res, req)
	assert.Equal(t, 200, res.Result().StatusCode)
	b, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, b)

}

func Test404UnRelablePath(t *testing.T) {
	sc := newStaticHandler("test", "firefly.core.yaml", "test")
	var handler http.HandlerFunc = sc.ServeHTTP
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test", nil)
	handler(res, req)
	assert.Equal(t, 404, res.Result().StatusCode)
}

func TestServeDefault(t *testing.T) {
	sc := newStaticHandler(configDir, "firefly.core.yaml", "/config")
	var handler http.HandlerFunc = sc.ServeHTTP
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/config", nil)
	handler(res, req)
	assert.Equal(t, 200, res.Result().StatusCode)
	b, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, b)
}

func TestServeNotFoundServeDefault(t *testing.T) {
	sc := newStaticHandler(configDir, "firefly.core.yaml", "/config")
	var handler http.HandlerFunc = sc.ServeHTTP
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/config/wrong", nil)
	handler(res, req)
	assert.Equal(t, 200, res.Result().StatusCode)
}

type fakeOS struct{}

func (f *fakeOS) Stat(name string) (os.FileInfo, error) {
	return nil, fmt.Errorf("pop")
}

func TestStatError(t *testing.T) {

	sc := &staticHandler{
		staticPath: "../../test",
		urlPrefix:  "/test",
		indexPath:  "firefly.core.yaml",
		os:         &fakeOS{},
	}
	var handler http.HandlerFunc = sc.ServeHTTP
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test/config", nil)
	handler(res, req)
	assert.Equal(t, 500, res.Result().StatusCode)
	b, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Regexp(t, "FF10185", string(b))
}
