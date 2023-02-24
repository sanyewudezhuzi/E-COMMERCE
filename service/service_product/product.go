package serviceproduct

import (
	"context"
	"mime/multipart"
	"strconv"
	"sync"

	daoproduct "github.com/sanyewudezhuzi/E-COMMERCE/dao/dao_product"
	daouser "github.com/sanyewudezhuzi/E-COMMERCE/dao/dao_user"
	"github.com/sanyewudezhuzi/E-COMMERCE/model"
	"github.com/sanyewudezhuzi/E-COMMERCE/pkg/e"
	"github.com/sanyewudezhuzi/E-COMMERCE/serializer"
	"github.com/sanyewudezhuzi/E-COMMERCE/service/upload"
)

type ProductService struct {
	ID            uint   `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	CategoryID    uint   `json:"category_id" form:"category_id"`
	Title         string `json:"title" form:"title"`
	Info          string `json:"info" form:"info"`
	ImgPath       string `json:"img_path" form:"img_path"`
	Price         string `json:"price" form:"price"`
	DiscountPrice string `json:"discount_price" form:"discount_price"`
	OnSale        bool   `json:"on_sale" form:"on_sale"`
	Num           int    `json:"num" form:"num"`
	model.BasePage
}

func (s *ProductService) Create(ctx context.Context, uid uint, files []*multipart.FileHeader) serializer.Response {
	code := e.Success
	userDao := daouser.NewUserDao(ctx)
	boss, err := userDao.GetUserByID(uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
		}
	}
	// 以第一张作为封面图
	tmp, _ := files[0].Open()
	path, err := upload.UploadProductToLocalStatic(tmp, uid, s.Name)
	if err != nil {
		code = e.ErrorProductImgLoad
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	product := model.Product{
		Name:          s.Name,
		CategoryID:    s.CategoryID,
		Title:         s.Title,
		Info:          s.Info,
		ImgPath:       path,
		Price:         s.Price,
		DiscountPrice: s.DiscountPrice,
		OnSale:        true,
		Num:           s.Num,
		BossID:        uid,
		BossName:      boss.NickName,
		BossAvatar:    boss.Avatar,
	}
	productDao := daoproduct.NewProductDao(ctx)
	if err := productDao.CreateProduct(&product); err != nil {
		code = e.ErrorProductImgLoad
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for i, file := range files {
		num := strconv.Itoa(i)
		productImgDao := daoproduct.NewProductImgDapByDB(productDao.DB)
		tmp, _ = file.Open()
		path, err = upload.UploadProductToLocalStatic(tmp, uid, s.Name+num)
		if err != nil {
			code = e.ErrorUploadFail
			return serializer.Response{
				StatusCode: code,
				Msg:        e.GetMsg(code),
				Error:      err,
			}
		}
		productImg := model.ProductImg{
			ProductID: product.ID,
			ImgPath:   path,
		}
		err = productImgDao.CreateProductImg(&productImg)
		if err != nil {
			code = e.Error
			return serializer.Response{
				StatusCode: code,
				Msg:        e.GetMsg(code),
				Error:      err,
			}
		}
		wg.Done()
	}
	wg.Wait()
	return serializer.Response{
		StatusCode: code,
		Data:       serializer.BuildProduct(product),
		Msg:        e.GetMsg(code),
	}
}

func (s *ProductService) ProductList(ctx context.Context) serializer.Response {
	code := e.Success
	if s.PageSize == 0 {
		s.PageSize = 5
	}
	if s.PageNum == 0 {
		s.PageNum = 1
	}
	condition := make(map[string]interface{})
	if s.CategoryID != 0 {
		condition["category_id"] = s.CategoryID
	}
	productDao := daoproduct.NewProductDao(ctx)
	var products []model.Product
	total, err := productDao.GetProductCountByCondition(condition)
	if err != nil {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = daoproduct.NewProductDaoByDB(productDao.DB)
		products, _ = productDao.GetProductsByCondition(condition, s.BasePage)
		wg.Done()
	}()
	wg.Wait()
	return serializer.BuildListResponse(serializer.BuildProducts(products), int(total))
}
