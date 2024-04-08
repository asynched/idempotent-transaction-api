package entities

import "time"

type Transaction struct {
	Id        string    `json:"id"`
	Amount    float64   `json:"amount"`
	Payer     string    `json:"payer"`
	Payee     string    `json:"payee"`
	CreatedAt time.Time `json:"created_at"`
}
