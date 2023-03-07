package lib

import (
	"time"

	"github.com/google/uuid"
)

type CreateDeviceRequest struct {
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

type UpdateDeviceRequest struct {
	Owner   string  `json:"owner"`
	Name    string  `json:"name"`
	PIN     string  `json:"pin"`
	TempSup float64 `json:"temp_sup"`
	TempMid float64 `json:"temp_mid"`
	TempSub float64 `json:"temp_sub"`
}

type Device struct {
	ID       uuid.UUID `json:"id"`
	Owner    string    `json:"owner"`
	Name     string    `json:"name"`
	PIN      string    `json:"pin"`
	TempSup  float64   `json:"temp_sup"`
	TempMid  float64   `json:"temp_mid"`
	TempSub  float64   `json:"temp_sub"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

func NewDevice(owner, name, pin string, sup, mid, sub float64) *Device {
	return &Device{
		ID:       uuid.New(),
		Owner:    owner,
		Name:     name,
		PIN:      pin,
		TempSup:  sup,
		TempMid:  mid,
		TempSub:  sub,
		Created:  time.Now().UTC(),
		Modified: time.Now().UTC(),
	}
}
