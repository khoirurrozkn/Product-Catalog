package request

type UserRegister struct {
	Email string `json:"email"`
	Nickname string `json:"nickname"`
	Password        string `json:"password"`
}