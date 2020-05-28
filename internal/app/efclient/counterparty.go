package efclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"syreclabs.com/go/faker"
	"syreclabs.com/go/faker/locales"
)

// Counterparty ...
type Counterparty struct {
	ID                    int                    `json:"id"`
	CreditLimit           int                    `json:"credit_limit"`
	User                  []User                 `json:"user"`
	Orders                []Order                `json:"orders"`
	CounterpartyDocuments []CounterpartyDocument `json:"counterparty_documents"`
	Title                 string                 `json:"title"`
}

// CounterpartyData ...
type CounterpartyData struct {
	CreditLimit int    `json:"credit_limit"`
	User        int    `json:"user"`
	Title       string `json:"title"`
}

// CreateCounterparty ...
func (c *Client) CreateCounterparty(counterpartyData *CounterpartyData) (*Counterparty, error) {

	requestData := map[string]interface{}{
		"credit_limit": &counterpartyData.CreditLimit,
		"user":         &counterpartyData.User,
		"title":        &counterpartyData.Title,
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

// CreateFakeCounterparties ...
func (c *Client) CreateFakeCounterparties(wg *sync.WaitGroup, count int) int {
	faker.Locale = locales.Ru
	ch := make(chan int, count)
	ch <- 0
	for i := 0; i < count; i++ {
		wg.Add(1)
		time.Sleep(time.Millisecond * 50)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			rand.Seed(time.Now().UnixNano())
			minCreditLimit := 500000
			maxCreditLimit := 5000000
			creditLimit := rand.Intn(maxCreditLimit-minCreditLimit+1) + minCreditLimit
			firstUserID := 1
			lastUserID := count / 2
			userID := rand.Intn(lastUserID-firstUserID+1) + firstUserID
			fakeCounterparty := CounterpartyData{
				Title:       faker.Company().Name() + " " + faker.Company().Suffix(),
				CreditLimit: creditLimit,
				User:        userID,
			}
			log.Println(fakeCounterparty)
			_, err := c.CreateCounterparty(&fakeCounterparty)
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
