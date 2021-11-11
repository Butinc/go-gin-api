package department_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/department_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Create(ctx core.Context, adminData *CreateDepartmentData) (id int32, err error)
	PageList(ctx core.Context, searchData *SearchData) (listData []*department_repo.Department, err error)
	PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error)
	Detail(ctx core.Context, searchOneData *SearchOneData) (info *department_repo.Department, err error)
	ModifyDepartmentInfo(ctx core.Context, id int32, modifyData *ModifyData) (err error)
	Delete(ctx core.Context, id int32) (err error)
	/*UpdateUsed(ctx core.Context, id int32, used int32) (err error)
	ResetPassword(ctx core.Context, id int32) (err error)
	ModifyPassword(ctx core.Context, id int32, newPassword string) (err error)

	CreateMenu(ctx core.Context, menuData *CreateMenuData) (err error)
	ListMenu(ctx core.Context, searchData *SearchListMenuData) (menuData []ListMenuData, err error)
	MyMenu(ctx core.Context, searchData *SearchMyMenuData) (menuData []ListMyMenuData, err error)
	MyAction(ctx core.Context, searchData *SearchMyActionData) (actionData []MyActionData, err error)*/
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
