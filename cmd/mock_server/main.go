package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type ResponseData struct {
	Data []Data `json:"data"`
}

type Data struct {
	Url            string  `json:"url"`
	Views          int64   `json:"views"`
	RelevanceScore float64 `json:"relevanceScore"`
}

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/assignment132/assignment/main").Subrouter()
	api.HandleFunc("/duckduckgo", GetDuckDuckGoAssignment).Methods("GET")
	api.HandleFunc("/google", GetGoogleAssignment).Methods("GET")
	api.HandleFunc("/wikipedia", GetWikipediaAssignment).Methods("GET")

	http.Handle("/", r)

	http.ListenAndServe(":3000", r)
}

func GetDuckDuckGoAssignment(w http.ResponseWriter, r *http.Request) {
	//declare response variable
	var response ResponseData

	//Retrieve data
	data := mockDuckduckgo()

	//assign person details to response
	response.Data = data

	//update content type
	w.Header().Set("Content-Type", "application/json")

	//specify HTTP status code
	w.WriteHeader(http.StatusOK)

	//convert struct to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	//update response
	w.Write(jsonResponse)
}

func mockDuckduckgo() []Data {
	var dataFromResource []Data
	var data Data

	data.Url = "www.yahoo.com/abc6"
	data.Views = 6000
	data.RelevanceScore = 0.6
	dataFromResource = append(dataFromResource, data)

	data.Url = "www.yahoo.com/abc7"
	data.Views = 7000
	data.RelevanceScore = 0.7
	dataFromResource = append(dataFromResource, data)

	data.Url = "www.yahoo.com/abc8"
	data.Views = 8000
	data.RelevanceScore = 0.8
	dataFromResource = append(dataFromResource, data)

	data.Url = "www.yahoo.com/abc9"
	data.Views = 9000
	data.RelevanceScore = 0.9
	dataFromResource = append(dataFromResource, data)

	data.Url = "www.yahoo.com/abc10"
	data.Views = 10000
	data.RelevanceScore = 1.0
	dataFromResource = append(dataFromResource, data)

	return dataFromResource
}


func GetGoogleAssignment(w http.ResponseWriter, r *http.Request) {
	//declare response variable
	var response ResponseData

	//Retrieve data
	data := mockGoogle()

	//assign person details to response
	response.Data = data

	//update content type
	w.Header().Set("Content-Type", "application/json")

	//specify HTTP status code
	w.WriteHeader(http.StatusOK)

	//convert struct to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	//update response
	w.Write(jsonResponse)
}


func mockGoogle() []Data {
	var dataFromResource []Data
	var data Data

	data.Url = "www.yahoo.com/abc6"
	data.Views = 6000
	data.RelevanceScore = 0.6
	dataFromResource = append(dataFromResource, data)

	data.Url = "www.yahoo.com/abc7"
	data.Views = 7000
	data.RelevanceScore = 0.7
	dataFromResource = append(dataFromResource, data)

	data.Url = "www.yahoo.com/abc8"
	data.Views = 8000
	data.RelevanceScore = 0.8
	dataFromResource = append(dataFromResource, data)

	data.Url = "www.yahoo.com/abc9"
	data.Views = 9000
	data.RelevanceScore = 0.9
	dataFromResource = append(dataFromResource, data)

	data.Url = "www.yahoo.com/abc10"
	data.Views = 10000
	data.RelevanceScore = 1.0
	dataFromResource = append(dataFromResource, data)

	return dataFromResource
}


func GetWikipediaAssignment(w http.ResponseWriter, r *http.Request) {
	//declare response variable
	var response ResponseData

	//Retrieve data
	data := mockWikipedia()

	//assign person details to response
	response.Data = data

	//update content type
	w.Header().Set("Content-Type", "application/json")

	//specify HTTP status code
	w.WriteHeader(http.StatusOK)

	//convert struct to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	//update response
	w.Write(jsonResponse)
}


func mockWikipedia() []Data {
	var dataFromResource []Data
	var data Data

	data.Url = "www.yahoo.com/abc6"
	data.Views = 6000
	data.RelevanceScore = 0.6
	dataFromResource = append(dataFromResource, data)

	data.Url = "www.yahoo.com/abc7"
	data.Views = 7000
	data.RelevanceScore = 0.7
	dataFromResource = append(dataFromResource, data)

	data.Url = "www.yahoo.com/abc8"
	data.Views = 8000
	data.RelevanceScore = 0.8
	dataFromResource = append(dataFromResource, data)

	data.Url = "www.yahoo.com/abc9"
	data.Views = 9000
	data.RelevanceScore = 0.9
	dataFromResource = append(dataFromResource, data)

	data.Url = "www.yahoo.com/abc10"
	data.Views = 10000
	data.RelevanceScore = 1.0
	dataFromResource = append(dataFromResource, data)

	return dataFromResource
}

