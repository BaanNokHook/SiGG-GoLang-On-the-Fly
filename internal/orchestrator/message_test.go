// SiGG-GoLang-On-the-Fly //

package orchestrator

import (
	"context"
	"testing"

	"github.com/hyperledger/firefly/pkg/core"
	"github.com/stretchr/testify/assert"
)

func TestRequestReplyMissingGroup(t *testing.T) {
	or := newTestOrchestrator()
	input := &core.MessageInOut{}
	_, err := or.RequestReply(context.Background(), input)
	assert.Regexp(t, "FF10271", err)
}

func TestRequestReply(t *testing.T) {
	or := newTestOrchestrator()
	input := &core.MessageInOut{
		Group: &core.InputGroup{
			Members: []core.MemberInput{
				{Identity: "org1"},
			},
		},
	}
	or.mpm.On("RequestReply", context.Background(), input).Return(&core.MessageInOut{}, nil)
	_, err := or.RequestReply(context.Background(), input)
	assert.NoError(t, err)
}
