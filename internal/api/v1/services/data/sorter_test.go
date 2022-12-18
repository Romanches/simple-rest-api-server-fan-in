package data

import (
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/models"
	"reflect"
	"testing"
)

func getItems() []models.Data {
	return []models.Data{
		models.Data{
			Url: "www.yahoo.com/abc6",
			Views: 6000,
			RelevanceScore: 0.6,
		},
		models.Data{
			Url: "www.example.com/abc1",
			Views: 1000,
			RelevanceScore: 0.1,
		},
		models.Data{
			Url: "www.wikipedia.com/abc2",
			Views: 12000,
			RelevanceScore: 0.2,
		},
	}
}

func Test_sortSlice(t *testing.T) {
	type args struct {
		items   []models.Data
		sortKey string
	}

	items := getItems()

	itemsSortedByUrl := []models.Data{
		models.Data{
			Url: "www.example.com/abc1",
			Views: 1000,
			RelevanceScore: 0.1,
		},
		models.Data{
			Url: "www.wikipedia.com/abc2",
			Views: 12000,
			RelevanceScore: 0.2,
		},
		models.Data{
			Url: "www.yahoo.com/abc6",
			Views: 6000,
			RelevanceScore: 0.6,
		},
	}

	itemsSortedByViews := []models.Data{
		models.Data{
			Url: "www.example.com/abc1",
			Views: 1000,
			RelevanceScore: 0.1,
		},
		models.Data{
			Url: "www.yahoo.com/abc6",
			Views: 6000,
			RelevanceScore: 0.6,
		},
		models.Data{
			Url: "www.wikipedia.com/abc2",
			Views: 12000,
			RelevanceScore: 0.2,
		},
	}

	itemsSortedByRelevanceScore := []models.Data{
		models.Data{
			Url: "www.example.com/abc1",
			Views: 1000,
			RelevanceScore: 0.1,
		},
		models.Data{
			Url: "www.wikipedia.com/abc2",
			Views: 12000,
			RelevanceScore: 0.2,
		},
		models.Data{
			Url: "www.yahoo.com/abc6",
			Views: 6000,
			RelevanceScore: 0.6,
		},
	}

	tests := []struct {
		name string
		args args
		expectedItems    []models.Data
	}{
		{
			name: "Empty sortKey, no sorting",
			args: args{
				items: items,
				sortKey: "",
			},
			expectedItems: items,
		},

		// sortKey = Url
		{
			name: "sort by Url, sortKey in lower-case ",
			args: args{
				items: items,
				sortKey: "url",
			},
			expectedItems: itemsSortedByUrl,
		},
		{
			name: "sort by Url, sortKey in upper-case ",
			args: args{
				items: items,
				sortKey: "URL",
			},
			expectedItems: itemsSortedByUrl,
		},
		{
			name: "sort by Url, sortKey in camel-case ",
			args: args{
				items: items,
				sortKey: "Url",
			},
			expectedItems: itemsSortedByUrl,
		},

		// sortKey = Views
		{
			name: "sort by Views, sortKey in lower-case ",
			args: args{
				items: items,
				sortKey: "views",
			},
			expectedItems: itemsSortedByViews,
		},
		{
			name: "sort by Views, sortKey in upper-case ",
			args: args{
				items: items,
				sortKey: "VIEWS",
			},
			expectedItems: itemsSortedByViews,
		},
		{
			name: "sort by Views, sortKey in camel-case ",
			args: args{
				items: items,
				sortKey: "Views",
			},
			expectedItems: itemsSortedByViews,
		},

		// sortKey = RelevanceScore
		{
			name: "sort by RelevanceScore, sortKey in lower-case ",
			args: args{
				items: items,
				sortKey: "relevancescore",
			},
			expectedItems: itemsSortedByRelevanceScore,
		},
		{
			name: "sort by RelevanceScore, sortKey in upper-case ",
			args: args{
				items: items,
				sortKey: "RELEVANCESCORE",
			},
			expectedItems: itemsSortedByRelevanceScore,
		},
		{
			name: "sort by RelevanceScore, sortKey in camel-case ",
			args: args{
				items: items,
				sortKey: "RelevanceScore",
			},
			expectedItems: itemsSortedByRelevanceScore,
		},

		// wrong sortKey
		{
			name: "wrong sortKey, no sorting",
			args: args{
				items: items,
				sortKey: "WrongSortKey",
			},
			expectedItems: items,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := getItems()

			sortSlice(got, tt.args.sortKey)

			if !reflect.DeepEqual(got, tt.expectedItems) {
				t.Errorf("Get() got = %v, want %v", got, tt.expectedItems)
			}
		})
	}
}
