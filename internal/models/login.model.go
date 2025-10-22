package models

import "github.com/golang-jwt/jwt/v5"

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=10,max=14"`
}

type LoginAPIRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Group    string `json:"group"`
}

type LoginAPIResponse struct {
	Token string `json:"token"`
}

type LoginSuccessResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Token      string `json:"token"`
}

type LoginClaims struct {
	Username    string `json:"username"`
	Permissions []int  `json:"permissions"`
	jwt.RegisteredClaims
}
