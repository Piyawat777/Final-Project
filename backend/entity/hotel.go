package entity

import (

	"gorm.io/gorm"
)

// Hotel represents a hotel entity.
type Hotel struct {
	gorm.Model
	Name        string      // Hotel name
	Address     string      // Street address
	City        string
	State       string
	Phone       string
	Email       string
	Website     string
	Description string
	Rating      float32
	// One hotel has many rooms.
	Rooms []Room `gorm:"foreignKey:HotelID"`
	// Many-to-many relationship with Amenity.
	Amenities []Amenity `gorm:"many2many:hotel_amenities;"`
}