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

type ImplementGovHandler struct {
	handlers.BaseHandler
	Svc service.ImplementGovService
}

func NewImplementGovHandler() *ImplementGovHandler {
	return &ImplementGovHandler{
		Svc: service.GetImplementGovService(),
	}
}

// Create godoc
// @Summary 创建实施库政府投资项目
// @Description 创建实施库政府投资项目
// @Tags 实施库 - 政府投资项目
// @Param parameters body vo.ImplementGovReq true "ImplementGovReq"
// @Success 200  "创建实施库政府投资项目成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/implement/gov/project [post]
func (ih *ImplementGovHandler) Create(ctx iris.Context) mvc.Result {
	req := &vo.ImplementGovReq{}
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
// @Summary 查看实施库政府投资项目
// @Description 查看实施库政府投资项目
// @Tags 实施库 - 政府投资项目
// @Param id path string true "项目id"
// @Success 200 {object} vo.ImplementGovResp "查询实施库政府投资项目成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/implement/gov/project/{id} [get]
func (ih *ImplementGovHandler) Get(ctx iris.Context) mvc.Result {
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
// @Summary 获取实施库政府投资项目列表
// @Description 获取实施库政府投资项目列表
// @Tags 实施库 - 政府投资项目
// @Param page query int false "请求页"
// @Param page_size query int false "页大小"
// @Param parameters body vo.ImplementGovFilterParam true "ImplementGovFilterParam"
// @Success 200 {object} vo.DataPagination{data=[]vo.ListImplementGovResp} "查询实施库政府投资项目列表成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/implement/gov/projects [post]
func (ih *ImplementGovHandler) List(ctx iris.Context) mvc.Result {
	page, ex := handlers.GetPageInfo(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	params := &vo.ImplementGovFilterParam{}
	if err := ctx.ReadJSON(params); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	resp, ex := ih.Svc.List(ih.UserName, params, page)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Create godoc
// @Summary 删除实施库政府投资项目
// @Description 删除实施库政府投资项目
// @Tags 实施库 - 政府投资项目
// @Param id path string true "实施库政府投资项目id"
// @Success 200 "删除实施库政府投资项目成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/implement/gov/project/{id} [delete]
func (ih *ImplementGovHandler) Delete(ctx iris.Context) mvc.Result {
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
// @Summary 批量删除实施库政府投资项目
// @Description 批量删除实施库政府投资项目
// @Tags 实施库 - 政府投资项目
// @Param ids query string true "实施库政府投资项目id, `,` 连接"
// @Success 200 "批量删除实施库政府投资项目成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/implement/gov/project/multi [delete]
func (ih *ImplementGovHandler) MultiDelete(ctx iris.Context) mvc.Result {
	ids := ctx.URLParam(constant.IDS)
	ex := ih.Svc.MultiDelete(ids)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// BeforeActivation 初始化路由
func (ih *ImplementGovHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, "/gov/project", "Create")
	b.Handle(iris.MethodGet, "/gov/project/{id:string}", "Get")
	b.Handle(iris.MethodPost, "/gov/projects", "List")
	b.Handle(iris.MethodDelete, "/gov/project/{id:string}", "Delete")
	b.Handle(iris.MethodDelete, "/gov/project/multi", "MultiDelete")
}
