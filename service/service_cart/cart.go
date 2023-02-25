package servicecart

import (
	"context"
	"strconv"

	daocart "github.com/sanyewudezhuzi/E-COMMERCE/dao/dao_cart"
	daoproduct "github.com/sanyewudezhuzi/E-COMMERCE/dao/dao_product"
	daouser "github.com/sanyewudezhuzi/E-COMMERCE/dao/dao_user"
	"github.com/sanyewudezhuzi/E-COMMERCE/pkg/e"
	"github.com/sanyewudezhuzi/E-COMMERCE/serializer"
)

// CartService 创建购物车
type CartService struct {
	Id        uint `form:"id" json:"id"`
	BossID    uint `form:"boss_id" json:"boss_id"`
	ProductId uint `form:"product_id" json:"product_id"`
	Num       uint `form:"num" json:"num"`
}

func (service *CartService) Create(ctx context.Context, uId uint) serializer.Response {
	code := e.Success

	// 判断有无这个商品
	productDao := daoproduct.NewProductDao(ctx)
	product, err := productDao.GetProductByPID(int(service.ProductId))
	if err != nil {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}

	// 创建购物车
	cartDao := daocart.NewCartDao(ctx)
	cart, status, _ := cartDao.CreateCart(service.ProductId, uId, service.BossID)
	if status == e.Error {
		return serializer.Response{
			StatusCode: status,
			Msg:        e.GetMsg(status),
		}
	}

	userDao := daouser.NewUserDao(ctx)
	boss, _ := userDao.GetUserByID(service.BossID)
	return serializer.Response{
		StatusCode: status,
		Msg:        e.GetMsg(status),
		Data:       serializer.BuildCart(cart, &product, boss),
	}
}

// Show 购物车
func (service *CartService) Show(ctx context.Context, uId string) serializer.Response {
	code := e.Success
	cartDao := daocart.NewCartDao(ctx)
	userId, _ := strconv.Atoi(uId)
	carts, err := cartDao.ListCartByUserId(uint(userId))
	if err != nil {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	return serializer.Response{
		StatusCode: code,
		Data:       serializer.BuildCarts(carts),
		Msg:        e.GetMsg(code),
	}
}

// Update 修改购物车信息
func (service *CartService) Update(ctx context.Context, cId string) serializer.Response {
	code := e.Success
	cartId, _ := strconv.Atoi(cId)

	cartDao := daocart.NewCartDao(ctx)
	err := cartDao.UpdateCartNumById(uint(cartId), service.Num)
	if err != nil {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}

	return serializer.Response{
		StatusCode: code,
		Msg:        e.GetMsg(code),
	}
}

// 删除购物车
func (service *CartService) Delete(ctx context.Context) serializer.Response {
	code := e.Success
	cartDao := daocart.NewCartDao(ctx)
	err := cartDao.DeleteCartById(service.Id)
	if err != nil {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	return serializer.Response{
		StatusCode: code,
		Msg:        e.GetMsg(code),
	}
}
