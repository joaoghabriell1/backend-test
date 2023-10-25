package controllers

import (
	"backend-test/helpers"
	"backend-test/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCepData(c *gin.Context) {

	cep := c.Param("cep")

	var CepResponse models.Cep

	uri := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	response, err := http.Get(uri)

	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err)
		return
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err)
		return
	}

	if err := json.Unmarshal(responseData, &CepResponse); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(c, http.StatusOK, CepResponse)
}
