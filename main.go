package main

import (
	"backend-test/config"
	"backend-test/database"
	"backend-test/models"
	"backend-test/router"
	"fmt"
)

func init() {
	config.LoadEnvVariables()
	database.SetupDatabase()
	models.MigrateModels()
}

func main() {
	r := router.SetupRouter()

	fmt.Printf("Listening on Port: %d", config.PORT)

	r.Run()
}
