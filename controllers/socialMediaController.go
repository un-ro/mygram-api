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

// CreateSocialMedia godoc
// @Summary Create Social media
// @Description Create a new social media for a user
// @Tags Social Media
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param SocialMedia body models.RequestSocialMedia true "create social media"
// @Success 201 {object} models.SocialMedia
// @Failure 400 {object} models.ResponseFailed
// @Failure 401 {object} models.ResponseFailedUnauthorized
// @Router /social-media [post]
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

// GetAllSocialMedia godoc
// @Summary Get details of all social media
// @Description Get details of all social media or add query parameter user_id for all social media from user_id (optional)
// @Tags Social Media
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_id query string false "Get all social media filter by user_id"
// @Success 200 {object} models.SocialMedia
// @Failure 400 {object} models.ResponseFailed
// @Failure 401 {object} models.ResponseFailedUnauthorized
// @Failure 404 {object} models.ResponseFailed
// @Router /social-media [get]
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

// GetOneSocialMedia godoc
// @Summary Get details for a given socialMediaID
// @Description Get details of social media corresponding to the input socialMediaID
// @Tags Social Media
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param socialMediaID path integer true "ID of the social media"
// @Success 200 {object} models.SocialMedia
// @Failure 400 {object} models.ResponseFailed
// @Failure 401 {object} models.ResponseFailedUnauthorized
// @Failure 404 {object} models.ResponseFailed
// @Router /social-media/{socialMediaID} [get]
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

// UpdateSocialMedia godoc
// @Summary Updated data social media
// @Description Update data social media by id
// @Tags Social Media
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param socialMediaID path integer true "socialMediaID of the data social media to be updated"
// @Param SocialMedia body models.RequestSocialMedia true "updated social media"
// @Success 200 {object} models.SocialMedia
// @Failure 400 {object} models.ResponseFailed
// @Failure 401 {object} models.ResponseFailedUnauthorized
// @Failure 404 {object} models.ResponseFailed
// @Router /social-media/{socialMediaID} [put]
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

// DeleteSocialMedia godoc
// @Summary Delete data social media
// @Description Delete data social media by id
// @Tags Social Media
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param socialMediaID path integer true "socialMediaID of the data social media to be deleted"
// @Success 200 {object} models.SocialMedia
// @Failure 400 {object} models.ResponseFailed
// @Failure 401 {object} models.ResponseFailedUnauthorized
// @Failure 404 {object} models.ResponseFailed
// @Router /social-media/{socialMediaID} [delete]
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
