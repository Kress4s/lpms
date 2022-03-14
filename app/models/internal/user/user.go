package user

import (
	"lpms/app/models/internal/common"
	"lpms/app/models/tables"
)

type User struct {
	common.Base `gorm:"embedded"`
	Username    string `gorm:"column:user_name;type:varchar(50);not null;comment:用户名"`
	Password    string `gorm:"column:password;type:varchar(60);not null;comment:密码"`
	ID          int64  `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	IsAdmin     bool   `gorm:"column:is_admin;type:boolean;not null;comment:是否是超管"`
	Status      bool   `gorm:"column:status;type:boolean;comment:状态"`
}

func (User) TableName() string {
	return tables.User
}
