package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const jsonType = "application/json"

func GetHeader(ctx *gin.Context) string {
	return ctx.Request.Header.Get("Content-Type")
}

func BindRequest(ctx *gin.Context, data interface{}) {
	contentType := GetHeader(ctx)

	// Check if request is json or form-data
	if contentType == jsonType {
		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := ctx.ShouldBind(&data); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
	}
}

func ConvertKeyToInt(ctx *gin.Context, key string, message string) int {
	id, err := strconv.Atoi(ctx.Query(key))
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, message)
		return 0
	}

	return id
}
