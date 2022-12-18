package data

import (
	"context"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/models"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/repository"
)

// Service
type DataService struct {
	resources      []string
	dataRepository repository.Data
}

// NewDataService creates a new instance of Data service
func NewDataService(resources []string, dataRepo repository.Data) *DataService {
	return &DataService{
		resources:      resources,
		dataRepository: dataRepo,
	}
}

// Service method for /data endpoint
func (s DataService) GetData(ctx context.Context, params models.QueryParams) (result models.ResponseData, err error) {
	dataSet := []models.Data{}
	result = models.ResponseData{
		Data: []models.Data{},
	}

	// Call repository instance to get statistic about web-links views and their relevance-score
	dataSet, err = s.dataRepository.GetStatistic(ctx, s.resources)
	if err != nil {

	}

	// If we have no data
	if len(dataSet) == 0 {
		return
	}

	// Total amount of items
	result.Count = len(dataSet)

	// Sorting the result
	sortSlice(dataSet, params.SortKey)

	result.Data = dataSet

	// Limiting
	if len(dataSet) > params.Limit {
		result.Data = dataSet[:params.Limit]
	}

	// Return all
	return result, nil
}
