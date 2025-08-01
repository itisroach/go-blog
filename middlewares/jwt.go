package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/itisroach/go-blog/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("authorization")

		if !strings.HasPrefix(authHeader,"Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer format should be used"})
			ctx.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.VerifyToken(token)

		if err != nil {
			ctx.JSON(err.Code, gin.H{"error": err.Message})
			ctx.Abort()
			return
		}

		ctx.Set("user", claims.Username)

		ctx.Next()
	}
}