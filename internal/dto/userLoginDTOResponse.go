package dto

type LoginResponse struct {
	Token string `json:"token"`
	ID    int    `json:"id"`
	Email string `json:"email"`
}
