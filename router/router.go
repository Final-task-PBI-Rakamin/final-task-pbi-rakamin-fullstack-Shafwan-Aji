package router

import (
	"myapp/controllers"
	"myapp/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", controllers.RegisterUser)
		userGroup.POST("/login", controllers.LoginUser)
		userGroup.PUT("/:userId", middlewares.AuthMiddleware(), controllers.UpdateUser)
		userGroup.DELETE("/:userId", middlewares.AuthMiddleware(), controllers.DeleteUser)
	}

	photoGroup := r.Group("/photos")
	{
		photoGroup.POST("", middlewares.AuthMiddleware(), controllers.CreatePhoto)
		photoGroup.GET("", middlewares.AuthMiddleware(), controllers.GetPhotos)
		photoGroup.PUT("/:photoId", middlewares.AuthMiddleware(), controllers.UpdatePhoto)
		photoGroup.DELETE("/:photoId", middlewares.AuthMiddleware(), controllers.DeletePhoto)
	}

	return r
}
