package main

import "net/http"

// registerRoutes registers the routes with the provided *http.ServeMux
func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", home)
	mux.HandleFunc("/login", loginPage)
	mux.HandleFunc("/api/submitContact", submitContact)
}
