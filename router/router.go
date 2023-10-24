package router

import (
	"backend-test/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(r *gin.Engine) *gin.Engine {

	user := r.Group("user")
	{
		user.GET("/", controllers.GetAllUsers)
		user.GET("/:userInfo", controllers.GetUsersByNameOrCpf)
		user.POST("/", controllers.CreateNewUser)
		user.PUT("/:userId", controllers.UpdateUserInfo)
		user.DELETE("/:userId", controllers.DeleteUser)
	}

	return r
}

func SetupRouter() *gin.Engine {

	r := gin.Default()

	return ConfigRoutes(r)
}
