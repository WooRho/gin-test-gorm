package model

import (
	"gorm.io/gorm"
)

type BaseField struct {
	gorm.Model
	//CreatorId uint `gorm:"column:creator_id;type:bigint(20) unsigned;not null;default:0;comment:创建人id"` // 创建人id
	//UpdaterId uint `gorm:"column:updater_id;type:bigint(20) unsigned;not null;default:0;comment:更新人id"` // 更新人id
	//DeletedId uint `gorm:"column:deleted_id;type:bigint(20) unsigned;not null;default:0;comment:删除人id"` // 删除人id
}

// 多条插入比单条插入性能好；
//func (model *BaseField) BeforeCreate() {
//	value, _ := metadata.UserInfo.Load("token-123456")
//	model.CreatorId = value.(uint)
//	model.UpdaterId = value.(uint)
//}

//func (model *BaseField) BeforeUpdate() {
//	value, _ := metadata.UserInfo.Load("token-123456")
//	model.UpdaterId = value.(uint)
//}

//func (model *BaseField) BeforeDel() {
//	value, _ := metadata.UserInfo.Load("token-123456")
//	model.DeletedId = value.(uint)
//}
