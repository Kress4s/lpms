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
	windowServiceInstance WindowService
	windowOnce            sync.Once
)

type windowServiceImpl struct {
	db   *gorm.DB
	repo repositories.WindowRepo
}

func GetWindowService() WindowService {
	windowOnce.Do(func() {
		windowServiceInstance = &windowServiceImpl{
			db:   database.GetDriver(),
			repo: repositories.GetWindowRepo(),
		}
	})
	return windowServiceInstance
}

type WindowService interface {
	Create(openID string, param *vo.WindowsReq) exception.Exception
	List() (*vo.WindowsResponse, exception.Exception)
	Update(openID string, id int64, param *vo.WindowsUpdateReq) exception.Exception
}

func (wsi *windowServiceImpl) Create(openID string, param *vo.WindowsReq) exception.Exception {
	return wsi.repo.Create(wsi.db, param.ToModel(openID))
}

func (wsi *windowServiceImpl) List() (*vo.WindowsResponse, exception.Exception) {
	res, ex := wsi.repo.List(wsi.db)
	if ex != nil {
		return nil, ex
	}
	if len(res) == 0 {
		return nil, nil
	}
	return vo.NewWindowsResponse(&res[0]), nil
}

func (wsi *windowServiceImpl) Update(openID string, id int64, param *vo.WindowsUpdateReq) exception.Exception {
	return wsi.repo.Update(wsi.db, id, param.ToMap(openID))
}
