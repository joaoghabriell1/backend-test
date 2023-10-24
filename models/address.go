package models

type Address struct {
	ID          uint   `gorm:"primaryKey"`
	Rua         string `binding:"required" json:"rua"`
	Numero      string `binding:"required" json:"numero"`
	Bairro      string `binding:"required" json:"bairro"`
	Complemento string `json:"complemento"`
	UFID        uint
	UF          UF
	CEPID       uint
	CEP         CEP
	UserID      uint
}
