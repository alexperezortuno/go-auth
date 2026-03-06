package user

import "github.com/google/uuid"

type Response struct {
	ID       uuid.UUID `json:"id" copier:"must"`
	FullName string    `json:"full_name"`
	Name     string    `json:"name"`
	LastName string    `json:"last_name"`
	Nickname string    `json:"nickname"`
	IdCard   string    `json:"id_card"`
	Email    string    `json:"email" copier:"must"`
}
