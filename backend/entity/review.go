package entity

import (

	"gorm.io/gorm"
)

type Review struct {
    gorm.Model
    BookingID  uint  `gorm:"not null"`  // FK
    CustomerID uint  `gorm:"not null"`  // FK
    Rating     int   `gorm:"not null;check:rating >= 1 AND rating <= 5"`  
    Comment    string
    Reply      string  
}
