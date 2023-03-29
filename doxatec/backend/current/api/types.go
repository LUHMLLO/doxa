package main

import (
	"time"

	"github.com/google/uuid"
)

type Clients struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

type Users struct {
	ID       uuid.UUID `json:"id"`
	Client   uuid.UUID `json:"client_id"`
	Avatar   string    `json:"avatar"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	Accessed time.Time `json:"accessed"`
}

type Devices struct {
	ID       uuid.UUID      `json:"id"`
	User     uuid.UUID      `json:"user_id"`
	Name     string         `json:"name"`
	PIN      string         `json:"pin"`
	TEMPS    []Temperatures `json:"temps"`
	Created  time.Time      `json:"created"`
	Modified time.Time      `json:"modified"`
}

type Temperatures struct {
	ID      uuid.UUID `json:"id"`
	Device  uuid.UUID `json:"device_id"`
	TempSup float64   `json:"temp_sup"`
	TempMid float64   `json:"temp_mid"`
	TempSub float64   `json:"temp_sub"`
	Created time.Time `json:"created"`
}

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

type Transfers struct {
	ID             uuid.UUID `json:"id"`
	Subscription   uuid.UUID `json:"subscription_id"`
	Amount         float64   `json:"amount"`
	InitialBalance float64   `json:"initial_balance"`
	FinalBalance   float64   `json:"final_balance"`
	Created        time.Time `json:"created"`
}
