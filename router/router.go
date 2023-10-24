package router

import (
	"backend-test/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(r *gin.Engine) *gin.Engine {

	users := r.Group("users")
	{
		users.GET("/", controllers.GetUsers)
		users.GET("/:userInfo", controllers.GetUsersByNameOrCpf)
		users.POST("/", controllers.CreateUser)
		users.PUT("/:userId", controllers.UpdateUser)
		users.DELETE("/:userId", controllers.DeleteUser)
	}

	ceps := r.Group("ceps")
	{
		ceps.GET("/:cep", controllers.GetCepData)
	}

	return r
}

func SetupRouter() *gin.Engine {

	r := gin.Default()

	return ConfigRoutes(r)
}
