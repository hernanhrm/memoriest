package domain

import "time"

type GetAll struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetAlls []*GetAll
