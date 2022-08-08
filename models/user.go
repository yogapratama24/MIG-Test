package models

import "time"

type UserRegister struct {
	Id        int       `json:"id"`
	UserName  string    `json:"user_name,omitempty" validate:"required"`
	Password  string    `json:"password,omitempty" validate:"required"`
	Email     string    `json:"email,omitempty" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLogin struct {
	Id       int    `json:"id"`
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}
