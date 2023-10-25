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

func GetAddressByUserId(id uint) (Address, error) {
	var Address Address

	err := database.DB.Where("user_id = ?", id).Find(&Address).Error

	if err != nil {
		return Address, err
	}

	return Address, nil
}
