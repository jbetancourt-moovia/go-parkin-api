package models

type Customer struct {
	ID        int64  `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Phone     string `json:"phone" db:"phone"`
	Email     string `json:"email" db:"email"`
}

type CustomerCreate struct {
	FirstName string `json:"first_name" db:"first_name" validate:"required,min=3,max=100"`
	LastName  string `json:"last_name" db:"last_name" validate:"required,min=3,max=100"`
	Phone     string `json:"phone" db:"phone" validate:"required,len=10,numeric"`
	Email     string `json:"email" db:"email" validate:"required,email"`
}

type ListCustomerResponse struct {
	StatusCode int        `json:"status_code"`
	Data       []Customer `json:"data"`
	Message    string     `json:"message"`
}

type CustomerResponse struct {
	StatusCode int       `json:"status_code"`
	Data       *Customer `json:"data"`
	Message    string    `json:"message"`
}

type BasicResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
