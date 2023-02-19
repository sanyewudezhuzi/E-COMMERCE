package apiuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	serviceuser "github.com/sanyewudezhuzi/E-COMMERCE/service/service_user"
)

func UserRegister(ctx *gin.Context) {
	// 创建服务
	var userRegister serviceuser.UserRegisterService

	// 绑定参数
	if err := ctx.ShouldBind(&userRegister); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		res := userRegister.Register(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	}
}

func UserLogin(ctx *gin.Context) {
	// 创建服务
	var userLogin serviceuser.UserRegisterService

	// 绑定参数
	if err := ctx.ShouldBind(&userLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		res := userLogin.Login(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	}
}
