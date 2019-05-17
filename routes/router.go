package routes

import (
	"github.com/MaJloe3Jlo/perx_test_done/handler"
	"github.com/gorilla/mux"
	"net/http"
)

// Routes in app.
var routs = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handler.Index,
	},
	Route{
		"Key",
		"GET",
		"/key",
		handler.Key,
	},
	Route{
		"KeyInfo",
		"GET",
		"/key/{keyID}",
		handler.KeyInfo,
	},
	Route{
		"KeyOff",
		"POST",
		"/key",
		handler.KeyOff,
	},
	Route{
		"Statistics",
		"GET",
		"/statistics",
		handler.Statistics,
	},
}

// func NewRouter is define to create router with parameters.
func Router() *mux.Router {
	// New router start here.
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routs {
		var handler http.Handler

		// Handle with log.
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}

	return router
}
