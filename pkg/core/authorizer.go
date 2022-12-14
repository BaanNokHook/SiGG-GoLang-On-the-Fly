// SiGG-GoLang-On-the-Fly //
package core

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
)

type Authorizer interface {
	Authorize(ctx context.Context, authReq *fftypes.AuthReq) error
}
