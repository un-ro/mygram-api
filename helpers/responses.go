package helpers

import "github.com/gin-gonic/gin"

func ErrorResponse(ctx *gin.Context, code int, message string) {
	ctx.AbortWithStatusJSON(code, gin.H{
		"status":  "error",
		"message": message,
	})
}

func SuccessResponse(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(code, gin.H{
		"status": "success",
		"data":   data,
	})
}

func LoginResponse(ctx *gin.Context, code int, token string) {
	ctx.JSON(code, gin.H{
		"status": "success",
		"token":  token,
	})
}
