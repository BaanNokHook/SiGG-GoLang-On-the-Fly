// SiGG-GoLang-On-the-Fly //
//go:build !cgo
// +build !cgo

package difactory

import (
	"github.com/hyperledger/firefly/internal/database/postgres"
	"github.com/hyperledger/firefly/pkg/database"
)

var pluginsByName = map[string]func() database.Plugin{
	(*postgres.Postgres)(nil).Name(): func() database.Plugin { return &postgres.Postgres{} },
}
