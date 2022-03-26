package vo

import (
	"lpms/app/models"
	"lpms/constant"
	"time"

	"github.com/goccy/go-json"
)

type ImplementGovFilterParam struct {
	//项目名称
	Name string `json:"name"`
	// 项目级别
	Level *int `json:"level"`
	// 项目类型
	ProjectType *int `json:"project_type"`
	// 建设主体  ***注意:（所有参数，有就传，无则不传）***
	ConstructSubject string `json:"construct_subject"`
	// 标签 0:省重点实施项目,1:省重点预备项目,2:省重大产业项目;3:省4+1项目;4:省6千亿项目;5:市重点实施项目;6:市重点预备项目;7:无重点类型;8: 152工程
	PointType *int `json:"point_type"`
	// 计划开工开始时间
	PlanBegin string `json:"plan_begin"`
	// 计划开工结束时间
	PlanEnd string `json:"plan_end"`
	// 开工开始时间
	StartTime string `json:"start_time"`
	// 开工结束时间
	EndTime string `json:"end_time"`
	// 状态(必传参数) -1: 当年全部(注意cur_year_all_begin和cur_year_all_end带上) 0:未开工, 2:开工建设; 4:竣工库(如果是 当年竣工，必须要带上cur_year_all_begin和cur_year_all_end参数)
	Status *int `json:"status"`
	// 当年 起始时间(闭区间) eg: 2022-01-01 00:00:00
	CurYearBegin string `json:"cur_year_all_begin"`
	// 当年 终止时间(开区间) eg : 2023-01-01 00:00:00
	CurYearEnd string `json:"cur_year_all_end"`
	// 总投资额范围开始
	BeginInvest *float64 `json:"begin_invest"`
	// 总投资额范围结束
	EndInvest *float64 `json:"end_invest"`
	// 责任单位
	DutyUnit string `json:"duty_unit"`
}

type ImplementGovCountFilter struct {
	//项目名称
	Name string `json:"name"`
	// 项目级别
	Level *int `json:"level"`
	// 项目类型
	ProjectType *int `json:"project_type"`
	// 建设主体  ***注意:（所有参数，有就传，无则不传）***
	ConstructSubject string `json:"construct_subject"`
	// 标签 0:省重点实施项目,1:省重点预备项目,2:省重大产业项目;3:省4+1项目;4:省6千亿项目;5:市重点实施项目;6:市重点预备项目;7:无重点类型;8: 152工程
	PointType *int `json:"point_type"`
	// 计划开工开始时间
	PlanBegin string `json:"plan_begin"`
	// 计划开工结束时间
	PlanEnd string `json:"plan_end"`
	// 开工开始时间
	StartTime string `json:"start_time"`
	// 开工结束时间
	EndTime string `json:"end_time"`
	// 总投资额范围开始
	BeginInvest *float64 `json:"begin_invest"`
	// 总投资额范围结束
	EndInvest *float64 `json:"end_invest"`
	// 责任单位
	DutyUnit string `json:"duty_unit"`
}

type ImplementGovReq struct {
	// 项目级别 0:区级,1:街镇级
	Level *int `json:"level"`
	// 项目名称
	Name string `json:"name"`
	// 建设主体
	ConstructSubject string `json:"construct_subject"`
	// 建设地点
	ConstructSite string `json:"construct_site"`
	// 项目类型; 0:安置房,1:道路交通,2:市政设施;3:提升整治;4:卫生;5:五水共治;6:学校;7:其他
	ProjectType *int `json:"project_type"`
	// 计划开工时间
	PlanBegin *time.Time `json:"plan_begin"`
	// 建设周期
	Period *int `json:"period"`
	// 重点类型; 0:省重点实施项目,1:省重点预备项目,2:省重大产业项目;3:省4+1项目;4:省6千亿项目;5:市重点实施项目;6:市重点预备项目;7:无重点类型
	PointType *int `json:"point_type"`
	// 实施类型 0:新开工,1:续建
	ImplementType *int `json:"implement_type"`
	// 建设内容及规模
	ConstructContentScope string `json:"construct_content_scope"`
	// 建设依据及必要性
	ConstructBasisNecessity string `json:"construct_basis_necessity"`
	// 入库类别 0:A类,1:B类;2:C类
	EnterDBType *int `json:"enter_db_type"`
	// 是否有用地情况
	IsLandUse *bool `json:"is_land_use"`
	// 总用亩
	Total *float64 `json:"total"`
	// 新增建设用地
	Add *float64 `json:"add"`
	// 不符合土地利用规划面积
	NoConformUsePlan *float64 `json:"no_conform_use_plan"`
	// 选址红线 0:有拆迁,1:无拆迁
	SiteRed *int `json:"site_red"`
	// 无拆迁照片
	SitePhoto string `json:"site_photo"`
	// 需征地面积
	NeedCollect *float64 `json:"need_collect"`
	// 需拆迁农户/居民数(人)
	NeedPeopleMove *int `json:"need_people_move"`
	// 企/事业单位(家)
	CompanyBusiness *int `json:"company_business"`
	// CAD文件ID(上传文件接口返回的ID)
	UploadCadID string `json:"upload_cad_id"`
	// 总投资
	TotalInvestment *float64 `json:"total_investment"`
	// 工程费用
	ProjectComsumption *float64 `json:"project_consumption"`
	// 征迁/土地费用
	MoveLandComsumption *float64 `json:"move_land_comsumption"`
	// 资金详情 eg:
	// "[{\\"type\\":0, \\"total\\":100, \\"detail\\":{\\"total\\": 100,\\"2022\\": 20,\\"comment\\":\\"xxx\\"}, {\\"total\\": 100,\\"2023\\": 30,\\"comment\\":\\"xxx\\"}, ...}, {}...]"
	// type说明： 0:区财政;1:自筹;2:其他
	InvestmentDetail string `json:"investment_detail"`
	// 前期工作联系人
	Contract string `json:"contract"`
	// 联系人手机号
	Phone string `json:"phone"`
	// 项目编码
	ProjectCode string `json:"project_code"`
	// 责任单位
	DutyUnit string `json:"duty_unit"`
	// 项目本质类型 1:政府项目 2：产业项目
	Type int `json:"type"`
}

func (r *ImplementGovReq) ToModel(openID string) *models.ImplementGov {
	return &models.ImplementGov{
		Level:                   r.Level,
		Name:                    r.Name,
		ConstructSubject:        r.ConstructSubject,
		ConstructSite:           r.ConstructSite,
		ProjectType:             r.ProjectType,
		PlanBegin:               r.PlanBegin,
		Period:                  r.Period,
		PointType:               r.PointType,
		ImplementType:           r.ImplementType,
		ConstructContentScope:   r.ConstructContentScope,
		ConstructBasisNecessity: r.ConstructBasisNecessity,
		EnterDBType:             r.EnterDBType,
		IsLandUse:               r.IsLandUse,
		Total:                   r.Total,
		Add:                     r.Add,
		NoConformUsePlan:        r.NoConformUsePlan,
		SiteRed:                 r.SiteRed,
		SitePhoto:               r.SitePhoto,
		NeedCollect:             r.NeedCollect,
		NeedPeopleMove:          r.NeedPeopleMove,
		CompanyBusiness:         r.CompanyBusiness,
		UploadCadID:             r.UploadCadID,
		TotalInvestment:         r.TotalInvestment,
		ProjectComsumption:      r.ProjectComsumption,
		MoveLandComsumption:     r.MoveLandComsumption,
		InvestmentDetail:        json.RawMessage([]byte(r.InvestmentDetail)),
		Contract:                r.Contract,
		Phone:                   r.Phone,
		Status:                  constant.UnStart,
		ProjectCode:             r.ProjectCode,
		DutyUint:                r.DutyUnit,
		Type:                    r.Type,
		Base: models.Base{
			UpdateBy: openID,
			CreateBy: openID,
		},
	}
}

type ImplementGovResp struct {
	// id
	ID int64 `json:"id"`
	// 项目级别 0:区级,1:街镇级
	Level *int `json:"level"`
	// 项目名称
	Name string `json:"name"`
	// 建设主体
	ConstructSubject string `json:"construct_subject"`
	// 建设地点
	ConstructSite string `json:"construct_site"`
	// 项目类型; 0:安置房,1:道路交通,2:市政设施;3:提升整治;4:卫生;5:五水共治;6:学校;7:其他
	ProjectType *int `json:"project_type"`
	// 计划开工时间
	PlanBegin *time.Time `json:"plan_begin"`
	// 建设周期
	Period *int `json:"period"`
	// 重点类型; 0:省重点实施项目,1:省重点预备项目,2:省重大产业项目;3:省4+1项目;4:省6千亿项目;5:市重点实施项目;6:市重点预备项目;7:无重点类型
	PointType *int `json:"point_type"`
	// 实施类型 0:新开工,1:续建
	ImplementType *int `json:"implement_type"`
	// 建设内容及规模
	ConstructContentScope string `json:"construct_content_scope"`
	// 建设依据及必要性
	ConstructBasisNecessity string `json:"construct_basis_necessity"`
	// 入库类别 0:A类,1:B类;2:C类
	EnterDBType *int `json:"enter_db_type"`
	// 是否有用地情况
	IsLandUse *bool `json:"is_land_use"`
	// 总用亩
	Total *float64 `json:"total"`
	// 新增建设用地
	Add *float64 `json:"add"`
	// 不符合土地利用规划面积
	NoConformUsePlan *float64 `json:"no_conform_use_plan"`
	// 选址红线 0:有拆迁,1:无拆迁
	SiteRed *int `json:"site_red"`
	// 无拆迁照片
	SitePhoto string `json:"site_photo"`
	// 需征地面积
	NeedCollect *float64 `json:"need_collect"`
	// 需拆迁农户/居民数(人)
	NeedPeopleMove *int `json:"need_people_move"`
	// 企/事业单位(家)
	CompanyBusiness *int `json:"company_business"`
	// CAD文件ID(上传文件接口返回的ID)
	UploadCadID string `json:"upload_cad_id"`
	// 总投资
	TotalInvestment *float64 `json:"total_investment"`
	// 工程费用
	ProjectComsumption *float64 `json:"project_consumption"`
	// 征迁/土地费用
	MoveLandComsumption *float64 `json:"move_land_comsumption"`
	// 资金详情
	InvestmentDetail string `json:"investment_detail"`
	// 前期工作联系人
	Contract string `json:"contract"`
	// 联系人手机号
	Phone string `json:"phone"`
	// 项目状态 0:未开工,1:开工待审核,2:已开工;3:竣工待审核;4:已竣工"
	Status int `json:"status"`
	// 项目编码
	ProjectCode string `json:"project_code"`
	// 责任单位
	DutyUnit string `json:"duty_unit"`
	// 项目本质类型 1:政府项目 2：产业项目
	Type int `json:"type"`
}

func NewImplementGovResponse(r *models.ImplementGov) (*ImplementGovResp, error) {
	return &ImplementGovResp{
		ID:                      r.ID,
		Level:                   r.Level,
		Name:                    r.Name,
		ConstructSubject:        r.ConstructSubject,
		ConstructSite:           r.ConstructSite,
		ProjectType:             r.ProjectType,
		PlanBegin:               r.PlanBegin,
		Period:                  r.Period,
		PointType:               r.PointType,
		ImplementType:           r.ImplementType,
		ConstructContentScope:   r.ConstructContentScope,
		ConstructBasisNecessity: r.ConstructBasisNecessity,
		EnterDBType:             r.EnterDBType,
		IsLandUse:               r.IsLandUse,
		Total:                   r.Total,
		Add:                     r.Add,
		NoConformUsePlan:        r.NoConformUsePlan,
		SiteRed:                 r.SiteRed,
		SitePhoto:               r.SitePhoto,
		NeedCollect:             r.NeedCollect,
		NeedPeopleMove:          r.NeedPeopleMove,
		CompanyBusiness:         r.CompanyBusiness,
		UploadCadID:             r.UploadCadID,
		TotalInvestment:         r.TotalInvestment,
		ProjectComsumption:      r.ProjectComsumption,
		MoveLandComsumption:     r.MoveLandComsumption,
		InvestmentDetail:        string(r.InvestmentDetail),
		Contract:                r.Contract,
		Phone:                   r.Phone,
		Status:                  r.Status,
		ProjectCode:             r.ProjectCode,
		DutyUnit:                r.DutyUint,
		Type:                    r.Type,
	}, nil
}

type ListImplementGovResp struct {
	// id
	ID int64 `json:"id"`
	// 名称
	Name string `json:"name"`
	// 项目级别
	Level *int `json:"level"`
	// 项目类型
	ProjectType *int `json:"project_type"`
	// 建设主体
	ConstructSubject string `json:"construct_subject"`
	// 计划开工时间
	PlanBegin *time.Time `json:"plan_begin"`
	// 竣工时间
	FinishTime *time.Time `json:"finish_time"`
	// 状态
	Status int `json:"status"`
	// 实际开工时间
	StartTime *time.Time `json:"start_time"`
	// 项目编码
	ProjectCode string `json:"project_code"`
	// 责任单位
	DutyUnit string `json:"duty_unit"`
	// 项目本质类型 1:政府项目 2：产业项目
	Type int `json:"type"`
}

type StatusCountResp struct {
	// 状态：-2：当年全部，-1：当年竣工，0：未开工，2：开工建设，4:竣工库
	Status int `json:"status"`
	// 数量
	Count int64 `json:"count"`
}
