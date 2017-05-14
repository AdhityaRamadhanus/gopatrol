package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	"github.com/AdhityaRamadhanus/gopatrol"
	"github.com/AdhityaRamadhanus/gopatrol/api"
	"github.com/AdhityaRamadhanus/gopatrol/api/helper"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

type EventsHandlers struct {
	EventService gopatrol.EventService
	CacheService gopatrol.CacheService
}

func (h *EventsHandlers) AddRoutes(router *mux.Router) {
	router.HandleFunc("/events/all", h.GetAllEvents).Methods("GET")
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
		size, _ = strconv.Atoi(queryStrings["limit"][0])
	}

	cacheKey := fmt.Sprintf("events:%d:%d", page, size)
	respBytes, err := h.CacheService.Get(cacheKey)
	if err == nil {
		helper.WriteGzipBytes(res, req, http.StatusOK, respBytes)
		return
	}

	events, _ := h.EventService.GetAllEvents(map[string]interface{}{
		"query":      bson.M{},
		"pagination": true,
		"page":       page,
		"limit":      size,
	})
	response := map[string]interface{}{
		"page":   page,
		"size":   size,
		"events": events,
	}
	respBytes, err = json.Marshal(response)
	if err != nil {
		log.Println(err)
		helper.WriteJSON(res, http.StatusInternalServerError, api.ErrInternalServerError)
		return
	}
	h.CacheService.Set(cacheKey, respBytes)
	helper.WriteGzipBytes(res, req, http.StatusOK, respBytes)
	return
}
