package apicarousel

import (
	"net/http"

	"github.com/gin-gonic/gin"
	servicecarousel "github.com/sanyewudezhuzi/E-COMMERCE/service/service_carousel"
	serviceproduct "github.com/sanyewudezhuzi/E-COMMERCE/service/service_product"
)

func ListCarousel(ctx *gin.Context) {
	var listCarousel servicecarousel.CarouselService
	if err := ctx.ShouldBind(&listCarousel); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		res := listCarousel.Show(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	}
}

func ProductList(ctx *gin.Context) {
	var productList serviceproduct.ProductService
	if err := ctx.ShouldBind(&productList); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		res := productList.ProductList(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	}
}

func ProductSearch(ctx *gin.Context) {
	var productSearch serviceproduct.ProductService
	if err := ctx.ShouldBind(&productSearch); err != nil {
		ctx.JSON(http.StatusOK, err)
	} else {
		res := productSearch.ProductSearch(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	}
}

func ProductShow(ctx *gin.Context) {
	var productShow serviceproduct.ProductService
	if err := ctx.ShouldBind(&productShow); err != nil {
		ctx.JSON(http.StatusOK, err)
	} else {
		res := productShow.ProductShow(ctx.Request.Context(), ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	}
}

func ProductImg(ctx *gin.Context) {
	var productImg serviceproduct.ProductService
	if err := ctx.ShouldBind(&productImg); err != nil {
		ctx.JSON(http.StatusOK, err)
	} else {
		res := productImg.ProductImg(ctx.Request.Context(), ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	}
}

func ProductCategory(ctx *gin.Context) {
	var productCategory serviceproduct.ProductService
	if err := ctx.ShouldBind(&productCategory); err != nil {
		ctx.JSON(http.StatusOK, err)
	} else {
		res := productCategory.ProductCategory(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	}
}
