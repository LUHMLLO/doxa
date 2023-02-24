package types

import (
	"time"

	"github.com/google/uuid"
)

type CreateDeviceRequest struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type Device struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	TempSup string    `json:"temp_sup"`
	TempMid string    `json:"temp_mid"`
	TempSub string    `json:"temp_sub"`

	Owner string `json:"owner"`

	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

func NewDevice(name, owner string) *Device {
	return &Device{
		ID:      uuid.New(),
		Name:    name,
		TempSup: "",
		TempMid: "",
		TempSub: "",
		Owner:   owner,

		Created:  time.Now().UTC(),
		Modified: time.Now().UTC(),
	}
}
