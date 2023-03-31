package main

import (
	"time"

	"github.com/google/uuid"
)

type Transfers struct {
	ID             uuid.UUID `json:"id"`
	Subscription   uuid.UUID `json:"subscription_id"`
	Amount         float64   `json:"amount"`
	InitialBalance float64   `json:"initial_balance"`
	FinalBalance   float64   `json:"final_balance"`
	Created        time.Time `json:"created"`
}
