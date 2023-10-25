package models

import (
	"backend-test/database"
)

func MigrationModels() {
	database.DB.Migrator().DropTable(&User{}, &Address{}, &UF{})
	database.DB.AutoMigrate(&User{}, &Address{}, &UF{})
	PopulateUfs()
}
