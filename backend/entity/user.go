package entity

import (

	"gorm.io/gorm"
)

// Customer represents a guest with login functionality.
type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Username string 
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Phone    string
	Role     string  //(Admin, Customer, Staff)
	// GenderID ทำหน้าที่เป็น FK
	GenderID uint	
	Gender   Genders `gorm:"foreignKey:GenderID"`
}

