package entity

import (
	"time"

	"gorm.io/gorm"
)

type Promotion struct {
	gorm.Model
	Name         string
	Description  string
	DiscountRate float64    // Alternatively, a fixed DiscountAmount can be used.
	StartDate    time.Time
	EndDate      time.Time
	Conditions   string
	// One promotion can apply to many bookings.
	Bookings []Booking `gorm:"foreignKey:PromotionID"`
}