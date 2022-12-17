package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//var (
//	Client HTTPClient
//)

//func init() {
//	Client = &http.Client{}
//}

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type HTTPClientMock struct {
	// DoFunc will be executed whenever Do function is executed
	// so we'll be able to create a custom response
	DoFunc func(*http.Request) (*http.Response, error)
}

func (H HTTPClientMock) Do(r *http.Request) (*http.Response, error) {
	return H.DoFunc(r)
}

func NewHttpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

// Rest-client's method to make Get request with context
func SendGetRequest(ctx context.Context, c *http.Client, method, endpoint string, into interface{}) error {
	//req, err := http.NewRequest(method, endpoint, bytes.NewBuffer([]byte("")))
	// To make a request with a specified context.Context, use NewRequestWithContext and Client.Do.

	req, err := http.NewRequestWithContext(ctx, method, endpoint, nil)
	if err != nil {
		log.Printf("Error Occurred. %+v\n", err)
		return err
	}

	response, err := c.Do(req)
	if err != nil {
		log.Printf("Error sending request to API endpoint. %+v\n", err)
		return err
	}

	// Close the connection to reuse it
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf(http.StatusText(response.StatusCode))
	}

	// Parse the answer
	err = json.NewDecoder(response.Body).Decode(into)
	if err != nil {
		log.Printf("Couldn't parse response body. %+v\n", err)
		return err
	}

	// Success
	return nil
}

