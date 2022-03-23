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
	List(db *gorm.DB, pageInfo *vo.PageInfo, params *vo.ImplementGovFilterParam) (int64, []models.ImplementGov,
		exception.Exception)
	Delete(db *gorm.DB, id int64) exception.Exception
	MultiDelete(db *gorm.DB, ids []int64) exception.Exception
}

func (igi *ImplementGovRepoImpl) Create(db *gorm.DB, impl *models.ImplementGov) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(impl).Error)
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

func (igi *ImplementGovRepoImpl) List(db *gorm.DB, pageInfo *vo.PageInfo, params *vo.ImplementGovFilterParam) (int64,
	[]models.ImplementGov, exception.Exception) {
	data := make([]models.ImplementGov, 0)
	tx := db.Table(tables.ImplementGov).Select("id, name, level, project_type, construct_subject, create_at, status").Where("status <> ? and status <> ?", constant.StartInspecting, constant.FinishInspect)
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
	if params.PlanBegin != "" && params.PlanEnd != "" {
		tx = tx.Where("create_at <= ? and create_at >= ", params.PlanEnd, params.PlanBegin)
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

func (igi *ImplementGovRepoImpl) Delete(db *gorm.DB, id int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.ImplementGov{}, id).Error)
}

func (igi *ImplementGovRepoImpl) MultiDelete(db *gorm.DB, ids []int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.ImplementGov{}, ids).Error)
}
