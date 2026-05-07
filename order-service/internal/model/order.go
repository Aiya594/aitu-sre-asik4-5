package model

type Order struct {
	ID      string  `json:"id"`
	UserID  int     `json:"user_id"`
	Product string  `json:"product"`
	Amount  float64 `json:"amount"`
}
