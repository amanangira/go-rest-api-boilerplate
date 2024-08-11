package dtos

import (
	"time"
)

type UserDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password,omitempty"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
