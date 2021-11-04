package mysql_table

func CreateDepartmentTableSql() (sql string) {
	sql = "CREATE TABLE `department` ("
	sql += "`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',"
	sql += "`name` varchar(32) NOT NULL DEFAULT '' COMMENT '部门名称',"
	sql += "`is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT '是否删除 1:是  -1:否',"
	sql += "`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',"
	sql += "`created_user` varchar(60) NOT NULL DEFAULT '' COMMENT '创建人',"
	sql += "`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',"
	sql += "`updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT '更新人',"
	sql += "PRIMARY KEY (`id`),"
	sql += "UNIQUE KEY `unique_name` (`name`)"
	sql += ") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='部门表';"

	return
}

func CreateDepartmentTableDataSql() (sql string) {
	sql = "INSERT INTO `department` (`id`, `name`, `created_user`) VALUES"
	sql += "(1, '财务部', 'init');"

	return
}
