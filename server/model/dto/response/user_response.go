package response

import "time"

type UserResponse struct {
	Id          string    `json:"id"`
	Email       string    `json:"email"`
	Nickname    string    `json:"nickname"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	AccessToken string    `json:"access_token"`
}
