package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/AdhityaRamadhanus/gopatrol"
	"github.com/AdhityaRamadhanus/gopatrol/api"
	"github.com/AdhityaRamadhanus/gopatrol/api/helper"
	log "github.com/Sirupsen/logrus"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

type CheckersHandler struct {
	CheckerService gopatrol.CheckersService
	CacheService   gopatrol.CacheService
}

func (h *CheckersHandler) AddRoutes(router *mux.Router) {
	router.HandleFunc("/checkers/all", h.GetAllEndpoints).Methods("GET")
	router.HandleFunc("/checkers/create", h.CreateChecker).Methods("POST")
	router.HandleFunc("/checkers/{slug}", h.GetOneBySlug).Methods("GET")
	router.HandleFunc("/checkers/{slug}/delete", h.DeleteOneBySlug).Methods("DELETE")
}

func (h *CheckersHandler) CreateChecker(res http.ResponseWriter, req *http.Request) {
	// Read Body, limit to 1 MB //
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		helper.WriteJSON(res, http.StatusInternalServerError, api.ErrFailedToReadBody)
		return
	}
	if err := req.Body.Close(); err != nil {
		helper.WriteJSON(res, http.StatusInternalServerError, api.ErrFailedToCloseBody)
		return
	}
	endpoint := struct {
		Type string `json:"type"`
	}{}
	// Deserialize
	if err := json.Unmarshal(body, &endpoint); err != nil {
		helper.WriteJSON(res, http.StatusInternalServerError, api.ErrFailedToUnmarshalJSON)
		return
	}
	switch endpoint.Type {
	case "http":
		httpChecker := gopatrol.HTTPChecker{
			Attempts:    5,
			LastChecked: time.Now(),
			LastStatus:  "",
			LastChange:  time.Now(),
		}
		// Deserialize
		if err := json.Unmarshal(body, &httpChecker); err != nil {
			helper.WriteJSON(res, http.StatusInternalServerError, api.ErrFailedToUnmarshalJSON)
			return
		}
		httpChecker.Slug = helper.Slugify(httpChecker.Name)
		if ok, err := govalidator.ValidateStruct(httpChecker); !ok || err != nil {
			helper.WriteJSON(res, http.StatusInternalServerError, api.ErrFailedToValidateStruct)
			return
		}
		if err := h.CheckerService.InsertChecker(httpChecker); err != nil {
			helper.WriteJSON(res, http.StatusInternalServerError, err)
			return
		}
		log.WithFields(log.Fields{
			"name": httpChecker.Name,
			"url":  httpChecker.URL,
		}).Info("Add Endpoint")

		helper.WriteJSON(res, http.StatusCreated, "Endpoint created")
		return
	case "tcp":
		tcpChecker := gopatrol.TCPChecker{
			Timeout:     time.Second * 10,
			Attempts:    5,
			LastChecked: time.Now(),
			LastStatus:  "",
			LastChange:  time.Now(),
		}
		// Deserialize
		if err := json.Unmarshal(body, &tcpChecker); err != nil {
			helper.WriteJSON(res, http.StatusInternalServerError, api.ErrFailedToUnmarshalJSON)
			return
		}
		tcpChecker.Slug = helper.Slugify(tcpChecker.Name)
		if ok, err := govalidator.ValidateStruct(tcpChecker); !ok || err != nil {
			helper.WriteJSON(res, http.StatusInternalServerError, api.ErrFailedToValidateStruct)
			return
		}
		if err := h.CheckerService.InsertChecker(tcpChecker); err != nil {
			helper.WriteJSON(res, http.StatusInternalServerError, err)
			return
		}

		log.WithFields(log.Fields{
			"name": tcpChecker.Name,
			"url":  tcpChecker.URL,
		}).Info("Add Endpoint")

		helper.WriteJSON(res, http.StatusCreated, "Endpoint created")
		return
	case "dns":
		dnsChecker := gopatrol.DNSChecker{
			Timeout:     3000000000,
			Attempts:    5,
			LastChecked: time.Now(),
			LastStatus:  "",
			LastChange:  time.Now(),
		}
		// Deserialize
		if err := json.Unmarshal(body, &dnsChecker); err != nil {
			helper.WriteJSON(res, http.StatusInternalServerError, api.ErrFailedToUnmarshalJSON)
			return
		}
		dnsChecker.Slug = helper.Slugify(dnsChecker.Name)
		if ok, err := govalidator.ValidateStruct(dnsChecker); !ok || err != nil {
			helper.WriteJSON(res, http.StatusInternalServerError, api.ErrFailedToValidateStruct)
			return
		}
		if err := h.CheckerService.InsertChecker(dnsChecker); err != nil {
			helper.WriteJSON(res, http.StatusInternalServerError, err)
			return
		}

		log.WithFields(log.Fields{
			"name": dnsChecker.Name,
			"url":  dnsChecker.URL,
		}).Info("Add Endpoint")

		helper.WriteJSON(res, http.StatusCreated, "Endpoint created")
		return
	default:
		helper.WriteJSON(res, http.StatusBadRequest, api.ErrUnknownEndpointTYpe)
		return
	}
}

func (h *CheckersHandler) GetAllEndpoints(res http.ResponseWriter, req *http.Request) {
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
	checkers, _ := h.CheckerService.GetAllCheckers(map[string]interface{}{
		"query":      bson.M{},
		"pagination": true,
		"page":       page,
		"limit":      size,
	})
	response := map[string]interface{}{
		"page":      page,
		"size":      size,
		"endpoints": checkers,
	}
	respBytes, err := json.Marshal(response)
	if err != nil {
		helper.WriteJSON(res, http.StatusInternalServerError, api.ErrFailedToMarshalJSON)
		return
	}
	helper.WriteGzipBytes(res, req, http.StatusOK, respBytes)
}

func (h *CheckersHandler) GetOneBySlug(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	slug := params["slug"]
	endpoint, err := h.CheckerService.GetCheckerBySlug(slug)
	if err != nil {
		switch err {
		case mgo.ErrNotFound:
			helper.WriteJSON(res, http.StatusNotFound, "Cannot find Checker")
			return
		default:
			log.WithError(err).Error("Failed to get checker")
			helper.WriteJSON(res, http.StatusInternalServerError, api.ErrInternalServerError)
			return
		}
		return
	}
	response := map[string]interface{}{
		"data": endpoint,
	}
	respBytes, err := json.Marshal(response)
	if err != nil {
		helper.WriteJSON(res, http.StatusInternalServerError, api.ErrFailedToMarshalJSON)
		return
	}
	helper.WriteGzipBytes(res, req, http.StatusOK, respBytes)
}

func (h *CheckersHandler) DeleteOneBySlug(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	slug := params["slug"]
	err := h.CheckerService.DeleteCheckerBySlug(slug)
	if err != nil {
		switch err {
		case mgo.ErrNotFound:
			helper.WriteJSON(res, http.StatusNotFound, "Cannot find Checker")
			return
		default:
			log.WithError(err).Error("Failed to get checker")
			helper.WriteJSON(res, http.StatusInternalServerError, api.ErrInternalServerError)
			return
		}
	}
	helper.WriteJSON(res, http.StatusOK, "Endpoint Deleted")
}
