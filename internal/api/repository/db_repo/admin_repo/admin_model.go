package admin_repo

import "gorm.io/gorm"

/**
 * 关于管理员的额外操作方法
 */

func (qb *adminRepoQueryBuilder) JoinDepartment(db *gorm.DB) *adminRepoQueryBuilder {
	qb.query.Joins("left join department on admin.department_id = department.id")
	return qb
}
