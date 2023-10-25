package models

type Cep struct {
	Cep         string `json:"cep,omitempty"`
	Logradouro  string `json:"logradouro,omitempty"`
	Complemento string `json:"complemento,omitempty"`
	Bairro      string `json:"bairro,omitempty"`
	Localidade  string `json:"localidade,omitempty"`
	UF          string `json:"uf,omitempty"`
	IBGE        string `json:"ibge,omitempty"`
	GIA         string `json:"gia,omitempty"`
	DDD         string `json:"ddd,omitempty"`
	SIAFI       string `json:"siafi,omitempty"`
}
