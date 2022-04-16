#### go_gin_api.department 
部门表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | 主键 | int(11) unsigned | PRI | NO | auto_increment |  |
| 2 | name | 部门名称 | varchar(32) | UNI | NO |  |  |
| 3 | is_deleted | 是否删除 1:是  -1:否 | tinyint(1) |  | NO |  | -1 |
| 4 | created_at | 创建时间 | timestamp |  | NO |  | CURRENT_TIMESTAMP |
| 5 | created_user | 创建人 | varchar(60) |  | NO |  |  |
| 6 | updated_at | 更新时间 | timestamp |  | NO | on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
| 7 | updated_user | 更新人 | varchar(60) |  | NO |  |  |
