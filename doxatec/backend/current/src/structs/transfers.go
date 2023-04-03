package entities

import (
	"time"

	"github.com/google/uuid"
)

type Transfer struct {
	ID           uuid.UUID `json:"id"`
	User         uuid.UUID `json:"user_id"`
	Subscription uuid.UUID `json:"subscription_id"`
	Amount       float64   `json:"amount"`
	Created      time.Time `json:"created"`
}
type NewTransfer struct {
	User         uuid.UUID `json:"user_id"`
	Subscription uuid.UUID `json:"subscription_id"`
	Amount       float64   `json:"amount"`
}
