package controllers

import (
	"backend-test/helpers"
	"encoding/json"
	"fmt"
	"io"
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
	SIAFI       string `json:"siafi,omitempty"`
}

func GetCepData(c *gin.Context) {

	cep := c.Param("cep")

	var CepResponse Cep

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
