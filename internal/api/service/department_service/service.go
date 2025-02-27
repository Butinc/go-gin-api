package department_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/admin_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Create(ctx core.Context, adminData *CreateDepartmentData) (id int32, err error)
	PageList(ctx core.Context, searchData *SearchData) (listData []*admin_repo.Department, err error)
	PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error)
	Detail(ctx core.Context, searchOneData *SearchOneData) (info *admin_repo.Department, err error)
	ModifyDepartmentInfo(ctx core.Context, id int32, modifyData *ModifyData) (err error)
	Delete(ctx core.Context, id int32) (err error)
}

type service struct {
	db    db.Repo
	cache cache.Repo
}

func New(db db.Repo, cache cache.Repo) Service {
	return &service{
		db:    db,
		cache: cache,
	}
}

func (s *service) i() {}
