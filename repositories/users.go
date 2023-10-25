package repositories

import (
	"backend-test/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateNewUser(u *models.User) error {

	u.CleanInputs()
	err := u.ValidateCPF()

	if err != nil {
		return errors.New("O CPF informado é inválido.")
	}

	UF, err := models.GetUfById(u.Address.UFID)

	if err != nil {
		return errors.New("UF Id inválido.")
	}

	u.Address.UF = UF

	err = r.DB.Create(&u).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UpdateUser(u *models.User) error {

	u.CleanInputs()
	err := u.ValidateCPF()

	if err != nil {
		return errors.New("O CPF informado é inválido.")
	}

	UF, err := models.GetUfById(u.Address.UFID)

	if err != nil {
		return errors.New("UF Id inválido.")
	}

	u.Address.UF = UF

	if u.Address.ID == 0 {
		return errors.New("Informe o ID do endereço.")
	}

	Address, err := models.GetAddressById(int(u.Address.ID))

	if err != nil {
		return errors.New("Não foi possível encontrar endereço com o ID informado.")
	}

	if Address.UserID != u.ID {
		return errors.New("O ID de endereço fornecido não pertence ao usuário.")
	}

	err = r.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&u).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUserById(u *models.User, uId uint) error {

	err := r.DB.Preload("Address").Preload("Address.UF").First(&u, uId).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUser(u *[]models.User, userInfo string) error {

	err := r.DB.Preload("Address").Preload("Address.UF").Where("users.cpf LIKE ?", fmt.Sprintf("%%%s%%", userInfo)).Or("users.nome LIKE ?", fmt.Sprintf("%%%s%%", userInfo)).Find(&u).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUsers(u *[]models.User) error {

	err := r.DB.Preload("Address").Preload("Address.UF").Preload("Address.CEP").Find(&u).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) DeleteUser(userId uint) error {

	var User models.User

	err := r.DB.First(&User, userId).Error

	if err != nil {
		return err
	}

	err = r.DB.Select(clause.Associations).Omit("User.Adress.UF").Delete(&User).Error

	if err != nil {
		return err
	}

	return nil
}
