#### go_gin_api.admin 
管理员表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | 主键 | int(11) unsigned | PRI | NO | auto_increment |  |
| 2 | department_id | 部门ID | int(11) unsigned |  | NO |  | 0 |
| 3 | username | 用户名 | varchar(32) | UNI | NO |  |  |
| 4 | password | 密码 | varchar(100) |  | NO |  |  |
| 5 | nickname | 昵称 | varchar(60) |  | NO |  |  |
| 6 | mobile | 手机号 | varchar(20) |  | NO |  |  |
| 7 | is_used | 是否启用 1:是  -1:否 | tinyint(1) |  | NO |  | 1 |
| 8 | is_deleted | 是否删除 1:是  -1:否 | tinyint(1) |  | NO |  | -1 |
| 9 | created_at | 创建时间 | timestamp |  | NO |  | CURRENT_TIMESTAMP |
| 10 | created_user | 创建人 | varchar(60) |  | NO |  |  |
| 11 | updated_at | 更新时间 | timestamp |  | NO | on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
| 12 | updated_user | 更新人 | varchar(60) |  | NO |  |  |
