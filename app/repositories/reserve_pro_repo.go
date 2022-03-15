package repositories

import (
	"lpms/app/models"
	"lpms/app/models/tables"
	"lpms/app/response"
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
	GetInvestDetail(db *gorm.DB, id int64) ([]models.InvestDetail, exception.Exception)
	Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception
	Delete(db *gorm.DB, id int64) exception.Exception
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
