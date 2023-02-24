package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apicarousel "github.com/sanyewudezhuzi/E-COMMERCE/controller/api_carousel"
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
	}

	return r
}
