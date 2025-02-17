package entity

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	BookingID   uint      // FK: refers to Booking.ID
	Details     string    // Breakdown of charges (could be JSON or plain text)
	TotalAmount float64
	InvoiceDate time.Time
	// Removed the Booking field to break the cyclic dependency.
}