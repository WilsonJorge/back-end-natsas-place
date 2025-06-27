package models

import (
	"time"

	"gorm.io/gorm"
)

type Visit struct {
	gorm.Model
	UserID        uint      `json:"user_id"`
	User          User      `json:"user" gorm:"foreignKey:UserID"`
	ScheduledDate time.Time `json:"scheduled_date" gorm:"not null"`
	Status        string    `json:"status"` // "requested", "scheduled", "completed", "cancelled", "no_show"
	Notes         string    `json:"notes" gorm:"type:text"`
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
}
