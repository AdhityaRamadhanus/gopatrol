package api

import (
	"net/http"
	"time"

	"github.com/AdhityaRamadhanus/gopatrol/config"
	"github.com/gorilla/mux"
)

type ApiHandler interface {
	AddRoutes(router *mux.Router)
}

type Api struct {
	Config   *config.Config
	router   *mux.Router
	Handlers []ApiHandler
}

func NewApi() *Api {
	router := mux.NewRouter().StrictSlash(true)
	return &Api{
		router: router,
		Config: config.GetDefaultConfig(),
	}
}

func (api *Api) InitHandler() {
	for _, handler := range api.Handlers {
		handler.AddRoutes(api.router)
	}
}

func (api *Api) CreateServer() *http.Server {
	srv := &http.Server{
		Handler:      api.router,
		Addr:         api.Config.Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	return srv
}
