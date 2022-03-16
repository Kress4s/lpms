/*
 * Copyright (c) 1998-2022 江苏斯菲尔电气股份有限公司
 * All rights reserved.
 *
 * Filename: reserve_vo.go
 * Description:
 *
 * Created by xiayoushuang at 2022/03/15 11:41
 * http://www.sfere-elec.com/
 */

package vo

import (
	"lpms/app/models"
	"time"

	"github.com/goccy/go-json"
)

// type LandUseReq struct {
// 	// 总用亩
// 	Total float64 `json:"total"`
// 	// 新增建设用地
// 	Add float64 `json:"add"`
// 	// 不符合土地利用规划面积
// 	NoConformUsePlan float64 `json:"no_conform_use_plan"`
// 	// 选址红线 0:有拆迁,1:无拆迁
// 	SiteRed int `json:"site_red"`
// 	// 需征地面积
// 	NeedCollect float64 `json:"need_collect"`
// 	// 需拆迁农户/居民数(人)
// 	NeedPeopleMove int `json:"need_people_move"`
// 	// 企/事业单位(家)
// 	CompanyBusiness int `json:"company_business"`
// 	// CAD文件ID(上传文件接口返回的ID)
// 	UploadCadID string `json:"upload_cad_id"`
// }

type ReserveReq struct {
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
	// "[{\\"type\\":0, \\"total\\":100, \\"detail\\":[{\\"total\\": 100,\\"year\\": \\"2022\\",\\"value\\":20,\\"comment\\":\\"xxx\\"}, {\\"total\\": 100,\\"year\\": \\"2023\\",\\"value\\":30,\\"comment\\":\\"xxx\\"}, ...]}, {}...]"
	// type说明： 0:区财政;1:自筹;2:其他
	InvestmentDetail string `json:"investment_detail"`
	// 前期工作联系人
	Contract string `json:"contract"`
	// 联系人手机号
	Phone string `json:"phone"`
}

func (r *ReserveReq) ToModel(openID string) *models.ReservePro {
	return &models.ReservePro{
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
		Status:                  0,
		Base: models.Base{
			UpdateBy: openID,
			CreateBy: openID,
		},
	}
}

type ReserveResp struct {
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
	InvestmentDetail []InvestmentDetail `json:"investment_detail"`
	// 前期工作联系人
	Contract string `json:"contract"`
	// 联系人手机号
	Phone string `json:"phone"`
	// 项目状态 0:草稿,1:已入库,2:前期计划;3:已发文"
	Status int `json:"status"`
	// 创建时间
	CreateAt string `json:"create_at"`
}

type InvestmentDetail struct {
	// 投资类型
	Type int `json:"type"`
	// 总投资
	Total float64 `json:"total"`
	// 投资情况数额细节
	Detail []InvestDetail `json:"detail"`
}

type InvestDetail struct {
	// 资金类别 总投资
	Total float64 `json:"total"`
	// 年份
	Year string `json:"year"`
	// 投资数额
	Value float64 `json:"value"`
	// 备注
	Comment string `json:"comment"`
}

func NewReserveProResponse(r *models.ReservePro, invests []models.InvestDetail) (*ReserveResp, error) {
	investment := make([]InvestmentDetail, 0, len(invests))
	for i := range invests {
		invest := InvestmentDetail{}
		if err := json.Unmarshal([]byte(invests[i].Info), &invest); err != nil {
			return nil, err
		}
		investment = append(investment, invest)
	}
	return &ReserveResp{
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
		InvestmentDetail:        investment,
		Contract:                r.Contract,
		Phone:                   r.Phone,
	}, nil
}

type ReserveFilterParam struct {
	// 以下所有参数，有就传，无则不传
	//项目名称
	Name string `json:"name"`
	// 项目级别
	Level *int `json:"level"`
	// 项目类型
	ProjectType *int `json:"project_type"`
	// 建设主体
	ConstructSubject string `json:"construct_subject"`
	// 计划开始时间
	PlanBegin string `json:"plan_begin"`
	// 计划周期(根据起止时间计算相差的月数)
	Period *int `json:"period"`
	// 状态
	Status *int `json:"status"`
}

type ListReserveProResp struct {
	ID               int64      `json:"id"`
	Name             string     `json:"name"`
	Level            *int       `json:"level"`
	ProjectType      *int       `json:"project_type"`
	ConstructSubject string     `json:"construct_subject"`
	CreateAt         *time.Time `json:"create_at"`
	Status           int        `json:"status"`
}

type ReserveUpdateReq struct {
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
	// "[{\\"type\\":0, \\"total\\":100, \\"detail\\":[{\\"total\\": 100,\\"year\\": \\"2022\\",\\"value\\":20,\\"comment\\":\\"xxx\\"}, {\\"total\\": 100,\\"year\\": \\"2023\\",\\"value\\":30,\\"comment\\":\\"xxx\\"}, ...]}, {}...]"
	// type说明： 0:区财政;1:自筹;2:其他
	InvestmentDetail string `json:"investment_detail"`
	// 前期工作联系人
	Contract string `json:"contract"`
	// 联系人手机号
	Phone string `json:"phone"`
}

func (r *ReserveUpdateReq) ToMap(openID string) map[string]interface{} {
	return map[string]interface{}{
		"level":                     r.Level,
		"name":                      r.Name,
		"construct_subject":         r.ConstructSubject,
		"construct_site":            r.ConstructSite,
		"project_type":              r.ProjectType,
		"plan_begin":                r.PlanBegin,
		"period":                    r.Period,
		"point_type":                r.PointType,
		"implement_type":            r.ImplementType,
		"construct_content_scope":   r.ConstructContentScope,
		"construct_basis_necessity": r.ConstructBasisNecessity,
		"enter_db_type":             r.EnterDBType,
		"is_land_use":               r.IsLandUse,
		"total":                     r.Total,
		"add":                       r.Add,
		"no_conform_use_plan":       r.NoConformUsePlan,
		"site_red":                  r.SiteRed,
		"site_photo":                r.SitePhoto,
		"need_collect":              r.NeedCollect,
		"need_people_move":          r.NeedPeopleMove,
		"company_business":          r.CompanyBusiness,
		"upload_cad_id":             r.UploadCadID,
		"total_investment":          r.TotalInvestment,
		"project_consumption":       r.ProjectComsumption,
		"move_land_comsumption":     r.MoveLandComsumption,
		"investment_detail":         r.InvestmentDetail,
		"contract":                  r.Contract,
		"phone":                     r.Phone,
		"create_by":                 openID,
	}
}
