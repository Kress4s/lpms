package vo

import "lpms/app/models"

type GovProgressReq struct {
	Info []GovProgressInfo `json:"info"`
	// // 项目ID
	// ProjectID int64 `json:"project_id"`
	// //月份
	// Month int `json:"month"`
	// //本月计划投资额(万)
	// PlanInvest *float64 `json:"plan_invest"`
	// //本月计划形象进度
	// PlanProgress string `json:"plan_progress"`
	// //本月完成投资额(万)
	// PlanInvested *float64 `json:"plan_invested"`
	// //当年累计投资额(万)
	// YearSumInvested *float64 `json:"year_sum_invested"`
	// //本月支付工程款(万)
	// PayProject *float64 `json:"pay_project"`
	// //上月完成固投(万)
	// LastMonthFixedInvested *float64 `json:"last_month_fixed_invested"`
	// //当年累计固投(万)
	// YearSumFixedInvested *float64 `json:"year_sum__fixed_invested"`
	// //本月完成形象进度
	// ActualProgress string `json:"actual_progress"`
	// //是否需协调解决问题
	// IsHelp *bool `json:"is_help"`
	// //问题类型
	// ProblemType *int `json:"problem_type"`
	// //问题详情
	// ProblemDetail string `json:"problem_detail"`
	// //是否本月产生联系单变更
	// IsChange *bool `json:"is_change"`
	// //变更类型;0:较小变更,1:一般变更,2:较大变更,3:重大变更
	// ChangeType *int `json:"change_type"`
	// //变更内容
	// ChangeContent string `json:"change_content"`
	// //变更金额
	// ChangeMoney *float64 `json:"change_money"`
	// //备注
	// Comment string `json:":comment"`
}

type GovProgressInfo struct {
	// id,**新添加的不需要传**
	ID *int64 `json:"id,omitempty"`
	// 项目ID
	ProjectID int64 `json:"project_id"`
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
				Month:        g.Info[i].Month,
				PlanInvest:   g.Info[i].PlanInvest,
				PlanProgress: g.Info[i].PlanProgress,
				// PlanInvested:           g.PlanInvested,
				// YearSumInvested:        g.YearSumInvested,
				// PayProject:             g.PayProject,
				// LastMonthFixedInvested: g.LastMonthFixedInvested,
				// YearSumFixedInvested:   g.YearSumFixedInvested,
				// ActualProgress:         g.ActualProgress,
				// IsHelp:                 g.IsHelp,
				// ProblemType:            g.ProblemType,
				// ProblemDetail:          g.ProblemDetail,
				// IsChange:               g.IsChange,
				// ChangeType:             g.ChangeType,
				// ChangeContent:          g.ChangeContent,
				// ChangeMoney:            g.ChangeMoney,
				// Comment:                g.Comment,
				Base: models.Base{
					UpdateBy: openID,
					CreateBy: openID,
				},
			})
		} else {
			gps = append(gps, models.GovProgress{
				ProjectID:    g.Info[i].ProjectID,
				Month:        g.Info[i].Month,
				PlanInvest:   g.Info[i].PlanInvest,
				PlanProgress: g.Info[i].PlanProgress,
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
	ID int64 `json:":id"`
	// 项目ID
	ProjectID int64 `json:"project_id"`
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
	//本月支付工程款(万)
	PayProject *float64 `json:"pay_project"`
	//上月完成固投(万)
	LastMonthFixedInvested *float64 `json:"last_month_fixed_invested"`
	//当年累计固投(万)
	YearSumFixedInvested *float64 `json:"year_sum__fixed_invested"`
	//本月完成形象进度
	ActualProgress string `json:"actual_progress"`
	//是否需协调解决问题
	IsHelp *bool `json:"is_help"`
	//问题类型
	ProblemType *int `json:"problem_type"`
	//问题详情
	ProblemDetail string `json:"problem_detail"`
	//是否本月产生联系单变更
	IsChange *bool `json:"is_change"`
	//变更类型;0:较小变更,1:一般变更,2:较大变更,3:重大变更
	ChangeType *int `json:"change_type"`
	//变更内容
	ChangeContent string `json:"change_content"`
	// //变更金额
	ChangeMoney *float64 `json:"change_money"`
	//备注
	Comment string `json:":comment"`
}

func NewGovProgressResponse(r *models.GovProgress) (*GovProgressResp, error) {
	return &GovProgressResp{
		ID:                     r.ID,
		ProjectID:              r.ProjectID,
		Month:                  r.Month,
		PlanInvest:             r.PlanInvest,
		PlanProgress:           r.PlanProgress,
		PlanInvested:           r.PlanInvested,
		YearSumInvested:        r.YearSumInvested,
		PayProject:             r.PayProject,
		LastMonthFixedInvested: r.LastMonthFixedInvested,
		YearSumFixedInvested:   r.YearSumFixedInvested,
		ActualProgress:         r.ActualProgress,
		IsHelp:                 r.IsHelp,
		ProblemType:            r.ProblemType,
		ProblemDetail:          r.ProblemDetail,
		IsChange:               r.IsChange,
		ChangeType:             r.ChangeType,
		ChangeContent:          r.ChangeContent,
		ChangeMoney:            r.ChangeMoney,
		Comment:                r.Comment,
	}, nil
}

type GovProgressUpdateReq struct {
	//本月完成投资额(万)
	PlanInvested *float64 `json:"plan_invested"`
	//当年累计投资额(万)
	YearSumInvested *float64 `json:"year_sum_invested"`
	//本月支付工程款(万)
	PayProject *float64 `json:"pay_project"`
	//上月完成固投(万)
	LastMonthFixedInvested *float64 `json:"last_month_fixed_invested"`
	//当年累计固投(万)
	YearSumFixedInvested *float64 `json:"year_sum__fixed_invested"`
	//本月完成形象进度
	ActualProgress string `json:"actual_progress"`
	//是否需协调解决问题
	IsHelp *bool `json:"is_help"`
	//问题类型
	ProblemType *int `json:"problem_type"`
	//问题详情
	ProblemDetail string `json:"problem_detail"`
	//是否本月产生联系单变更
	IsChange *bool `json:"is_change"`
	//变更类型;0:较小变更,1:一般变更,2:较大变更,3:重大变更
	ChangeType *int `json:"change_type"`
	//变更内容
	ChangeContent string `json:"change_content"`
	//变更金额
	ChangeMoney *float64 `json:"change_money"`
	//备注
	Comment string `json:":comment"`
}

func (g *GovProgressUpdateReq) ToMap(openID string) map[string]interface{} {
	return map[string]interface{}{
		"plan_invested":             g.PlanInvested,
		"year_sum_invested":         g.YearSumInvested,
		"pay_project":               g.PayProject,
		"last_month_fixed_invested": g.LastMonthFixedInvested,
		"year_sum__fixed_invested":  g.YearSumFixedInvested,
		"actual_progress":           g.ActualProgress,
		"is_help":                   g.IsHelp,
		"problem_type":              g.ProblemType,
		"problem_detail":            g.ProblemDetail,
		"is_change":                 g.IsChange,
		"change_type":               g.ChangeType,
		"change_content":            g.ChangeContent,
		"change_money":              g.ChangeMoney,
		"comment":                   g.Comment,
		"update_by":                 openID,
	}
}

type ListGovProgressPlan struct {
	// id
	ID int64 `json:"id"`
	// 月份数
	Month int `json:"month"`
	//本月计划投资额(万)
	PlanInvest *float64 `json:"plan_invest"`
	//本月计划形象进度
	PlanProgress string `json:"plan_progress"`
}

type GovProgressCompare struct {
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
