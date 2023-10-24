package models

type CEP struct {
	ID  uint   `gorm:"primaryKey"`
	CEP string `json:"cep" binding:"required"`
}
