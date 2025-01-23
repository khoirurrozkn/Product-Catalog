package model

type UserCredential struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}
