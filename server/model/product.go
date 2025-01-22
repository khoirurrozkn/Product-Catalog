package model

import "time"

type Product struct{
	Id string `json:"id"`
	ImgUrl string `json:"img_url"`
	Price int `json:"price"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
