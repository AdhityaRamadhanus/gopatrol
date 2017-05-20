package handlers

import (
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	"github.com/AdhityaRamadhanus/gopatrol"
	"github.com/AdhityaRamadhanus/gopatrol/api/middlewares"
	"github.com/AdhityaRamadhanus/gopatrol/api/render"
	"github.com/gorilla/mux"
)

type EventsHandlers struct {
	EventService gopatrol.EventService
	CacheService gopatrol.CacheService
}

func (h *EventsHandlers) AddRoutes(router *mux.Router, isUnixDomain bool) {
	if !isUnixDomain {
		router.HandleFunc("/events/all", middlewares.AuthenticateToken(middlewares.Gzip(http.HandlerFunc(h.GetAllEvents)), 2)).Methods("GET")
	} else {
		router.HandleFunc("/events/all", h.GetAllEvents).Methods("GET")
	}
}

func (h *EventsHandlers) GetAllEvents(res http.ResponseWriter, req *http.Request) {
	page := 0
	size := 10
	// Take Querystring
	queryStrings := req.URL.Query()
	if len(queryStrings["page"]) > 0 {
		page, _ = strconv.Atoi(queryStrings["page"][0])
	}
	if len(queryStrings["size"]) > 0 {
		size, _ = strconv.Atoi(queryStrings["size"][0])
	}

	// cacheKey := fmt.Sprintf("events:%d:%d", page, size)
	// respBytes, err := h.CacheService.Get(cacheKey)
	// if err == nil {
	// 	helper.WriteGzipBytes(res, req, http.StatusOK, respBytes)
	// 	return
	// }
	query := bson.M{}
	counts, _ := h.EventService.CountEvents(query)
	events, _ := h.EventService.GetAllEvents(map[string]interface{}{
		"query":      query,
		"pagination": true,
		"page":       page,
		"limit":      size,
	})
	response := render.JSON{
		"pagination": map[string]int{
			"total": counts,
			"page":  page,
			"size":  size,
		},
		"events": events,
	}
	// h.CacheService.Set(cacheKey, respBytes)
	render.WriteJSON(res, http.StatusOK, response)
}
