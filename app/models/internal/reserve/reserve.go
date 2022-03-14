package reserve

import (
	"encoding/json"
	"lpms/app/models/internal/common"
	"time"
)

type ReservePro struct {
	common.Base             `gorm:"embedded"`
	ID                      int64           `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	Level                   int             `gorm:"column:level;type:integer;comment:项目级别 0:区级,1:街镇级"`
	Name                    string          `gorm:"column:name;type:varchar(60);not null;comment:项目名称"`
	ConstructSubject        string          `gorm:"column:construct_subject;type:varchar(60);comment:建设主体"`
	ConstructSite           string          `gorm:"column:construct_site;type:varchar(200);comment:建设地点"`
	ProjectType             int             `gorm:"column:project_type;type:integer;comment:项目类型 0:安置房,1:道路交通,2:市政设施;3:提升整治;4:卫生;5:五水共治;6:学校;7:其他"`
	PlanBegin               time.Time       `gorm:"column:plan_begin;type:timestamp;comment:计划开工时间"`
	PointType               int             `gorm:"column:point_type;type:integer;comment:重点类型 0:省重点实施项目,1:省重点预备项目,2:省重大产业项目;3:省4+1项目;4:省6千亿项目;5:市重点实施项目;6:市重点预备项目;7:无重点类型"`
	ImplementType           int             `gorm:"column:implement_type;type:integer;comment:实施类型 0:新开工,1:续建"`
	ConstructContentScope   string          `gorm:"column:construct_content_scope;type:text;comment:建设内容及规模"`
	ConstructBasisNecessity string          `gorm:"column:construct_basis_necessity;type:text;comment:建设依据及必要性"`
	EnterDBType             int             `gorm:"column:enter_db_type;type:integer;comment:图库类别 0:A类,1:B类;2:C类"`
	IsLandUse               bool            `gorm:"column:is_land_use;type:boolean;comment:是否土地信息"`
	LandUseID               int64           `gorm:"column:land_use_id;comment:用地情况记录的ID"`
	TotalInvestment         int             `gorm:"column:total_investment;type:integer;comment:总投资"`
	ProjectComsumption      int             `gorm:"column:project_consumption;type:integer;comment:工程费用"`
	MoveLandComsumption     int             `gorm:"column:project_consumption;type:integer;comment:征迁/土地费用"`
	InvestmentDetail        json.RawMessage `gorm:"column:investment_detail;type:jsonb;comment:资金详情"`
	Contract                string          `gorm:"column:name;type:varchar(20);comment:前期工作联系人"`
	Phone                   string          `gorm:"column:name;type:varchar(20);comment:联系人手机号"`
	IsDraft                 bool            `gorm:"column:is_draft;type:boolean;comment:是否是草稿"`
}

/*
[{
"type":0,
"total":100,
"detail":[
{
"time": "2022",
"value":20
},
{
"time":"2023",
"value":80
}
]
},
{
"type":1,
"total":100,
"detail":[
{
"time": "2022",
"value":20
},
{
"time":"2023",
"value":80
}
]
}
]
*/
