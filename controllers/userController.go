package controllers

import (
	"MyGram/database"
	"MyGram/helpers"
	"MyGram/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register godoc
// @Summary Register User
// @Description Register user for MyGram
// @Tags User
// @Accept json
// @Produce json
// @Param UserRegister body dto.RequestUserRegister true "User Register"
// @Success 200 {object} models.User
// @Failure 400 {object} dto.ResponseFailed
// @Router /user/register [post]
func Register(ctx *gin.Context) {
	var user models.User
	var checkUser models.User

	db := database.GetDB()

	// Check if request is json or form-data
	helpers.BindRequest(ctx, &user)

	// Check if username or email is already registered
	db.Debug().Find(&checkUser, "username = ?", user.Username)
	if checkUser.Username != "" {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "Username is already registered")
		return
	}

	db.Debug().Find(&checkUser, "email = ?", user.Email)
	if checkUser.Email != "" {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "Email is already registered")
		return
	}

	// Create user
	err := db.Debug().Create(&user).Error
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, http.StatusCreated, gin.H{
		"id":       user.ID,
		"age":      user.Age,
		"email":    user.Email,
		"username": user.Username,
	})
}

// Login godoc
// @Summary Login User
// @Description Login user to get token (jwt)
// @Tags User
// @Accept json
// @Produce json
// @Param UserLogin body dto.RequestUserLogin true "User Login"
// @Success 200 {object} models.User
// @Failure 400 {object} dto.ResponseFailed
// @Failure 401 {object} dto.ResponseFailed
// @Router /user/login [post]
func Login(ctx *gin.Context) {
	var user models.User

	db := database.GetDB()

	// Check if request is json or form-data
	helpers.BindRequest(ctx, &user)

	// Check if email is registered
	password := user.Password
	err := db.Debug().Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusUnauthorized, "Email is not registered")
		return
	}

	compare := helpers.ComparePassword([]byte(user.Password), []byte(password))
	if !compare {
		helpers.ErrorResponse(ctx, http.StatusUnauthorized, "Password is wrong")
		return
	}

	token, err := helpers.GenerateToken(user.ID, user.Username)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.LoginResponse(ctx, http.StatusOK, token)
}
