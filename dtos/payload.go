package dtos

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                   uint           `json:"id"`
	Name                 string         `json:"name" validate:"required"`
	Email                string         `json:"email" validate:"required,email"`
	Nip                  string         `json:"nip"`
	Address              string         `json:"address"`
	PhoneNumber          string         `json:"phone_number"`
	Password             *string        `json:"password,omitempty" validate:"required,eqfield=PasswordConfirmation"`
	PasswordConfirmation *string        `gorm:"-" json:"password_confirmation,omitempty" validate:"required"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
	EmailVerifiedAt      time.Time      `json:"email_verified_at"`
	DeletedAt            gorm.DeletedAt `json:"deleted_at"`
}

type UpdateUser struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type UpdateUserPassword struct {
	CurrentPassword      *string `json:"current_password,omitempty" validate:"required"`
	Password             *string `json:"password,omitempty" validate:"required,nefield=CurrentPassword,eqfield=PasswordConfirmation"`
	PasswordConfirmation *string `json:"password_confirmation,omitempty" validate:"required"`
}

type Login struct {
	Email    string  `json:"email" validate:"required,email"`
	Password *string `json:"password,omitempty" validate:"required"`
}
