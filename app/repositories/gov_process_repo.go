package repositories

import (
	"lpms/app/models"
	"lpms/app/models/tables"
	"lpms/app/response"
	"lpms/exception"
	"sync"

	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

var (
	govProgressRepoInstance GovProgressRepo
	govProgressOnce         sync.Once
)

type GovProgressRepoImpl struct{}

func GetGovProgressRepo() GovProgressRepo {
	govProgressOnce.Do(func() {
		govProgressRepoInstance = &GovProgressRepoImpl{}
	})
	return govProgressRepoInstance
}

type GovProgressRepo interface {
	Create(db *gorm.DB, impl []models.GovProgress) exception.Exception
	ListProgressPlan(db *gorm.DB, projectID int64, year int) ([]models.ListGovProgressPlan, exception.Exception)
	Get(db *gorm.DB, id int64, year, month int) (*models.GovProgress, exception.Exception)
	Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception
	ListGovProgressCompare(db *gorm.DB, projectID int64, year int) ([]models.GovProgressCompare, exception.Exception)
	DeleteByProjectID(db *gorm.DB, projectID ...int64) exception.Exception
	ListInvested(db *gorm.DB, projectID int64, year int) ([]models.GovProgress, exception.Exception)
	FormNowInvested(db *gorm.DB, projectID int64, year, month int) (float64, exception.Exception)
	StartFormNowInvestedAndFixed(db *gorm.DB, projectID int64, startYear, startMonth, year, month int) (float64, float64, exception.Exception)
}

func (grr *GovProgressRepoImpl) Create(db *gorm.DB, govProgress []models.GovProgress) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"plan_invest", "plan_progress"}),
	}).Create(&govProgress).Error)
}

func (grr *GovProgressRepoImpl) ListProgressPlan(db *gorm.DB, projectID int64, year int) ([]models.ListGovProgressPlan, exception.Exception) {
	lpg := make([]models.ListGovProgressPlan, 0)
	tx := db.Table(tables.GovProgress).Where("project_id = ?", projectID).Where("year = ?", year).Find(&lpg)
	return lpg, exception.Wrap(response.ExceptionDatabase, tx.Error)
}

func (grr *GovProgressRepoImpl) ListInvested(db *gorm.DB, projectID int64, year int) ([]models.GovProgress, exception.Exception) {
	lpg := make([]models.GovProgress, 0)
	tx := db.Table(tables.GovProgress).Select("plan_invested, last_month_fixed_invested").Where("project_id = ?", projectID).Where("year = ?", year).Find(&lpg)
	return lpg, exception.Wrap(response.ExceptionDatabase, tx.Error)
}

// 开工至今累计投资额 和 开工至今累计固投
func (grr *GovProgressRepoImpl) StartFormNowInvestedAndFixed(db *gorm.DB, projectID int64, startYear, startMonth, year, month int) (float64, float64, exception.Exception) {
	lpg := make([]models.GovProgress, 0)
	tx := db.Table(tables.GovProgress).Select("plan_invested, last_month_fixed_invested").Where("project_id = ?", projectID).
		Where("year <= ? and month <= ?", year, month).Where("year >= ? and month >= ?", startYear, startMonth).Find(&lpg)
	if tx.Error != nil {
		return 0, 0, exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	invested := float64(0)
	fixedSum := float64(0)
	for i := range lpg {
		if lpg[i].PlanInvested != nil {
			invested += *lpg[i].PlanInvested
		}
		if lpg[i].LastMonthFixedInvested != nil {
			fixedSum += *lpg[i].LastMonthFixedInvested
		}
	}
	return invested, fixedSum, nil
}

// 一月至本月计划累计完成投资额
func (grr *GovProgressRepoImpl) FormNowInvested(db *gorm.DB, projectID int64, year, month int) (float64, exception.Exception) {
	lpg := make([]models.GovProgress, 0)
	tx := db.Table(tables.GovProgress).Select("plan_invested").Where("project_id = ?", projectID).Where("year = ? and month <= ?", year, month).Find(&lpg)
	if tx.Error != nil {
		return 0, exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	total := float64(0)
	for i := range lpg {
		if lpg[i].PlanInvested != nil {
			total += *lpg[i].PlanInvested
		}
	}
	return total, exception.Wrap(response.ExceptionDatabase, tx.Error)
}

func (grr *GovProgressRepoImpl) Get(db *gorm.DB, id int64, year, month int) (*models.GovProgress, exception.Exception) {
	govProgress := models.GovProgress{}
	res := db.Where(&models.GovProgress{ProjectID: id, Month: month, Year: year}).Find(&govProgress)
	if res.RowsAffected == 0 {
		return nil, exception.New(response.ExceptionRecordNotFound, "recode not found")
	}
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return &govProgress, nil
}

func (rri *GovProgressRepoImpl) Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.GovProgress{}).Where(&models.GovProgress{ID: id}).Updates(param).Error)
}

func (rri *GovProgressRepoImpl) ListPlan(db *gorm.DB, projectID int64) ([]models.ListGovProgressPlan, exception.Exception) {
	res := make([]models.ListGovProgressPlan, 0)
	// tx := db.Table(tables.GovProgress).Select("id, plan_invest, plan_progress, month").Where("project_id = ?", projectID).Find(&res)
	return res, exception.Wrap(response.ExceptionDatabase, db.Table(tables.GovProgress).Select("id, plan_invest, plan_progress, month").Where("project_id = ?", projectID).Find(&res).Error)
}

func (grr *GovProgressRepoImpl) ListGovProgressCompare(db *gorm.DB, projectID int64, year int) ([]models.GovProgressCompare, exception.Exception) {
	lpg := make([]models.GovProgressCompare, 0)
	tx := db.Table(tables.GovProgress).Where("project_id = ?", projectID).Where("year = ?", year).Select("year, month, plan_invest, plan_progress, plan_invested, actual_progress").Find(&lpg)
	return lpg, exception.Wrap(response.ExceptionDatabase, tx.Error)
}

func (grr *GovProgressRepoImpl) DeleteByProjectID(db *gorm.DB, projectID ...int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Where("project_id in (?)", projectID).Delete(&models.GovProgress{}).Error)
}
