package routers

import (
	"MyGram/controllers"
	"MyGram/middlewares"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/user")
	{
		userRouter.POST("/register", controllers.Register)
		userRouter.POST("/login", controllers.Login)
	}

	r.Static("/img", "./assets")
	photoRouter := r.Group("/photo")
	{
		photoRouter.Use(middlewares.Auth())
		photoRouter.POST("/post", controllers.CreatePhoto)
		photoRouter.GET("/getAll", controllers.GetAllPhotos)
		photoRouter.GET("/getOne/:photoID", controllers.GetPhoto)
		photoRouter.PUT("/update/:photoID", middlewares.PhotoAuth(), controllers.UpdatePhoto)
		photoRouter.DELETE("/delete/:photoID", middlewares.PhotoAuth(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comment")
	{
		commentRouter.Use(middlewares.Auth())
		commentRouter.POST("/create", controllers.CreateComment)
		commentRouter.GET("/getAll", controllers.GetAllComment)
		commentRouter.GET("/getOne/:commentID", controllers.GetOneComment)
		commentRouter.PUT("/update/:commentID", middlewares.CommentAuth(), controllers.UpdateComment)
		commentRouter.DELETE("/delete/:commentID", middlewares.CommentAuth(), controllers.DeleteComment)
	}

	socialMediaRouter := r.Group("/social-media")
	{
		socialMediaRouter.Use(middlewares.Auth())
		socialMediaRouter.POST("/create", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/getAll", controllers.GetAllSocialMedia)
		socialMediaRouter.GET("/getOne/:socialMediaID", controllers.GetOneSocialMedia)
		socialMediaRouter.PUT("/update/:socialMediaID", middlewares.SocialMediaAuth(), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/delete/:socialMediaID", middlewares.SocialMediaAuth(), controllers.DeleteSocialMedia)
	}

	return r
}
