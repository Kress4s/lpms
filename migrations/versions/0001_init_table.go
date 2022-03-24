package versions

import (
	"lpms/app/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// V0001InitTables init tables
var V0001InitTables = &gormigrate.Migration{
	ID: "0001_init_tables",
	Migrate: func(tx *gorm.DB) error {
		if err := tx.AutoMigrate(
			// 用户
			models.User{},
			// 储备库
			models.ReservePro{},
			// 文件库
			models.Object{},
			// 实施库-政府项目
			models.ImplementGov{},
			// 实施库-产业项目
			models.ImpleIndustry{},
			// 窗口设置
			models.WindowSetting{},
		); err != nil {
			return err
		}
		return nil
	},
}
