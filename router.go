package main

import (
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	var prefix string
	if os.Getenv("PREFIX") == "" {
		prefix = "/mongo-backup"
	} else {
		prefix = path.Join("/", os.Getenv("PREFIX"))
	}

	router := mux.NewRouter().StrictSlash(true)
	s := router.PathPrefix(prefix).Subrouter()
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger(handler, route.Name)

		s.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return s
}
