package main

import (
	"fmt"

	"github.com/RODO94/petty-tracker-api/initializers"
)

func init() {

	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
fmt.Println("Migration")
	// initializers.DB.AutoMigrate(&models.User{},&models.Home{},&models.Task{})
	type Task struct {
		Status string
	}
	initializers.DB.Migrator().AddColumn(&Task{}, "Status")}