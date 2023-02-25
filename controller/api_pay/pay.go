package apipay

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanyewudezhuzi/E-COMMERCE/pkg/util"
	servicepay "github.com/sanyewudezhuzi/E-COMMERCE/service/service_pay"
)

func OrderPay(ctx *gin.Context) {
	claims, ok := ctx.Get("claims")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, "Failed to get claims.")
	}
	var orderPay servicepay.OrderPayService
	if err := ctx.ShouldBind(&orderPay); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		res := orderPay.Pay(ctx.Request.Context(), claims.(*util.Claims).ID)
		ctx.JSON(http.StatusOK, res)
	}
}
