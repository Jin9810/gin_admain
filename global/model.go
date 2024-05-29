package global

import (
	"gorm.io/gorm"
	"time"
)

// 默认字段
type GVA_MODEL struct {
	ID        uint      // 主键
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
	DeletedAt gorm.DeletedAt
}
