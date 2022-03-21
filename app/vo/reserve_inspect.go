package vo

type ReserveInspectParam struct {
	//项目名称
	Name string `json:"name"`
	// 项目级别
	Level *int `json:"level"`
	// 项目类型
	ProjectType *int `json:"project_type"`
	// 建设主体 ***注意:（所有参数，有就传，无则不传）***
	ConstructSubject string `json:"construct_subject"`
	// 计划开始时间
	PlanBegin string `json:"plan_begin"`
	// 计划结束时间
	PlanEnd string `json:"plan_end"`
}
