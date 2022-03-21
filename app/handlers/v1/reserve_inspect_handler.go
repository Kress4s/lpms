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

type ReserveInspectHandler struct {
	handlers.BaseHandler
	Svc service.ReserveInspectService
}

func NewReserveInspectHandler() *ReserveInspectHandler {
	return &ReserveInspectHandler{
		Svc: service.GetReserveInspectService(),
	}
}

// Create godoc
// @Summary 获取储备库前期计划项目列表
// @Description 获取储备库前期计划项目列表
// @Tags 审批中心 - 项目审核 - 储备库审核
// @Param page query int false "请求页"
// @Param page_size query int false "页大小"
// @Param parameters body vo.ReserveInspectParam true "ReserveInspectParam"
// @Success 200 {object} vo.DataPagination{data=[]vo.ListReserveProResp} "查询储备库前期计划项目列表成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/inspect/reserve/early-plan/list [post]
func (rh *ReserveInspectHandler) EarlyPlanList(ctx iris.Context) mvc.Result {
	page, ex := handlers.GetPageInfo(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	params := &vo.ReserveInspectParam{}
	if err := ctx.ReadJSON(params); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	resp, ex := rh.Svc.EarlyPlanList(params, page)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Create godoc
// @Summary 获取储备库出库审核项目列表
// @Description 获取储备库出库审核项目列表
// @Tags 审批中心 - 项目审核 - 储备库审核
// @Param page query int false "请求页"
// @Param page_size query int false "页大小"
// @Param parameters body vo.ReserveInspectParam true "ReserveInspectParam"
// @Success 200 {object} vo.DataPagination{data=[]vo.ListReserveProResp} "查询储备库出库审核项目列表成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/inspect/reserve/out-storage/list [post]
func (rh *ReserveInspectHandler) OutStorageInspList(ctx iris.Context) mvc.Result {
	page, ex := handlers.GetPageInfo(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	params := &vo.ReserveInspectParam{}
	if err := ctx.ReadJSON(params); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	resp, ex := rh.Svc.OutStorageInspList(params, page)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Create godoc
// @Summary 前期计划-审核通过(发文)
// @Description 前期计划-审核通过(发文)
// @Tags 审批中心 - 项目审核 - 储备库审核
// @Param id path string true "储备库项目id"
// @Success 200  "前期计划-审核通过成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/inspect/reserve/{id}/early-plan/pass [put]
func (rh *ReserveInspectHandler) EarlyPlanPass(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	if ex := rh.Svc.EarlyPlanPass(rh.UserName, id); ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 出库-审核通过
// @Description 出库-审核通过
// @Tags 审批中心 - 项目审核 - 储备库审核
// @Param id path string true "储备库项目id"
// @Success 200  "出库-审核通过成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/inspect/reserve/{id}/out-storage/pass [put]
func (rh *ReserveInspectHandler) OutStoragePass(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	if ex := rh.Svc.OutStoragePass(rh.UserName, id); ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 前期计划/出库 - 驳回
// @Description 前期计划/出库 - 驳回
// @Tags 审批中心 - 项目审核 - 储备库审核
// @Param id path string true "储备库项目id"
// @Success 200  "前期计划/出库 - 驳回 成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/inspect/reserve/{id}/refuse [put]
func (rh *ReserveInspectHandler) Refuse(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	if ex := rh.Svc.Refuse(rh.UserName, id); ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// BeforeActivation 初始化路由
func (rh *ReserveInspectHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, "/reserve/early-plan/list", "EarlyPlanList")
	b.Handle(iris.MethodPost, "/reserve/out-storage/list", "OutStorageInspList")
	b.Handle(iris.MethodPut, "/reserve/{id:string}/early-plan/pass", "EarlyPlanPass")
	b.Handle(iris.MethodPut, "/reserve/{id:string}/out-storage/pass", "OutStoragePass")
	b.Handle(iris.MethodPut, "/reserve/{id:string}/refuse", "Refuse")
}
