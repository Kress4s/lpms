package repositories

import (
	"lpms/app/models"
	"lpms/app/models/tables"
	"lpms/app/response"
	"lpms/app/vo"
	"lpms/constant"
	"lpms/exception"
	"sync"

	"gorm.io/gorm"
)

var (
	impleIndustryRepoInstance ImpleIndustryRepo
	impleIndustryOnce         sync.Once
)

type ImpleIndustryRepoImpl struct{}

func GetImpleIndustryRepo() ImpleIndustryRepo {
	impleIndustryOnce.Do(func() {
		impleIndustryRepoInstance = &ImpleIndustryRepoImpl{}
	})
	return impleIndustryRepoInstance
}

type ImpleIndustryRepo interface {
	Create(db *gorm.DB, impl *models.ImpleIndustry) exception.Exception
	Get(db *gorm.DB, id int64) (*models.ImpleIndustry, exception.Exception)
	List(db *gorm.DB, pageInfo *vo.PageInfo, params *vo.ImpleIndustryFilterParam, user string) (int64, []models.ImpleIndustry,
		exception.Exception)
	Delete(db *gorm.DB, id int64) exception.Exception
	MultiDelete(db *gorm.DB, ids []int64) exception.Exception
}

func (igi *ImpleIndustryRepoImpl) Create(db *gorm.DB, impl *models.ImpleIndustry) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(impl).Error)
}

func (igi *ImpleIndustryRepoImpl) Get(db *gorm.DB, id int64) (*models.ImpleIndustry, exception.Exception) {
	reserve := models.ImpleIndustry{}
	res := db.Where(&models.ImpleIndustry{ID: id}).Find(&reserve)
	if res.RowsAffected == 0 {
		return nil, exception.New(response.ExceptionRecordNotFound, "recode not found")
	}
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return &reserve, nil
}

func (igi *ImpleIndustryRepoImpl) List(db *gorm.DB, pageInfo *vo.PageInfo, params *vo.ImpleIndustryFilterParam, user string) (int64,
	[]models.ImpleIndustry, exception.Exception) {
	data := make([]models.ImpleIndustry, 0)
	tx := db.Table(tables.ImplementIndustry).Select("id, name, level, project_type, construct_subject, create_at, status, start_time, finish_time").
		Where("status <> ? and status <> ?", constant.StartInspecting, constant.FinishInspect)
	if user != "admin" {
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
		tx = tx.Where("create_at <= ? and create_at >= ?", params.PlanEnd, params.PlanBegin)
	}
	if params.CurYearBegin != "" && params.CurYearEnd != "" {
		tx = tx.Where("finish_time < ? and finish_time >= ?", params.PlanEnd, params.PlanBegin)
	}
	if params.Status != nil {
		tx = tx.Where("status = ?", params.Status)
	}
	count := int64(0)
	tx = tx.Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).
		Scan(&data).Limit(-1).Offset(-1).Count(&count)
	return count, data, exception.Wrap(response.ExceptionDatabase, tx.Error)
}

func (igi *ImpleIndustryRepoImpl) Delete(db *gorm.DB, id int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.ImpleIndustry{}, id).Error)
}

func (igi *ImpleIndustryRepoImpl) MultiDelete(db *gorm.DB, ids []int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.ImpleIndustry{}, ids).Error)
}
