package repositories

import (
	"lpms/app/models"
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
	Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception
	Delete(db *gorm.DB, id int64) exception.Exception
}

func (rri *ReserveRepoImpl) Create(db *gorm.DB, reserve *models.ReservePro) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(reserve).Error)
}

func (rri *ReserveRepoImpl) Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.ReservePro{}).Where(&models.ReservePro{ID: id}).Updates(param).Error)
}

func (rri *ReserveRepoImpl) Delete(db *gorm.DB, id int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.ReservePro{}, id).Error)
}
