package response

type UpdatedProductResponse struct {
	Id string `json:"id"`
	ImgUrl string `json:"img_url"`
	Price int `json:"price"`
	Name string `json:"name"`
}