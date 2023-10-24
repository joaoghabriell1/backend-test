package router

import "github.com/gin-gonic/gin"

func ConfigRoutes(r *gin.Engine) *gin.Engine {

	user := r.Group("user")
	{
		user.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, map[string]any{
				"res": "oi",
			})
		})
	}

	return r
}

func SetupRouter() *gin.Engine {

	r := gin.Default()

	return ConfigRoutes(r)
}
