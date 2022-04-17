package department_repo

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/admin_repo"
	"time"
)

// Department 部门表
//go:generate gormgen -structs Department -input .
type Department struct {
	Id          int32     // 主键
	Name        string    // 部门名称
	IsDeleted   int32     // 是否删除 1:是  -1:否
	CreatedAt   time.Time `gorm:"time"` // 创建时间
	CreatedUser string    // 创建人
	UpdatedAt   time.Time `gorm:"time"` // 更新时间
	UpdatedUser string    // 更新人
	Admins      []admin_repo.Admin
}
