package daouser

import (
	"context"

	"github.com/sanyewudezhuzi/E-COMMERCE/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{model.NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// 根据 account 判断用户是否已注册
func (dao *UserDao) ExistOrNotByAccount(account string) bool {
	var user model.User
	model.DB.Model(&model.User{}).Where("account = ?", account).First(&user)
	if user.ID == 0 {
		return false
	}
	return true
}

// 创建用户
func (dao *UserDao) CreateUser(user *model.User) error {
	return model.DB.Model(&model.User{}).Create(user).Error
}
