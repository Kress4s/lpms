package versions

import (
	"lpms/app/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// V0003InitProgressTables init tables
var V0003InitProgressTables = &gormigrate.Migration{
	ID: "0003_init_progress",
	Migrate: func(tx *gorm.DB) error {
		if err := tx.AutoMigrate(
			// 实施库-政府项目-进度
			models.GovProgress{},
		); err != nil {
			return err
		}
		return nil
	},
}
