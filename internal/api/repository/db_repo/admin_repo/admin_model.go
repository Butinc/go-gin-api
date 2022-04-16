package admin_repo

import (
	"gorm.io/gorm"
)

/**
 * query all admins with department
 */

func (qb *adminRepoQueryBuilder) QueryAllWithDepartment(db *gorm.DB) ([]*Admin, error) {
	var ret []*Admin
	err := qb.buildQuery(db).Preload("Department").Find(&ret).Error
	return ret, err
}
