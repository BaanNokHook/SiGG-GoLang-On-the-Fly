// SiGG-GoLang-On-the-Fly //
package tbd

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/hyperledger/firefly/pkg/identity"
)

// TBD is a null implementation of the Identity Interface to avoid breaking configuration created with the previous "onchain" plugin
type TBD struct {
	capabilities *identity.Capabilities
}

func (tbd *TBD) Name() string {
	return "onchain" // For backwards compatibility with previous config that might have specified "onchain"
}

func (tbd *TBD) Init(ctx context.Context, config config.Section) (err error) {
	tbd.capabilities = &identity.Capabilities{}
	return nil
}

func (tbd *TBD) SetHandler(namespace string, handler identity.Callbacks) {
}

func (tbd *TBD) Start() error {
	return nil
}

func (tbd *TBD) Capabilities() *identity.Capabilities {
	return tbd.capabilities
}
