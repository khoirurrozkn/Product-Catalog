package model

import "time"

type Favorite struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	ProductId string    `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
}
