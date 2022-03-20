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
	ListProgressPlan(db *gorm.DB, projectID int64) ([]models.ListGovProgressPlan, exception.Exception)
	Get(db *gorm.DB, id int64, month int) (*models.GovProgress, exception.Exception)
	Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception
	ListGovProgressCompare(db *gorm.DB, projectID int64) ([]models.GovProgressCompare, exception.Exception)
}

func (grr *GovProgressRepoImpl) Create(db *gorm.DB, govProgress []models.GovProgress) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"plan_invest", "plan_progress"}),
	}).Create(&govProgress).Error)
}

func (grr *GovProgressRepoImpl) ListProgressPlan(db *gorm.DB, projectID int64) ([]models.ListGovProgressPlan, exception.Exception) {
	lpg := make([]models.ListGovProgressPlan, 0)
	tx := db.Table(tables.GovProgress).Where("project_id = ?", projectID).Find(&lpg)
	return lpg, exception.Wrap(response.ExceptionDatabase, tx.Error)
}

func (grr *GovProgressRepoImpl) Get(db *gorm.DB, id int64, month int) (*models.GovProgress, exception.Exception) {
	govProgress := models.GovProgress{}
	res := db.Where(&models.GovProgress{ProjectID: id, Month: month}).Find(&govProgress)
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

func (grr *GovProgressRepoImpl) ListGovProgressCompare(db *gorm.DB, projectID int64) ([]models.GovProgressCompare, exception.Exception) {
	lpg := make([]models.GovProgressCompare, 0)
	tx := db.Table(tables.GovProgress).Where("project_id = ?", projectID).Select("month, plan_invest, plan_progress, plan_invested, actual_progress").Find(&lpg)
	return lpg, exception.Wrap(response.ExceptionDatabase, tx.Error)
}
