package entities

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       float64    `json:"amount"`
	Transfers   []Transfer `json:"transfers"`
	Created     time.Time  `json:"created"`
	Modified    time.Time  `json:"modified"`
}
type NewSubscription struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"amount"`
}
