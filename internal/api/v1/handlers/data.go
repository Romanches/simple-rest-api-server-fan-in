package handlers

import (
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/errs"
	"github.com/Romanches/simple-rest-api-server-fan-in/pkg/render"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/models"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/services"
	"log"
	"net/http"
	"strconv"
)

// DataHandler provides /data handler
type DataHandler struct {
	dataService services.Data
}

// NewDataHandler creates a new instance of DataHandler
func NewDataHandler(dataService services.Data) *DataHandler {
	return &DataHandler{dataService: dataService}
}

// Get is a handler for /data rest-api endpoint
func (h *DataHandler) Get(w http.ResponseWriter, r *http.Request) {
	params := models.QueryParams{}

	// Grab query params
	params.SortKey = r.URL.Query().Get("sortKey")

	qLimit := r.URL.Query().Get("limit")
	if qLimit == "" {
		// Set limit as 20 by default
		params.Limit = 20
	} else {
		limit, err := strconv.Atoi(qLimit)
		if err != nil {
			render.Error(w, errs.ErrNotValid)
			return
		}
		params.Limit = limit
	}

	// Validate parameters
	if err := params.Validate(); err != nil {
		log.Println(err.Error())

		// Respond 401
		render.Error(w, errs.ErrNotValid)
		return
	}

	////u, err := r.Get("YourHandler").URL("id", id, "key", key)
	////if err != nil {
	////	http.Error(w, err.Error(), 500)
	////	return
	////}
	//fmt.Println(sortKey3)
	//fmt.Println(limit3)

	dataResult, err := h.dataService.GetData(r.Context(), params)
	if err != nil {
		render.Error(w, err)
		return
	}

	// Ok! 200
	render.OK(w, dataResult)
}
