package apiserver

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFavIcon16(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/favicon-16x16.png", nil)
	var handler http.HandlerFunc = favIcons
	handler(res, req)
	assert.Equal(t, 200, res.Result().StatusCode)
	b, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, b)
	assert.Equal(t, b, ffLogo16)
}

func TestFavIcon32(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/favicon-32x32.png", nil)
	var handler http.HandlerFunc = favIcons
	handler(res, req)
	assert.Equal(t, 200, res.Result().StatusCode)
	b, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, b)
	assert.Equal(t, b, ffLogo32)
}
