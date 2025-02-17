package entity

import (

	"gorm.io/gorm"
)

type Staff struct {
	gorm.Model
	FullName    string
	Position    string
	Phone       string
	Email       string
	Username    string
	Password    string
	WorkSchedule string // Can be stored as a JSON string or separate table if needed.
}