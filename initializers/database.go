package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	DB, err = gorm.Open(postgres.Open("host=ep-proud-tooth-a2f52sxt.eu-central-1.aws.neon.tech user=petty-tracker_owner password=Z7jauN4GvmgR dbname=petty-tracker port=5432 sslmode=require"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to Database")
	}
}