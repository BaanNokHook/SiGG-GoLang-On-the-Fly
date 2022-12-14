// SiGG-GoLang-On-the-Fly //
//go:build cgo
// +build cgo

package sqlite3

import (
	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/hyperledger/firefly/internal/database/sqlcommon"
)

const (
	defaultConnectionLimitSQLite = 1
)

func (sqlite *SQLite3) InitConfig(config config.Section) {
	sqlite.SQLCommon.InitConfig(sqlite, config)
	config.SetDefault(sqlcommon.SQLConfMaxConnections, defaultConnectionLimitSQLite)
}
