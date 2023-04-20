package routers

import (
	"MyGram/controllers"
	"MyGram/middlewares"

	_ "MyGram/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userRouter := r.Group("/user")
	{
		userRouter.POST("/register", controllers.Register)
		userRouter.POST("/login", controllers.Login)
	}

	r.Static("/img", "./assets")
	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Auth())
		photoRouter.POST("", controllers.CreatePhoto)
		photoRouter.GET("", controllers.GetAllPhotos)
		photoRouter.GET("/:photoID", controllers.GetPhoto)
		photoRouter.PUT("/:photoID", middlewares.PhotoAuth(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoID", middlewares.PhotoAuth(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Auth())
		commentRouter.POST("/:photoID", controllers.CreateComment)
		commentRouter.GET("/photo/:photoID", controllers.GetAllComment)
		commentRouter.GET("/:commentID", controllers.GetOneComment)
		commentRouter.PUT("/:commentID", middlewares.CommentAuth(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentID", middlewares.CommentAuth(), controllers.DeleteComment)
	}

	socialMediaRouter := r.Group("/social-media")
	{
		socialMediaRouter.Use(middlewares.Auth())
		socialMediaRouter.POST("", controllers.CreateSocialMedia)
		socialMediaRouter.GET("", controllers.GetAllSocialMedia)
		socialMediaRouter.GET("/:socialMediaID", controllers.GetOneSocialMedia)
		socialMediaRouter.PUT("/:socialMediaID", middlewares.SocialMediaAuth(), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaID", middlewares.SocialMediaAuth(), controllers.DeleteSocialMedia)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
