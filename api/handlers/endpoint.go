package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/AdhityaRamadhanus/gopatrol"
	"github.com/AdhityaRamadhanus/gopatrol/api/helper"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

type EndpointHandler struct {
	EndpointService gopatrol.EndpointService
	CacheService    gopatrol.CacheService
}

func (h *EndpointHandler) AddRoutes(router *mux.Router) {
	router.
		PathPrefix("/endpoints").
		Path("/").
		Methods("GET").
		HandlerFunc(h.GetAllEndpoints)

	router.
		PathPrefix("/endpoints").
		Path("/").
		Methods("POST").
		HandlerFunc(h.CreateEndpoint)

	router.
		PathPrefix("/endpoints").
		Path("/{slug}").
		Methods("GET").
		HandlerFunc(h.GetOneBySlug)

	router.
		PathPrefix("/endpoints").
		Path("/{slug}").
		Methods("DELETE").
		HandlerFunc(h.DeleteOneBySlug)
}

func (h *EndpointHandler) CreateEndpoint(res http.ResponseWriter, req *http.Request) {
	// Read Body, limit to 1 MB //
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		log.Println(err)
		helper.WriteJSON(res, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	if err := req.Body.Close(); err != nil {
		log.Println(err)
		helper.WriteJSON(res, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	endpoint := gopatrol.Endpoint{
		Attempts:     5,
		ThresholdRTT: 2000,
	}
	// Deserialize
	if err := json.Unmarshal(body, &endpoint); err != nil {
		log.Println(err)
		helper.WriteJSON(res, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	switch endpoint.Type {
	case "http":
		httpChecker := gopatrol.HTTPChecker{}
		// Deserialize
		if err := json.Unmarshal(body, &httpChecker); err != nil {
			log.Println(err)
			helper.WriteJSON(res, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		httpChecker.Slug = helper.Slugify(httpChecker.Name)
		if ok, err := govalidator.ValidateStruct(httpChecker); !ok || err != nil {
			helper.WriteJSON(res, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		if err := h.EndpointService.InsertEndpoint(endpoint); err != nil {
			log.Println(err)
			helper.WriteJSON(res, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		helper.WriteJSON(res, http.StatusCreated, "endpoint Created")
	case "tcp":
		tcpChecker := gopatrol.TCPChecker{}
		// Deserialize
		if err := json.Unmarshal(body, &tcpChecker); err != nil {
			log.Println(err)
			helper.WriteJSON(res, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		tcpChecker.Slug = helper.Slugify(tcpChecker.Name)
		if ok, err := govalidator.ValidateStruct(tcpChecker); !ok || err != nil {
			helper.WriteJSON(res, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		if err := h.EndpointService.InsertEndpoint(tcpChecker); err != nil {
			log.Println(err)
			helper.WriteJSON(res, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		helper.WriteJSON(res, http.StatusCreated, "endpoint Created")
	default:
		helper.WriteJSON(res, http.StatusBadRequest, "Unknown Endpoint Type")
	}
}

func (h *EndpointHandler) GetAllEndpoints(res http.ResponseWriter, req *http.Request) {
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
	endpoints, err := h.EndpointService.GetAllEndpoints(
		&bson.M{},
		page,
		size)
	if err != nil {
		log.Println(err)
		helper.WriteJSON(res, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	// redisKey := fmt.Sprintf("products:%d:%d", page, size)
	// respBytes, err := h.CacheService.GetKey(redisKey)
	// if err == nil {
	// 	helper.WriteGzipBytes(res, req, http.StatusOK, respBytes)
	// 	return
	// }
	response := map[string]interface{}{
		"page":      page,
		"size":      size,
		"endpoints": endpoints,
	}
	respBytes, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		helper.WriteJSON(res, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	// h.CacheService.SetKey(redisKey, respBytes, time.Second*5)
	helper.WriteGzipBytes(res, req, http.StatusOK, respBytes)
}

func (h *EndpointHandler) GetOneBySlug(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	slug := params["slug"]
	endpoint, err := h.EndpointService.GetEndpointBySlug(slug)
	if err != nil {
		log.Println(err)
		helper.WriteJSON(res, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	response := map[string]interface{}{
		"data": endpoint,
	}
	respBytes, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		helper.WriteJSON(res, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	// h.CacheService.SetKey(redisKey, respBytes, time.Second*5)
	helper.WriteGzipBytes(res, req, http.StatusOK, respBytes)
}

func (h *EndpointHandler) DeleteOneBySlug(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	slug := params["slug"]
	err := h.EndpointService.DeleteEndpointBySlug(slug)
	if err != nil {
		log.Println(err)
		helper.WriteJSON(res, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	// h.CacheService.SetKey(redisKey, respBytes, time.Second*5)
	helper.WriteJSON(res, http.StatusOK, "Endpoint Deleted")
}
