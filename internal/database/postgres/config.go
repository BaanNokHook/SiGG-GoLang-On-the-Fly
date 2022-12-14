// SiGG-GoLang-On-the-Fly //
package postgres

import (
	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/hyperledger/firefly/internal/database/sqlcommon"
)

const (
	defaultConnectionLimitPostgreSQL = 50
)

func (psql *Postgres) InitConfig(config config.Section) {
	psql.SQLCommon.InitConfig(psql, config)
	config.SetDefault(sqlcommon.SQLConfMaxConnections, defaultConnectionLimitPostgreSQL)
}
