package main

import (
	"time"

	"github.com/google/uuid"
)

type Devices struct {
	ID       uuid.UUID      `json:"id"`
	User     uuid.UUID      `json:"user_id"`
	Name     string         `json:"name"`
	PIN      string         `json:"pin"`
	TEMPS    []Temperatures `json:"temps"`
	Created  time.Time      `json:"created"`
	Modified time.Time      `json:"modified"`
}
