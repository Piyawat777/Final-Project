package entity

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	CustomerID  uint      // FK: refers to Customer.ID
	Customer    Customer  `gorm:"foreignKey:CustomerID"`
	RoomID      uint      // FK: refers to Room.ID
	Room        Room      `gorm:"foreignKey:RoomID"`
	PromotionID *uint     // FK: refers to Promotion.ID
	Promotion   Promotion `gorm:"foreignKey:PromotionID"`
	CheckInDate  time.Time // Date of check-in
	CheckOutDate time.Time // Date of check-out
	BookingDate  time.Time // Date when booking was made
	Status       string    // e.g., Confirmed, Pending Payment, Cancelled
	TotalAmount  float64
	// One booking is linked to one invoice.
	Invoice Invoice `gorm:"foreignKey:BookingID"`
	// One booking may have one payment record.
	Payment Payment `gorm:"foreignKey:BookingID"`
	// Optional review can be linked to a booking.
	Review Review `gorm:"foreignKey:BookingID"`
}
