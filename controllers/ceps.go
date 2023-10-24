package controllers

import (
	"backend-test/helpers"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Cep struct {
	Cep         string `json:"cep,omitempty"`
	Logradouro  string `json:"logradouro,omitempty"`
	Complemento string `json:"compleme nto,omitempty"`
	Bairro      string `json:"bairro,omitempty"`
	Localidade  string `json:"localidade,omitempty"`
	UF          string `json:"uf,omitempty"`
	IBGE        string `json:"ibge,omitempty"`
	GIA         string `json:"gia,omitempty"`
	DDD         string `json:"ddd,omitempty"`
	SIAFI       string `json:"siafi omitempty"`
}

func GetCepData(c *gin.Context) {

	var Cep Cep

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
