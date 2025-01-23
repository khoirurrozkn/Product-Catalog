package model

import "time"

type User struct {
	Id        string    `json:"id"`
	Email     string    `json:"email"`
	Nickname  string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
