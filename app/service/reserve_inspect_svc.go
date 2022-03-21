package service

import (
	"lpms/app/repositories"
	"lpms/app/response"
	"lpms/app/vo"
	"lpms/commom/drivers/database"
	"lpms/constant"
	"lpms/exception"
	"sync"

	"gorm.io/gorm"
)

var (
	reserveInspectServiceInstance ReserveInspectService
	reserveInspectOnce            sync.Once
)

type reserveInspectServiceImpl struct {
	db          *gorm.DB
	repo        repositories.ReserveInspectRepo
	reserveRepo repositories.ReserveRepo
	GovRepo     repositories.ImplementGovRepo
}

func GetReserveInspectService() ReserveInspectService {
	reserveInspectOnce.Do(func() {
		reserveInspectServiceInstance = &reserveInspectServiceImpl{
			db:          database.GetDriver(),
			repo:        repositories.GetReserveInspectRepo(),
			reserveRepo: repositories.GetReserveRepo(),
			GovRepo:     repositories.GetImplementGovRepo(),
		}
	})
	return reserveInspectServiceInstance
}

type ReserveInspectService interface {
	EarlyPlanList(params *vo.ReserveInspectParam, pageInfo *vo.PageInfo) (*vo.DataPagination, exception.Exception)
	OutStorageInspList(params *vo.ReserveInspectParam, pageInfo *vo.PageInfo) (*vo.DataPagination, exception.Exception)
	EarlyPlanPass(openID string, id int64) exception.Exception
	OutStoragePass(openID string, id int64) exception.Exception
	Refuse(openID string, id int64) exception.Exception
}

func (ris *reserveInspectServiceImpl) EarlyPlanList(params *vo.ReserveInspectParam, pageInfo *vo.PageInfo) (*vo.DataPagination,
	exception.Exception) {
	count, projects, ex := ris.repo.EarlyPlanList(ris.db, pageInfo, params)
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

func (ris *reserveInspectServiceImpl) OutStorageInspList(params *vo.ReserveInspectParam, pageInfo *vo.PageInfo) (
	*vo.DataPagination, exception.Exception) {
	count, projects, ex := ris.repo.OutStorageInspList(ris.db, pageInfo, params)
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

func (ris *reserveInspectServiceImpl) EarlyPlanPass(openID string, id int64) exception.Exception {
	return ris.repo.Pass(ris.db, id, map[string]interface{}{
		"update_by": openID,
		"status":    constant.Posted,
	})
}

func (ris *reserveInspectServiceImpl) OutStoragePass(openID string, id int64) exception.Exception {
	pro, ex := ris.reserveRepo.Get(ris.db, id)
	if ex != nil {
		return ex
	}
	tx := ris.db.Begin()
	if tx.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	defer tx.Rollback()
	if ex := ris.repo.Pass(tx, id, map[string]interface{}{
		"update_by": openID,
		"status":    constant.OutStorage,
	}); ex != nil {
		return ex
	}
	gov := pro.ToGovReserveModel(openID)
	if ex = ris.GovRepo.Create(ris.db, gov); ex != nil {
		return ex
	}
	if err := tx.Commit(); err.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, err.Error)
	}
	return nil
}

func (ris *reserveInspectServiceImpl) Refuse(openID string, id int64) exception.Exception {
	return ris.repo.Refuse(ris.db, id, map[string]interface{}{
		"update_by": openID,
		"status":    constant.Draft,
	})
}
