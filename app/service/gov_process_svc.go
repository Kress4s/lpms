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
	govProgressServiceInstance GovProgressService
	govProgressOnce            sync.Once
)

type govProgressServiceImpl struct {
	db   *gorm.DB
	repo repositories.GovProgressRepo
}

func GetGovProgressService() GovProgressService {
	govProgressOnce.Do(func() {
		govProgressServiceInstance = &govProgressServiceImpl{
			db:   database.GetDriver(),
			repo: repositories.GetGovProgressRepo(),
		}
	})
	return govProgressServiceInstance
}

type GovProgressService interface {
	Create(openID string, param *vo.GovProgressReq) exception.Exception
	Get(id int64, month int) (*vo.GovProgressResp, exception.Exception)
	Update(openID string, id int64, param *vo.GovProgressUpdateReq) exception.Exception
	ListPlan(projectID int64) ([]vo.ListGovProgressPlan, exception.Exception)
}

func (gsi *govProgressServiceImpl) Create(openID string, param *vo.GovProgressReq) exception.Exception {
	govProgress := param.ToModel(openID)
	return gsi.repo.Create(gsi.db, govProgress)
}

func (gsi *govProgressServiceImpl) Get(id int64, month int) (*vo.GovProgressResp, exception.Exception) {
	govProgress, ex := gsi.repo.Get(gsi.db, id, month)
	if ex != nil {
		return nil, ex
	}
	resp, err := vo.NewGovProgressResponse(govProgress)
	if err != nil {
		return nil, exception.Wrap(response.ExceptionUnmarshalJSON, err)
	}

	return resp, nil
}

func (gsi *govProgressServiceImpl) Update(openID string, id int64, param *vo.GovProgressUpdateReq) exception.Exception {
	return gsi.repo.Update(gsi.db, id, param.ToMap(openID))
}

func (gsi *govProgressServiceImpl) ListPlan(projectID int64) ([]vo.ListGovProgressPlan, exception.Exception) {
	res, ex := gsi.repo.ListProgressPlan(gsi.db, projectID)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.ListGovProgressPlan, 0, len(res))
	for i := range res {
		resp = append(resp, vo.ListGovProgressPlan{
			ID:           res[i].ID,
			Month:        res[i].Month,
			PlanInvest:   res[i].PlanInvest,
			PlanProgress: res[i].PlanProgress,
		})
	}
	return resp, nil
}
