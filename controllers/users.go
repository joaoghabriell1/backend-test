package controllers

import (
	"backend-test/database"
	"backend-test/helpers"
	"backend-test/models"
	"backend-test/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateNewUser(c *gin.Context) {

	var User models.User

	if err := c.BindJSON(&User); err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err)
		return
	}

	repo := repositories.NewUserRepository(database.DB)

	err := repo.CreateNewUser(&User)

	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(c, http.StatusNoContent, User)
}

func UpdateUserInfo(c *gin.Context) {

	var User models.User

	userId, err := strconv.Atoi(c.Param("userId"))

	if err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err)
		return
	}

	if err := c.BindJSON(&User); err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err)
		return
	}

	User.ID = uint(userId)

	repo := repositories.NewUserRepository(database.DB)

	err = repo.UpdateUserInfo(&User)

	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err)
		return
	}

}

func GetAllUsers(c *gin.Context) {

}

func GetUser(c *gin.Context) {

}
