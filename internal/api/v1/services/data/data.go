package data

import (
	"context"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/models"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/repository"
)

type DataService struct {
	resources []string
	dataRepository repository.Data
}

// NewDataService creates a new instance of Data service
func NewDataService(resources []string, dataRepo repository.Data) *DataService {
	return &DataService{
		resources: resources,
		dataRepository: dataRepo,
	}
}


func (s DataService) GetData(ctx context.Context, params models.QueryParams) (result []models.Data, err error) {
	// Poll all resources from the list
	for _, url := range s.resources {

		// Call repository
		dataSet, err := s.dataRepository.GetStatistic(ctx, url)
		if err != nil {
			return result, err
		}

		if len(dataSet.Data) > 0 {
			result = append(result, dataSet.Data...)
		}
	}

	// Sorting the result
	sortSlice(result, params.SortKey)

	// Limiting
	if len(result) > params.Limit {
		return result[:params.Limit], nil
	}

	// Return all
	return result, nil
}
