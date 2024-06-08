package model

import "gin-test-gorm/model/base_mysql"

type User struct {
	base_mysql.BaseField
	Name     string `gorm:"column:name;default:;NOT NULL;comment:'成品编号'"`
	Password string `gorm:"column:password;default:;NOT NULL;comment:'密码'"`
	Phone    string `gorm:"column:phone;default:;NOT NULL;comment:'手机号码'"`
}

func (u *User) TableName() string {
	return "user"
}
