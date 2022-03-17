package service

import (
	"lpms/app/repositories"
	"lpms/app/response"
	"lpms/app/vo"
	"lpms/commom/drivers/database"
	"lpms/exception"
	"strconv"
	"strings"
	"sync"

	"gorm.io/gorm"
)

var (
	implementGovServiceInstance ImplementGovService
	implementGovOnce            sync.Once
)

type implementGovServiceImpl struct {
	db      *gorm.DB
	repo    repositories.ImplementGovRepo
	objRepo repositories.ObjectRepo
}

func GetImplementGovService() ImplementGovService {
	implementGovOnce.Do(func() {
		implementGovServiceInstance = &implementGovServiceImpl{
			db:      database.GetDriver(),
			repo:    repositories.GetImplementGovRepo(),
			objRepo: repositories.GetObjectRepo(),
		}
	})
	return implementGovServiceInstance
}

type ImplementGovService interface {
	Create(openID string, param *vo.ImplementGovReq) exception.Exception
	Get(id int64) (*vo.ImplementGovResp, exception.Exception)
	List(params *vo.ImplementGovFilterParam, pageInfo *vo.PageInfo) (*vo.DataPagination, exception.Exception)
	Delete(id int64) exception.Exception
	MultiDelete(ids string) exception.Exception
}

func (isi *implementGovServiceImpl) Create(openID string, param *vo.ImplementGovReq) exception.Exception {
	res := param.ToModel(openID)
	return isi.repo.Create(isi.db, res)
}

func (isi *implementGovServiceImpl) Get(id int64) (*vo.ImplementGovResp, exception.Exception) {
	impl, ex := isi.repo.Get(isi.db, id)
	if ex != nil {
		return nil, ex
	}
	resp, err := vo.NewImplementGovResponse(impl)
	if err != nil {
		return nil, exception.Wrap(response.ExceptionUnmarshalJSON, err)
	}

	return resp, nil
}

func (isi *implementGovServiceImpl) List(params *vo.ImplementGovFilterParam, pageInfo *vo.PageInfo) (*vo.DataPagination,
	exception.Exception) {
	count, projects, ex := isi.repo.List(isi.db, pageInfo, params)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.ListImplementGovResp, 0, len(projects))
	for i := range projects {
		resp = append(resp, vo.ListImplementGovResp{
			ID:               projects[i].ID,
			Name:             projects[i].Name,
			Level:            projects[i].Level,
			ProjectType:      projects[i].ProjectType,
			ConstructSubject: projects[i].ConstructSubject,
			PlanBegin:        projects[i].PlanBegin,
			FinishTime:       projects[i].FinishTime,
			Status:           projects[i].Status,
		})
	}
	return vo.NewDataPagination(count, resp, pageInfo), nil
}

func (isi *implementGovServiceImpl) Delete(id int64) exception.Exception {
	pro, ex := isi.repo.Get(isi.db, id)
	if ex != nil {
		return ex
	}
	if pro.SitePhoto != "" {
		if exx := isi.objRepo.Delete(isi.db, pro.SitePhoto); exx != nil {
			return exx
		}
	}
	if pro.UploadCadID != "" {
		if exx := isi.objRepo.Delete(isi.db, pro.UploadCadID); exx != nil {
			return exx
		}
	}
	return isi.repo.Delete(isi.db, id)
}

func (isi *implementGovServiceImpl) MultiDelete(ids string) exception.Exception {
	idslice := strings.Split(ids, ",")
	if len(idslice) == 0 {
		return exception.New(response.ExceptionInvalidRequestParameters, "无效参数")
	}
	did := make([]int64, 0, len(idslice))
	for i := range idslice {
		id, err := strconv.ParseUint(idslice[i], 10, 0)
		if err != nil {
			return exception.Wrap(response.ExceptionParseStringToInt64Error, err)
		}
		did = append(did, int64(id))
	}
	for i := range did {
		pro, ex := isi.repo.Get(isi.db, did[i])
		if ex != nil {
			return ex
		}
		if pro.SitePhoto != "" {
			if exx := isi.objRepo.Delete(isi.db, pro.SitePhoto); exx != nil {
				return exx
			}
		}
		if pro.UploadCadID != "" {
			if exx := isi.objRepo.Delete(isi.db, pro.UploadCadID); exx != nil {
				return exx
			}
		}
	}
	return isi.repo.MultiDelete(isi.db, did)
}
