package models

import "backend-test/database"

type Address struct {
	ID          uint   `gorm:"primaryKey"`
	Rua         string `json:"rua" binding:"required"`
	Numero      string `json:"numero" binding:"required"`
	Bairro      string `json:"bairro" binding:"required"`
	Complemento string `json:"complemento"`
	Cidade      string `json:"cidade" binding:"required"`
	CEP         string `json:"cep" binding:"required"`
	UserID      uint
	UFID        uint
	UF          UF
}

func GetAddressById(id int) (Address, error) {
	var Address Address

	err := database.DB.First(&Address, id).Error

	if err != nil {
		return Address, err
	}

	return Address, nil
}
