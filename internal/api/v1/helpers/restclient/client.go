package restclient

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
//
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
	client := &http.Client{Timeout: 60 * time.Second}
	return client
}

func SendGetRequest(ctx context.Context, c *http.Client, method, endpoint string, into interface{}) error {

	//req, err := http.NewRequest(method, endpoint, bytes.NewBuffer([]byte("")))
	req, err := http.NewRequestWithContext(ctx, method, endpoint, nil)
	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
		return err
	}
	//log.Println(req.URL)

	response, err := c.Do(req)
	//response, err := NewHttpClient().Do(req)
	if err != nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
		return err
	}

	// Close the connection to reuse it
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf(http.StatusText(response.StatusCode))
	}

	// 1
	//log.Println(endpoint)
	//body, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	log.Fatalf("Couldn't parse response body. %+v", err)
	//	return err
	//}
	//log.Println("Response Body:", string(body))
	//io.Copy(ioutil.Discard, response.Body)


	// 2
	err = json.NewDecoder(response.Body).Decode(into)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
		return err
	}

	// 3
	//err = helpers.ParsePayload(response.Body, into)
	//if err != nil {
	//	log.Fatalf("Couldn't parse response body. %+v", err)
	//	return err
	//}

	//log.Println(into)
	//log.Println(body)


	//body, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	log.Fatalf("Couldn't parse response body. %+v", err)
	//	return nil, err
	//}
	//
	//return body, nil

	return nil
}

//func main() {
//	// c should be re-used for further calls
//	c := HttpClient()
//	endpoint := "https://httpbin.org/post"
//	response, err := SendRequest(c, http.MethodPost, endpoint)
//	if err != nil {
//		log.Fatalf("Error: %+v", err)
//	}
//	log.Println("Response Body:", string(response))
//}

