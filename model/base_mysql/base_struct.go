package base_mysql

import (
	"gorm.io/gorm"
)

type BaseField struct {
	gorm.Model
	CreatorId uint64 `gorm:"column:creator_id;type:bigint(20) unsigned;not null;default:0;comment:创建人id"` // 创建人id
	UpdaterId uint64 `gorm:"column:updater_id;type:bigint(20) unsigned;not null;default:0;comment:更新人id"` // 更新人id
	DeletedId uint64 `gorm:"column:deleted_id;type:bigint(20) unsigned;not null;default:0;comment:删除人id"` // 删除人id
}

// 多条插入比单条插入性能好；
