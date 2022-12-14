// SiGG-GoLang-On-the-Fly //
package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdempotencyKeyNilValuer(t *testing.T) {
	v, err := ((IdempotencyKey)("")).Value()
	assert.NoError(t, err)
	assert.Nil(t, v)

	v, err = ((IdempotencyKey)("testValue")).Value()
	assert.NoError(t, err)
	assert.Equal(t, "testValue", v.(string))
}

func TestIdempotencyKeyNilScanner(t *testing.T) {
	var ik *IdempotencyKey
	err := ik.Scan(nil)
	assert.NoError(t, err)
	assert.Nil(t, ik)

	ik = new(IdempotencyKey)
	err = ik.Scan("testValue")
	assert.NoError(t, err)
	assert.Equal(t, "testValue", (string)(*ik))

	ik = new(IdempotencyKey)
	err = ik.Scan([]byte("testValue2"))
	assert.NoError(t, err)
	assert.Equal(t, "testValue2", (string)(*ik))

	ik = new(IdempotencyKey)
	err = ik.Scan(12345)
	assert.Regexp(t, "FF00105", err)
}
