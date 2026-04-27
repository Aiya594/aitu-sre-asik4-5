package model

type Order struct {
	ID      int     `json:"id"`
	UserID  int     `json:"user_id"`
	Product string  `json:"product"`
	Amount  float64 `json:"amount"`
}
