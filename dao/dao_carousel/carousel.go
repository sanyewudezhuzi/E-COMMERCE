package daocarousel

import (
	"context"

	"github.com/sanyewudezhuzi/E-COMMERCE/dao"
	"github.com/sanyewudezhuzi/E-COMMERCE/model"
	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB
}

func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{dao.NewDBClient(ctx)}
}

func NewCarouselDaoByDB(db *gorm.DB) *CarouselDao {
	return &CarouselDao{db}
}

func (dao *CarouselDao) ShowCarousel() ([]model.Carousel, error) {
	var carousels []model.Carousel
	err := dao.DB.Model(&model.Carousel{}).Find(&carousels).Error
	return carousels, err
}
