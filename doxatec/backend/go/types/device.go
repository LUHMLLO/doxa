package types

import (
	"time"

	"github.com/google/uuid"
)

type Device struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Owner   string    `json:"owner"`
	TempSup float64   `json:"temp_sup"`
	TempMid float64   `json:"temp_mid"`
	TempSub float64   `json:"temp_sub"`
	Date    time.Time `json:"date"`
}

func ValidateDevice(d *Device) bool { return true }
