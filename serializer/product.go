package serializer

import (
	"github.com/sanyewudezhuzi/E-COMMERCE/conf"
	"github.com/sanyewudezhuzi/E-COMMERCE/model"
)

type Product struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	CategoryID    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	View          int64  `json:"view"`
	CreateAt      int64  `json:"create_at"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
	BossID        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
	BossAvatar    string `json:"boss_avatar"`
}

func BuildProduct(product model.Product) Product {
	return Product{
		ID:            product.ID,
		Name:          product.Name,
		CategoryID:    product.CategoryID,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		View:          product.View(),
		CreateAt:      product.CreatedAt.Unix(),
		Num:           product.Num,
		OnSale:        product.OnSale,
		BossID:        product.BossID,
		BossName:      product.BossName,
		BossAvatar:    product.BossAvatar,
	}
}
