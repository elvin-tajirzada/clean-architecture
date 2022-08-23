package models

import "time"

type Users struct {
	ID        uint8     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name" validate:"required"`
	Email     string    `json:"email" db:"email" validate:"required,email"`
	Password  string    `json:"password" db:"password" validate:"required,min=8,max=20"`
	CreatedAt time.Time `json:"created_at" db:"created_at" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" validate:"required"`
}
