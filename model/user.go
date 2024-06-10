package model

type User struct {
	BaseField
	Name     string `gorm:"column:name;default:;NOT NULL;comment:'成品编号'"`
	Password string `gorm:"column:password;default:;NOT NULL;comment:'密码'"`
	Phone    string `gorm:"column:phone;default:;NOT NULL;comment:'手机号码'"`
}

type UserParams struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

func (u UserParams) AdjustParam() {
	//TODO implement me
	//panic("implement me")
}

func (u UserParams) ValidateParam() error {
	//TODO implement me
	//panic("implement me")
	return nil
}

func (u *User) TableName() string {
	return "user"
}
