package types

import "github.com/google/uuid"

type Owner struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Phone string    `json:"phone"`
	Email string    `json:"email"`
}

func ValidateOwner(o *Owner) bool { return true }
