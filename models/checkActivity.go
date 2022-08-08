package models

import "time"

type CheckIn struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	DateCheckIn time.Time `json:"date_check_in"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Activity struct {
	Id          int       `json:"id"`
	CheckInId   int       `json:"check_in_id"`
	Description string    `json:"description" validate:"required"`
	Date        string    `json:"date,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Attendance struct {
	Id          int    `json:"id"`
	DateCheckIn string `json:"date_check_in"`
	UserId      int    `json:"user_id"`
	UserName    string `json:"user_name"`
}
