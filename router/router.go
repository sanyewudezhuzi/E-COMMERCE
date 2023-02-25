package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apiaddress "github.com/sanyewudezhuzi/E-COMMERCE/controller/api_address"
	apicarousel "github.com/sanyewudezhuzi/E-COMMERCE/controller/api_carousel"
	apicarts "github.com/sanyewudezhuzi/E-COMMERCE/controller/api_carts"
	apifavorite "github.com/sanyewudezhuzi/E-COMMERCE/controller/api_favorite"
	apipay "github.com/sanyewudezhuzi/E-COMMERCE/controller/api_pay"
	apiproduct "github.com/sanyewudezhuzi/E-COMMERCE/controller/api_product"
	apiuser "github.com/sanyewudezhuzi/E-COMMERCE/controller/api_user"
	"github.com/sanyewudezhuzi/E-COMMERCE/middleware"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	E := r.Group("E-COMMERCE")
	{
		// 心跳测试
		E.GET("ping", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, "pong") })

		// unlogin
		unlogin := E.Group("unlogin")
		{
			unlogin.POST("register", apiuser.UserRegister)
			unlogin.POST("login", apiuser.UserLogin)
		}

		// carousel
		carousel := E.Group("carousel")
		{
			carousel.GET("show", apicarousel.ListCarousel)
			carousel.GET("productlist", apicarousel.ProductList)
			carousel.POST("productsearch", apicarousel.ProductSearch)
			carousel.POST("productshow/:id", apicarousel.ProductShow)
			carousel.POST("productimg/:id", apicarousel.ProductImg)
			carousel.GET("productcategories", apicarousel.ProductCategory)
		}

		// middleware
		E.Use(middleware.JWT())

		// user
		user := E.Group("user")
		{
			user.PUT("update", apiuser.UserUpdate)
			user.POST("avatar", apiuser.UploadAvatar)
			user.POST("sending-email", apiuser.SendEmail)
			user.POST("valid-email", apiuser.ValidEmail)
			user.POST("money", apiuser.ShowMoney)
		}

		// product
		product := E.Group("product")
		{
			product.POST("create", apiproduct.CreateProduct)
		}

		// favorite
		favorite := E.Group("favorite")
		{
			favorite.GET("list", apifavorite.FavoriteList)
			favorite.POST("create", apifavorite.FavoriteCreate)
			favorite.DELETE("delete/:id", apifavorite.FavoriteDelete)
		}

		// carts
		carts := E.Group("carts")
		{
			carts.POST("create", apicarts.CreateCart)
			carts.GET("show/:id", apicarts.ShowCarts)
			carts.PUT("update/:id", apicarts.UpdateCart)
			carts.DELETE("delete/:id", apicarts.DeleteCart)
		}

		// address
		address := E.Group("address")
		{
			address.POST("create", apiaddress.CreateAddress)
			address.GET("get/:id", apiaddress.GetAddress)
			address.GET("list", apiaddress.ListAddress)
			address.PUT("update/:id", apiaddress.UpdateAddress)
			address.DELETE("delete/:id", apiaddress.DeleteAddress)
		}

		// pay
		pay := E.Group("pay")
		{
			pay.POST("order", apipay.OrderPay)
		}
	}

	return r
}
