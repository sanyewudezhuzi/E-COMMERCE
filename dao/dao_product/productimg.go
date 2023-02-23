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
