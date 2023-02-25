package servicefavorite

import (
	"context"
	"strconv"

	daofavorite "github.com/sanyewudezhuzi/E-COMMERCE/dao/dao_favorite"
	daoproduct "github.com/sanyewudezhuzi/E-COMMERCE/dao/dao_product"
	daouser "github.com/sanyewudezhuzi/E-COMMERCE/dao/dao_user"
	"github.com/sanyewudezhuzi/E-COMMERCE/model"
	"github.com/sanyewudezhuzi/E-COMMERCE/pkg/e"
	"github.com/sanyewudezhuzi/E-COMMERCE/serializer"
)

type FavoriteService struct {
	PID        uint `json:"pid" form:"pid"`
	BID        uint `json:"bid" form:"bid"`
	FavoriteID uint `json:"favorite_id" form:"favorite_id"`
	model.BasePage
}

func (s *FavoriteService) List(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	favoriteDao := daofavorite.NewFavoriteDao(ctx)
	list, err := favoriteDao.FavoriteList(uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	return serializer.BuildListResponse(serializer.BuildFavorites(list), len(list))
}

func (s *FavoriteService) Create(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	favoriteDao := daofavorite.NewFavoriteDao(ctx)
	exist, err := favoriteDao.FavoriteExist(s.PID, uid)
	if exist || err != nil {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	userDao := daouser.NewUserDao(ctx)
	user, err := userDao.GetUserByID(uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	bossDao := daouser.NewUserDao(ctx)
	boss, err := bossDao.GetUserByID(s.BID)
	if err != nil {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	productDao := daoproduct.NewProductDao(ctx)
	product, err := productDao.GetProductByPID(int(s.PID))
	if err != nil {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	favorite := &model.Favorite{
		User:      *user,
		UserID:    uid,
		Product:   product,
		ProductID: s.PID,
		Boss:      *boss,
		BossID:    s.BID,
	}
	err = favoriteDao.FavoriteCreate(*favorite)
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
		Data:       serializer.BuildFavorite(*favorite),
		Msg:        e.GetMsg(code),
	}
}

func (s *FavoriteService) Delete(ctx context.Context, uid uint, id string) serializer.Response {
	code := e.Success
	fid, _ := strconv.Atoi(id)
	favoriteDao := daofavorite.NewFavoriteDao(ctx)
	err := favoriteDao.FavoriteDelete(uint(fid), uid)
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
