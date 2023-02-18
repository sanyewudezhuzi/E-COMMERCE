package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account        string `gorm:"unique"`
	Email          string
	PasswordDigest string
	NickName       string
	Status         string
	Avatar         string
	Money          string
}

const (
	PasswordCost        = 12       // 密码加密难度
	Active       string = "active" // 激活用户
)

// 密码加密
func (u *User) Bcrypt(pwd string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	if err != nil {
		return err
	}
	u.PasswordDigest = string(hash)
	return nil
}

// 密码校验
func (u *User) Check(pwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(pwd)) == nil
}
