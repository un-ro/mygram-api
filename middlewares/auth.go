package middlewares

import (
	"MyGram/helpers"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := helpers.VerifyToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}

		ctx.Set("userData", token)

		ctx.Next()
	}
}
