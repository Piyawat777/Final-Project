package entity

import (
	"time"

	"gorm.io/gorm"
)

// Customer represents a guest with login functionality.
type Customer struct {
	gorm.Model
	// Login Credentials
	Username  string    `gorm:"uniqueIndex" valid:"required~Username is required"`               // Unique username for login
	Email     string    `gorm:"uniqueIndex" valid:"required~Email is required, email~Email is invalid"` // Email used for login
	Password  string    `valid:"required~Password is required"`                                  // Hashed password

	// Personal Information
	FullName  string    `valid:"required~FullName is required"`
	Phone     string    `valid:"required~Phone is required"`
	Address   string    `valid:"required~Address is required"`
	IDNumber  string    `valid:"required~ID/Passport number is required"` // Identification document

	// Additional Information
	LoyaltyInfo string  // Optional: Loyalty program details
	LastLogin   *time.Time // Records the last login time

	// Relationships
	Bookings []Booking `gorm:"foreignKey:CustomerID"` // One customer can have many bookings
	Reviews  []Review  `gorm:"foreignKey:CustomerID"` // One customer can leave many reviews
}
