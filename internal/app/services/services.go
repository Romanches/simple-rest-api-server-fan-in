package services

import (
	"context"
)

// Data is the service-layer interface that describes "/data" endpoint
type Data interface {
	Get(context.Context, string, int) (string, error)
}

