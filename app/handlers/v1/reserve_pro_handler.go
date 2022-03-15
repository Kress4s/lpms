package v1

import (
	"lpms/app/handlers"
	"lpms/app/response"
	"lpms/app/service"
	"lpms/app/vo"
	"lpms/exception"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ReserveHandler struct {
	handlers.BaseHandler
	Svc service.ReserveService
}

func NewReserveHandler() *ReserveHandler {
	return &ReserveHandler{
		Svc: service.GetReserveService(),
	}
}

// Create godoc
// @Summary 创建储备库项目
// @Description 创建储备库项目
// @Tags 储备库 - 项目
// @Param parameters body vo.ReserveReq true "ReserveReq"
// @Success 200  "创建储备库项目成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/reserve/project [post]
func (rh *ReserveHandler) Create(ctx iris.Context) mvc.Result {
	req := &vo.ReserveReq{}
	if err := ctx.ReadJSON(req); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := rh.Svc.Create(rh.UserName, req)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// BeforeActivation 初始化路由
func (rh *ReserveHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, "/project", "Create")
}
