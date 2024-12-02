package response

type UpdatedProductResponse struct {
	Id string `json:"id"`
	Price int `json:"price"`
	Name string `json:"name"`
}