package reserve

import (
	"lpms/app/models/internal/common"
	"lpms/app/models/internal/implement"
	"lpms/app/models/tables"
	"lpms/constant"
	"time"

	"github.com/goccy/go-json"

	"gorm.io/gorm"
)

type ReservePro struct {
	common.Base             `gorm:"embedded"`
	ID                      int64           `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	Level                   *int            `gorm:"column:level;type:integer;comment:项目级别 0:区级,1:街镇级"`
	Name                    string          `gorm:"column:name;type:varchar(60);not null;comment:项目名称"`
	ConstructSubject        string          `gorm:"column:construct_subject;type:varchar(60);comment:建设主体"`
	ConstructSite           string          `gorm:"column:construct_site;type:varchar(200);comment:建设地点"`
	ProjectType             *int            `gorm:"column:project_type;type:integer;comment:项目类型 0:安置房,1:道路交通,2:市政设施;3:提升整治;4:卫生;5:五水共治;6:学校;7:其他"`
	PlanBegin               *time.Time      `gorm:"column:plan_begin;type:timestamp;comment:计划开工时间"`
	Period                  *int            `gorm:"column:period;type:integer;comment:建设周期"`
	PointType               *int            `gorm:"column:point_type;type:integer;comment:重点类型 0:省重点实施项目,1:省重点预备项目,2:省重大产业项目;3:省4+1项目;4:省6千亿项目;5:市重点实施项目;6:市重点预备项目;7:无重点类型"`
	ImplementType           *int            `gorm:"column:implement_type;type:integer;comment:实施类型 0:新开工,1:续建"`
	ConstructContentScope   string          `gorm:"column:construct_content_scope;type:text;comment:建设内容及规模"`
	ConstructBasisNecessity string          `gorm:"column:construct_basis_necessity;type:text;comment:建设依据及必要性"`
	EnterDBType             *int            `gorm:"column:enter_db_type;type:integer;comment:入库类别 0:A类,1:B类;2:C类"`
	IsLandUse               *bool           `gorm:"column:is_land_use;type:boolean;comment:是否有用地情况"`
	Total                   *float64        `gorm:"column:total;type:numeric;comment:总用亩"`
	Add                     *float64        `gorm:"column:add;type:numeric;comment:新增建设用地"`
	NoConformUsePlan        *float64        `gorm:"column:no_conform_use_plan;type:numeric;comment:不符合土地利用规划面积"`
	SiteRed                 *int            `gorm:"column:site_red;type:integer;comment:选址红线 0:有拆迁,1:无拆迁"`
	SitePhoto               string          `gorm:"column:site_photo;type:varchar(40);comment:无拆迁照片"`
	NeedCollect             *float64        `gorm:"column:need_collect;type:numeric;comment:需征地面积"`
	NeedPeopleMove          *int            `gorm:"column:need_people_move;type:integer;comment:需拆迁农户/居民数(人)"`
	CompanyBusiness         *int            `gorm:"column:company_business;type:integer;comment:企/事业单位(家)"`
	UploadCadID             string          `gorm:"column:upload_cad_id;type:varchar(40);comment:CAD文件ID"`
	TotalInvestment         *float64        `gorm:"column:total_investment;type:numeric;comment:总投资"`
	ProjectComsumption      *float64        `gorm:"column:project_consumption;type:numeric;comment:工程费用"`
	MoveLandComsumption     *float64        `gorm:"column:move_land_comsumption;type:numeric;comment:征迁/土地费用"`
	InvestmentDetail        json.RawMessage `gorm:"column:investment_detail;type:jsonb;comment:资金详情"`
	Contract                string          `gorm:"column:contract;type:varchar(20);comment:前期工作联系人"`
	Phone                   string          `gorm:"column:phone;type:varchar(20);comment:联系人手机号"`
	Status                  int             `gorm:"column:status;type:integer;comment:项目状态 0:草稿,1:已入库,2:前期计划;3:已发文"`
	IsCaseFinish            *bool           `gorm:"column:is_case_finish;type:boolean;comment:方案是否完成"`
	IsResearch              *int            `gorm:"column:is_research;type:integer;comment:是否可研编制; 0:编制中 1:已完成"`
}

type InvestDetail struct {
	Info json.RawMessage `gorm:"column:info"`
}

type ListReservePro struct {
	ID               int64     `gorm:"column:id"`
	Name             string    `gorm:"column:name"`
	Level            *int      `gorm:"column:level"`
	ProjectType      *int      `gorm:"column:project_type"`
	ConstructSubject string    `gorm:"column:construct_subject"`
	CreateAt         time.Time `gorm:"column:create_at"`
	Status           int       `gorm:"column:status"`
}

func (b *ReservePro) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	b.CreateAt = now
	b.UpdateAt = now
	return nil
}

func (b *ReservePro) BeforeUpdate(tx *gorm.DB) error {
	b.UpdateAt = time.Now()
	return nil
}

func (ReservePro) TableName() string {
	return tables.Reserve
}

func (r *ReservePro) ToGovReserveModel(openID string) *implement.ImplementGov {
	return &implement.ImplementGov{
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
		Base: common.Base{
			UpdateBy: openID,
			CreateBy: openID,
		},
	}
}

type ReserveAnalysis1 struct {
	Bucket     string `gorm:"column:bucket"`
	Total      int64  `gorm:"column:total"`
	Add        int64  `gorm:"column:add"`
	OutStorage int64  `gorm:"column:out_storage"`
	EnteredDB  int64  `gorm:"column:enter_ddb"`
}

type ReserveAnalysis struct {
	Bucket string `gorm:"column:bucket"`
	// Total      int64  `gorm:"column:total"`
	// Add        int64  `gorm:"column:add"`
	// OutStorage int64  `gorm:"column:out_storage"`
	// EnteredDB  int64  `gorm:"column:enter_ddb"`
	Status int `gorm:"column:status"`
	Count  int64 `gorm:"column:count"`
}
