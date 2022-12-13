package migrate

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteList(t *testing.T) {
	value := map[interface{}]interface{}{
		"values": []interface{}{"test1", "test2"},
	}
	config := &ConfigItem{value: value, writer: os.Stdout}
	config.Get("values").Each().Delete()
	assert.Equal(t, 0, len(value))
}

func TestNoRename(t *testing.T) {
	value := map[interface{}]interface{}{
		"key1": "val1",
		"key2": "val2",
	}
	config := &ConfigItem{value: value, writer: os.Stdout}
	config.Get("key1").RenameTo("key2")
	assert.Equal(t, map[interface{}]interface{}{
		"key2": "val2",
	}, value)
}
