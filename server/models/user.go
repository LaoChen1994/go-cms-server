package models

import (
	"errors"
	"fmt"
)

type User struct {
	Model
	Mobile   string `gorm:"column:mobile" json:"mobile"`
	Account  string `gorm:"column:account" json:"account" gorm:"index:nickname_idx"`
	Password string `gorm:"column:password" json:"password"`
	NickName string `gorm:"column:nickname" json:"nick_name" gorm:"index:nickname_idx"`
	Email    string `gorm:"column:email" json:"email"`
}

func CreateUser(user *User) (err error) {
	err = nil

	if isValidUserByName(user.Account) {
		err = errors.New("已存在相同用户名用户创建失败")
		return
	}

	if err = DB.Model(&User{}).Create(user).Error; err != nil {
		return
	}

	return
}

func IsValidUser(user *User) bool {
	DB.Model(&User{}).Where(user).First(user)

	if user.ID > 0 {
		return true
	}

	return false
}

func isValidUserByName(account string) bool {
	var user = User{
		Account: account,
	}
	DB.Model(&User{}).Select("id").Where(&user).First(&user)

	fmt.Println(user)

	if user.ID > 0 {
		return true
	}

	return false
}

func GetUserById(id uint) (user *User, err error) {
	err = DB.Model(&User{}).Where("id = ?", id).First(&user).Error

	return
}
