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
