// SiGG-GoLang-On-the-Fly //

package core

import (
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/stretchr/testify/assert"
)

func TestNewEvent(t *testing.T) {

	ref := fftypes.NewUUID()
	tx := fftypes.NewUUID()
	e := NewEvent(EventTypeMessageConfirmed, "ns1", ref, tx, "topic1")
	assert.Equal(t, EventTypeMessageConfirmed, e.Type)
	assert.Equal(t, "ns1", e.Namespace)
	assert.Equal(t, *ref, *e.Reference)
	assert.Equal(t, *tx, *e.Transaction)
	assert.Equal(t, "topic1", e.Topic)

	e.Sequence = 12345
	var ls LocallySequenced = e
	assert.Equal(t, int64(12345), ls.LocalSequence())

}
