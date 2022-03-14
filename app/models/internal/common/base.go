package common

import (
	"time"

	"gorm.io/gorm"
)

// Base 基础模型定义
type Base struct {
	CreateAt time.Time `gorm:"column:create_at;type:timestamp;not null;comment:创建时间" json:"create_at"`
	UpdateAt time.Time `gorm:"column:update_at;type:timestamp;not null;comment:最后一次更新时间" json:"update_at"`
	CreateBy string    `gorm:"column:create_by;type:varchar(40);not null;comment:创建者ID" json:"create_by"`
	UpdateBy string    `gorm:"column:update_by;type:varchar(40);not null;comment:最后一次更新者ID" json:"update_by"`
}

type LandUse struct {
	ID               int64  `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	Total            int64  `gorm:"column:total;type:bigint;comment:总用亩"`
	Add              int64  `gorm:"column:add;type:bigint;comment:新增建设用地"`
	NoConformUsePlan int64  `gorm:"column:no_conform_use_plan;type:bigint;comment:不符合土地利用规划面积"`
	SiteRed          int    `gorm:"column:site_red;type:integer;comment:选址红线 0:有拆迁,1:无拆迁"`
	NeedCollect      int64  `gorm:"column:need_collect;type:bigint;comment:需征地面积"`
	NeedPeopleMove   int64  `gorm:"column:need_people_move;type:bigint;comment:需拆迁农户/居民数(人)"`
	CompanyBusiness  int64  `gorm:"column:company_business;type:bigint;comment:企/事业单位(家)"`
	UploadCadID      string `gorm:"column:upload_cad_id;type:varchar(40);comment:CAD文件ID"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	b.CreateAt = now
	b.UpdateAt = now
	return nil
}

func (b *Base) BeforeUpdate(tx *gorm.DB) error {
	b.UpdateAt = time.Now()
	return nil
}
