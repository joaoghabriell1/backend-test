package helpers

import "github.com/gin-gonic/gin"

func JSON(c *gin.Context, status int, res interface{}) {

	if res != nil {
		c.JSON(status, map[string]any{
			"data": res,
		})
	}
}

func Error(c *gin.Context, status int, err error) {
	JSON(c, status, err.Error())
}
