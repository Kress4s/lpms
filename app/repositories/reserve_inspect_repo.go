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
	reserveInspectRepoInstance ReserveInspectRepo
	reserveInspectOnce         sync.Once
)

type ReserveInspectRepoImpl struct{}

func GetReserveInspectRepo() ReserveInspectRepo {
	reserveInspectOnce.Do(func() {
		reserveInspectRepoInstance = &ReserveInspectRepoImpl{}
	})
	return reserveInspectRepoInstance
}

type ReserveInspectRepo interface {
	EarlyPlanList(db *gorm.DB, pageInfo *vo.PageInfo, params *vo.ReserveInspectParam) (int64, []models.ReservePro,
		exception.Exception)
	Pass(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception
	Refuse(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception
	OutStorageInspList(db *gorm.DB, pageInfo *vo.PageInfo, params *vo.ReserveInspectParam) (int64,
		[]models.ReservePro, exception.Exception)
}

func (rir *ReserveInspectRepoImpl) EarlyPlanList(db *gorm.DB, pageInfo *vo.PageInfo, params *vo.ReserveInspectParam) (int64,
	[]models.ReservePro, exception.Exception) {
	data := make([]models.ReservePro, 0)
	tx := db.Table(tables.Reserve).Select("id, name, level, project_type, construct_subject, create_at, status").
		Where("status = ?", constant.EarlyPlan)
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
		tx = tx.Where("create_at <= ? and create_at >= ?", params.PlanEnd, params.PlanBegin)
	}
	count := int64(0)
	tx = tx.Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).
		Scan(&data).Limit(-1).Offset(-1).Count(&count)
	return count, data, exception.Wrap(response.ExceptionDatabase, tx.Error)
}

func (rir *ReserveInspectRepoImpl) OutStorageInspList(db *gorm.DB, pageInfo *vo.PageInfo,
	params *vo.ReserveInspectParam) (int64, []models.ReservePro, exception.Exception) {
	data := make([]models.ReservePro, 0)
	tx := db.Table(tables.Reserve).Select("id, name, level, project_type, construct_subject, create_at, status").
		Where("status = ?", constant.OutStorageInspect)
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
		tx = tx.Where("create_at <= ? and create_at >= ?", params.PlanEnd, params.PlanBegin)
	}
	count := int64(0)
	tx = tx.Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).
		Scan(&data).Limit(-1).Offset(-1).Count(&count)
	return count, data, exception.Wrap(response.ExceptionDatabase, tx.Error)
}

func (rir *ReserveInspectRepoImpl) Pass(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.ReservePro{}).Where(&models.ReservePro{ID: id}).Updates(param).Error)
}

func (rir *ReserveInspectRepoImpl) Refuse(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.ReservePro{}).Where(&models.ReservePro{ID: id}).Updates(param).Error)
}
