package middlewares

import (
	"MyGram/database"
	"MyGram/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
)

func SocialMediaAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var socialMedia models.SocialMedia
		db := database.GetDB()

		socialMediaID, err := strconv.Atoi(ctx.Param("socialMediaID"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid parameter",
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		err = db.Debug().Select("user_id").First(&socialMedia, uint(socialMediaID)).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Social media doesn't exist",
			})
			return
		}

		if socialMedia.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data social media",
			})
			return
		}

		ctx.Next()
	}
}
