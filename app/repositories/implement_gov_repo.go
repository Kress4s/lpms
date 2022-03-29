package repositories

import (
	"lpms/app/models"
	"lpms/app/models/tables"
	"lpms/app/response"
	"lpms/app/vo"
	"lpms/constant"
	"lpms/exception"
	"sync"
	"time"

	"gorm.io/gorm"
)

var (
	implementGovRepoInstance ImplementGovRepo
	implementGovOnce         sync.Once
)

type ImplementGovRepoImpl struct{}

func GetImplementGovRepo() ImplementGovRepo {
	implementGovOnce.Do(func() {
		implementGovRepoInstance = &ImplementGovRepoImpl{}
	})
	return implementGovRepoInstance
}

type ImplementGovRepo interface {
	Create(db *gorm.DB, impl *models.ImplementGov) exception.Exception
	Get(db *gorm.DB, id int64) (*models.ImplementGov, exception.Exception)
	List(db *gorm.DB, pageInfo *vo.PageInfo, params *vo.ImplementGovFilterParam, isAdmin bool, user string) (int64, []models.ImplementGov,
		exception.Exception)
	Delete(db *gorm.DB, id int64) exception.Exception
	MultiDelete(db *gorm.DB, ids []int64) exception.Exception
	ListStatusCount(db *gorm.DB, params *vo.ImplementGovCountFilter, isAdmin bool, user string) ([]ListCountModel, exception.Exception)
	ProgressLight(db *gorm.DB, projectID int64, year, month int) (int, exception.Exception)
}

func (igi *ImplementGovRepoImpl) Create(db *gorm.DB, impl *models.ImplementGov) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(impl).Error)
}

func (igi *ImplementGovRepoImpl) ProgressLight(db *gorm.DB, projectID int64, year, month int) (int, exception.Exception) {
	count := int64(0)
	tx := db.Table(tables.GovProgress).Where("project_id = ? and year = ? and month <= ? and status = ?", projectID, year, month, 1).
		Count(&count)
	if tx.Error != nil {
		return 0, exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	if int(count) == month {
		return constant.Green, nil
	}
	return constant.Red, nil
}

func (igi *ImplementGovRepoImpl) Get(db *gorm.DB, id int64) (*models.ImplementGov, exception.Exception) {
	reserve := models.ImplementGov{}
	res := db.Where(&models.ImplementGov{ID: id}).Find(&reserve)
	if res.RowsAffected == 0 {
		return nil, exception.New(response.ExceptionRecordNotFound, "recode not found")
	}
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return &reserve, nil
}

func (igi *ImplementGovRepoImpl) List(db *gorm.DB, pageInfo *vo.PageInfo, params *vo.ImplementGovFilterParam, isAdmin bool, user string) (int64,
	[]models.ImplementGov, exception.Exception) {
	data := make([]models.ImplementGov, 0)
	tx := db.Table(tables.ImplementGov).Where("status <> ? and status <> ?", constant.StartInspecting, constant.FinishInspect)
	if !isAdmin {
		tx = tx.Where("create_by = ?", user)
	}
	if params.Name != "" {
		tx = tx.Where("name = ?", params.Name)
	}
	if params.Level != nil {
		tx = tx.Where("level = ?", params.Level)
	}
	if params.ProjectType != nil {
		tx = tx.Where("project_type = ?", params.ProjectType)
	}
	if params.ConstructSubject != "" {
		tx = tx.Where("construct_subject = ?", params.ConstructSubject)
	}
	if params.PointType != nil {
		tx = tx.Where("point_type = ?", params.PointType)
	}
	if params.PlanBegin != "" && params.PlanEnd != "" {
		tx = tx.Where("plan_begin <= ? and plan_begin >= ?", params.PlanEnd, params.PlanBegin)
	}
	if params.StartTime != "" && params.EndTime != "" {
		tx = tx.Where("start_time <= ? and start_time >= ?", params.EndTime, params.StartTime)
	}
	if params.CurYearBegin != "" && params.CurYearEnd != "" {
		if params.Status != nil {
			if *params.Status == -1 {
				tx = tx.Where("create_at < ? and create_at >= ?", params.CurYearEnd, params.CurYearBegin)
			} else if *params.Status == 4 {
				tx = tx.Where("finish_time < ? and finish_time >= ?", params.CurYearEnd, params.CurYearBegin)
			}
		}
	}
	if params.Status != nil {
		if *params.Status != -1 {
			tx = tx.Where("status = ?", params.Status)
		}
	}
	if params.BeginInvest != nil && params.EndInvest != nil {
		tx = tx.Where("total_investment <= ? and total_investment >= ?", params.EndInvest, params.BeginInvest)
	}
	if params.DutyUnit != "" {
		tx = tx.Where("duty_unit = ?", params.DutyUnit)
	}
	if params.Type != nil {
		tx = tx.Where("type = ?", params.Type)
	}
	count := int64(0)
	tx = tx.Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).Order("type ASC").Order("project_code ASC").
		Scan(&data).Limit(-1).Offset(-1).Count(&count)
	return count, data, exception.Wrap(response.ExceptionDatabase, tx.Error)
}

func (igi *ImplementGovRepoImpl) Delete(db *gorm.DB, id int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.ImplementGov{}, id).Error)
}

func (igi *ImplementGovRepoImpl) MultiDelete(db *gorm.DB, ids []int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.ImplementGov{}, ids).Error)
}

type ListCountModel struct {
	Status int   `json:"status"`
	Count  int64 `json:"count"`
}

// 统计数量 未开工、开工建设、 竣工
func (igi *ImplementGovRepoImpl) ListStatusCount(db *gorm.DB, params *vo.ImplementGovCountFilter, isAdmin bool, user string) ([]ListCountModel, exception.Exception) {
	res := make([]ListCountModel, 0)
	subTx := db.Table(tables.ImplementGov).Select("status, create_at").
		Where("status in (?, ?, ?)", constant.UnStart, constant.Started, constant.Finished)
	if !isAdmin {
		subTx = subTx.Where("create_by = ?", user)
	}
	if params.Name != "" {
		subTx = subTx.Where("name = ?", params.Name)
	}
	if params.Level != nil {
		subTx = subTx.Where("level = ?", params.Level)
	}
	if params.ProjectType != nil {
		subTx = subTx.Where("project_type = ?", params.ProjectType)
	}
	if params.ConstructSubject != "" {
		subTx = subTx.Where("construct_subject = ?", params.ConstructSubject)
	}
	if params.PointType != nil {
		subTx = subTx.Where("point_type = ?", params.PointType)
	}
	if params.PlanBegin != "" && params.PlanEnd != "" {
		subTx = subTx.Where("plan_begin <= ? and plan_begin >= ?", params.PlanEnd, params.PlanBegin)
	}
	if params.StartTime != "" && params.EndTime != "" {
		subTx = subTx.Where("start_time <= ? and start_time >= ?", params.PlanEnd, params.PlanBegin)
	}
	if params.BeginInvest != nil && params.EndInvest != nil {
		subTx = subTx.Where("total_investment <= ? and total_investment >= ?", params.EndInvest, params.BeginInvest)
	}
	if params.DutyUnit != "" {
		subTx = subTx.Where("duty_unit = ?", params.DutyUnit)
	}
	if params.Type != nil {
		subTx = subTx.Where("type = ?", params.Type)
	}
	tx := db.Table("(?) AS sub", subTx).Select("sub.status AS status, count(*) AS count").Group("sub.status").Find(&res)
	if tx.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, tx.Error)
	}

	now := time.Now()
	CurYearBegin := time.Date(now.Year(), 1, 0, 0, 0, 0, 0, time.Local).AddDate(0, 0, 1)
	CurYearEnd := CurYearBegin.AddDate(1, 0, 0)

	count := int64(0)
	tx1 := db.Table("(?) AS sub", subTx).Select("*").Where("sub.create_at < ? and sub.create_at >= ?", CurYearEnd, CurYearBegin).Limit(-1).Offset(-1).Count(&count)
	if tx1.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, tx1.Error)
	}
	res = append(res, ListCountModel{
		Status: -2,
		Count:  count,
	})

	subTx = subTx.Where("status = ?", constant.Finished).Where("finish_time < ? and finish_time >= ?", CurYearEnd, CurYearBegin)
	count1 := int64(0)
	tx2 := db.Table("(?) AS sub", subTx).Limit(-1).Offset(-1).Count(&count1)
	if tx2.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, tx2.Error)
	}
	res = append(res, ListCountModel{
		Status: -1,
		Count:  count1,
	})
	return res, nil
}
