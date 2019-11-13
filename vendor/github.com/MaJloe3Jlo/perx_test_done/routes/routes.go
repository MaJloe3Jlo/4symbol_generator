package routes

import "net/http"

// type Route is an define:
// * Name of route - E.g. "Index"
// * Method of route - GET or POST
// * Pattern of route - E.g. "/index"
// * And func to handle this - all handlers in path "handler"
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// slice of Routes.
type Routes []Route
