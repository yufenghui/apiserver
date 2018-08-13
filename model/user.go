package model

import "github.com/yufenghui/apiserver/pkg/auth"

// User represents a registered user.
type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null"`
	Password string `json:"password" gorm:"column:password;not null"`
}

func (u *UserModel) TableName() string {
	return "tb_users"
}

func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

func DeleteUser(id uint64) error {
	user := UserModel{}
	user.BaseModel.Id = id

	return DB.Self.Delete(user).Error
}

func (u *UserModel) Update() error {
	return DB.Self.Save(u).Error
}

func GetUser(username string) (UserModel, error) {
	u := UserModel{}
	d := DB.Self.Where("username = ?", username).First(&u)

	return u, d.Error
}

// List all users
func List() ([]UserModel, error) {
	users := make([]UserModel, 0)
	err := DB.Self.Find(&users).Error

	if err != nil {
		return users, err
	}

	return users, nil
}
