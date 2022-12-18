package data

import (
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/models"
	"log"
	"sort"
	"strings"
)

func sortSlice(items []models.Data, sortKey string) {

	// Choose sorting-method
	switch strings.ToLower(sortKey) {
	case "url":
		SortByUrl(items)
	case "views":
		SortByViews(items)
	case "relevancescore":
		SortByRelevanceScore(items)
	case "":
	default:
		// We never reach there, we validated parameters earlier
		log.Println("Error: unexpected sortKey")
	}
}

// SortByViews sorts a slice of items by views.
// If items are equally viewed, we can sort it by the score.
func SortByViews(items []models.Data) {

	sort.Slice(items, func(i, j int) bool {
		//var sortedByViews, sortedByRelevanceScore bool
		var sortedByViews bool

		// sort by Views
		sortedByViews = items[i].Views < items[j].Views

		//// sort by lowest RelevanceScore
		//if items[i].Views == items[j].Views {
		//	sortedByRelevanceScore = items[i].RelevanceScore < items[j].RelevanceScore
		//	return sortedByRelevanceScore
		//}
		return sortedByViews
	})
}

// Sorting by Score
func SortByRelevanceScore(items []models.Data) {

	sort.Slice(items, func(i, j int) bool {
		var sortedByRelevanceScore bool

		sortedByRelevanceScore = items[i].RelevanceScore < items[j].RelevanceScore

		return sortedByRelevanceScore
	})
}

// Sorting by URL
func SortByUrl(items []models.Data) {

	sort.Slice(items, func(i, j int) bool {
		var sortedByUrl bool

		sortedByUrl = items[i].Url < items[j].Url

		return sortedByUrl
	})
}
