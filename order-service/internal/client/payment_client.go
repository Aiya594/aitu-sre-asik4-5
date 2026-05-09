package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type PaymentClient interface {
	ProcessPayment(
		orderID string,
		userID int,
		amount float64,
	) error
}

type PaymentClientHTTP struct {
	BaseURL string
}

func NewPaymentClient(url string) *PaymentClientHTTP {
	return &PaymentClientHTTP{
		BaseURL: url,
	}
}

func (c *PaymentClientHTTP) ProcessPayment(
	orderID string,
	userID int,
	amount float64,
) error {

	body, _ := json.Marshal(map[string]interface{}{
		"order_id": orderID,
		"user_id":  userID,
		"amount":   amount,
	})

	req, err := http.NewRequest(
		"POST",
		c.BaseURL+"/payment",
		bytes.NewBuffer(body),
	)

	if err != nil {
		return err
	}

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// if resp.StatusCode != 200 {
	// 	return errors.New("payment failed")
	// }

	return nil
}
