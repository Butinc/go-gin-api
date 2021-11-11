package department_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/department_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type ModifyData struct {
	Name string // 部门
}

func (s *service) ModifyDepartmentInfo(ctx core.Context, id int32, modifyData *ModifyData) (err error) {
	data := map[string]interface{}{
		"name":         modifyData.Name,
		"updated_user": ctx.UserName(),
	}

	qb := department_repo.NewQueryBuilder()
	qb.WhereId(db_repo.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	return
}
