package auth

import (
	"lpms/app/response"
	"lpms/app/service"
	"lpms/app/vo"
	"lpms/exception"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type LoginHandler struct {
	Svc service.LoginService
}

func NewLoginHandler() *LoginHandler {
	return &LoginHandler{
		Svc: service.GetLoginService(),
	}
}

// Create godoc
// @Summary 用户登录
// @Description 用户登录
// @Tags 登录
// @Param parameters body vo.LoginReq true "LoginReq"
// @Success 200 {object} vo.LoginResponse "响应成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Router /auth/login [post]
func (lh *LoginHandler) Login(ctx iris.Context) mvc.Result {
	user := &vo.LoginReq{}
	if err := ctx.ReadJSON(user); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	res, ex := lh.Svc.Login(user.UserName, user.Password)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(res)
}

// BeforeActivation 初始化路由
func (u *LoginHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, "/login", "Login")
}
