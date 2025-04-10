package models

import "time"

type User struct {
	ID        int       `json:"id" db:"id"`
	Firstname string    `json:"firstname" db:"firstname"`
	Lastname  string    `json:"lastname" db:"lastname"`
	Age       uint8     `json:"age" db:"age"`
	Email     string    `json:"email" db:"email"`
	Country   string    `json:"country" db:"country"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
