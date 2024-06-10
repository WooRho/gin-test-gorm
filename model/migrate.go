package model

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&User{},
	)
	//fmt.Println(err.Error())
}
