package router

import (
	"backend-test/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(r *gin.Engine) *gin.Engine {

	user := r.Group("user")
	{
		user.POST("/", controllers.CreateNewUser)
		user.PUT("/:userId", controllers.UpdateUserInfo)
		user.GET("/", controllers.GetAllUsers)
		user.GET("/:info", controllers.GetUser)
	}

	return r
}

func SetupRouter() *gin.Engine {

	r := gin.Default()

	return ConfigRoutes(r)
}
