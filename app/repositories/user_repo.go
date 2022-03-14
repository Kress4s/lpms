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
	CheckPassword(db *gorm.DB, account, password string) (bool, bool, int64, exception.Exception)
}

func (u *UserRepoImpl) CheckPassword(db *gorm.DB, username, password string) (bool, bool, int64, exception.Exception) {
	user := &models.User{}
	res := db.Where(&models.User{Username: username, Password: password}).Find(user)
	if res.Error != nil {
		return false, false, 0, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	if res.RowsAffected == 0 {
		return false, false, 0, exception.New(response.ExceptionInvalidUserPassword, "用户名/密码错误")
	}
	return true, user.Status, user.ID, nil
}
