package routers

import (
	"MyGram/controllers"
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
		photoRouter.POST("/post", controllers.CreatePhoto)
		photoRouter.GET("/getAll", controllers.GetAllPhotos)
		photoRouter.GET("/getOne/:photoID", controllers.GetPhoto)
		photoRouter.PUT("/update/:photoID", controllers.UpdatePhoto)
		photoRouter.DELETE("/delete/:photoID", controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comment")
	{
		commentRouter.POST("/create", controllers.CreateComment)
		commentRouter.GET("/getAll", controllers.GetAllComment)
		commentRouter.GET("/getOne/:commentID", controllers.GetOneComment)
		commentRouter.PUT("/update/:commentID", controllers.UpdateComment)
		commentRouter.DELETE("/delete/:commentID", controllers.DeleteComment)
	}

	socialMediaRouter := r.Group("/social-media")
	{
		socialMediaRouter.POST("/create", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/getAll", controllers.GetAllSocialMedia)
		socialMediaRouter.GET("/getOne/:socialMediaID", controllers.GetOneSocialMedia)
		socialMediaRouter.PUT("/update/:socialMediaID", controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/delete/:socialMediaID", controllers.DeleteSocialMedia)
	}

	return r
}
