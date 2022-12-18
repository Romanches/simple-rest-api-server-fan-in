package services

import (
	"context"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/models"
)

// Data is the service-layer interface that describes "/views" endpoint
type Data interface {
	GetData(context.Context, models.QueryParams) (models.ResponseData, error)
}

