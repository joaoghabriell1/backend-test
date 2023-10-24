package repositories

import (
	"backend-test/models"

	"gorm.io/gorm"
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
