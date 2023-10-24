package models

type Address struct {
	ID          uint   `gorm:"primaryKey"`
	Rua         string `json:"rua" binding:"required"`
	Numero      string `json:"numero" binding:"required"`
	Bairro      string `json:"bairro" binding:"required"`
	Complemento string `json:"complemento"`
	Cidade      string `json:"cidade" binding:"required"`
	UFID        uint
	UF          UF
	CEPID       uint
	CEP         CEP
	UserID      uint
}
