package database

import (
	"backend-test/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase() {

	var err error

	DB, err = gorm.Open(postgres.Open(config.DSN), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

}
