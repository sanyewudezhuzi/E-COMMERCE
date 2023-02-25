package daopay

import (
	"context"

	"github.com/sanyewudezhuzi/E-COMMERCE/dao"
	"github.com/sanyewudezhuzi/E-COMMERCE/model"
	"gorm.io/gorm"
)

type PayDao struct {
	*gorm.DB
}

func NewPayDao(ctx context.Context) *PayDao {
	return &PayDao{dao.NewDBClient(ctx)}
}

func (dao *PayDao) GetOrderByOrderID(oid, uid uint) (model.Order, error) {
	var order model.Order
	err := dao.DB.Model(&model.Order{}).Where("id = ? and user_id = ?", oid, uid).First(&order).Error
	return order, err
}
