package productclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
)

type ProductClient interface {
	DecreaseStock(productID int, amount int) error
}

type ProductClientHTTP struct {
	BaseURL string
}

func NewProductClient(url string) *ProductClientHTTP {

	return &ProductClientHTTP{BaseURL: url}
}

type stockRequest struct {
	Amount int `json:"amount"`
}

func (c *ProductClientHTTP) DecreaseStock(productID int, amount int) error {

	log.Printf("[ProductClient] DecreaseStock product=%d amount=%d", productID, amount)

	body, _ := json.Marshal(stockRequest{Amount: amount})

	req, err := http.NewRequest(
		"PUT",
		c.BaseURL+"/products/"+strconv.Itoa(productID)+"/decrease-stock",
		bytes.NewBuffer(body),
	)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Printf("[ProductClient][ERROR] request failed: %v", err)
		return errors.New("product service unreachable")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("[ProductClient][ERROR] stock rejected status=%d", resp.StatusCode)
		return errors.New("insufficient stock")
	}

	log.Printf("[ProductClient] stock reserved successfully")
	return nil
}
