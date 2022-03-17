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
		// 创建 操作人员表，角色表, 操作人员角色关联表，用户登录记录表
		if err := tx.AutoMigrate(
			models.User{},
			models.ReservePro{},
			models.Object{},
			models.ImplementGov{},
		); err != nil {
			return err
		}
		return nil
	},
}
