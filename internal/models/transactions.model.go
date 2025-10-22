package models

import "time"

type StartTiming struct {
	Placa       string `json:"placa" validate:"required"`
	CustomerDNI string `json:"customer_dni" validate:"required"`
}

type GetTiming struct {
	Placa string `json:"placa" validate:"required"`
}

type PaymentInfo struct {
	ID          int64     `json:"id" db:"id"`
	Placa       string    `json:"placa"`
	CustomerDNI string    `json:"customer_dni"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	StartsAt    time.Time `json:"starts_at" db:"starts_at"`
	EndsAt      time.Time `json:"ends_at" db:"ends_at"`
	Minutes     float64   `json:"minutes"`
	Fee         float64   `json:"minute_fee"`
	TotalDue    float64   `json:"total_due" db:"total_due"`
}

type ConfirmPayment struct {
	ID             int64   `json:"id" db:"id"`
	AmountDue      float64 `json:"amount_due" db:"amount_due"`
	AmountReceived float64 `json:"amount_received" db:"amount_received" validate:"required,min=0"`
	AmountChange   float64 `json:"amount_change" db:"amount_change"`
}
