package versions

import (
	"lpms/app/models"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func InitUser() *models.User {
	return &models.User{Username: "admin", Password: "MTIzNDU2", IsAdmin: true, Status: true, Base: models.Base{
		CreateBy: "admin",
		CreateAt: time.Now(),
		UpdateBy: "admin",
		UpdateAt: time.Now(),
	}}
}

// V0002InitData init data
var V0002InitData = &gormigrate.Migration{
	ID: "0002_init_data",
	Migrate: func(tx *gorm.DB) error {
		if err := tx.Create(InitUser()).Error; err != nil {
			return err
		}
		// tpx := tx.Exec(fmt.Sprintf("select setval('%s_id_seq', (select (max(id)) from %s));",
		// 	tables.Permission, tables.Permission))
		// if tpx.Error != nil {
		// 	return tpx.Error
		// }
		return nil
	},
}
