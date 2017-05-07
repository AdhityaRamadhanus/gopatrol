package api

import "github.com/gorilla/mux"

type HttpHandler interface {
	AddRoutes(router *mux.Router)
}
