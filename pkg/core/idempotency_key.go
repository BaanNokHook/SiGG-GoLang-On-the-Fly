// SiGG-GoLang-On-the-Fly //
package core

import (
	"context"
	"database/sql/driver"

	"github.com/hyperledger/firefly-common/pkg/i18n"
)

// IdempotencyKey is accessed in Go as a string, but when persisted to storage it will be stored as a null
// to allow multiple entries in a unique index to exist with the same un-set idempotency key.
type IdempotencyKey string

func (ik IdempotencyKey) Value() (driver.Value, error) {
	if ik == "" {
		return nil, nil
	}
	return (string)(ik), nil
}

func (ik *IdempotencyKey) Scan(src interface{}) error {
	switch src := src.(type) {
	case nil:
		return nil
	case []byte:
		*ik = IdempotencyKey(src)
		return nil
	case string:
		*ik = IdempotencyKey(src)
		return nil
	default:
		return i18n.NewError(context.Background(), i18n.MsgTypeRestoreFailed, src, ik)
	}
}
