package v1

import (
	"lpms/app/handlers"
	"lpms/app/response"
	"lpms/app/service"
	"lpms/app/vo"
	"lpms/constant"
	"lpms/exception"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type GovProgressHandler struct {
	handlers.BaseHandler
	Svc service.GovProgressService
}

func NewGovProgressHandler() *GovProgressHandler {
	return &GovProgressHandler{
		Svc: service.GetGovProgressService(),
	}
}

// Create godoc
// @Summary 实施库政府投资项目 进度计划 保存(修改)
// @Description 实施库政府投资项目**进度计划**保存
// @Tags 实施库 - 政府投资项目 - 进度
// @Param parameters body vo.GovProgressReq true "GovProgressReq"
// @Success 201  "实施库政府投资项目进度保存成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/implement/gov/progress [post]
func (ih *GovProgressHandler) Create(ctx iris.Context) mvc.Result {
	req := &vo.GovProgressReq{}
	if err := ctx.ReadJSON(req); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := ih.Svc.Create(ih.UserName, req)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 查看实施库政府投资 项目进度
// @Description 查看实施库政府投资项目进度
// @Tags 实施库 - 政府投资项目 - 进度
// @Param project_id query string true "所属项目id"
// @Param month query string true "项目进度所属月份 eg: 1月 --> 1"
// @Success 200 {object} vo.GovProgressResp "查询实施库政府投资项目进度成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/implement/gov/progress [get]
func (ih *GovProgressHandler) Get(ctx iris.Context) mvc.Result {
	id, err := ctx.URLParamInt64(constant.ProjectID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	month, err := ctx.URLParamInt(constant.Month)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	resp, ex := ih.Svc.Get(id, month)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Create godoc
// @Summary 查看实施库政府投资项目 进度计划
// @Description 查看实施库政府投资项目 **进度计划**
// @Tags 实施库 - 政府投资项目 - 进度
// @Param project_id path string true "所属项目id"
// @Success 200 {array} vo.ListGovProgressPlan "查询实施库政府投资项目进度计划"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/implement/gov/progress/{project_id}/list [get]
func (ih *GovProgressHandler) ListPlan(ctx iris.Context) mvc.Result {
	project_id, err := ctx.Params().GetInt64(constant.ProjectID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	resp, ex := ih.Svc.ListPlan(project_id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Create godoc
// @Summary 修改项目进度
// @Description 修改项目进度
// @Tags 实施库 - 政府投资项目 - 进度
// @Param id path string true "项目进度记录id"
// @Param parameters body vo.GovProgressUpdateReq true "GovProgressUpdateReq"
// @Success 200  "修改项目进度成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/implement/gov/progress/{id} [put]
func (ih *GovProgressHandler) Update(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	param := &vo.GovProgressUpdateReq{}
	if err := ctx.ReadJSON(param); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	if ex := ih.Svc.Update(ih.UserName, id, param); ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// BeforeActivation 初始化路由
func (ih *GovProgressHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, "/gov/progress", "Create")
	b.Handle(iris.MethodGet, "/gov/progress", "Get")
	b.Handle(iris.MethodPut, "/gov/progress/{id:string}", "Update")
	b.Handle(iris.MethodGet, "/gov/progress/{project_id:string}/list", "ListPlan")
}
