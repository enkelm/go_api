package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type router struct{ *mux.Router }

func (r *router) handleGet(path string, fn http.HandlerFunc) *mux.Router {
	r.Router.HandleFunc(path, fn).Methods("GET")
	return r.Router
}

func (r *router) handlePost(path string, fn http.HandlerFunc) *mux.Router {
	r.Router.HandleFunc(path, fn).Methods("POST")
	return r.Router
}

func (r *router) handlePut(path string, fn http.HandlerFunc) *mux.Router {
	r.Router.HandleFunc(path, fn).Methods("PUT")
	return r.Router
}
