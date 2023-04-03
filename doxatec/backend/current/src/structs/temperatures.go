package entities

import (
	"time"

	"github.com/google/uuid"
)

type Temperature struct {
	ID      uuid.UUID `json:"id"`
	Device  uuid.UUID `json:"device_id"`
	TempSup float64   `json:"temp_sup"`
	TempMid float64   `json:"temp_mid"`
	TempSub float64   `json:"temp_sub"`
	Created time.Time `json:"created"`
}
type NewTemperature struct {
	Device  uuid.UUID `json:"device_id"`
	TempSup float64   `json:"temp_sup"`
	TempMid float64   `json:"temp_mid"`
	TempSub float64   `json:"temp_sub"`
}
