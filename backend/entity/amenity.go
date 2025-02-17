package entity

import (

	"gorm.io/gorm"
)

type Amenity struct {
	gorm.Model
	Name        string  // e.g., Spa, Breakfast, Shuttle
	Description string
	Price       float64
	Category    string
	Status      string  // Active, Unavailable, etc.
	// Many-to-many relationship with Hotel.
	Hotels []Hotel `gorm:"many2many:hotel_amenities;"`
}