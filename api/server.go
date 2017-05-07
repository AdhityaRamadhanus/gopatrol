package api

import (
	"net/http"
	"time"

	"github.com/AdhityaRamadhanus/gopatrol/config"
	"github.com/gorilla/mux"
)

type Server struct {
	Config   *config.Config
	router   *mux.Router
	Handlers []HttpHandler
}

func NewApiServer() *Server {
	router := mux.NewRouter()
	return &Server{
		router: router,
		Config: config.GetDefaultConfig(),
	}
}

func (s *Server) AddHandler(handler HttpHandler) {
	s.Handlers = append(s.Handlers, handler)
}

func (s *Server) InitHandler() {
	for _, handler := range s.Handlers {
		handler.AddRoutes(s.router)
	}
}

func (s *Server) ListenAndServe() error {
	srv := &http.Server{
		Handler:      s.router,
		Addr:         s.Config.Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	return srv.ListenAndServe()
}
