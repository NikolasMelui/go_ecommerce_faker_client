package efclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// User ...
type User struct {
	ID                 int            `json:"id"`
	Username           string         `json:"username"`
	Email              string         `json:"email"`
	Provider           string         `json:"provider"`
	Password           string         `json:"password"`
	ResetPasswordToken string         `json:"resetPasswordToken"`
	Confirmed          bool           `json:"confirmed"`
	Blocked            bool           `json:"blocked"`
	Role               string         `json:"role"`
	Counterparties     []Counterparty `json:"counterparties"`
	Phone              string         `json:"phone"`
}

// UserData ...
type UserData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

// CreateUser ...
func (c *Client) CreateUser(userData *UserData) (*User, error) {

	requestData := map[string]interface{}{
		"username": &userData.Username,
		"email":    &userData.Email,
		"password": &userData.Password,
		"phone":    &userData.Phone,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth/local/register", c.BaseURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	var res User
	if err := c.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
