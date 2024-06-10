package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type ITable interface {
	TableName() string
	GetUserId()
}

func CreateModel(ctx context.Context, db *gorm.DB, model ITable) {
	model.GetUserId()
	err := db.Debug().WithContext(ctx).Create(model).Error
	if err != nil {
		fmt.Println(err)
	}
}
