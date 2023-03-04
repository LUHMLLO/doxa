package lib

import (
	"time"

	"github.com/google/uuid"
)

type CreateDeviceRequest struct {
	JWT     string  `json:"jwt"`
	PIN     string  `json:"pin"`
	Owner   string  `json:"owner"`
	Name    string  `json:"name"`
	TempSup float64 `json:"temp_sup"`
	TempMid float64 `json:"temp_mid"`
	TempSub float64 `json:"temp_sub"`
}

type Device struct {
	ID       uuid.UUID `json:"id"`
	JWT      string    `json:"jwt"`
	PIN      string    `json:"pin"`
	Owner    string    `json:"owner"`
	Name     string    `json:"name"`
	TempSup  float64   `json:"temp_sup"`
	TempMid  float64   `json:"temp_mid"`
	TempSub  float64   `json:"temp_sub"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

func NewDevice(jwt, pin, owner, name string, sup, mid, sub float64) *Device {
	return &Device{
		ID:       uuid.New(),
		JWT:      jwt,
		PIN:      pin,
		Owner:    owner,
		Name:     name,
		TempSup:  sup,
		TempMid:  mid,
		TempSub:  sub,
		Created:  time.Now().UTC(),
		Modified: time.Now().UTC(),
	}
}

type SigninDeviceRequest struct {
	Owner string `json:"username"`
	PIN   string `json:"password"`
}
type SigninDevice struct {
	Owner string `json:"username"`
	PIN   string `json:"password"`
}
