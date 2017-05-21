package api

import (
	"net/http"
	"time"

	"github.com/AdhityaRamadhanus/gopatrol/api/middlewares"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type ApiHandler interface {
	AddRoutes(router *mux.Router, isUnixDomain bool)
}

type Api struct {
	router   *mux.Router
	Handlers []ApiHandler
}

func NewApi() *Api {
	router := mux.NewRouter().StrictSlash(true)
	apirouter := router.PathPrefix("/api/v1").Subrouter()
	return &Api{
		router: apirouter,
	}
}

func (api *Api) InitHandler(isUnixDomain bool) {
	for _, handler := range api.Handlers {
		handler.AddRoutes(api.router, isUnixDomain)
	}
}

func (api *Api) CreateServer(isUnixDomain bool) *http.Server {
	var srv *http.Server
	if !isUnixDomain {
		srv = &http.Server{
			Handler:      cors.Default().Handler(middlewares.HTTPReqLogger(api.router)),
			Addr:         ":3000",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  5 * time.Second,
		}
	} else {
		srv = &http.Server{
			Handler:      middlewares.HTTPReqLogger(api.router),
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  5 * time.Second,
		}
	}
	return srv
}
