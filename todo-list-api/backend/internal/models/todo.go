package models

import (
	"time"
)

type Todo struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"varchar(300)" json:"title"`
	Description string    `gorm:"varchar(300)" json:"description"`
	Completed   bool      `gorm:"default:false" json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      int       `gorm:"not null" json:"user_id"`
	//User        User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
}
