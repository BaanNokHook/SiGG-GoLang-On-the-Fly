// SiGG-GoLang-On-the-Fly //
package data

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/pkg/core"
)

type Validator interface {
	Validate(ctx context.Context, data *core.Data) error
	ValidateValue(ctx context.Context, value *fftypes.JSONAny, expectedHash *fftypes.Bytes32) error
	Size() int64 // for cache management
}
