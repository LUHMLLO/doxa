package main

import (
	"time"

	"github.com/google/uuid"
)

type Temperatures struct {
	ID      uuid.UUID `json:"id"`
	Device  uuid.UUID `json:"device_id"`
	TempSup float64   `json:"temp_sup"`
	TempMid float64   `json:"temp_mid"`
	TempSub float64   `json:"temp_sub"`
	Created time.Time `json:"created"`
}
