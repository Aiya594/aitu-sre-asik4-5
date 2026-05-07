package model

type Payment struct {
	ID      int     `json:"id"`
	OrderID int     `json:"order_id"`
	UserID  int     `json:"user_id"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"`
}
