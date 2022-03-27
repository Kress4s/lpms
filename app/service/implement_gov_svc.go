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
	db             *gorm.DB
	repo           repositories.ImplementGovRepo
	objRepo        repositories.ObjectRepo
	GovProcessRepo repositories.GovProgressRepo
	userRepo       repositories.UserRepo
}

func GetImplementGovService() ImplementGovService {
	implementGovOnce.Do(func() {
		implementGovServiceInstance = &implementGovServiceImpl{
			db:             database.GetDriver(),
			repo:           repositories.GetImplementGovRepo(),
			objRepo:        repositories.GetObjectRepo(),
			GovProcessRepo: repositories.GetGovProgressRepo(),
			userRepo:       repositories.GetUserRepo(),
		}
	})
	return implementGovServiceInstance
}

type ImplementGovService interface {
	Create(openID string, param *vo.ImplementGovReq) exception.Exception
	Get(id int64) (*vo.ImplementGovResp, exception.Exception)
	List(user string, params *vo.ImplementGovFilterParam, pageInfo *vo.PageInfo) (*vo.DataPagination, exception.Exception)
	Delete(id int64) exception.Exception
	MultiDelete(ids string) exception.Exception
	ListStatusCount(user string, params *vo.ImplementGovCountFilter) ([]vo.StatusCountResp, exception.Exception)
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

func (isi *implementGovServiceImpl) List(user string, params *vo.ImplementGovFilterParam, pageInfo *vo.PageInfo) (*vo.DataPagination,
	exception.Exception) {
	userInfo, ex := isi.userRepo.Get(isi.db, user)
	if ex != nil {
		return nil, ex
	}
	count, projects, ex := isi.repo.List(isi.db, pageInfo, params, userInfo.IsAdmin, user)
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
			BeginImpl:        projects[i].BeginImpl,
			FinishTime:       projects[i].FinishTime,
			Status:           projects[i].Status,
			StartTime:        projects[i].StartTime,
			ProjectCode:      projects[i].ProjectCode,
			DutyUnit:         projects[i].DutyUint,
			Type:             projects[i].Type,
		})
	}
	return vo.NewDataPagination(count, resp, pageInfo), nil
}

func (isi *implementGovServiceImpl) Delete(id int64) exception.Exception {
	pro, ex := isi.repo.Get(isi.db, id)
	if ex != nil {
		return ex
	}
	tx := isi.db.Begin()
	if tx.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}

	defer tx.Rollback()
	if pro.SitePhoto != "" {
		if exx := isi.objRepo.Delete(tx, pro.SitePhoto); exx != nil {
			return exx
		}
	}
	if pro.UploadCadID != "" {
		if exx := isi.objRepo.Delete(tx, pro.UploadCadID); exx != nil {
			return exx
		}
	}

	if ex := isi.GovProcessRepo.DeleteByProjectID(tx, id); ex != nil {
		return ex
	}
	if ex != isi.repo.Delete(tx, id) {
		return ex
	}
	if err := tx.Commit(); err.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, err.Error)
	}
	return nil
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
	tx := isi.db.Begin()
	if tx.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}

	for i := range did {
		pro, ex := isi.repo.Get(tx, did[i])
		if ex != nil {
			return ex
		}
		if pro.SitePhoto != "" {
			if exx := isi.objRepo.Delete(tx, pro.SitePhoto); exx != nil {
				return exx
			}
		}
		if pro.UploadCadID != "" {
			if exx := isi.objRepo.Delete(tx, pro.UploadCadID); exx != nil {
				return exx
			}
		}
	}
	if ex := isi.GovProcessRepo.DeleteByProjectID(tx, did...); ex != nil {
		return ex
	}
	if ex := isi.repo.MultiDelete(isi.db, did); ex != nil {
		return ex
	}
	if err := tx.Commit(); err.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, err.Error)
	}
	return nil
}

func (isi *implementGovServiceImpl) ListStatusCount(user string, params *vo.ImplementGovCountFilter) ([]vo.StatusCountResp, exception.Exception) {
	userInfo, ex := isi.userRepo.Get(isi.db, user)
	if ex != nil {
		return nil, ex
	}
	res, ex := isi.repo.ListStatusCount(isi.db, params, userInfo.IsAdmin, user)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.StatusCountResp, 0, len(res))
	for i := range res {
		resp = append(resp, vo.StatusCountResp{
			Status: res[i].Status,
			Count:  res[i].Count,
		})
	}
	return resp, nil
}
