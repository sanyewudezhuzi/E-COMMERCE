package daoproduct

import (
	"context"

	"github.com/sanyewudezhuzi/E-COMMERCE/dao"
	"github.com/sanyewudezhuzi/E-COMMERCE/model"
	"gorm.io/gorm"
)

type ProductImgDao struct {
	*gorm.DB
}

func NewProductImgDao(ctx context.Context) *ProductImgDao {
	return &ProductImgDao{dao.NewDBClient(ctx)}
}

func NewProductImgDapByDB(db *gorm.DB) *ProductImgDao {
	return &ProductImgDao{db}
}

func (dao *ProductImgDao) CreateProductImg(productImg *model.ProductImg) error {
	return dao.DB.Model(&model.ProductImg{}).Create(&productImg).Error
}

func (dao *ProductImgDao) GetProductImgsAndTotalByPID(pid int) ([]model.ProductImg, int64, error) {
	var productImgs []model.ProductImg
	var count int64
	err := dao.DB.Model(&model.ProductImg{}).Where("product_id = ?", pid).Find(&productImgs).Count(&count).Error
	return productImgs, count, err
}
