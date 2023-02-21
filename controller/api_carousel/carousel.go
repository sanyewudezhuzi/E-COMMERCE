package apicarousel

import (
	"net/http"

	"github.com/gin-gonic/gin"
	servicecarousel "github.com/sanyewudezhuzi/E-COMMERCE/service/service_carousel"
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
