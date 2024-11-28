package response

import "server/model/dto"

type Status struct{
	Code int `json:"code"`
	Description string `json:"description"`
}

type PagedResponse struct {
	Status Status `json:"status"`
	Data   []interface{} `json:"data"`
	Paging dto.Paging `json:"paging"`
}

type SingleResponse struct{
	Status Status `json:"status"`
	Data interface{} `json:"data"`
}