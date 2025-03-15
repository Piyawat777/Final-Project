package entity

import "gorm.io/gorm"

type Genders struct {
	gorm.Model
	Gender string

	User []User `gorm:"foreignKey:GenderID"`
}