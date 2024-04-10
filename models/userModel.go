package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uuid.UUID
	Name string
	Email string
	Password string
}