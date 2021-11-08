package department_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/department_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type SearchOneData struct {
	Id   int32  // 用户ID
	Name string // 名称
}

func (s *service) Detail(ctx core.Context, searchOneData *SearchOneData) (info *department_repo.Department, err error) {

	qb := department_repo.NewQueryBuilder()
	qb.WhereIsDeleted(db_repo.EqualPredicate, -1)

	if searchOneData.Id != 0 {
		qb.WhereId(db_repo.EqualPredicate, searchOneData.Id)
	}

	if searchOneData.Name != "" {
		qb.WhereName(db_repo.EqualPredicate, searchOneData.Name)
	}

	info, err = qb.QueryOne(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
