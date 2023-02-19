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
func (dao *UserDao) ExistOrNotByAccount(account string) (*model.User, bool) {
	var user model.User
	dao.DB.Model(&model.User{}).Where("account = ?", account).First(&user)
	if user.ID == 0 {
		return nil, false
	}
	return &user, true
}

// 创建用户
func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(user).Error
}

// 根据用户 id 获取 user
func (dao *UserDao) GetUserByID(id uint) (*model.User, error) {
	var user *model.User
	err := dao.DB.Model(&model.User{}).Where("id = ?", id).First(&user).Error
	return user, err
}

// 根据用户 id 更改 user
func (dao *UserDao) UpdateUserByID(id uint, user *model.User) error {
	return dao.DB.Model(&model.User{}).Where("id = ?", id).Updates(&user).Error
}
