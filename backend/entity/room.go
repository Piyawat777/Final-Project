package entity

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	RoomNumber  string  // Room number/identifier
	HotelID     uint    // FK: refers to Hotel.ID
	Hotel       Hotel   `gorm:"foreignKey:HotelID"`
	RoomType    string  // e.g., Standard, Deluxe, Suite
	Capacity    int     // Number of guests that can be accommodated
	Price       float64 // Price per night
	Status      string  // Available, Booked, Under Cleaning, etc.
	Description string
	// One room can have multiple images.
	Images []RoomImage `gorm:"foreignKey:RoomID"`
	// One room can have many maintenance records.
	Maintenances []Maintenance `gorm:"foreignKey:RoomID"`
	// One room can have many bookings.
	Bookings []Booking `gorm:"foreignKey:RoomID"`
}

type RoomImage struct {
	gorm.Model
	RoomID uint   // FK: refers to Room.ID
	URL    string // URL of the image
}