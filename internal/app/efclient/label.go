package efclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Label ...
type Label struct {
	ID int `json:"id"`
}

// LabelData ...
type LabelData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// CreateLabel ...
func (c *Client) CreateLabel(labelData *LabelData) (*Label, error) {

	requestData := map[string]interface{}{
		"title":       &labelData.Title,
		"description": &labelData.Description,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/product-labels", c.BaseURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	var res Label
	if err := c.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
