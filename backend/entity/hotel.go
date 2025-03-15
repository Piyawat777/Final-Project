package entity

import (

	"gorm.io/gorm"
)

// Hotel represents a hotel entity.
type Hotel struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Location string
	Phone    string
	Email    string
	// One hotel has many rooms.
	Rooms []Room `gorm:"foreignKey:HotelID"`
	// Many-to-many relationship with Amenity.
	Amenities []Amenity `gorm:"many2many:hotel_amenities;"`
}