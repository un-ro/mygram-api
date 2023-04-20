package controllers

import (
	"MyGram/database"
	"MyGram/helpers"
	"MyGram/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func CreateSocialMedia(ctx *gin.Context) {
	var socialMedia models.SocialMedia

	db := database.GetDB()

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	helpers.BindRequest(ctx, &socialMedia)

	socialMedia.UserID = userID

	err := db.Debug().Create(&socialMedia).Error
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, http.StatusCreated, socialMedia)
}

func GetAllSocialMedia(ctx *gin.Context) {
	var socialMedia []models.SocialMedia

	db := database.GetDB()

	if _, ok := ctx.GetQuery("user_id"); ok {
		userId := helpers.ConvertKeyToInt(ctx, "user_id", "user_id must be integer")

		err := db.Debug().Order("id").Where("user_id = ?", userId).Find(&socialMedia).Error
		if err != nil {
			helpers.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}

		if len(socialMedia) == 0 {
			helpers.ErrorResponse(ctx, http.StatusBadRequest, fmt.Sprintf("user_id %d doesn't have social media", userId))
			return
		}
	} else {
		err := db.Debug().Order("id").Find(&socialMedia).Error
		if err != nil {
			helpers.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
	}

	helpers.SuccessResponse(ctx, http.StatusOK, socialMedia)
}

func GetOneSocialMedia(ctx *gin.Context) {
	var socialMedia models.SocialMedia

	db := database.GetDB()

	socialMediaID := helpers.ConvertKeyToInt(ctx, "socialMediaID", "socialMediaID must be integer")

	err := db.Debug().Where("id = ?", socialMediaID).First(&socialMedia).Error
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "Social media not found")
		return
	}

	helpers.SuccessResponse(ctx, http.StatusOK, socialMedia)
}

func UpdateSocialMedia(ctx *gin.Context) {
	var socialMedia, findSocialMedia models.SocialMedia

	db := database.GetDB()

	socialMediaID := helpers.ConvertKeyToInt(ctx, "socialMediaID", "socialMediaID must be integer")

	helpers.BindRequest(ctx, &socialMedia)

	err := db.Debug().Where("id = ?", socialMediaID).First(&findSocialMedia).Error
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "Social media not found")
		return
	}

	socialMedia = models.SocialMedia{
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
	}

	socialMedia.ID = uint(socialMediaID)
	socialMedia.CreatedAt = findSocialMedia.CreatedAt
	socialMedia.UserID = findSocialMedia.UserID

	err = db.Debug().Model(&socialMedia).Where("id = ?", socialMediaID).Updates(socialMedia).Error
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, http.StatusOK, socialMedia)
}

func DeleteSocialMedia(ctx *gin.Context) {
	var socialMedia models.SocialMedia

	db := database.GetDB()
	socialMediaID := helpers.ConvertKeyToInt(ctx, "socialMediaID", "socialMediaID must be integer")

	err := db.Debug().Where("id = ?", socialMediaID).First(&socialMedia).Delete(&socialMedia).Error
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, http.StatusOK, fmt.Sprintf("Social Media %s successfully deleted", socialMedia.Name))
}
