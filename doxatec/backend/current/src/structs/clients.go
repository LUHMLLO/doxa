package main

import (
	"time"

	"github.com/google/uuid"
)

type Clients struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}
