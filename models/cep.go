package models

type CEP struct {
	ID  uint `gorm:"primaryKey"`
	CEP string
}
