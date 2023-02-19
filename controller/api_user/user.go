package apiuser

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanyewudezhuzi/E-COMMERCE/pkg/util"
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

func UserUpdate(ctx *gin.Context) {
	claims, ok := ctx.Get("claims")
	if !ok || claims == nil {
		ctx.JSON(http.StatusInternalServerError, "Failed to get claims.")
	}
	fmt.Println("claim:", claims)
	// 创建服务
	var userUpdate serviceuser.UserRegisterService

	// 绑定参数
	if err := ctx.ShouldBind(&userUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		res := userUpdate.Update(ctx.Request.Context(), claims.(*util.Claims).ID)
		ctx.JSON(http.StatusOK, res)
	}
}
