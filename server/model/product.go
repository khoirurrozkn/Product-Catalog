package model

type Product struct{
	Id string `json:"id"`
	Price int `json:"price"`
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
}