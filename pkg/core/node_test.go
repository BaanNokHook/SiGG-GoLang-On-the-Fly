// SiGG-GoLang-On-the-Fly //
package core

import (
	"testing"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/stretchr/testify/assert"
)

func TestNodeigration(t *testing.T) {

	node := DeprecatedNode{
		ID:          fftypes.NewUUID(),
		Name:        "node1",
		Description: "Node 1",
		DX: DeprecatedDXInfo{
			Peer: "ignored",
			Endpoint: fftypes.JSONObject{
				"id": "peer1",
			},
		},
	}
	parentID := fftypes.NewUUID()
	assert.Equal(t, &IdentityClaim{
		Identity: &Identity{
			IdentityBase: IdentityBase{
				ID:        node.ID,
				Type:      IdentityTypeNode,
				DID:       "did:firefly:node/node1",
				Namespace: LegacySystemNamespace,
				Name:      "node1",
				Parent:    parentID,
			},
			IdentityProfile: IdentityProfile{
				Description: "Node 1",
				Profile: fftypes.JSONObject{
					"id": "peer1",
				},
			},
		},
	}, node.AddMigratedParent(parentID))

	assert.Equal(t, "14c4157d50d35470b15a6576affa62adea1b191e8238f2273a099d1ef73fb335", node.Topic())

	msg := &Message{
		Header: MessageHeader{
			ID: fftypes.NewUUID(),
		},
	}
	node.SetBroadcastMessage(msg.Header.ID)
	assert.Equal(t, msg.Header.ID, node.Migrated().Identity.Messages.Claim)

}
