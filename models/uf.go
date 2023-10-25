package models

import (
	"backend-test/database"
	"errors"
)

type UF struct {
	ID    uint   `gorm:"primaryKey"`
	Nome  string `gorm:"type:varchar(30)"`
	Sigla string `gorm:"type:varchar(2)"`
}

func GetUfById(id uint) (UF, error) {

	var UF UF

	err := database.DB.First(&UF, id).Error

	if err != nil {
		return UF, errors.New("UF não encontrada.")
	}

	return UF, nil
}
