package models

import (
	"backend-test/database"
	"errors"
	"fmt"
	"strings"

	"github.com/klassmann/cpfcnpj"
)

type User struct {
	ID               uint   `gorm:"primaryKey" json:"id"`
	Nome             string `gorm:"type:varchar(60)" json:"nome" binding:"required"`
	CPF              string `gorm:"unique; not null; type:varchar(11)" json:"cpf" binding:"required"`
	DataDeNascimento string `gorm:"type:varchar(10)" json:"data_de_nascimento"`
	Telefone         string `gorm:"type:varchar(20)" json:"telefone" binding:"required"`
	Email            string `gorm:"type:varchar(30); unique" json:"email" binding:"required,email"`
	Address          Address
}

func MigrateModels() {
	database.DB.AutoMigrate(&User{}, &Address{}, &UF{})
}

func (u *User) CleanInputs() {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Email = strings.TrimSpace(u.Email)
	u.Telefone = strings.TrimSpace(u.Telefone)
	u.Address.Rua = strings.TrimSpace(u.Address.Rua)
	u.Address.Numero = strings.TrimSpace(u.Address.Numero)
	u.Address.Bairro = strings.TrimSpace(u.Address.Bairro)
	u.Address.Complemento = strings.TrimSpace(u.Address.Complemento)
	u.Address.Cidade = strings.TrimSpace(u.Address.Cidade)
}

func (u *User) ValidateCPF() error {
	cpf := cpfcnpj.NewCPF(u.CPF)

	fmt.Println(cpf)

	if !cpf.IsValid() {
		return errors.New("CPF inválido.")
	}

	u.CPF = fmt.Sprintf("%v", cpf)

	return nil
}
