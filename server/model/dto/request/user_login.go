package request

type UserLogin struct {
	EmailOrNickname string `json:"email_or_nickname"`
	Password        string `json:"password"`
}
