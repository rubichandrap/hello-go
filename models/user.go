package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              uint           `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name            string         `gorm:"not null" json:"name" validate:"required"`
	Email           string         `gorm:"unique;not null" json:"email" validate:"required,email"`
	Nip             string         `json:"nip"`
	Address         string         `gorm:"type:text" json:"address"`
	PhoneNumber     string         `json:"phone_number"`
	Password        *string        `gorm:"not null" json:"password,omitempty" validate:"required,eqfield=PasswordConfirmation"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	EmailVerifiedAt time.Time      `json:"email_verified_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
