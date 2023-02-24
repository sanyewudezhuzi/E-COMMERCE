package daoproduct

import (
	"context"

	"github.com/sanyewudezhuzi/E-COMMERCE/dao"
	"github.com/sanyewudezhuzi/E-COMMERCE/model"
	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{dao.NewDBClient(ctx)}
}

func NewProductDaoByDB(db *gorm.DB) *ProductDao {
	return &ProductDao{db}
}

func (dao *ProductDao) CreateProduct(product *model.Product) error {
	return dao.DB.Model(&model.Product{}).Create(&product).Error
}

func (dao *ProductDao) GetProductCountByCondition(condition map[string]interface{}) (int64, error) {
	var count int64
	err := dao.DB.Model(&model.Product{}).Where(condition).Count(&count).Error
	return count, err
}

func (dao *ProductDao) GetProductsByCondition(condition map[string]interface{}, page model.BasePage) ([]model.Product, error) {
	var products []model.Product
	err := dao.DB.Model(&model.Product{}).Where(condition).Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&products).Error
	return products, err
}
