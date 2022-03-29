package implement

import (
	"lpms/app/models/internal/common"
	"lpms/app/models/tables"
	"time"

	"github.com/goccy/go-json"

	"gorm.io/gorm"
)

type GovProgress struct {
	common.Base            `gorm:"embedded"`
	ID                     int64           `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	ProjectID              int64           `gorm:"column:project_id;type:bigint;not null;comment:项目ID"`
	Year                   int             `gorm:"column:year;type:integer;not null;comment:年份"`
	Month                  int             `gorm:"column:month;type:integer;not null;comment:月份"`
	PlanInvest             *float64        `gorm:"column:plan_invest;type:numeric;not null;comment:本月计划投资额(万)"`
	PlanProgress           string          `gorm:"column:plan_progress;type:text;not null;comment:本月计划形象进度"`
	PlanInvested           *float64        `gorm:"column:plan_invested;type:numeric;comment:本月完成投资额(万)"`
	YearSumInvested        *float64        `gorm:"column:year_sum_invested;type:numeric;comment:当年累计投资额(万)"`
	LastMonthFixedInvested *float64        `gorm:"column:last_month_fixed_invested;type:numeric;comment:上月完成固投(万)"`
	YearSumFixedInvested   *float64        `gorm:"column:year_sum__fixed_invested;type:numeric;comment:当年累计固投(万)"`
	ActualProgress         string          `gorm:"column:actual_progress;type:text;comment:本月完成形象进度"`
	ProblemDetail          json.RawMessage `gorm:"column:problem_detail;type:jsonb;comment:需协调问题详情"`
	ChangeContent          json.RawMessage `gorm:"column:change_content;type:jsonb;comment:本月产生联系单变更"`
	Contracts              json.RawMessage `gorm:"column:contracts;type:jsonb;comment:本月新增合同信息"`
	Status                 int             `gorm:"column:status;type:integer;default(0);comment:填报状态 0:未提交, 1:已提交"`
	Comment                string          `gorm:"column:comment;type:text;comment:备注"`
}

func (GovProgress) TableName() string {
	return tables.GovProgress
}

func (b *GovProgress) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	b.CreateAt = now
	b.UpdateAt = now
	return nil
}

func (b *GovProgress) BeforeUpdate(tx *gorm.DB) error {
	b.UpdateAt = time.Now()
	return nil
}

type ListGovProgressPlan struct {
	ID           int64    `gorm:"column:id"`
	Year         int      `gorm:"column:year"`
	Month        int      `gorm:"column:month"`
	PlanInvest   *float64 `gorm:"column:plan_invest"`
	PlanProgress string   `gorm:"column:plan_progress"`
}

type GovProgressCompare struct {
	// 年份
	Year int `gorm:"column:year"`
	// 月份
	Month int `gorm:"column:month"`
	// 本月计划投资额
	PlanInvest *float64 `gorm:"column:plan_invest"`
	// 本月计划形象进度
	PlanProgress string `gorm:"column:plan_progress"`
	// 本月完成投资额
	PlanInvested *float64 `gorm:"column:plan_invested"`
	// 本月完成形象进度
	ActualProgress string `gorm:"column:actual_progress"`
}
