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

	"github.com/nikolasMelui/go_ecommerce_faker_client/internal/app/helper"
	"syreclabs.com/go/faker"
	"syreclabs.com/go/faker/locales"
)

// CounterpartyDocument ...
type CounterpartyDocument struct {
	ID           int          `json:"id"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	URL          string       `json:"url"`
	Counterparty Counterparty `json:"counterparty"`
}

// CounterpartyDocumentData ...
type CounterpartyDocumentData struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	URL          string `json:"url"`
	Counterparty int    `json:"counterparty"`
}

// CreateCounterpartyDocument ...
func (c *Client) CreateCounterpartyDocument(counterpartyDocumentData *CounterpartyDocumentData) (*CounterpartyDocument, error) {
	if err := helper.CreateFakeFile(counterpartyDocumentData.Title, 1e3); err != nil {
		return nil, err
	}
	log.Println("File was created seccessfully")
	requestData := map[string]interface{}{
		"title":        &counterpartyDocumentData.Title,
		"description":  &counterpartyDocumentData.Description,
		"url":          &counterpartyDocumentData.URL,
		"counterparty": &counterpartyDocumentData.Counterparty,
	}
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/counterparty-documents", c.BaseURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	var res CounterpartyDocument
	if err := c.SendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// CreateFakeCounterpartyDocuments ...
func (c *Client) CreateFakeCounterpartyDocuments(wg *sync.WaitGroup, count int) int {
	faker.Locale = locales.Ru
	ch := make(chan int, count)
	ch <- 0
	for i := 0; i < count; i++ {
		wg.Add(1)
		time.Sleep(time.Millisecond * 50)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			rand.Seed(time.Now().UnixNano())
			maxCounterpartyID := count / 2
			counterpartyID := rand.Intn(maxCounterpartyID) + 1
			fileExtension := "pdf"
			fileName := faker.Lorem().Characters(10) + "." + fileExtension
			fileDescription := faker.Lorem().Sentence(10)
			fakeCounterpartyDocument := CounterpartyDocumentData{
				Title:        fileName,
				Description:  fileDescription,
				URL:          "https://localhost:1337/" + fileName,
				Counterparty: counterpartyID,
			}
			log.Println(fakeCounterpartyDocument)
			_, err := c.CreateCounterpartyDocument(&fakeCounterpartyDocument)
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
