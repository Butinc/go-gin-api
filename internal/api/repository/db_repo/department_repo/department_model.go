package department_repo

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/related_repo"
	"gorm.io/gorm"
)

/**
 * query all departments with admins
 */

func (qb *departmentRepoQueryBuilder) QueryAllWithAdmins(db *gorm.DB) ([]*related_repo.DepartmentWithAdmins, error) {
	var ret []*related_repo.DepartmentWithAdmins
	err := qb.buildQuery(db).Preload("Admins").Find(&ret).Error
	return ret, err
}
