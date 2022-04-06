#### go_gin_api.order 
订单表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | 主键 | int unsigned | PRI | NO | auto_increment |  |
| 2 | order_no | 订单号 | char(32) | UNI | NO |  |  |
| 3 | order_fee | 订单金额(分) | int unsigned |  | NO |  | 0 |
| 4 | status | 订单状态 1:已创建  2:已取消 | tinyint unsigned |  | NO |  | 1 |
| 5 | is_deleted | 是否删除 1:是  -1:否 | tinyint(1) |  | NO |  | -1 |
| 6 | created_at | 创建时间 | timestamp |  | NO | DEFAULT_GENERATED | CURRENT_TIMESTAMP |
| 7 | created_user | 创建人 | varchar(60) |  | NO |  |  |
| 8 | updated_at | 更新时间 | timestamp |  | NO | DEFAULT_GENERATED on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
| 9 | updated_user | 更新人 | varchar(60) |  | NO |  |  |
