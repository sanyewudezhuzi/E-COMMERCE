package daofavorite

import (
	"context"

	"github.com/sanyewudezhuzi/E-COMMERCE/dao"
	"github.com/sanyewudezhuzi/E-COMMERCE/model"
	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{dao.NewDBClient(ctx)}
}

func (dao *FavoriteDao) FavoriteList(uid uint) ([]model.Favorite, error) {
	// 可添加分页功能 此处省略
	var list []model.Favorite
	err := dao.DB.Model(&model.Favorite{}).Where("user_id = ?", uid).Find(&list).Error
	return list, err
}

func (dao *FavoriteDao) FavoriteCreate(favorite model.Favorite) error {
	return dao.DB.Model(&model.Favorite{}).Create(&favorite).Error
}

func (dao *FavoriteDao) FavoriteDelete(fid, uid uint) error {
	return dao.DB.Model(&model.Favorite{}).Where("id = ? and user_id = ?", fid, uid).Delete(&model.Favorite{}).Error
}

func (dao *FavoriteDao) FavoriteExist(pid, uid uint) (bool, error) {
	var count int64
	err := dao.DB.Model(&model.Favorite{}).Where("product_id = ? and user_id = ?", pid, uid).Count(&count).Error
	if err != nil || count != 0 {
		return true, err
	}
	return false, nil
}
