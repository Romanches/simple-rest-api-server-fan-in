package rest

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"
)

// A backoff schedule for when and how often to retry failed HTTP
// requests. The first element is the time to wait after the
// first failure, the second the time to wait after the second
// failure, etc. After reaching the last element, retries stop
// and the request is considered failed.
var backoffSchedule = []time.Duration{
	1 * time.Second,
	3 * time.Second,
	10 * time.Second,
	//30 * time.Second,
}

func GetWithRetryAsync(ctx context.Context, c *http.Client, method, endpoint string, into interface{}) {
	GetWithRetry(ctx, c, method, endpoint, into)
}

// GetWithRetry makes pauses and retries requests for unavailable endpoints according to the schedule
func GetWithRetry(ctx context.Context, c *http.Client, method, url string, into interface{}) error {
	var err error

	for _, backoff := range backoffSchedule {

		err = SendGetRequest(ctx, c, method, url, into)
		// Success connection
		if err == nil {
			break
		}

		log.Printf("Request error: %s %+v. Retrying in %v\n", url, err, backoff)

		// Wait before try again
		time.Sleep(backoff)
	}

	// All retries failed
	if err != nil {
		return errors.New(url + " " + err.Error())
	}

	// Success
	return nil
}
