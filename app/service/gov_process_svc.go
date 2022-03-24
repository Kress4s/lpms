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
	Get(id int64, year, month int) (*vo.GovProgressResp, exception.Exception)
	Update(openID string, id int64, param *vo.GovProgressUpdateReq) exception.Exception
	ListPlan(projectID int64, year int) ([]vo.ListGovProgressPlan, exception.Exception)
	ListGovProgressCompare(projectID int64, year int) ([]vo.GovProgressCompare, exception.Exception)
}

func (gsi *govProgressServiceImpl) Create(openID string, param *vo.GovProgressReq) exception.Exception {
	govProgress := param.ToModel(openID)
	return gsi.repo.Create(gsi.db, govProgress)
}

func (gsi *govProgressServiceImpl) Get(id int64, year, month int) (*vo.GovProgressResp, exception.Exception) {
	govProgress, ex := gsi.repo.Get(gsi.db, id, year, month)
	if ex != nil {
		return nil, ex
	}
	resp := &vo.GovProgressResp{}
	var err error
	if govProgress == nil {
		return resp, nil
	}
	info, ex := gsi.repo.ListInvested(gsi.db, id, year)
	if ex != nil {
		return nil, ex
	}
	investd := float64(0)
	fixInvested := float64(0)
	for i := range info {
		if info[i].PlanInvested != nil {
			investd += *info[i].PlanInvested
		}
		if info[i].LastMonthFixedInvested != nil {
			fixInvested += *info[i].LastMonthFixedInvested
		}
	}
	govProgress.YearSumFixedInvested = &fixInvested
	govProgress.YearSumInvested = &investd
	resp, err = vo.NewGovProgressResponse(govProgress)
	if err != nil {
		return nil, exception.Wrap(response.ExceptionUnmarshalJSON, err)
	}
	return resp, nil
}

func (gsi *govProgressServiceImpl) Update(openID string, id int64, param *vo.GovProgressUpdateReq) exception.Exception {
	return gsi.repo.Update(gsi.db, id, param.ToMap(openID))
}

func (gsi *govProgressServiceImpl) ListPlan(projectID int64, year int) ([]vo.ListGovProgressPlan, exception.Exception) {
	res, ex := gsi.repo.ListProgressPlan(gsi.db, projectID, year)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.ListGovProgressPlan, 0, len(res))
	for i := range res {
		resp = append(resp, vo.ListGovProgressPlan{
			ID:           res[i].ID,
			Year:         res[i].Year,
			Month:        res[i].Month,
			PlanInvest:   res[i].PlanInvest,
			PlanProgress: res[i].PlanProgress,
		})
	}
	return resp, nil
}

func (gsi *govProgressServiceImpl) ListGovProgressCompare(projectID int64, year int) ([]vo.GovProgressCompare, exception.Exception) {
	res, ex := gsi.repo.ListGovProgressCompare(gsi.db, projectID, year)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.GovProgressCompare, 0, len(res))
	for i := range res {
		compare := float64(0)
		if res[i].PlanInvested != nil && res[i].PlanInvest != nil {
			compare = *res[i].PlanInvested / *res[i].PlanInvest
		}
		resp = append(resp, vo.GovProgressCompare{
			Month:          res[i].Month,
			PlanInvest:     res[i].PlanInvest,
			PlanProgress:   res[i].PlanProgress,
			PlanInvested:   res[i].PlanInvested,
			ActualProgress: res[i].ActualProgress,
			Completeness:   compare,
		})
	}
	return resp, nil
}
