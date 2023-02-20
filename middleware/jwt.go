package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sanyewudezhuzi/E-COMMERCE/pkg/e"
	"github.com/sanyewudezhuzi/E-COMMERCE/pkg/util"
	"github.com/sanyewudezhuzi/E-COMMERCE/serializer"
)

// JWT token 中间件
func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := e.Success
		tokenStr := ctx.GetHeader("Authorization")
		if tokenStr == "" {
			code = e.StatusNotFound
		} else {
			claims, err := util.ParseToken(tokenStr)
			if err != nil {
				code = e.ErrorAuthToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			} else {
				ctx.Set("claims", claims)
			}
		}
		if code != e.Success {
			ctx.JSON(http.StatusNotFound, serializer.Response{
				StatusCode: code,
				Msg:        e.GetMsg(code),
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
