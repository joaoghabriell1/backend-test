package controllers

import (
	"backend-test/helpers"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCepData(c *gin.Context) {
	response, err := http.Get("https://viacep.com.br/ws/01001000/json/")
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err)
		return
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	helpers.JSON(c, http.StatusOK, string(responseData))
}
