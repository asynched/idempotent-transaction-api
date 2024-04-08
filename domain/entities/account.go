package entities

import "time"

type Account struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
