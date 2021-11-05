package department_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/department_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type CreateDepartmentData struct {
	Name string // 部门
}

func (s *service) Create(ctx core.Context, departmentData *CreateDepartmentData) (id int32, err error) {
	model := department_repo.NewModel()
	model.Name = departmentData.Name
	model.CreatedUser = ctx.UserName()
	model.IsDeleted = -1

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
