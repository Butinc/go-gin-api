package admin_repo

import (
	"gorm.io/gorm"
)

/**
 * query all departments with admins
 */

func (qb *departmentRepoQueryBuilder) QueryAllWithAdmins(db *gorm.DB) ([]*Department, error) {
	var ret []*Department
	err := qb.buildQuery(db).Preload("Admins").Find(&ret).Error
	return ret, err
}
