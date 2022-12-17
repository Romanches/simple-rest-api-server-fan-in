package data

import (
	"context"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/helpers/rest"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/models"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/repository"
	"log"
	"net/http"
)

// Repository describes this layer parameters
type Repository struct {
	// Isolate HTTP-client and make it available from this layer only
	httpClient *http.Client
}

// NewRepository creates new instance of Repository
func NewRepository(httpClient *http.Client) repository.Data {
	return &Repository{
		httpClient: httpClient,
	}
}

// GetStatistic gets data from remote sources
func (r *Repository) GetStatistic(ctx context.Context, urls []string) (dataSet []models.Data, err error) {
	// Expected response from HTTP client
	responseData := models.ClientResponseData{}

	// Channel to read data from goroutines
	chOut := make(chan models.ClientResponseData, 1)

	// Poll all resources from the list
	for _, url := range urls {

		// Run worker
		go worker(ctx, r.httpClient, http.MethodGet, url, chOut)

	}

	// Read all responses from the channel
	for range urls {
		responseData = <-chOut

		// Add received items to the general list
		if len(responseData.Data) > 0 {
			dataSet = append(dataSet, responseData.Data...)
		}

	}

	// Return the list of received items
	return
}

func worker(ctx context.Context, c *http.Client, method, endpoint string, chOut chan models.ClientResponseData) {
	resp := models.ClientResponseData{}

	// Init http-connection and do request
	err := rest.GetWithRetry(ctx, c, method, endpoint, &resp)
	if err != nil {
		log.Println(err)
	}

	// Write http-response into output-channel
	chOut <- resp
}
