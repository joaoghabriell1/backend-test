package models

import (
	"backend-test/database"
)

func MigrationModels() {
	database.DB.AutoMigrate(&User{}, &Address{}, &UF{})
	PopulateUfs()
}
