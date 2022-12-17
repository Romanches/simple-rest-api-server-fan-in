package mock_server

import (
	"encoding/json"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/models"
	"github.com/gorilla/mux"
	"net/http"
)

type ResponseData struct {
	Data []models.Data `json:"data"`
}

//type Data struct {
//	Url            string  `json:"url"`
//	Views          int64   `json:"views"`
//	RelevanceScore float64 `json:"relevanceScore"`
//}

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
	data := MockDuckduckgo()

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

func MockDuckduckgo() []models.Data {
	var dataFromResource []models.Data
	var data models.Data

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
	data := MockGoogle()

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


func MockGoogle() []models.Data {
	var dataFromResource []models.Data
	var data models.Data

	data.Url = "www.example.com/abc1"
	data.Views = 1000
	data.RelevanceScore = 0.1
	dataFromResource = append(dataFromResource, data)

	data.Url = "www.example.com/abc2"
	data.Views = 2000
	data.RelevanceScore = 0.2
	dataFromResource = append(dataFromResource, data)

	data.Url = "www.example.com/abc4"
	data.Views = 4000
	data.RelevanceScore = 0.4
	dataFromResource = append(dataFromResource, data)

	data.Url = "www.example.com/abc3"
	data.Views = 3000
	data.RelevanceScore = 0.3
	dataFromResource = append(dataFromResource, data)

	data.Url = "www.example.com/abc5"
	data.Views = 5000
	data.RelevanceScore = 0.5
	dataFromResource = append(dataFromResource, data)

	return dataFromResource
}


func GetWikipediaAssignment(w http.ResponseWriter, r *http.Request) {
	//declare response variable
	var response ResponseData

	//Retrieve data
	data := MockWikipedia()

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


func MockWikipedia() []models.Data {
	var dataFromResource []models.Data
	var data models.Data

	data.Url = "www.wikipedia.com/abc1"
	data.Views = 11000
	data.RelevanceScore = 0.1
	dataFromResource = append(dataFromResource, data)

	data.Url = "www.wikipedia.com/abc2"
	data.Views = 12000
	data.RelevanceScore = 0.2
	dataFromResource = append(dataFromResource, data)

	data.Url = "www.wikipedia.com/abc4"
	data.Views = 13000
	data.RelevanceScore = 0.4
	dataFromResource = append(dataFromResource, data)

	data.Url = "www.wikipedia.com/abc3"
	data.Views = 14000
	data.RelevanceScore = 0.3
	dataFromResource = append(dataFromResource, data)

	data.Url = "www.wikipedia.com/abc5"
	data.Views = 15000
	data.RelevanceScore = 0.5
	dataFromResource = append(dataFromResource, data)

	return dataFromResource
}

