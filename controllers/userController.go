package controllers

import (
	"MyGram/database"
	"MyGram/helpers"
	"MyGram/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(ctx *gin.Context) {
	var user models.User

	db := database.GetDB()

	contentType := helpers.GetHeader(ctx)

	if contentType == helpers.JsonType {
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		}
	} else {
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		}
	}

	err := db.Debug().Create(&user).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"age":      user.Age,
		"email":    user.Email,
		"username": user.Username,
	})
}

func Login(ctx *gin.Context) {
	var user models.User

	db := database.GetDB()

	contentType := helpers.GetHeader(ctx)

	if contentType == helpers.JsonType {
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		}
	} else {
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		}
	}

	password := user.Password
	err := db.Debug().Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}

	compare := helpers.ComparePassword([]byte(user.Password), []byte(password))
	if !compare {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Password is wrong")
		return
	}

	token, err := helpers.GenerateToken(user.ID, user.Username)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
