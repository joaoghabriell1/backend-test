package models

import (
	"backend-test/database"
	"time"
)

type User struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Nome             string    `gorm:"type:varchar(60)" json:"nome" binding:"required"`
	CPF              string    `gorm:"unique; not null; type:varchar(11)" json:"cpf" binding:"required"`
	DataDeNascimento time.Time `json:"data_de_nascimento"`
	Telefone         string    `gorm:"type:varchar(20)" json:"telefone" binding:"required"`
	Email            string    `gorm:"type:varchar(30)" json:"email" binding:"required,email"`
	Address          Address
}

func (u *User) ValidateCpf() {

}

func MigrateModels() {
	database.DB.AutoMigrate(&User{}, &Address{}, &CEP{}, &UF{})
}
