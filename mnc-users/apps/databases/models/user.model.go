package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	ID          uuid.UUID  `json:"user_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	PhoneNumber string     `json:"phone_number" gorm:"unique;not null"`
	Address     string     `json:"address"`
	PIN         string     `json:"-"`
	CreatedAt   time.Time  `json:"created_date"`
}

func (user *Users) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	return nil
}