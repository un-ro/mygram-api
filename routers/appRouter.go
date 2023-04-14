package routers

import "github.com/gin-gonic/gin"

func StartServer() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/user")
	{
		userRouter.POST("/register")
		userRouter.POST("/login")
	}

	r.Static("/img", "./assets")
	photoRouter := r.Group("/photo")
	{
		photoRouter.POST("/post")
		photoRouter.GET("/getAll")
		photoRouter.GET("/getOne/:photoID")
		photoRouter.PUT("/update/:photoID")
		photoRouter.DELETE("/delete/:photoID")
	}

	commentRouter := r.Group("/comment")
	{
		commentRouter.POST("/create")
		commentRouter.GET("/getAll")
		commentRouter.GET("/getOne/:commentID")
		commentRouter.PUT("/update/:commentID")
		commentRouter.DELETE("/delete/:commentID")
	}

	socialMediaRouter := r.Group("/social-media")
	{
		socialMediaRouter.POST("/create")
		socialMediaRouter.GET("/getAll")
		socialMediaRouter.GET("/getOne/:socialMediaID")
		socialMediaRouter.PUT("/update/:socialMediaID")
		socialMediaRouter.DELETE("/delete/:socialMediaID")
	}

	return r
}
