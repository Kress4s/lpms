package repositories

import (
	"lpms/app/models"
	"lpms/app/models/tables"
	"lpms/app/response"
	"lpms/app/vo"
	"lpms/exception"
	"sync"

	"gorm.io/gorm"
)

var (
	reserveRepoInstance ReserveRepo
	reserveOnce         sync.Once
)

type ReserveRepoImpl struct{}

func GetReserveRepo() ReserveRepo {
	reserveOnce.Do(func() {
		reserveRepoInstance = &ReserveRepoImpl{}
	})
	return reserveRepoInstance
}

type ReserveRepo interface {
	Create(db *gorm.DB, reserve *models.ReservePro) exception.Exception
	Get(db *gorm.DB, id int64) (*models.ReservePro, exception.Exception)
	List(db *gorm.DB, pageInfo *vo.PageInfo, params *vo.ReserveFilterParam) (int64, []models.ReservePro, exception.Exception)
	GetInvestDetail(db *gorm.DB, id int64) ([]models.InvestDetail, exception.Exception)
	Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception
	Delete(db *gorm.DB, id int64) exception.Exception
	MultiDelete(db *gorm.DB, ids []int64) exception.Exception
	Refer(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception
	Submission(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception
	MultiSubmission(db *gorm.DB, ids []int64, param map[string]interface{}) exception.Exception
}

func (rri *ReserveRepoImpl) Create(db *gorm.DB, reserve *models.ReservePro) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(reserve).Error)
}

func (rri *ReserveRepoImpl) Get(db *gorm.DB, id int64) (*models.ReservePro, exception.Exception) {
	reserve := models.ReservePro{}
	res := db.Where(&models.ReservePro{ID: id}).Find(&reserve)
	if res.RowsAffected == 0 {
		return nil, exception.New(response.ExceptionRecordNotFound, "recode not found")
	}
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return &reserve, nil
}

func (rri *ReserveRepoImpl) List(db *gorm.DB, pageInfo *vo.PageInfo, params *vo.ReserveFilterParam) (int64, []models.ReservePro, exception.Exception) {
	data := make([]models.ReservePro, 0)
	tx := db.Table(tables.Reserve).Select("id, name, level, project_type, construct_subject, create_at, status")
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
		tx = tx.Where("plan_begin <= ? and plan_begin >= ", params.PlanEnd, params.PlanBegin)
	}
	if params.Status != nil {
		tx = tx.Where("status = ?", params.Status)
	}
	count := int64(0)
	tx = tx.Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).
		Scan(&data).Limit(-1).Offset(-1).Count(&count)
	return count, data, exception.Wrap(response.ExceptionDatabase, tx.Error)
}

func (rri *ReserveRepoImpl) GetInvestDetail(db *gorm.DB, id int64) ([]models.InvestDetail, exception.Exception) {
	info := make([]models.InvestDetail, 0)
	tx := db.Table(tables.Reserve).Select("json_array_elements (investment_detail::json) as info").Where("id = ?", id).Scan(&info)
	if tx.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return info, nil
}

func (rri *ReserveRepoImpl) Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.ReservePro{}).Where(&models.ReservePro{ID: id}).Updates(param).Error)
}

func (rri *ReserveRepoImpl) Delete(db *gorm.DB, id int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.ReservePro{}, id).Error)
}

func (rri *ReserveRepoImpl) MultiDelete(db *gorm.DB, ids []int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.ReservePro{}, ids).Error)
}

// 提交 : 0(草稿) -> 1(已入库)
func (rri *ReserveRepoImpl) Refer(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.ReservePro{}).Where(&models.ReservePro{ID: id}).Updates(param).Error)
}

// 提报 : 1(已入库) -> 2(前期计划）
func (rri *ReserveRepoImpl) Submission(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.ReservePro{}).Where(&models.ReservePro{ID: id}).Updates(param).Error)
}

// 批量提报:  1(已入库) -> 2(前期计划）
func (rri *ReserveRepoImpl) MultiSubmission(db *gorm.DB, ids []int64, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.ReservePro{}).Where("id in (?)", ids).Updates(param).Error)
}
