package entities

import (
	"time"

	"github.com/google/uuid"
)

type Device struct {
	ID       uuid.UUID     `json:"id"`
	User     uuid.UUID     `json:"user_id"`
	Name     string        `json:"name"`
	PIN      string        `json:"pin"`
	TEMPS    []Temperature `json:"temps"`
	Created  time.Time     `json:"created"`
	Modified time.Time     `json:"modified"`
}
type NewDevice struct {
	User uuid.UUID `json:"user_id"`
	Name string    `json:"name"`
	PIN  string    `json:"pin"`
}
