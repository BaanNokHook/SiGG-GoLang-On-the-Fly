// SiGG-GoLang-On-the-Fly //
package sharedstorage

import (
	"context"
	"io"

	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/hyperledger/firefly/pkg/core"
)

// Plugin is the interface implemented by each Shared Storage plugin
type Plugin interface {
	core.Named

	// InitConfig initializes the set of configuration options that are valid, with defaults. Called on all plugins.
	InitConfig(config config.Section)

	// Init initializes the plugin, with configuration
	Init(ctx context.Context, config config.Section) error

	// SetHandler registers a handler to receive callbacks
	// Plugin will attempt (but is not guaranteed) to deliver events only for the given namespace
	SetHandler(namespace string, handler Callbacks)

	// Capabilities returns capabilities - not called until after Init
	Capabilities() *Capabilities

	// UploadData publishes data to the Shared Storage, and returns a payload reference ID
	UploadData(ctx context.Context, data io.Reader) (payloadRef string, err error)

	// DownloadData reads data back from IPFS using the payload reference format returned from UploadData
	DownloadData(ctx context.Context, payloadRef string) (data io.ReadCloser, err error)
}

type Callbacks interface {
}

type Capabilities struct {
}
