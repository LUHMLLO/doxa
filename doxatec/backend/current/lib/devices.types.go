package lib

import (
	"time"

	"github.com/google/uuid"
)

type CreateDeviceRequest struct {
	Owner   string `json:"owner"`
	Name    string `json:"name"`
	TempSup string `json:"temp_sup"`
	TempMid string `json:"temp_mid"`
	TempSub string `json:"temp_sub"`
}

type Device struct {
	ID       uuid.UUID `json:"id"`
	Owner    string    `json:"owner"`
	Name     string    `json:"name"`
	TempSup  string    `json:"temp_sup"`
	TempMid  string    `json:"temp_mid"`
	TempSub  string    `json:"temp_sub"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

func NewDevice(owner, name, sup, mid, sub string) *Device {
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
