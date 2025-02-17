package entity

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	CustomerID uint      // FK: refers to Customer.ID
	BookingID  uint      // FK: refers to Booking.ID
	Rating     int       // e.g., 1 to 5 stars
	Comments   string
	ReviewDate time.Time
	// Removed the Booking field.
}