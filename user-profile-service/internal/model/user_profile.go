package model

type UserProfile struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
