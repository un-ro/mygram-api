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

func CreateComment(ctx *gin.Context) {
	var comment models.Comment
	var photo models.Photo

	db := database.GetDB()

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	photoID := helpers.ConvertKeyToInt(ctx, "photoID", "Invalid photo id")

	helpers.BindRequest(ctx, &comment)

	err := db.Debug().Where("id = ?", photoID).First(&photo).Error
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusNotFound, "Photo not found")
		return
	}

	comment.UserID = userID
	comment.PhotoID = uint(photoID)

	err = db.Debug().Create(&comment).Error
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, http.StatusCreated, comment)
}

func GetAllComment(ctx *gin.Context) {
	var comments []models.Comment
	var photo models.Photo

	db := database.GetDB()
	photoID := helpers.ConvertKeyToInt(ctx, "photoID", "Invalid photo id")

	if photoID != 0 {
		err := db.Debug().Where("id = ?", photoID).First(&photo).Error
		if err != nil {
			helpers.ErrorResponse(ctx, http.StatusNotFound, "Photo not found")
			return
		}

		err = db.Debug().Order("id").Where("photo_id = ?", photoID).Find(&comments).Error
		if err != nil {
			helpers.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}

		if len(comments) == 0 {
			helpers.SuccessResponse(ctx, http.StatusOK, "No comments found")
			return
		}
	} else {
		err := db.Debug().Order("id").Find(&comments).Error
		if err != nil {
			helpers.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
	}

	helpers.SuccessResponse(ctx, http.StatusOK, comments)
}

func GetOneComment(ctx *gin.Context) {
	var comment models.Comment

	db := database.GetDB()
	commentID := helpers.ConvertKeyToInt(ctx, "commentID", "Invalid comment id")

	err := db.Debug().Where("id = ?", commentID).First(&comment).Error
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusNotFound, "Comment not found")
		return
	}

	helpers.SuccessResponse(ctx, http.StatusOK, comment)
}

func UpdateComment(ctx *gin.Context) {
	var comment, findComment models.Comment

	db := database.GetDB()
	commentID := helpers.ConvertKeyToInt(ctx, "commentID", "Invalid comment id")

	helpers.BindRequest(ctx, &comment)

	err := db.Debug().Where("id = ?", commentID).First(&findComment).Error
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusNotFound, fmt.Sprintf("Comment with id %d not found", commentID))
		return
	}

	comment = models.Comment{
		Message: comment.Message,
	}

	comment.ID = uint(commentID)
	comment.CreatedAt = findComment.CreatedAt
	comment.PhotoID = findComment.PhotoID
	comment.UserID = findComment.UserID

	err = db.Debug().Model(&comment).Where("id = ?", commentID).Updates(comment).Error
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, http.StatusOK, comment)
}

func DeleteComment(ctx *gin.Context) {
	var comment models.Comment

	db := database.GetDB()
	commentID := helpers.ConvertKeyToInt(ctx, "commentID", "Invalid comment id")

	err := db.Debug().Where("id = ?", commentID).First(&comment).Delete(&comment).Error
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusNotFound, fmt.Sprintf("Comment with id %d not found", commentID))
		return
	}

	helpers.SuccessResponse(ctx, http.StatusOK, fmt.Sprintf("Comment message '%s' successfully deleted", comment.Message))
}
