package repository

import (
	"context"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/models"
)

// Data is the interface that wraps work with repository
type Data interface {
	GetStatistic(ctx context.Context, urls []string) ([]models.Data, error)
}
