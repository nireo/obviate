package router

import (
	"net/http"
)

type HandlerFunc = func(http.ResponseWriter, *http.Request)

type Router struct {
	URL      string
	Handlers []Handler
}

type Handler struct {
	Method  string
	Handler HandlerFunc
}

func NewRouter() *Router {
	var handlers []Handler
	return &Router{
		URL:      "",
		Handlers: handlers,
	}
}

func (r *Router) NewHandler(method string, handler HandlerFunc) {
	newHandler := Handler{
		Method:  method,
		Handler: handler,
	}

	r.Handlers = append(r.Handlers, newHandler)
}

func (r *Router) SetUrl(newUrl string) error {
	r.URL = newUrl
	return nil
}

func validMethod(method string) bool {
	switch method {
	case http.MethodPost:
		return true
	case http.MethodConnect:
		return true
	case http.MethodGet:
		return true
	case http.MethodDelete:
		return true
	case http.MethodPut:
		return true
	case http.MethodOptions:
		return true
	case http.MethodPatch:
		return true
	case http.MethodTrace:
		return true
	default:
		// since a match wasn't found the method isn't valid
		return false
	}
}
