package serializer

import "github.com/sanyewudezhuzi/E-COMMERCE/model"

type Carousel struct {
	ID        uint   `json:"id"`
	ImgPath   string `json:"img_path"`
	ProductID uint   `json:"product_id"`
	CreateAt  int64  `json:"create_at"`
}

func BuildCarousel(item *model.Carousel) Carousel {
	return Carousel{
		ID:        item.ID,
		ImgPath:   item.ImgPath,
		ProductID: item.ProductID,
		CreateAt:  item.CreatedAt.Unix(),
	}
}

func BuildCarousels(item []model.Carousel) []Carousel {
	var carousels []Carousel = make([]Carousel, len(item))
	for k, v := range item {
		carousels[k] = BuildCarousel(&v)
	}
	return carousels
}
