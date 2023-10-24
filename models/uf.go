package models

type UF struct {
	ID uint   `gorm:"primaryKey"`
	UF string `gorm:"type:varchar(2)" json:"uf" binding:"required"`
}
