package service

import (
	"lpms/app/repositories"
	"lpms/app/vo"
	"lpms/commom/drivers/database"
	"lpms/exception"
	"sync"

	"gorm.io/gorm"
)

var (
	reserveServiceInstance ReserveService
	reserveOnce            sync.Once
)

type reserveServiceImpl struct {
	db   *gorm.DB
	repo repositories.ReserveRepo
}

func GetReserveService() ReserveService {
	reserveOnce.Do(func() {
		reserveServiceInstance = &reserveServiceImpl{
			db:   database.GetDriver(),
			repo: repositories.GetReserveRepo(),
		}
	})
	return reserveServiceInstance
}

type ReserveService interface {
	Create(openID string, param *vo.ReserveReq) exception.Exception
	// Update(openID string, id int64, param *vo.BlackIPUpdateReq) exception.Exception
}

func (rsi *reserveServiceImpl) Create(openID string, param *vo.ReserveReq) exception.Exception {
	reserve := param.ToModel(openID)
	return rsi.repo.Create(rsi.db, reserve)
}
