package apiuser

import (
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

func UploadAvatar(ctx *gin.Context) {
	claims, ok := ctx.Get("claims")
	if !ok || claims == nil {
		ctx.JSON(http.StatusInternalServerError, "Failed to get claims.")
	}
	file, fileHeader, _ := ctx.Request.FormFile("file")
	fileSize := fileHeader.Size
	var uploadAvatar serviceuser.UserRegisterService
	if err := ctx.ShouldBind(&uploadAvatar); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		res := uploadAvatar.Upload(ctx.Request.Context(), claims.(*util.Claims).ID, file, fileSize)
		ctx.JSON(http.StatusOK, res)
	}
}

func SendEmail(ctx *gin.Context) {
	claims, ok := ctx.Get("claims")
	if !ok || claims == nil {
		ctx.JSON(http.StatusInternalServerError, "Failed to get claims.")
	}
	var sendEmail serviceuser.SendEmailService
	if err := ctx.ShouldBind(&sendEmail); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		res := sendEmail.Send(ctx.Request.Context(), claims.(*util.Claims).ID)
		ctx.JSON(http.StatusOK, res)
	}
}

func ValidEmail(ctx *gin.Context) {
	var validEmail serviceuser.ValidEmailService
	if err := ctx.ShouldBind(&validEmail); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		res := validEmail.Valid(ctx.Request.Context(), ctx.GetHeader("Authorization"))
		ctx.JSON(http.StatusOK, res)
	}
}

func ShowMoney(ctx *gin.Context) {
	claims, ok := ctx.Get("claims")
	if !ok || claims == nil {
		ctx.JSON(http.StatusInternalServerError, "Failed to get claims.")
	}
	var showMoney serviceuser.ShowMoneyService
	if err := ctx.ShouldBind(&showMoney); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		res := showMoney.ShowMoney(ctx.Request.Context(), claims.(*util.Claims).ID)
		ctx.JSON(http.StatusOK, res)
	}
}
