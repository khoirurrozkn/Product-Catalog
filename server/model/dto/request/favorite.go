package request

type Favorite struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	ProductId string    `json:"product_id"`
}