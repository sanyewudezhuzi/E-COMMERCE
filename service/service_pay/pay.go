package servicepay

import (
	"context"
	"fmt"
	"strconv"

	daopay "github.com/sanyewudezhuzi/E-COMMERCE/dao/dao_pay"
	daoproduct "github.com/sanyewudezhuzi/E-COMMERCE/dao/dao_product"
	daouser "github.com/sanyewudezhuzi/E-COMMERCE/dao/dao_user"
	"github.com/sanyewudezhuzi/E-COMMERCE/pkg/e"
	"github.com/sanyewudezhuzi/E-COMMERCE/pkg/util"
	"github.com/sanyewudezhuzi/E-COMMERCE/serializer"
)

type OrderPayService struct {
	OrderID      uint    `json:"order_id" form:"order_id"`
	Money        float64 `json:"money" form:"money"`
	OrderNO      string  `json:"order_no" form:"order_no"`
	ProductID    uint    `json:"product_id" form:"product_id"`
	PayTime      string  `json:"pay_time" form:"pay_time"`
	Sign         string  `json:"sign" form:"sign"`
	BossID       uint    `json:"boss_id" form:"boss_id"`
	BossNickname string  `json:"Boss_nickname" form:"boss_nickname"`
	Num          int     `json:"num" form:"num"`
	Key          string  `json:"key" form:"key"`
}

func (s *OrderPayService) Pay(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	util.Encrypt.SetKey(s.Key)
	payDao := daopay.NewPayDao(ctx)

	tx := payDao.Begin()
	order, err := payDao.GetOrderByOrderID(s.OrderID, uid)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	pay_money := order.Money * float64(order.Num)

	userDao := daouser.NewUserDao(ctx)
	user, err := userDao.GetUserByID(uid)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	user_money, _ := strconv.ParseFloat(util.Encrypt.AesDecoding(user.Money), 64)
	if user_money-pay_money < 0.0 {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	fin_money := fmt.Sprintf("%f", user_money-pay_money)
	user.Money = util.Encrypt.AesEncoding(fin_money)

	boss, err := userDao.GetUserByID(s.BossID)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	boss_money, _ := strconv.ParseFloat(util.Encrypt.AesDecoding(boss.Money), 64)
	fin_money = fmt.Sprintf("%f", boss_money+pay_money)
	boss.Money = util.Encrypt.AesEncoding(fin_money)
	if err := userDao.UpdateUserByID(s.BossID, boss); err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	if err := userDao.UpdateUserByID(uid, user); err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}

	productDao := daoproduct.NewProductDao(ctx)
	product, err := productDao.GetProductByPID(int(s.ProductID))
	product.Num -= order.Num
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	if err := productDao.UpdateProduct(product); err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	tx.Commit()

	return serializer.Response{
		StatusCode: code,
		Msg:        e.GetMsg(code),
	}
}
