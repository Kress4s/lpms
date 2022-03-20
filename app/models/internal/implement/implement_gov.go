package implement

import (
	"encoding/json"
	"lpms/app/models/internal/common"
	"lpms/app/models/tables"
	"time"

	"gorm.io/gorm"
)

type ImplementGov struct {
	common.Base             `gorm:"embedded"`
	ID                      int64           `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	Level                   *int            `gorm:"column:level;type:integer;comment:项目级别 0:区级,1:街镇级"`
	Name                    string          `gorm:"column:name;type:varchar(60);not null;comment:项目名称"`
	ConstructSubject        string          `gorm:"column:construct_subject;type:varchar(60);comment:建设主体"`
	ConstructSite           string          `gorm:"column:construct_site;type:varchar(200);comment:建设地点"`
	ProjectType             *int            `gorm:"column:project_type;type:integer;comment:项目类型 0:安置房,1:道路交通,2:市政设施;3:提升整治;4:卫生;5:五水共治;6:学校;7:其他"`
	PlanBegin               *time.Time      `gorm:"column:plan_begin;type:timestamp;comment:计划开工时间"`
	BeginImpl               *time.Time      `gorm:"column:plan_begin;type:timestamp;comment:开工时间"`
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
	Status                  int             `gorm:"column:status;type:integer;;not null;comment:项目状态 0:未开工,1:开工待审核,2:已开工;3:竣工待审核;4:已竣工"`
	FinishTime              *time.Time      `gorm:"column:finish_time;type:timestamp;comment:竣工时间"`
}

func (ImplementGov) TableName() string {
	return tables.ImplementGov
}

func (b *ImplementGov) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	b.CreateAt = now
	b.UpdateAt = now
	return nil
}

func (b *ImplementGov) BeforeUpdate(tx *gorm.DB) error {
	b.UpdateAt = time.Now()
	return nil
}