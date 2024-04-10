package models

import "gorm.io/gorm"

type Home struct {
	gorm.Model
	ID string
	Name string
	UserID string
}