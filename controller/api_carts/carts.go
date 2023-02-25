package apicarts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanyewudezhuzi/E-COMMERCE/pkg/util"
	servicecart "github.com/sanyewudezhuzi/E-COMMERCE/service/service_cart"
)

func CreateCart(ctx *gin.Context) {
	claims, ok := ctx.Get("claims")
	if !ok || claims == nil {
		ctx.JSON(http.StatusInternalServerError, "Failed to get claims.")
	}
	var createCartService servicecart.CartService
	if err := ctx.ShouldBind(&createCartService); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		res := createCartService.Create(ctx.Request.Context(), claims.(*util.Claims).ID)
		ctx.JSON(http.StatusOK, res)
	}
}

func ShowCarts(ctx *gin.Context) {
	var showCartsService servicecart.CartService
	if err := ctx.ShouldBind(&showCartsService); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		res := showCartsService.Show(ctx.Request.Context(), ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	}
}

func UpdateCart(ctx *gin.Context) {
	updateCartService := servicecart.CartService{}
	if err := ctx.ShouldBind(&updateCartService); err == nil {
		res := updateCartService.Update(ctx.Request.Context(), ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
}

func DeleteCart(ctx *gin.Context) {
	deleteCartService := servicecart.CartService{}
	if err := ctx.ShouldBind(&deleteCartService); err == nil {
		res := deleteCartService.Delete(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
}
