package model

import "time"

type UserRefreshToken struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"updated_at"`
}
