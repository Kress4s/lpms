package vo

import (
	"lpms/app/models"
	"time"

	"github.com/goccy/go-json"
)

type GovProgressReq struct {
	Info []GovProgressInfo `json:"info"`
}

type GovProgressInfo struct {
	// id,**新添加的不需要传**
	ID *int64 `json:"id,omitempty"`
	// 项目ID
	ProjectID int64 `json:"project_id"`
	// 年份
	Year int `json:"year"`
	//月份
	Month int `json:"month"`
	//本月计划投资额(万)
	PlanInvest *float64 `json:"plan_invest"`
	//本月计划形象进度
	PlanProgress string `json:"plan_progress"`
}

func (g *GovProgressReq) ToModel(openID string) []models.GovProgress {
	gps := make([]models.GovProgress, 0, len(g.Info))
	for i := range g.Info {
		if g.Info[i].ID != nil {
			gps = append(gps, models.GovProgress{
				ID:           *g.Info[i].ID,
				ProjectID:    g.Info[i].ProjectID,
				Year:         g.Info[i].Year,
				Month:        g.Info[i].Month,
				PlanInvest:   g.Info[i].PlanInvest,
				PlanProgress: g.Info[i].PlanProgress,
				Base: models.Base{
					UpdateBy: openID,
					CreateBy: openID,
				},
			})
		} else {
			gps = append(gps, models.GovProgress{
				ProjectID:    g.Info[i].ProjectID,
				Year:         g.Info[i].Year,
				Month:        g.Info[i].Month,
				PlanInvest:   g.Info[i].PlanInvest,
				PlanProgress: g.Info[i].PlanProgress,
				Status:       0,
				Base: models.Base{
					UpdateBy: openID,
					CreateBy: openID,
				},
			})
		}
	}
	return gps
}

type GovProgressResp struct {
	// id
	ID int64 `json:"id"`
	// 项目ID
	ProjectID int64 `json:"project_id"`
	// 年份
	Year int `json:"year"`
	//月份
	Month int `json:"month"`
	//本月计划投资额(万)
	PlanInvest *float64 `json:"plan_invest"`
	//本月计划形象进度
	PlanProgress string `json:"plan_progress"`
	//本月完成投资额(万)
	PlanInvested *float64 `json:"plan_invested"`
	//当年累计投资额(万)
	YearSumInvested *float64 `json:"year_sum_invested"`
	//上月完成固投(万)
	LastMonthFixedInvested *float64 `json:"last_month_fixed_invested"`
	//当年累计固投(万)
	YearSumFixedInvested *float64 `json:"year_sum_fixed_invested"`
	//本月完成形象进度
	ActualProgress string `json:"actual_progress"`
	//需协调问题详情
	ProblemDetail string `json:"problem_detail"`
	//本月产生联系单变更
	ChangeContent string `json:"change_content"`
	// 本月新增合同信息
	Contracts string `json:"contracts"`
	//备注
	Comment string `json:"comment"`
	// 一月至本月计划累计完成投资额
	TotalPlanInvested float64 `json:"total_plan_invested"`
	// 开工至今累计投资额
	StartSumInvested float64 `json:"start_sum_invest"`
	// 开工至今累计固投
	StartFixedInvested float64 `json:"start_fixed_invested"`
	// 填报状态  0:未提交，1:已提交
	Status int `json:"status"`
}

func NewGovProgressResponse(r *models.GovProgress, total_plan_invested, start_sum, start_fixed float64) (*GovProgressResp, error) {
	return &GovProgressResp{
		ID:                     r.ID,
		ProjectID:              r.ProjectID,
		Year:                   r.Year,
		Month:                  r.Month,
		PlanInvest:             r.PlanInvest,
		PlanProgress:           r.PlanProgress,
		PlanInvested:           r.PlanInvested,
		YearSumInvested:        r.YearSumInvested,
		LastMonthFixedInvested: r.LastMonthFixedInvested,
		YearSumFixedInvested:   r.YearSumFixedInvested,
		ActualProgress:         r.ActualProgress,
		ProblemDetail:          string(r.ProblemDetail),
		ChangeContent:          string(r.ChangeContent),
		Contracts:              string(r.Contracts),
		Comment:                r.Comment,
		TotalPlanInvested:      total_plan_invested,
		StartSumInvested:       start_sum,
		StartFixedInvested:     start_fixed,
		Status:                 r.Status,
	}, nil
}

type GovProgressUpdateReq struct {
	//本月完成投资额(万)
	PlanInvested *float64 `json:"plan_invested"`
	//当年累计投资额(万)
	YearSumInvested *float64 `json:"year_sum_invested"`
	//上月完成固投(万)
	LastMonthFixedInvested *float64 `json:"last_month_fixed_invested"`
	//当年累计固投(万)
	YearSumFixedInvested *float64 `json:"year_sum__fixed_invested"`
	//本月完成形象进度
	ActualProgress string `json:"actual_progress"`
	//需协调问题详情
	ProblemDetail string `json:"problem_detail"`
	//本月产生联系单变更
	ChangeContent string `json:"change_content"`
	// 本月新增合同信息
	Contracts string `json:"contracts"`
	//备注
	Comment string `json:"comment"`
	// 操作方式：1：保存 2：提交
	Method int `json:"method"`
}

func (g *GovProgressUpdateReq) ToMap(openID string) map[string]interface{} {
	status := int(0)
	if g.Method == 2 {
		status = 1
	}
	return map[string]interface{}{
		"plan_invested":             g.PlanInvested,
		"year_sum_invested":         g.YearSumInvested,
		"last_month_fixed_invested": g.LastMonthFixedInvested,
		"year_sum__fixed_invested":  g.YearSumFixedInvested,
		"actual_progress":           g.ActualProgress,
		"problem_detail":            json.RawMessage([]byte(g.ProblemDetail)),
		"change_content":            json.RawMessage([]byte(g.ChangeContent)),
		"contracts":                 json.RawMessage([]byte(g.Contracts)),
		"comment":                   g.Comment,
		"update_by":                 openID,
		"status":                    status,
	}
}

type ListGovProgressPlan struct {
	// id
	ID int64 `json:"id"`
	// 年份
	Year int `json:"year"`
	// 月份数
	Month int `json:"month"`
	//本月计划投资额(万)
	PlanInvest *float64 `json:"plan_invest"`
	//本月计划形象进度
	PlanProgress string `json:"plan_progress"`
	// 计划提交状态 0:未提交，1:已提交
	Status int `json:"status"`
}

type GovProgressCompare struct {
	// 年份
	Year int `json:"year"`
	// 月份
	Month int `json:"month"`
	// 本月计划投资额
	PlanInvest *float64 `json:"plan_invest"`
	// 本月计划形象进度
	PlanProgress string `json:"plan_progress"`
	// 本月完成投资额
	PlanInvested *float64 `json:"plan_invested"`
	// 本月完成形象进度
	ActualProgress string `json:"actual_progress"`
	// 完成度
	Completeness float64 `json:"completeness"`
}

func MultiAddProcess(projectID int64) []models.GovProgress {
	res := make([]models.GovProgress, 0, 12)
	now := time.Now()
	for i := 1; i <= 12; i++ {
		res = append(res, models.GovProgress{
			ProjectID: projectID,
			Year:      now.Year(),
			Month:     i,
		})
	}
	return res
}
