package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/AdhityaRamadhanus/gopatrol"
	"github.com/AdhityaRamadhanus/gopatrol/api"
	"github.com/AdhityaRamadhanus/gopatrol/api/helper"
	"github.com/AdhityaRamadhanus/gopatrol/api/middlewares"
	"github.com/AdhityaRamadhanus/gopatrol/api/render"
	log "github.com/Sirupsen/logrus"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CheckersHandler struct {
	CheckerService gopatrol.CheckersService
	CacheService   gopatrol.CacheService
}

func (h *CheckersHandler) AddRoutes(router *mux.Router, isUnixDomain bool) {
	if !isUnixDomain {
		router.HandleFunc("/checkers/all", middlewares.AuthenticateToken(middlewares.Gzip(http.HandlerFunc(h.GetAllEndpoints)), 2)).Methods("GET")
		// router.HandleFunc("/checkers/all", middlewares.AuthenticateToken(http.HandlerFunc(h.GetAllEndpoints), 2)).Methods("GET")
		router.HandleFunc("/checkers/create", middlewares.AuthenticateToken(http.HandlerFunc(h.CreateChecker), 1)).Methods("POST")
		router.HandleFunc("/checkers/{slug}", middlewares.AuthenticateToken(http.HandlerFunc(h.GetOneBySlug), 2)).Methods("GET")
		router.HandleFunc("/checkers/{slug}/delete", middlewares.AuthenticateToken(http.HandlerFunc(h.DeleteOneBySlug), 1)).Methods("DELETE")
	} else {
		router.HandleFunc("/checkers/all", h.GetAllEndpoints).Methods("GET")
		router.HandleFunc("/checkers/create", h.CreateChecker).Methods("POST")
		router.HandleFunc("/checkers/stats", h.GetStatisticCheckers).Methods("GET")
		router.HandleFunc("/checkers/{slug}/delete", h.DeleteOneBySlug).Methods("POST")
	}
}

func (h *CheckersHandler) CreateChecker(res http.ResponseWriter, req *http.Request) {
	// Read Body, limit to 1 MB //
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		render.WriteJSON(res, http.StatusInternalServerError, render.JSON{
			"error": api.ErrFailedToReadBody,
		})
		return
	}
	if err := req.Body.Close(); err != nil {
		render.WriteJSON(res, http.StatusInternalServerError, render.JSON{
			"error": api.ErrFailedToCloseBody,
		})
		return
	}
	endpoint := struct {
		Type string `json:"type"`
	}{}
	// Deserialize
	if err := json.Unmarshal(body, &endpoint); err != nil {
		render.WriteJSON(res, http.StatusInternalServerError, render.JSON{
			"error": api.ErrFailedToUnmarshalJSON,
		})
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
			render.WriteJSON(res, http.StatusInternalServerError, render.JSON{
				"error": api.ErrFailedToUnmarshalJSON,
			})
			return
		}
		httpChecker.Slug = helper.Slugify(httpChecker.Name)
		if ok, err := govalidator.ValidateStruct(httpChecker); !ok || err != nil {
			render.WriteJSON(res, http.StatusBadRequest, render.JSON{
				"error": api.ErrFailedToValidateStruct,
			})
			return
		}
		if err := h.CheckerService.InsertChecker(httpChecker); err != nil {
			render.WriteJSON(res, http.StatusInternalServerError, render.JSON{
				"error": err,
			})
			return
		}
		log.WithFields(log.Fields{
			"name": httpChecker.Name,
			"url":  httpChecker.URL,
		}).Info("Add Endpoint")

		render.WriteJSON(res, http.StatusCreated, render.JSON{
			"message": "Endpoint created",
		})
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
			render.WriteJSON(res, http.StatusInternalServerError, render.JSON{
				"error": api.ErrFailedToUnmarshalJSON,
			})
			return
		}
		tcpChecker.Slug = helper.Slugify(tcpChecker.Name)
		if ok, err := govalidator.ValidateStruct(tcpChecker); !ok || err != nil {
			render.WriteJSON(res, http.StatusBadRequest, render.JSON{
				"error": api.ErrFailedToValidateStruct,
			})
			return
		}
		if err := h.CheckerService.InsertChecker(tcpChecker); err != nil {
			render.WriteJSON(res, http.StatusInternalServerError, render.JSON{
				"error": err,
			})
			return
		}

		log.WithFields(log.Fields{
			"name": tcpChecker.Name,
			"url":  tcpChecker.URL,
		}).Info("Add Endpoint")

		render.WriteJSON(res, http.StatusCreated, render.JSON{
			"message": "Endpoint Created",
		})
		return
	case "dns":
		dnsChecker := gopatrol.DNSChecker{
			Timeout:     time.Second * 10,
			Attempts:    5,
			LastChecked: time.Now(),
			LastStatus:  "",
			LastChange:  time.Now(),
		}
		// Deserialize
		if err := json.Unmarshal(body, &dnsChecker); err != nil {
			render.WriteJSON(res, http.StatusInternalServerError, render.JSON{
				"error": api.ErrFailedToUnmarshalJSON,
			})
			return
		}
		dnsChecker.Slug = helper.Slugify(dnsChecker.Name)
		if ok, err := govalidator.ValidateStruct(dnsChecker); !ok || err != nil {
			render.WriteJSON(res, http.StatusBadRequest, render.JSON{
				"error": api.ErrFailedToValidateStruct,
			})
			return
		}
		if err := h.CheckerService.InsertChecker(dnsChecker); err != nil {
			render.WriteJSON(res, http.StatusInternalServerError, render.JSON{
				"error": err,
			})
			return
		}

		log.WithFields(log.Fields{
			"name": dnsChecker.Name,
			"url":  dnsChecker.URL,
		}).Info("Add Endpoint")

		render.WriteJSON(res, http.StatusCreated, render.JSON{
			"message": "Endpoint created",
		})
		return
	default:
		render.WriteJSON(res, http.StatusBadRequest, render.JSON{
			"error": api.ErrUnknownEndpointTYpe,
		})
		return
	}
}

func (h *CheckersHandler) GetAllEndpoints(res http.ResponseWriter, req *http.Request) {
	page := 0
	size := 10
	response := map[string]interface{}{}
	// Take Querystring
	queryStrings := req.URL.Query()
	query := bson.M{}

	if len(queryStrings["status"]) > 0 {
		query["status"] = queryStrings["status"][0]
	}

	// Pagination is true
	if val, ok := queryStrings["pagination"]; ok && val[0] == "false" {
		checkers, _ := h.CheckerService.GetAllCheckers(map[string]interface{}{
			"query":      query,
			"pagination": false,
			"select":     bson.M{"name": 1, "laststatus": 1, "type": 1, "url": 1, "lastchecked": 1, "slug": 1},
		})
		response = render.JSON{
			"checkers": checkers,
		}
	} else {
		if len(queryStrings["page"]) > 0 {
			page, _ = strconv.Atoi(queryStrings["page"][0])
		}
		if len(queryStrings["size"]) > 0 {
			size, _ = strconv.Atoi(queryStrings["size"][0])
		}
		counts, _ := h.CheckerService.CountCheckers(query)
		checkers, _ := h.CheckerService.GetAllCheckers(map[string]interface{}{
			"query":      query,
			"pagination": true,
			"page":       page,
			"limit":      size,
			"select":     bson.M{"name": 1, "laststatus": 1, "type": 1, "url": 1, "lastchecked": 1, "slug": 1},
		})
		response = render.JSON{
			"pagination": map[string]int{
				"total": counts,
				"page":  page,
				"size":  size,
			},
			"checkers": checkers,
		}
	}
	render.WriteJSON(res, http.StatusOK, response)
}

func (h *CheckersHandler) GetOneBySlug(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	slug := params["slug"]
	endpoint, err := h.CheckerService.GetCheckerBySlug(slug)
	if err != nil {
		switch err {
		case mgo.ErrNotFound:
			render.WriteJSON(res, http.StatusNotFound, render.JSON{
				"error": "Cannot find Checker",
			})
			return
		default:
			log.WithError(err).Error("Failed to get checker")
			render.WriteJSON(res, http.StatusInternalServerError, render.JSON{
				"error": api.ErrInternalServerError,
			})
			return
		}
		return
	}
	render.WriteJSON(res, http.StatusOK, render.JSON{
		"checker": endpoint,
	})
}

func (h *CheckersHandler) DeleteOneBySlug(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	slug := params["slug"]
	err := h.CheckerService.DeleteCheckerBySlug(slug)
	if err != nil {
		switch err {
		case mgo.ErrNotFound:
			render.WriteJSON(res, http.StatusNotFound, render.JSON{
				"error": "Cannot find Checker",
			})
			return
		default:
			log.WithError(err).Error("Failed to get checker")
			render.WriteJSON(res, http.StatusInternalServerError, render.JSON{
				"error": api.ErrInternalServerError,
			})
			return
		}
	}
	render.WriteJSON(res, http.StatusOK, render.JSON{
		"message": "Endpoint Deleted",
	})
}

func (h *CheckersHandler) GetStatisticCheckers(res http.ResponseWriter, req *http.Request) {
	countsHealthy, _ := h.CheckerService.CountCheckers(bson.M{
		"laststatus": "healthy",
	})
	countsDown, _ := h.CheckerService.CountCheckers(bson.M{
		"laststatus": "down",
	})
	countsAll, _ := h.CheckerService.CountCheckers(bson.M{})

	response := render.JSON{
		"checkers": map[string]int{
			"healthy": countsHealthy,
			"down":    countsDown,
			"all":     countsAll,
			"unknown": countsAll - countsHealthy - countsDown,
		},
	}

	// h.CacheService.Set(cacheKey, respBytes)
	render.WriteJSON(res, http.StatusOK, response)
}
