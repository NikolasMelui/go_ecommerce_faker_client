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

// CreateFakeLabels ...
func (c *Client) CreateFakeLabels(wg *sync.WaitGroup, count int) int {
	faker.Locale = locales.Ru
	ch := make(chan int, count)
	ch <- 0
	for i := 0; i < count; i++ {
		wg.Add(1)
		time.Sleep(time.Millisecond * 50)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			fakeLabel := LabelData{
				Title:       faker.Commerce().Color() + " " + faker.Name().LastName(),
				Description: faker.Lorem().Sentence(10),
			}
			log.Println(fakeLabel)
			_, err := c.CreateLabel(&fakeLabel)
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
