package repository

import (
	"context"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/models"
)

// Data is the interface that wraps work with database or external data source
type Data interface {
	GetStatistic(ctx context.Context, url string) (models.ResponseData, error)
}
