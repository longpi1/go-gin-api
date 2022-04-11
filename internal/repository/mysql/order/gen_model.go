package order

import "time"

// Order 订单表
//go:generate gormgen -structs Order -input .
type Order struct {
	Id          int32     // 主键
	OrderNo     string    // 订单号
	OrderFee    int32     // 订单金额(分)
	Status      int32     // 订单状态 1:已创建  2:已取消
	IsDeleted   int32     // 是否删除 1:是  -1:否
	CreatedAt   time.Time `gorm:"time"` // 创建时间
	CreatedUser string    // 创建人
	UpdatedAt   time.Time `gorm:"time"` // 更新时间
	UpdatedUser string    // 更新人
}
