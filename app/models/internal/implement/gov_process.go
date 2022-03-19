package implement

import (
	"lpms/app/models/internal/common"
	"lpms/app/models/tables"
	"time"

	"gorm.io/gorm"
)

type GovProgress struct {
	common.Base            `gorm:"embedded"`
	ID                     int64    `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	ProjectID              int64    `gorm:"column:project_id;type:bigint;not null;uniqueIndex:P_M;comment:项目ID"`
	Month                  int      `gorm:"column:month;type:integer;not null;uniqueIndex:P_M;comment:月份"`
	PlanInvest             *float64 `gorm:"column:plan_invest;type:numeric;not null;comment:本月计划投资额(万)"`
	PlanProgress           string   `gorm:"column:plan_progress;type:text;not null;comment:本月计划形象进度"`
	PlanInvested           *float64 `gorm:"column:plan_invested;type:numeric;comment:本月完成投资额(万)"`
	YearSumInvested        *float64 `gorm:"column:year_sum_invested;type:numeric;comment:当年累计投资额(万)"`
	PayProject             *float64 `gorm:"column:pay_project;type:numeric;comment:本月支付工程款(万)"`
	LastMonthFixedInvested *float64 `gorm:"column:last_month_fixed_invested;type:numeric;comment:上月完成固投(万)"`
	YearSumFixedInvested   *float64 `gorm:"column:year_sum__fixed_invested;type:numeric;comment:当年累计固投(万)"`
	ActualProgress         string   `gorm:"column:actual_progress;type:text;comment:本月完成形象进度"`
	IsHelp                 *bool    `gorm:"column:is_help;type:boolean;comment:是否需协调解决问题"`
	ProblemType            *int     `gorm:"column:problem_type;type:integer;comment:问题类型"`
	ProblemDetail          string   `gorm:"column:problem_detail;type:varchar(500);comment:问题详情"`
	IsChange               *bool    `gorm:"column:is_change;type:boolean;comment:是否本月产生联系单变更"`
	ChangeType             *int     `gorm:"column:change_type;type:integer;comment:变更类型;0:较小变更,1:一般变更,2:较大变更,3:重大变更"`
	ChangeContent          string   `gorm:"column:change_content;type:varchar(500);comment:变更内容"`
	ChangeMoney            *float64 `gorm:"column:change_money;type:numeric;comment:变更金额(万)"`
	// ContractName           string   `gorm:"column:contract_name;type:varchar(100);comment:合同名称"`
	// ContractMoney          *float64 `gorm:"column:contract_money;type:numeric;comment:变更金额(万)"`
	// ContractComment string `gorm:"column:contract_name;type:varchar(255);comment:合同备注"`
	Comment string `gorm:"column:comment;type:text;comment:备注"`
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
	Month        int      `gorm:"column:month"`
	PlanInvest   *float64 `gorm:"column:plan_invest"`
	PlanProgress string   `gorm:"column:plan_progress"`
}
