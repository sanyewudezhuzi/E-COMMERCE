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

		// user
		user := E.Group("user")
		{
			user.POST("register", apiuser.UserRegister)
			user.POST("login", apiuser.UserLogin)
		}
	}
	return r
}
