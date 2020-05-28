package efclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Counterparty ...
type Counterparty struct {
	ID                    int                    `json:"id"`
	CreditLimit           int                    `json:"credit_limit"`
	User                  []User                 `json:"user"`
	Orders                []Order                `json:"orders"`
	CounterpartyDocuments []CounterpartyDocument `json:"counterparty_documents"`
	Name                  string                 `json:"name"`
}

// CounterpartyData ...
type CounterpartyData struct {
	CreditLimit int    `json:"credit_limit"`
	User        int    `json:"user"`
	Name        string `json:"name"`
}

// CreateCounterparty ...
func (c *Client) CreateCounterparty(counterpartyData *CounterpartyData) (*Counterparty, error) {

	requestData := map[string]interface{}{
		"credit_limit": &counterpartyData.CreditLimit,
		"user":         &counterpartyData.User,
		"name":         &counterpartyData.Name,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/counterparties", c.BaseURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	var res Counterparty

	if err := c.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
