package models

type Address struct {
	Rua         string `binding:"required" json:"rua"`
	Numero      string `binding:"required" json:"numero"`
	Bairro      string `binding:"required" json:"bairro"`
	Complemento string `json:"complemento"`
	UFID        uint
	UF          UF
	CEPID       uint
	CEP         CEP
}

type UF struct {
	ID uint `gorm:"primaryKey"`
	UF string
}

type CEP struct {
	ID  uint `gorm:"primaryKey"`
	CEP string
}
