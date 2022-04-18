package admin_repo

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/related_repo"
	"gorm.io/gorm"
)

/**
 * query all admins with department
 */

func (qb *adminRepoQueryBuilder) QueryAllWithDepartment(db *gorm.DB) ([]*related_repo.AdminWithDepartment, error) {
	var ret []*related_repo.AdminWithDepartment
	err := qb.buildQuery(db).Preload("Department").Find(&ret).Error
	return ret, err
}
