package repositories

import (
	"lpms/app/models"
	"lpms/app/response"
	"lpms/exception"
	"sync"

	"gorm.io/gorm"
)

var (
	windowRepoInstance WindowRepo
	windowOnce         sync.Once
)

type WindowRepoImpl struct{}

func GetWindowRepo() WindowRepo {
	windowOnce.Do(func() {
		windowRepoInstance = &WindowRepoImpl{}
	})
	return windowRepoInstance
}

type WindowRepo interface {
	Create(db *gorm.DB, window *models.WindowSetting) exception.Exception
	List(db *gorm.DB) ([]models.WindowSetting, exception.Exception)
	Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception
}

func (wri *WindowRepoImpl) Create(db *gorm.DB, window *models.WindowSetting) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(window).Error)
}

func (wri *WindowRepoImpl) List(db *gorm.DB) ([]models.WindowSetting, exception.Exception) {
	window := make([]models.WindowSetting, 0)
	return window, exception.Wrap(response.ExceptionDatabase, db.Model(&models.WindowSetting{}).Find(&window).Error)
}

func (wri *WindowRepoImpl) Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.WindowSetting{}).Where(&models.WindowSetting{ID: id}).Updates(param).Error)
}
