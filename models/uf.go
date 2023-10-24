package models

type UF struct {
	ID    uint   `gorm:"primaryKey"`
	Nome  string `gorm:"type:varchar(30)"`
	Sigla string `gorm:"type:varchar(2)"`
}
