package efclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"syreclabs.com/go/faker"
	"syreclabs.com/go/faker/locales"
)

// User ...
type User struct {
	ID                 int             `json:"id"`
	Username           string          `json:"username"`
	Email              string          `json:"email"`
	Provider           string          `json:"provider"`
	Password           string          `json:"password"`
	ResetPasswordToken string          `json:"resetPasswordToken"`
	Confirmed          bool            `json:"confirmed"`
	Blocked            bool            `json:"blocked"`
	Role               string          `json:"role"`
	Counterparties     json.RawMessage `json:"counterparties"`
	Phone              string          `json:"phone"`
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

// CreateFakeUsers ...
func (c *Client) CreateFakeUsers(wg *sync.WaitGroup, count int) int {
	faker.Locale = locales.Ru
	ch := make(chan int, count)
	ch <- 0
	for i := 0; i < count; i++ {
		wg.Add(1)
		time.Sleep(time.Millisecond * 50)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			fakeUser := UserData{
				Username: faker.Internet().UserName(),
				Email:    faker.Internet().Email(),
				Password: "password",
				Phone:    "+7" + faker.PhoneNumber().String(),
			}
			log.Println(fakeUser)
			_, err := c.CreateUser(&fakeUser)
			if err != nil {
				log.Print(fmt.Errorf("%v", err))
				// log.Fatal(err)
			} else {
				counter := <-ch
				ch <- counter + 1
			}
		}(wg)
	}
	wg.Wait()
	close(ch)
	return <-ch
}
