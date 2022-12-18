package models

import (
	"errors"
	"reflect"
	"strings"
)

// Represents the structure of data received from external platforms
type ClientResponseData struct {
	// The list of Items
	Data  []Data `json:"data"`
}

// Item
type Data struct {
	Url            string  `json:"url"`
	Views          int64   `json:"views"`
	RelevanceScore float64 `json:"relevanceScore"`
}

// Response model for GET:/data endpoint
type ResponseData struct {
	Data  []Data `json:"data"`
	Count int    `json:"count"`
}

// Request parameters for GET:/data endpoint
type QueryParams struct {
	SortKey string
	Limit   int
}

// Request parameters validator
func (p *QueryParams) Validate() error {
	var errs string

	// Limit should be greater than 1 less than 200
	if p.Limit <= 1 || p.Limit >= 200 {
		//errs = append(errs, "invalid limit parameter")
		errs = errs + "invalid limit parameter"
	}

	structExample := new(Data)

	// SortKey can be empty
	if p.SortKey != "" && !keyExists(structExample, p.SortKey) {
		//errs = append(errs, "invalid sortKey parameter")
		errs = errs + "invalid sortKey parameter"
	}

	if len(errs) > 0 {
		return errors.New(errs)
	}

	return nil
}

func keyExists(d *Data, key string) bool {
	metaValue := reflect.ValueOf(d).Elem()

	field := caseInsensitiveFieldByName(reflect.Indirect(metaValue), key)

	return field != (reflect.Value{})
}

func caseInsensitiveFieldByName(v reflect.Value, name string) reflect.Value {
	name = strings.ToLower(name)
	return v.FieldByNameFunc(func(n string) bool {
		return strings.ToLower(n) == name
	})
}
