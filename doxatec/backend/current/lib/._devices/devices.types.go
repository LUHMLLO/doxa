package lib

import (
	"time"

	"github.com/google/uuid"
)

type CreateDeviceRequest struct {
	Owner string `json:"owner"`
	Name  string `json:"name"`
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

func NewDevice(owner, name string) *Device {
	return &Device{
		ID:       uuid.New(),
		Owner:    owner,
		Name:     name,
		PIN:      "",
		TempSup:  0,
		TempMid:  0,
		TempSub:  0,
		Created:  time.Now().UTC(),
		Modified: time.Now().UTC(),
	}
}
