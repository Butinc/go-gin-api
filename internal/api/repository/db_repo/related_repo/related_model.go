package related_repo

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/admin_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/department_repo"
)

type AdminWithDepartment struct {
	admin_repo.Admin
	Department department_repo.Department
}

type DepartmentWithAdmins struct {
	department_repo.Department
	Admins []admin_repo.Admin
}
