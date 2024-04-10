package main

import (
	"fmt"

	"github.com/RODO94/petty-tracker-api/initializers"
	"github.com/RODO94/petty-tracker-api/models"
)

func init() {

	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
fmt.Println("Migration")
	initializers.DB.AutoMigrate(&models.User{},&models.Home{},&models.Task{})

}