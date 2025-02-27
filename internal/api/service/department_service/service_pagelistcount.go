package department_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/admin_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

func (s *service) PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error) {
	qb := admin_repo.NewDepartmentQueryBuilder()
	qb = qb.WhereIsDeleted(db_repo.EqualPredicate, -1)

	if searchData.Name != "" {
		qb.WhereName(db_repo.LikePredicate, "%"+searchData.Name+"%")
	}

	total, err = qb.Count(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}

	return
}
