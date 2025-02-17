package entity

import (
	"time"

	"gorm.io/gorm"
)

type Maintenance struct {
	gorm.Model
	RoomID         uint      // FK: refers to Room.ID
	Room           Room      `gorm:"foreignKey:RoomID"`
	ReportDate     time.Time // Date reported
	Description    string
	Status         string      // e.g., Pending, In Progress, Completed
	CompletionDate *time.Time  // Optional: when the maintenance was finished
}