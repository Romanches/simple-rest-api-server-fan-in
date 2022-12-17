package data

import (
	"context"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/models"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/repository"
	"log"
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


func (s DataService) GetData(ctx context.Context, params models.QueryParams) (result models.ResponseData, err error) {
	dataSet := []models.Data{}
	result = models.ResponseData{
		Data: []models.Data{},
	}

	// Poll all resources from the list
	for _, url := range s.resources {

		// Call repository
		repoResponse, err := s.dataRepository.GetStatistic(ctx, url)
		if err != nil {
			//return result, err
			log.Println("Error:", err)
			continue
		}

		if len(repoResponse.Data) > 0 {
			dataSet = append(dataSet, repoResponse.Data...)
		}
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
