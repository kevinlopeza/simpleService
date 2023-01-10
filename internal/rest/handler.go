package rest

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"simpleService/internal"
	"simpleService/internal/cache"
	"time"
)

// handler implements the http.Handler interface
type handler struct {
	cache cache.Cache
}

// NewHandler is a constructor to inject a cache implementation to
// a handler struct
func NewHandler(c cache.Cache) http.Handler {
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
		return
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

func (h handler) closeResponse(writer http.ResponseWriter, HTTPStatus int, result []*cache.Holiday) {

	writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	if HTTPStatus != http.StatusOK {
		writer.WriteHeader(HTTPStatus)
	}

	err := json.NewEncoder(writer).Encode(result)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

}
