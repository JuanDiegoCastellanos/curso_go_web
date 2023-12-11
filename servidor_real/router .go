package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

// ServeHTTP implements http.Handler
//
//	func (*Router) ServeHTTP(http.ResponseWriter, *http.Request) {
//		panic("unimplemented")
//	}
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, methodExist, exist := r.FindHandler(request.URL.Path, request.Method)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	handler(w, request)
}

func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, methodExist, exist

}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}
