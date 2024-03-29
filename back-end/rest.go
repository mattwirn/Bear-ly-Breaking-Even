package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	hnd "github.com/mattwirn/Bear-ly-Breaking-Even/back-end/handlers"
	"github.com/mattwirn/Bear-ly-Breaking-Even/back-end/middleware"
)

// httpHandler creates the backend HTTP router for queries, types,
// and serving the React frontend.
func httpHandler() http.Handler {
	router := mux.NewRouter()
	// Your REST API requests go here

	router.HandleFunc("/signup", hnd.Signup).Methods("POST")
	router.HandleFunc("/login", hnd.Login).Methods("POST")

	val := router.PathPrefix("/dashboard").Subrouter()
	val.HandleFunc("/get", hnd.Dashboard).Methods("GET")
	val.HandleFunc("/expense/add", hnd.AddExpense).Methods("POST")
	val.HandleFunc("/expense/update", hnd.UpdateExpense).Methods("PUT")
	val.HandleFunc("/expense/delete", hnd.DeleteExpense).Methods("DELETE")
	val.HandleFunc("/income/update", hnd.InputIncome).Methods("POST")
	val.Use(middleware.RequireAuth)
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
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}),
			handlers.AllowedOrigins([]string{"http://localhost:3000"}),
			handlers.ExposedHeaders([]string{"DNT", "Keep-Alive", "User-Agent",
				"X-Requested-With", "If-Modified-Since", "Cache-Control",
				"Content-Type", "Content-Range", "Range", "Content-Disposition"}),
			handlers.MaxAge(86400),
		)(router))
}
