package panel

import (
	"fmt"

	"github.com/gorilla/mux"
)

// BindRoutes binds all the panel urls with the logic (handlers)
func BindRoutes(router *mux.Router) *mux.Router {
	prefix := fmt.Sprint("/p/")

	// Create the router if not new
	if router == nil {
		router = mux.NewRouter()
	}

	s := router.PathPrefix(prefix).Subrouter()

	// Handy routes
	s.HandleFunc("/", RootHandler).Methods("GET")

	return router
}
