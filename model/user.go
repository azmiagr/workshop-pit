package model

import "github.com/google/uuid"

type UserParam struct {
	UserID uuid.UUID `json:"-"`
	Name   string    `json:"-"`
	Email  string    `json:"-"`
}

type UserRegisterParam struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=8"`
	PhoneNumber string `json:"phone_number" binding:"required,min=10,max=13"`
}

type UserRegisterResponse struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type UserLoginParam struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}
