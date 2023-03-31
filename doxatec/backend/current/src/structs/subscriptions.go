package main

import (
	"time"

	"github.com/google/uuid"
)

type Subscriptions struct {
	ID              uuid.UUID   `json:"id"`
	User            uuid.UUID   `json:"user_id"`
	Transfers       []Transfers `json:"transfers"`
	Amount          float64     `json:"amount"`
	InterestRate    float64     `json:"interest_rate"`
	TermDuration    int         `json:"term_duration"`
	TermFrequency   string      `json:"term_frequency"`
	LoanStart       time.Time   `json:"loan_start"`
	LoanEnd         time.Time   `json:"loan_end"`
	TotalInterest   float64     `json:"total_interest"`
	TotalLoan       float64     `json:"total_loan"`
	MonthlyInterest float64     `json:"monthly_interest"`
	MonthylLoan     float64     `json:"monthly_loan"`
	MonthlyFee      float64     `json:"monthly_fee"`
	Approval        bool        `json:"approval"`
	Created         time.Time   `json:"created"`
	Modified        time.Time   `json:"modified"`
}
