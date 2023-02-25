package serializer

import "github.com/sanyewudezhuzi/E-COMMERCE/model"

type Favorite struct {
	UserID        uint   `json:"user_id"`
	ProductID     uint   `json:"product_id"`
	CreateAt      int64  `json:"create_at"`
	Name          string `json:"name"`
	CategoryID    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	BossID        uint   `json:"boss_id"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
}

func BuildFavorite(favorite model.Favorite) Favorite {
	return Favorite{
		// 这里需要将 boss 和 product 传进来 此处省略了
		UserID:        favorite.UserID,
		ProductID:     favorite.ProductID,
		CreateAt:      favorite.CreatedAt.Unix(),
		Name:          favorite.Product.Name,
		CategoryID:    favorite.Product.CategoryID,
		Title:         favorite.Product.Title,
		Info:          favorite.Product.Info,
		ImgPath:       favorite.Product.ImgPath,
		Price:         favorite.Product.Price,
		DiscountPrice: favorite.Product.DiscountPrice,
		BossID:        favorite.BossID,
		Num:           favorite.Product.Num,
		OnSale:        favorite.Product.OnSale,
	}
}

func BuildFavorites(lists []model.Favorite) []Favorite {
	var favorites []Favorite = make([]Favorite, len(lists))
	for k, v := range lists {
		favorites[k] = BuildFavorite(v)
	}
	return favorites
}
