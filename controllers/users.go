package controllers

import (
	"backend-test/database"
	"backend-test/helpers"
	"backend-test/models"
	"backend-test/repositories"
	"fmt"
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

	helpers.JSON(c, http.StatusOK, User)
}

func GetAllUsers(c *gin.Context) {

	var Users []models.User

	repo := repositories.NewUserRepository(database.DB)

	err := repo.GetAllUsers(&Users)

	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(c, http.StatusOK, Users)
	return
}

func GetUsersByNameOrCpf(c *gin.Context) {

	var Users []models.User

	userInfo := c.Param("userInfo")

	repo := repositories.NewUserRepository(database.DB)
	fmt.Println(userInfo)

	err := repo.GetUser(&Users, userInfo)

	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(c, http.StatusOK, Users)
}

func DeleteUser(c *gin.Context) {

	userId, err := strconv.ParseUint(c.Param("userId"), 10, 64)

	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err)
		return
	}

	repo := repositories.NewUserRepository(database.DB)

	err = repo.DeleteUser(uint(userId))

	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(c, http.StatusOK, userId)
}
