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

type ImpleIndustryHandler struct {
	handlers.BaseHandler
	Svc service.ImpleIndustryService
}

func NewImpleIndustryHandler() *ImpleIndustryHandler {
	return &ImpleIndustryHandler{
		Svc: service.GetImpleIndustryService(),
	}
}

// Create godoc
// @Summary 创建实施库产业项目
// @Description 创建实施库产业项目
// @Tags 实施库 - 产业项目
// @Param parameters body vo.ImpleIndustryReq true "ImpleIndustryReq"
// @Success 201  "创建实施库产业项目成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/implement/indust/project [post]
func (ih *ImpleIndustryHandler) Create(ctx iris.Context) mvc.Result {
	req := &vo.ImpleIndustryReq{}
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
// @Summary 查看实施库产业项目
// @Description 查看实施库产业项目
// @Tags 实施库 - 产业项目
// @Param id path string true "项目id"
// @Success 200 {object} vo.ImpleIndustryResp "查询实施库产业项目成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/implement/indust/project/{id} [get]
func (ih *ImpleIndustryHandler) Get(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	resp, ex := ih.Svc.Get(id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Create godoc
// @Summary 获取实施库产业项目列表
// @Description 获取实施库产业项目列表
// @Tags 实施库 - 产业项目
// @Param page query int false "请求页"
// @Param page_size query int false "页大小"
// @Param parameters body vo.ImpleIndustryFilterParam true "ImpleIndustryFilterParam"
// @Success 200 {object} vo.DataPagination{data=[]vo.ListImpleIndustryResp} "查询实施库产业项目列表成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/implement/indust/projects [post]
func (ih *ImpleIndustryHandler) List(ctx iris.Context) mvc.Result {
	page, ex := handlers.GetPageInfo(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	params := &vo.ImpleIndustryFilterParam{}
	if err := ctx.ReadJSON(params); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	resp, ex := ih.Svc.List(params, page)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Create godoc
// @Summary 删除实施库产业项目
// @Description 删除实施库产业项目
// @Tags 实施库 - 产业项目
// @Param id path string true "实施库产业项目id"
// @Success 200 "删除实施库产业项目成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/implement/indust/project/{id} [delete]
func (ih *ImpleIndustryHandler) Delete(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	ex := ih.Svc.Delete(id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 批量删除实施库产业项目
// @Description 批量删除实施库产业项目
// @Tags 实施库 - 产业项目
// @Param ids query string true "实施库产业项目id, `,` 连接"
// @Success 200 "批量删除实施库产业项目成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/implement/indust/project/multi [delete]
func (ih *ImpleIndustryHandler) MultiDelete(ctx iris.Context) mvc.Result {
	ids := ctx.URLParam(constant.IDS)
	ex := ih.Svc.MultiDelete(ids)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// BeforeActivation 初始化路由
func (ih *ImpleIndustryHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, "/indust/project", "Create")
	b.Handle(iris.MethodGet, "/indust/project/{id:string}", "Get")
	b.Handle(iris.MethodPost, "/indust/projects", "List")
	b.Handle(iris.MethodDelete, "/indust/project/{id:string}", "Delete")
	b.Handle(iris.MethodDelete, "/indust/project/multi", "MultiDelete")
}
