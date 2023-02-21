package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
		// middleware
		E.Use(middleware.JWT())
		// user
		user := E.Group("user")
		{
			user.PUT("update", apiuser.UserUpdate)
			user.POST("avatar", apiuser.UploadAvatar)
			user.POST("sending-email", apiuser.SendEmail)
			user.POST("valid-email", apiuser.ValidEmail)
		}
	}

	return r
}
