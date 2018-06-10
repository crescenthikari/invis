package router

import (
	"github.com/crescenthikari/invis/http/middleware"
	"github.com/gorilla/mux"
	"github.com/crescenthikari/invis/http/handler"
)

func CreateRoutes() *mux.Router {
	r := mux.NewRouter()
	// Add your routes as needed
	r.HandleFunc("/", handler.RootHandler)
	r.Use(middleware.LoggingMiddleware)

	return r
}
