package serializer

import (
	"context"

	daoproduct "github.com/sanyewudezhuzi/E-COMMERCE/dao/dao_product"
	daouser "github.com/sanyewudezhuzi/E-COMMERCE/dao/dao_user"
	"github.com/sanyewudezhuzi/E-COMMERCE/model"
)

// 购物车
type Cart struct {
	ID            uint   `json:"id"`
	UserID        uint   `json:"user_id"`
	ProductID     uint   `json:"product_id"`
	CreateAt      int64  `json:"create_at"`
	Num           uint   `json:"num"`
	MaxNum        uint   `json:"max_num"`
	Check         bool   `json:"check"`
	Name          string `json:"name"`
	ImgPath       string `json:"img_path"`
	DiscountPrice string `json:"discount_price"`
	BossId        uint   `json:"boss_id"`
	BossNickName  string `json:"boss_nick_name"`
}

func BuildCart(cart *model.Cart, product *model.Product, boss *model.User) Cart {
	return Cart{
		ID:            cart.ID,
		UserID:        cart.UserID,
		ProductID:     cart.ProductID,
		CreateAt:      cart.CreatedAt.Unix(),
		Num:           cart.Num,
		MaxNum:        cart.MaxNum,
		Check:         cart.Check,
		Name:          product.Name,
		ImgPath:       product.ImgPath,
		DiscountPrice: product.DiscountPrice,
		BossId:        boss.ID,
		BossNickName:  boss.NickName,
	}
}

func BuildCarts(items []*model.Cart) (carts []Cart) {
	for _, item1 := range items {
		product, err := daoproduct.NewProductDao(context.Background()).
			GetProductByPID(int(item1.ProductID))
		if err != nil {
			continue
		}
		boss, err := daouser.NewUserDao(context.Background()).
			GetUserByID(item1.BossID)
		if err != nil {
			continue
		}
		cart := BuildCart(item1, &product, boss)
		carts = append(carts, cart)
	}
	return carts
}
