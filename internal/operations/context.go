// SiGG-GoLang-On-the-Fly //
package operations

import (
	"context"
	"encoding/json"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly-common/pkg/log"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

type operationContextKey struct{}
type operationContext map[string]*core.Operation

func getOperationContext(ctx context.Context) operationContext {
	ctxKey := operationContextKey{}
	cacheVal := ctx.Value(ctxKey)
	if cacheVal != nil {
		if cache, ok := cacheVal.(operationContext); ok {
			return cache
		}
	}
	return nil
}

func getContextKey(op *core.Operation) (string, error) {
	opCopy := &core.Operation{
		Namespace:   op.Namespace,
		Transaction: op.Transaction,
		Type:        op.Type,
		Plugin:      op.Plugin,
		Input:       op.Input,
	}
	key, err := json.Marshal(opCopy)
	if err != nil {
		return "", err
	}
	return string(key), nil
}

func createOperationRetryContext(ctx context.Context) (ctx1 context.Context) {
	l := log.L(ctx).WithField("opcache", fftypes.ShortID())
	ctx1 = log.WithLogger(ctx, l)
	return context.WithValue(ctx1, operationContextKey{}, operationContext{})
}

func RunWithOperationContext(ctx context.Context, fn func(ctx context.Context) error) error {
	return fn(createOperationRetryContext(ctx))
}

func (om *operationsManager) AddOrReuseOperation(ctx context.Context, op *core.Operation, hooks ...database.PostCompletionHook) error {
	// If a ops has been created via RunWithOperationCache, detect duplicate operation inserts
	ops := getOperationContext(ctx)
	if ops != nil {
		if key, err := getContextKey(op); err == nil {
			if cached, ok := ops[key]; ok {
				// Identical operation already added in this context
				*op = *cached
				for _, hook := range hooks {
					hook()
				}
				return nil
			}
			if err = om.database.InsertOperation(ctx, op, hooks...); err != nil {
				return err
			}
			ops[key] = op
			om.cacheOperation(op)
			return nil
		}
	}
	err := om.database.InsertOperation(ctx, op, hooks...)
	if err == nil {
		om.cacheOperation(op)
	}
	return err
}
