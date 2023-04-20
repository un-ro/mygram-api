package controllers

import (
	"MyGram/database"
	"MyGram/helpers"
	"MyGram/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

// CreatePhoto godoc
// @Summary Post Photo
// @Description Post a new Photo from user
// @Tags Photo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param PostPhoto body models.RequestPhoto true "Post photo"
// @Success 201 {object} models.Photo
// @Failure 400 {object} models.ResponseFailed
// @Failure 401 {object} models.ResponseFailedUnauthorized
// @Router /photos [post]
func CreatePhoto(ctx *gin.Context) {
	var photo models.Photo

	db := database.GetDB()

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	username := userData["username"].(string)

	_ = ctx.ShouldBind(&photo)
	file, err := ctx.FormFile("photo_url")
	if err == nil {
		charset := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
		stringRandom := make([]byte, rand.Intn(100))
		for i := range stringRandom {
			stringRandom[i] = charset[rand.Intn(len(charset))]
		}

		ext := strings.Split(file.Filename, ".")[1]

		log.Println("file ext ->", ext)

		if ext == "jpg" || ext == "jpeg" || ext == "png" || ext == "webp" {
			dst := "./img/" + username + "-" + string(stringRandom) + "." + "jpg"
			err := ctx.SaveUploadedFile(file, dst)
			if err != nil {
				helpers.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
				return
			}

			urlPhoto := fmt.Sprintf("http://%s/img/%s-%s.jpg", ctx.Request.Host, username, string(stringRandom))
			photo.PhotoUrl = urlPhoto
		} else {
			helpers.ErrorResponse(ctx, http.StatusBadRequest, "File is not image/photo")
			return
		}
	}

	photo.UserID = userID

	err = db.Debug().Create(&photo).Error
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, http.StatusCreated, photo)
}

// GetAllPhotos godoc
// @Summary Get details of All photo
// @Description Get details of all photo or add query parameter user_id for all photo from user_id (optional)
// @Tags Photo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_id query integer false "Get all photo filter by user_id"
// @Success 200 {object} models.Photo
// @Failure 400 {object} models.ResponseFailed
// @Failure 401 {object} models.ResponseFailedUnauthorized
// @Failure 404 {object} models.ResponseFailed
// @Router /photos [get]
func GetAllPhotos(ctx *gin.Context) {
	var photos []models.Photo

	db := database.GetDB()

	if _, ok := ctx.GetQuery("user_id"); ok {
		userId := helpers.ConvertKeyToInt(ctx, "user_id", "user_id must be integer")

		err := db.Debug().Preload("Comments").Order("id").Where("user_id = ?", userId).Find(&photos).Error
		if err != nil {
			helpers.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
	} else {
		err := db.Debug().Preload("Comments").Order("id").Find(&photos).Error
		if err != nil {
			helpers.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
	}

	if len(photos) == 0 {
		helpers.SuccessResponse(ctx, http.StatusOK, "No photos found")
		return
	}

	helpers.SuccessResponse(ctx, http.StatusOK, photos)
}

// GetPhoto godoc
// @Summary Get details of photo by id
// @Description Get details of photo by id
// @Tags Photo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param photoID path integer true "ID of the photo"
// @Success 200 {object} models.Photo
// @Failure 400 {object} models.ResponseFailed
// @Failure 401 {object} models.ResponseFailedUnauthorized
// @Failure 404 {object} models.ResponseFailed
// @Router /photos/{photoID} [get]
func GetPhoto(ctx *gin.Context) {
	var photo models.Photo

	db := database.GetDB()

	photoId := helpers.ConvertKeyToInt(ctx, "photoID", "photoID must be integer")

	err := db.Debug().Preload("Comments").Where("id = ?", photoId).First(&photo).Error
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, http.StatusOK, photo)
}

// UpdatePhoto godoc
// @Summary Update detail of photo by id
// @Description Update detail of photo by id
// @Tags Photo
// @Accept json
// @Produce json
// @Param photoID path integer true "photoID of the data photo to be updated"
// @Param UpdatePhoto body models.RequestPhoto true "Update photo"
// @Success 200 {object} models.Photo
// @Failure 400 {object} models.ResponseFailed
// @Failure 401 {object} models.ResponseFailedUnauthorized
// @Failure 404 {object} models.ResponseFailed
// @Router /photos/{photoID} [put]
func UpdatePhoto(ctx *gin.Context) {
	var photo, findPhoto models.Photo

	db := database.GetDB()

	photoId := helpers.ConvertKeyToInt(ctx, "photoID", "photoID must be integer")

	err := db.Debug().Where("id = ?", photoId).First(&findPhoto).Error
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.BindRequest(ctx, &photo)

	photo = models.Photo{
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: findPhoto.PhotoUrl,
	}

	photo.ID = uint(photoId)
	photo.CreatedAt = findPhoto.CreatedAt
	photo.UserID = findPhoto.UserID

	err = db.Debug().Model(&photo).Where("id = ?", photoId).Updates(photo).Error
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = db.Debug().Preload("Comments").Where("id = ?", photoId).First(&photo).Error
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, http.StatusOK, photo)
}

// DeletePhoto godoc
// @Summary Delete data photo by id
// @Description Delete data photo by id
// @Tags Photo
// @Accept json
// @Produce json
// @Security
// @Param photoID path integer true "photoID of the data photo to be deleted"
// @Success 200 {object} models.Photo
// @Failure 400 {object} models.ResponseFailed
// @Failure 401 {object} models.ResponseFailedUnauthorized
// @Failure 404 {object} models.ResponseFailed
// @Router /photos/{photoID} [delete]
func DeletePhoto(ctx *gin.Context) {
	var photo models.Photo

	db := database.GetDB()

	photoId := helpers.ConvertKeyToInt(ctx, "photoID", "photoID must be integer")

	err := db.Debug().Where("id = ?", photoId).First(&photo).Delete(&photo).Error
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, http.StatusOK, fmt.Sprintf("Photo with title '%s' successfully deleted", photo.Title))
}
