package data

import (
	"context"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/helpers/restclient"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/models"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/repository"
	"net/http"
	"time"
)

// Repository describes this layer parameters
type Repository struct {
	httpClient *http.Client
	timeout time.Duration
}

// NewRepository creates new instance of Repository
func NewRepository(httpClient *http.Client, timeout time.Duration) repository.Data {
	return &Repository{
		httpClient: httpClient,
		timeout: timeout,
	}
}

// GetStatistic gets data from remote sources
func (r *Repository) GetStatistic(ctx context.Context, url string) (data models.ResponseData, err error) {
	context, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	// Send http request to get data from remote source
	err = restclient.SendGetRequest(context, r.httpClient, http.MethodGet, url, &data)
	if err != nil {
		return data, err
	}

	//if err := helpers.ParsePayload(response, &statistic); err != nil {
	//	render.Error(w, err)
	//	return
	//}

	return data, nil
}

