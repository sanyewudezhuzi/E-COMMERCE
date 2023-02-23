package apiproduct

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sanyewudezhuzi/E-COMMERCE/pkg/util"
	serviceproduct "github.com/sanyewudezhuzi/E-COMMERCE/service/service_product"
)

func CreateProduct(ctx *gin.Context) {
	claims, ok := ctx.Get("claims")
	if !ok || claims == nil {
		ctx.JSON(http.StatusInternalServerError, "Failed to get claims.")
	}
	form, _ := ctx.MultipartForm()
	files := form.File["file"]
	var createProduct serviceproduct.ProductService
	if err := ctx.ShouldBind(&createProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		res := createProduct.Create(ctx, claims.(*util.Claims).ID, files)
		ctx.JSON(http.StatusOK, res)
	}
}
