package controllers

import (
	"backend-test/helpers"
	"backend-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNewUser(c *gin.Context) {

	var User models.User

	if err := c.BindJSON(&User); err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err)
		return
	}

	helpers.JSON(c, http.StatusNoContent, User)
}
