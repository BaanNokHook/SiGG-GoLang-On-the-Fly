// SiGG-GoLang-On-the-Fly //

package sqlcommon

import (
	"context"
	"database/sql"
	"io/ioutil"
	"os"
	"testing"

	sq "github.com/Masterminds/squirrel"
	migratedb "github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/hyperledger/firefly-common/pkg/dbsql"
	"github.com/hyperledger/firefly/mocks/databasemocks"
	"github.com/hyperledger/firefly/pkg/database"
	"github.com/stretchr/testify/assert"

	// Import SQLite driver
	_ "github.com/mattn/go-sqlite3"
)

// sqliteGoTestProvider uses QL in-memory database
type sqliteGoTestProvider struct {
	SQLCommon

	config       config.Section
	t            *testing.T
	callbacks    *databasemocks.Callbacks
	capabilities *database.Capabilities
}

// newTestProvider creates a real in-memory database provider for e2e testing
func newSQLiteTestProvider(t *testing.T) (*sqliteGoTestProvider, func()) {
	conf := config.RootSection("unittest.db")
	conf.AddKnownKey("url", "test")
	tp := &sqliteGoTestProvider{
		t:            t,
		callbacks:    &databasemocks.Callbacks{},
		capabilities: &database.Capabilities{},
		config:       conf,
	}
	tp.SQLCommon.InitConfig(tp, tp.config)
	dir, err := ioutil.TempDir("", "")
	assert.NoError(t, err)
	tp.config.Set(SQLConfDatasourceURL, "file::memory:")
	tp.config.Set(SQLConfMigrationsAuto, true)
	tp.config.Set(SQLConfMigrationsDirectory, "../../../db/migrations/sqlite")
	tp.config.Set(SQLConfMaxConnections, 1)

	err = tp.Init(context.Background(), tp, tp.config, tp.capabilities)
	assert.NoError(tp.t, err)
	tp.SetHandler(database.GlobalHandler, tp.callbacks)

	return tp, func() {
		tp.Close()
		_ = os.RemoveAll(dir)
	}
}

func (tp *sqliteGoTestProvider) Name() string {
	return "sqlite3"
}

func (tp *sqliteGoTestProvider) MigrationsDir() string {
	return "sqlite"
}

func (mp *sqliteGoTestProvider) SequenceColumn() string {
	return "seq"
}

func (psql *sqliteGoTestProvider) Features() dbsql.SQLFeatures {
	features := dbsql.DefaultSQLProviderFeatures()
	features.PlaceholderFormat = sq.Dollar
	features.UseILIKE = false // Not supported
	return features
}

func (tp *sqliteGoTestProvider) ApplyInsertQueryCustomizations(insert sq.InsertBuilder, requestConflictEmptyResult bool) (sq.InsertBuilder, bool) {
	// Nothing required - QL supports the query for returning the generated ID, and we use that for the sequence
	return insert, false
}

func (tp *sqliteGoTestProvider) Open(url string) (*sql.DB, error) {
	return sql.Open("sqlite3", url)
}

func (tp *sqliteGoTestProvider) GetMigrationDriver(db *sql.DB) (migratedb.Driver, error) {
	return sqlite3.WithInstance(db, &sqlite3.Config{})
}
