package response

type UserCredential struct {
	User UserResponse `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
