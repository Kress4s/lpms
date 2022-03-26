package repositories

import (
	"lpms/app/models"
	"lpms/app/response"
	"lpms/exception"
	"sync"

	"gorm.io/gorm"
)

var (
	userRepoInstance UserRepo
	userOnce         sync.Once
)

type UserRepoImpl struct{}

func GetUserRepo() UserRepo {
	userOnce.Do(func() {
		userRepoInstance = &UserRepoImpl{}
	})
	return userRepoInstance
}

type UserRepo interface {
	CheckPassword(db *gorm.DB, account, password string) (bool, bool, bool, int64, exception.Exception)
	Get(db *gorm.DB, username string) (*models.User, exception.Exception)
}

func (u *UserRepoImpl) CheckPassword(db *gorm.DB, username, password string) (bool, bool, bool, int64, exception.Exception) {
	user := &models.User{}
	res := db.Where(&models.User{Username: username, Password: password}).Find(user)
	if res.Error != nil {
		return false, false, false, 0, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	if res.RowsAffected == 0 {
		return false, false, false, 0, exception.New(response.ExceptionInvalidUserPassword, "用户名/密码错误")
	}
	return true, user.Status, user.IsAdmin, user.ID, nil
}

func (u *UserRepoImpl) Get(db *gorm.DB, username string) (*models.User, exception.Exception) {
	user := models.User{}
	res := db.Where(&models.User{Username: username}).Find(&user)
	if res.RowsAffected == 0 {
		return nil, exception.New(response.ExceptionRecordNotFound, "recode not found")
	}
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return &user, nil
}
