package repositories

import (
	"backend-test/models"
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

	err := r.DB.Create(&u).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UpdateUser(u *models.User) error {

	err := r.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&u).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUser(u *[]models.User, userInfo string) error {

	err := r.DB.Preload("Address").Preload("Address.UF").Preload("Address.CEP").Where("users.cpf LIKE ?", fmt.Sprintf("%%%s%%", userInfo)).Or("users.nome LIKE ?", fmt.Sprintf("%%%s%%", userInfo)).Find(&u).Error

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
