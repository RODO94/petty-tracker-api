package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name string
	UserID string
	HomeID string
	Weighting int16
}