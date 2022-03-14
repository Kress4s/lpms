package service

import (
	"lpms/app/repositories"
	"lpms/app/response"
	"lpms/app/vo"
	"lpms/commom/drivers/database"
	"lpms/commom/tools"
	"lpms/constant"
	"lpms/exception"
	"sync"

	"gorm.io/gorm"
)

var (
	loginInstance LoginService
	loginOnce     sync.Once
)

type loginServiceImpl struct {
	db   *gorm.DB
	repo repositories.UserRepo
}

func GetLoginService() LoginService {
	loginOnce.Do(func() {
		loginInstance = &loginServiceImpl{
			db:   database.GetDriver(),
			repo: repositories.GetUserRepo(),
		}
	})
	return loginInstance
}

type LoginService interface {
	Login(username, password string) (*vo.LoginResponse, exception.Exception)
}

func (ls *loginServiceImpl) Login(username, password string) (*vo.LoginResponse, exception.Exception) {
	password = string(tools.Base64Encode([]byte(password)))
	ok, status, userID, ex := ls.repo.CheckPassword(ls.db, username, password)
	if ex != nil || !ok {
		return nil, ex
	}
	if !status {
		return nil, exception.New(response.ExceptionUserClose, "对不起 您的账号已被冻结")
	}
	// token
	token, exp := tools.Token(userID, username)
	return &vo.LoginResponse{
		AccessToken: token,
		TokenType:   constant.Authorization,
		Expiry:      exp,
	}, nil
}
