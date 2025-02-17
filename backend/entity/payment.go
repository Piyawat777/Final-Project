package entity

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	BookingID     uint      // FK: refers to Booking.ID
	Amount        float64
	PaymentMethod string    // e.g., Credit Card, Bank Transfer, Cash
	PaymentDate   time.Time
	PaymentStatus string    // e.g., Completed, Pending, Failed
	// Removed the Booking field.
}