package migrations

import (
	"lpms/commom/drivers/database"
	"lpms/migrations/versions"

	"github.com/go-gormigrate/gormigrate/v2"
)

var migrations = []*gormigrate.Migration{
	// init table
	versions.V0001InitTables,
	// init data
	versions.V0002InitData,
}

func Migrate() error {
	return gormigrate.New(database.GetDriver(), &gormigrate.Options{
		TableName:                 "lpms_migrations",
		IDColumnName:              "id",
		IDColumnSize:              255,
		UseTransaction:            true,
		ValidateUnknownMigrations: true,
	}, migrations).Migrate()
}
