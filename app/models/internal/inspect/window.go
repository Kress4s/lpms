package inspect

import (
	"lpms/app/models/internal/common"
	"lpms/app/models/tables"
	"time"

	"gorm.io/gorm"

	"github.com/goccy/go-json"
)

type WindowSetting struct {
	common.Base     `gorm:"embedded"`
	ID              int64           `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	ReserveSetting  json.RawMessage `gorm:"column:reserve_setting;type:jsonb;comment:储备库填报"`
	ProgressSetting json.RawMessage `gorm:"column:progress_setting;type:jsonb;comment:项目进度填报"`
	ProPlanSetting  json.RawMessage `gorm:"column:pro_plan_setting;type:jsonb;comment:项目计划填报"`
}

func (WindowSetting) TableName() string {
	return tables.Window
}

func (b *WindowSetting) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	b.CreateAt = now
	b.UpdateAt = now
	return nil
}

func (b *WindowSetting) BeforeUpdate(tx *gorm.DB) error {
	b.UpdateAt = time.Now()
	return nil
}
