package models

import "time"

type Vehicle struct {
	Tipo      string    `json:"tipo" db:"tipo"`
	Marca     string    `json:"marca" db:"marca"`
	Modelo    string    `json:"modelo" db:"modelo"`
	Placa     string    `json:"placa" db:"placa"`
	Status    bool      `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type VehicleCreate struct {
	Placa  string `json:"placa" db:"placa"`
	Tipo   string `json:"tipo" db:"tipo"`
	Marca  string `json:"marca" db:"marca"`
	Modelo string `json:"modelo" db:"modelo"`
}

type ListVehicleResponse struct {
	StatusCode int       `json:"status_code"`
	Data       []Vehicle `json:"data"`
	Message    string    `json:"message"`
}

type VehicleResponse struct {
	StatusCode int      `json:"status_code"`
	Data       *Vehicle `json:"data"`
	Message    string   `json:"message"`
}
