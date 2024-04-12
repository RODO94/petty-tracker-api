package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Home struct {
	gorm.Model
	ID uuid.UUID
	Name string
	UserID string
}