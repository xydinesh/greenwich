package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes struct holds mapping for Route to handler
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes arrary with Route to Handler mapping
type Routes []Route

// NewRouter returns a router with routes populated
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

// Adding routes to handler mapping
var routes = Routes{
	Route{
		"Health",
		"GET",
		"/",
		Health,
	},
	Route{
		"TodoIndex",
		"GET",
		"/timezone/{coordinates}",
		TimeZoneShow,
	},
}
