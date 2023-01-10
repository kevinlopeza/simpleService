package rest

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"simpleService/internal"
	"simpleService/internal/process"
	"time"
)

type handler struct {
	cache process.Cache
}

func NewHandler(c process.Cache) http.Handler {
	return &handler{
		cache: c,
	}
}

func (h handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// Get query params from request
	holidayType := request.FormValue("type")
	beginDateParam := request.FormValue("beginDate")
	endDateParam := request.FormValue("endDate")

	if holidayType == "" || beginDateParam == "" || endDateParam == "" {
		log.Info("Empty query param detected")
		h.closeResponse(writer, http.StatusBadRequest, nil)
	}

	// Parse query params to dates
	beginDate, err := time.Parse(internal.Layout, beginDateParam)
	if err != nil {
		log.Info("Error parsing beginDate param", err.Error())
		h.closeResponse(writer, http.StatusBadRequest, nil)
		return
	}
	endDate, err := time.Parse(internal.Layout, endDateParam)
	if err != nil {
		log.Info("Error parsing endDate param", err.Error())
		h.closeResponse(writer, http.StatusBadRequest, nil)
		return
	}

	result := h.cache.GetHolidays(holidayType, beginDate, endDate)
	h.closeResponse(writer, http.StatusOK, result)
}

func (h handler) closeResponse(writer http.ResponseWriter, HTTPStatus int, result []*process.Holiday) {

	writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	err := json.NewEncoder(writer).Encode(result)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if HTTPStatus != http.StatusOK {
		writer.WriteHeader(HTTPStatus)
	}
}
