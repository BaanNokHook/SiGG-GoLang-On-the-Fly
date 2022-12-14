// SiGG-GoLang-On-the-Fly //
package sqlcommon

import (
	"fmt"

	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/hyperledger/firefly-common/pkg/dbsql"
)

const (
	// SQLConfMigrationsAuto enables automatic migrations
	SQLConfMigrationsAuto = "migrations.auto"
	// SQLConfMigrationsDirectory is the directory containing the numerically ordered migration DDL files to apply to the database
	SQLConfMigrationsDirectory = "migrations.directory"
	// SQLConfDatasourceURL is the datasource connection URL string
	SQLConfDatasourceURL = "url"
	// SQLConfMaxConnections maximum connections to the database
	SQLConfMaxConnections = "maxConns"
	// SQLConfMaxConnIdleTime maximum connections to the database
	SQLConfMaxConnIdleTime = "maxConnIdleTime"
	// SQLConfMaxIdleConns maximum connections to the database
	SQLConfMaxIdleConns = "maxIdleConns"
	// SQLConfMaxConnLifetime maximum connections to the database
	SQLConfMaxConnLifetime = "maxConnLifetime"
)

const (
	defaultMigrationsDirectoryTemplate = "./db/migrations/%s"
)

func (s *SQLCommon) InitConfig(provider dbsql.Provider, config config.Section) {
	config.AddKnownKey(SQLConfMigrationsAuto, false)
	config.AddKnownKey(SQLConfDatasourceURL)
	config.AddKnownKey(SQLConfMigrationsDirectory, fmt.Sprintf(defaultMigrationsDirectoryTemplate, provider.MigrationsDir()))
	config.AddKnownKey(SQLConfMaxConnections) // some providers set a default
	config.AddKnownKey(SQLConfMaxConnIdleTime, "1m")
	config.AddKnownKey(SQLConfMaxIdleConns) // defaults to the max connections
	config.AddKnownKey(SQLConfMaxConnLifetime)
}
