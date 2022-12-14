// SiGG-GoLang-On-the-Fly //
package core

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
)

// DeprecatedOrganization is the data structure we used to use prior to FIR-9.
// Now we use the common Identity structure throughout
type DeprecatedOrganization struct {
	ID          *fftypes.UUID      `json:"id"`
	Message     *fftypes.UUID      `json:"message,omitempty"`
	Parent      string             `json:"parent,omitempty"`
	Identity    string             `json:"identity,omitempty"`
	Name        string             `json:"name,omitempty"`
	Description string             `json:"description,omitempty"`
	Profile     fftypes.JSONObject `json:"profile,omitempty"`
	Created     *fftypes.FFTime    `json:"created,omitempty"`

	identityClaim *IdentityClaim
}

// Migrate creates and maintains a migrated IdentityClaim object, which
// is used when processing an old-style organization broadcast received when
// joining an existing network
func (org *DeprecatedOrganization) Migrated() *IdentityClaim {
	if org.identityClaim != nil {
		return org.identityClaim
	}
	org.identityClaim = &IdentityClaim{
		Identity: &Identity{
			IdentityBase: IdentityBase{
				ID:        org.ID,
				Type:      IdentityTypeOrg,
				Namespace: LegacySystemNamespace,
				Name:      org.Name,
				Parent:    nil, // No support for child identity migration (see FIR-9 for details)
			},
			IdentityProfile: IdentityProfile{
				Description: org.Description,
				Profile:     org.Profile,
			},
		},
	}
	org.identityClaim.Identity.DID, _ = org.identityClaim.Identity.GenerateDID(context.Background())
	return org.identityClaim
}

func (org *DeprecatedOrganization) Topic() string {
	return org.Migrated().Topic()
}

func (org *DeprecatedOrganization) SetBroadcastMessage(msgID *fftypes.UUID) {
	org.Migrated().SetBroadcastMessage(msgID)
}
