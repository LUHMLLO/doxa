package lib

import (
	"time"

	"github.com/google/uuid"
)

type CreateDeviceRequest struct {
	Owner   string  `json:"owner"`
	Name    string  `json:"name"`
	TempSup float64 `json:"temp_sup"`
	TempMid float64 `json:"temp_mid"`
	TempSub float64 `json:"temp_sub"`
}

type Device struct {
	ID       uuid.UUID `json:"id"`
	Owner    string    `json:"owner"`
	Name     string    `json:"name"`
	TempSup  float64   `json:"temp_sup"`
	TempMid  float64   `json:"temp_mid"`
	TempSub  float64   `json:"temp_sub"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

func NewDevice(owner, name string, sup, mid, sub float64) *Device {
	return &Device{
		ID:       uuid.New(),
		Owner:    owner,
		Name:     name,
		TempSup:  sup,
		TempMid:  mid,
		TempSub:  sub,
		Created:  time.Now().UTC(),
		Modified: time.Now().UTC(),
	}
}
