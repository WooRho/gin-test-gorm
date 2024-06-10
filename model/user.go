package model

type User struct {
	BaseField
	Password string `gorm:"column:password;"`
	Phone    string `gorm:"column:phone;"`
	Name     string `gorm:"column:name;"`
}

type UserParams struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

func (u *User) TableName() string {
	return "user"
}
