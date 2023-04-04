package app

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID       uuid.UUID  `json:"id"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Phone    string     `json:"phone"`
	Created  time.Time  `json:"created"`
	Modified *time.Time `json:"modified"`
}
type NewClient struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
type UpdateClient struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Modified time.Time `json:"modified"`
}

type User struct {
	ID       uuid.UUID  `json:"id"`
	Client   uuid.UUID  `json:"client_id"`
	Avatar   string     `json:"avatar"`
	Username string     `json:"username"`
	Password string     `json:"password"`
	Role     string     `json:"role"`
	Created  time.Time  `json:"created"`
	Modified *time.Time `json:"modified"`
	Accessed *time.Time `json:"accessed"`
}
type NewUser struct {
	Client   uuid.UUID `json:"client_id"`
	Avatar   string    `json:"avatar"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
}

type Device struct {
	ID       uuid.UUID     `json:"id"`
	User     uuid.UUID     `json:"user_id"`
	Name     string        `json:"name"`
	PIN      string        `json:"pin"`
	TEMPS    []Temperature `json:"temps"`
	Created  time.Time     `json:"created"`
	Modified *time.Time    `json:"modified"`
}
type NewDevice struct {
	User uuid.UUID `json:"user_id"`
	Name string    `json:"name"`
	PIN  string    `json:"pin"`
}

type Temperature struct {
	ID      uuid.UUID `json:"id"`
	Device  uuid.UUID `json:"device_id"`
	TempSup float64   `json:"temp_sup"`
	TempMid float64   `json:"temp_mid"`
	TempSub float64   `json:"temp_sub"`
	Created time.Time `json:"created"`
}
type NewTemperature struct {
	Device  uuid.UUID `json:"device_id"`
	TempSup float64   `json:"temp_sup"`
	TempMid float64   `json:"temp_mid"`
	TempSub float64   `json:"temp_sub"`
}

type Subscription struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       float64    `json:"amount"`
	Transfers   []Transfer `json:"transfers"`
	Created     time.Time  `json:"created"`
	Modified    *time.Time `json:"modified"`
}
type NewSubscription struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"amount"`
}

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
