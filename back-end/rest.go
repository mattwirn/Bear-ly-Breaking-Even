package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	hnd "github.com/mattwirn/Bear-ly-Breaking-Even/back-end/handlers"
)

// httpHandler creates the backend HTTP router for queries, types,
// and serving the React frontend.
func httpHandler() http.Handler {
	router := mux.NewRouter()
	// Your REST API requests go here
	//router.Use(middleware.RequireAuth)

	router.HandleFunc("/signup", hnd.Signup).Methods("POST")
	router.HandleFunc("/login", hnd.Login).Methods("POST")
	router.HandleFunc("/validate", hnd.Validate).Methods("GET")
	// Add your routes here.
	// WARNING: this route must be the last route defined.

	router.PathPrefix("/").Handler(ReactHandler).Methods("GET")

	/**
	 * We need some headers to be statically prepended to every response.
	 */
	return handlers.LoggingHandler(os.Stdout,
		handlers.CORS(
			handlers.AllowCredentials(),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization",
				"DNT", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since",
				"Cache-Control", "Content-Range", "Range"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"http://localhost:8080"}),
			handlers.ExposedHeaders([]string{"DNT", "Keep-Alive", "User-Agent",
				"X-Requested-With", "If-Modified-Since", "Cache-Control",
				"Content-Type", "Content-Range", "Range", "Content-Disposition"}),
			handlers.MaxAge(86400),
		)(router))
}
