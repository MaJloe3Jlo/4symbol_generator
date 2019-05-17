package routes

import (
	"log"
	"net/http"
	"time"
)

// func Logger is defined to logging all processes in app.
func Logger(inside http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Define starting time.
		begin := time.Now()

		// Serve inside func.
		inside.ServeHTTP(w, r)

		// Logging with parameters.
		log.Printf("%s\t%s\t%s\t%s", r.Method, r.RequestURI, name, time.Since(begin))
	})
}
