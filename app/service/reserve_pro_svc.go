package service

import (
	"lpms/app/repositories"
	"lpms/app/response"
	"lpms/app/vo"
	"lpms/commom/drivers/database"
	"lpms/exception"
	"sync"

	"gorm.io/gorm"
)

var (
	reserveServiceInstance ReserveService
	reserveOnce            sync.Once
)

type reserveServiceImpl struct {
	db   *gorm.DB
	repo repositories.ReserveRepo
}

func GetReserveService() ReserveService {
	reserveOnce.Do(func() {
		reserveServiceInstance = &reserveServiceImpl{
			db:   database.GetDriver(),
			repo: repositories.GetReserveRepo(),
		}
	})
	return reserveServiceInstance
}

type ReserveService interface {
	Create(openID string, param *vo.ReserveReq) exception.Exception
	Get(id int64) (*vo.ReserveResp, exception.Exception)
	List(params *vo.ReserveFilterParam, pageInfo *vo.PageInfo) (*vo.DataPagination, exception.Exception)
	// Update(openID string, id int64, param *vo.BlackIPUpdateReq) exception.Exception
}

func (rsi *reserveServiceImpl) Create(openID string, param *vo.ReserveReq) exception.Exception {
	reserve := param.ToModel(openID)
	return rsi.repo.Create(rsi.db, reserve)
}

func (rsi *reserveServiceImpl) Get(id int64) (*vo.ReserveResp, exception.Exception) {
	reserve, ex := rsi.repo.Get(rsi.db, id)
	if ex != nil {
		return nil, ex
	}
	investInfo, ex := rsi.repo.GetInvestDetail(rsi.db, id)
	if ex != nil {
		return nil, ex
	}
	resp, err := vo.NewReserveProResponse(reserve, investInfo)
	if err != nil {
		return nil, exception.Wrap(response.ExceptionUnmarshalJSON, err)
	}

	return resp, nil
}

func (rsi *reserveServiceImpl) List(params *vo.ReserveFilterParam, pageInfo *vo.PageInfo) (*vo.DataPagination, exception.Exception) {
	count, projects, ex := rsi.repo.List(rsi.db, pageInfo, params)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.ListReserveProResp, 0, len(projects))
	for i := range projects {
		resp = append(resp, vo.ListReserveProResp{
			ID:               projects[i].ID,
			Name:             projects[i].Name,
			Level:            projects[i].Level,
			ProjectType:      projects[i].ProjectType,
			ConstructSubject: projects[i].ConstructSubject,
			CreateAt:         projects[i].CreateAt,
			Status:           projects[i].Status,
		})
	}
	return vo.NewDataPagination(count, resp, pageInfo), nil
}
