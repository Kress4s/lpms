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
// @Success 201  "创建储备库项目成功"
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

// Create godoc
// @Summary 获取储备库项目
// @Description 获取储备库项目
// @Tags 储备库 - 项目
// @Param id path string true "储备库项目id"
// @Success 200 {object} vo.ReserveResp "查询储备库项目成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/reserve/project/{id} [get]
func (rh *ReserveHandler) Get(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	resp, ex := rh.Svc.Get(id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Create godoc
// @Summary 获取储备库项目列表
// @Description 获取储备库项目列表
// @Tags 储备库 - 项目
// @Param page query int false "请求页"
// @Param page_size query int false "页大小"
// @Param parameters body vo.ReserveFilterParam true "ReserveFilterParam"
// @Success 200 {object} vo.DataPagination{data=[]vo.ListReserveProResp} "查询储备库项目列表成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/reserve/projects [post]
func (rh *ReserveHandler) List(ctx iris.Context) mvc.Result {
	page, ex := handlers.GetPageInfo(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	params := &vo.ReserveFilterParam{}
	if err := ctx.ReadJSON(params); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	resp, ex := rh.Svc.List(params, page)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Create godoc
// @Summary 修改储备库项目
// @Description 修改储备库项目
// @Tags 储备库 - 项目
// @Param id path string true "储备库项目id"
// @Param parameters body vo.ReserveUpdateReq true "ReserveUpdateReq"
// @Success 200  "修改储备库项目成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/reserve/project/{id} [put]
func (rh *ReserveHandler) Update(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	param := &vo.ReserveUpdateReq{}
	if err := ctx.ReadJSON(param); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	if ex := rh.Svc.Update(rh.UserName, id, param); ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 删除储备库项目
// @Description 删除储备库项目
// @Tags 储备库 - 项目
// @Param id path string true "储备库项目id"
// @Success 200 "删除储备库项目成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/reserve/project/{id} [delete]
func (rh *ReserveHandler) Delete(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	ex := rh.Svc.Delete(id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 批量删除储备库项目
// @Description 批量删除储备库项目
// @Tags 储备库 - 项目
// @Param ids query string true "储备库项目id, `,` 连接"
// @Success 200 "批量删除储备库项目成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/reserve/project/multi [delete]
func (rh *ReserveHandler) MultiDelete(ctx iris.Context) mvc.Result {
	ids := ctx.URLParam(constant.IDS)
	ex := rh.Svc.MultiDelete(ids)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 提交储备库项目
// @Description 提交储备库项目
// @Tags 储备库 - 项目
// @Param id path string true "储备库项目id"
// @Success 200 "提交储备库项目成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/reserve/project/{id}/refer [patch]
func (rh *ReserveHandler) Refer(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	if ex := rh.Svc.Refer(rh.UserName, id); ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 提报储备库项目
// @Description 提报储备库项目
// @Tags 储备库 - 项目
// @Param id path string true "储备库项目id"
// @Success 200 "提报储备库项目成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/reserve/project/{id}/submit [patch]
func (rh *ReserveHandler) Submission(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	if ex := rh.Svc.Submission(rh.UserName, id); ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 批量提报储备库项目
// @Description 批量提报储备库项目
// @Tags 储备库 - 项目
// @Param ids query string true "储备库项目id, `,` 连接"
// @Success 200 "批量提报储备库项目成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/reserve/project/submit/multi [patch]
func (rh *ReserveHandler) MultiSubmission(ctx iris.Context) mvc.Result {
	ids := ctx.URLParam(constant.IDS)
	if ex := rh.Svc.MultiSubmission(rh.UserName, ids); ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 出库储备库项目至审核
// @Description 出库储备库项目至审核
// @Tags 储备库 - 项目
// @Param id path string true "储备库项目id"
// @Success 200 "出库储备库项目成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/reserve/project/{id}/out-storage [patch]
func (rh *ReserveHandler) OutStorage(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	if ex := rh.Svc.OutStorage(rh.UserName, id); ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// BeforeActivation 初始化路由
func (rh *ReserveHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, "/project", "Create")
	b.Handle(iris.MethodGet, "/project/{id:string}", "Get")
	b.Handle(iris.MethodPost, "/projects", "List")
	b.Handle(iris.MethodDelete, "/project/{id:string}", "Delete")
	b.Handle(iris.MethodPut, "/project/{id:string}", "Update")
	b.Handle(iris.MethodDelete, "/project/multi", "MultiDelete")
	b.Handle(iris.MethodPatch, "/project/{id:string}/refer", "Refer")
	b.Handle(iris.MethodPatch, "/project/{id:string}/submit", "Submission")
	b.Handle(iris.MethodPatch, "/project/submit/multi", "MultiSubmission")
	b.Handle(iris.MethodPatch, "/project/{id:string}/out-storage", "OutStorage")
}
