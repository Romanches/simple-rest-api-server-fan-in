package data

import (
	"context"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/helpers/rest"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/models"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/repository"
	"net/http"
	"time"
)

// Repository describes this layer parameters
type Repository struct {
	// Isolate HTTP-client and make it available from this layer only
	httpClient *http.Client

	// TODO: Do we need it here?
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
func (r *Repository) GetStatistic(ctx context.Context, url string) (data models.ClientResponseData, err error) {
	data = models.ClientResponseData{}

	// TODO: Do we need it here?
	ctx, cancel := context.WithTimeout(ctx, r.timeout)

	defer cancel()

	// Send http request to get data from remote source
	//err = restclient.SendGetRequest(context, r.httpClient, http.MethodGet, url, &data)
	err = rest.GetWithRetry(ctx, r.httpClient, http.MethodGet, url, &data)
	if err != nil {
		//return data, err
		return data, nil
	}

	//if err := helpers.ParsePayload(response, &statistic); err != nil {
	//	render.Error(w, err)
	//	return
	//}

	//_, body, err := restclient.GetURLDataWithRetries(url)
	//if err != nil {
	//	return data, err
	//}
	//fmt.Printf("response = %v\n", string(body))

	return data, nil
}

