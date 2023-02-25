package apifavorite

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanyewudezhuzi/E-COMMERCE/pkg/util"
	servicefavorite "github.com/sanyewudezhuzi/E-COMMERCE/service/service_favorite"
)

func FavoriteList(ctx *gin.Context) {
	claims, ok := ctx.Get("claims")
	if !ok || claims == nil {
		ctx.JSON(http.StatusInternalServerError, "Failed to get claims.")
	}
	var favoriteListService servicefavorite.FavoriteService
	if err := ctx.ShouldBind(&favoriteListService); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		res := favoriteListService.List(ctx.Request.Context(), claims.(*util.Claims).ID)
		ctx.JSON(http.StatusOK, res)
	}
}

func FavoriteCreate(ctx *gin.Context) {
	claims, ok := ctx.Get("claims")
	if !ok || claims == nil {
		ctx.JSON(http.StatusInternalServerError, "Failed to get claims.")
	}
	var favoriteCreateService servicefavorite.FavoriteService
	if err := ctx.ShouldBind(&favoriteCreateService); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		res := favoriteCreateService.Create(ctx.Request.Context(), claims.(*util.Claims).ID)
		ctx.JSON(http.StatusOK, res)
	}
}

func FavoriteDelete(ctx *gin.Context) {
	claims, ok := ctx.Get("claims")
	if !ok || claims == nil {
		ctx.JSON(http.StatusInternalServerError, "Failed to get claims.")
	}
	var favoriteDeleteService servicefavorite.FavoriteService
	if err := ctx.ShouldBind(&favoriteDeleteService); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		res := favoriteDeleteService.Delete(ctx.Request.Context(), claims.(*util.Claims).ID, ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	}
}
