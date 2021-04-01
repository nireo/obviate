package obviate

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Router struct {
	handlers map[string]http.HandlerFunc
}

type Route interface {
	Head() string
	Logger() bool
	Tester() bool
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]http.HandlerFunc),
	}
}

func combine(method string, path string) string {
	return fmt.Sprintf("%s:%s", method, path)
}

func (r *Router) GET(path string, handler http.HandlerFunc) {
	r.handlers[combine(http.MethodGet, path)] = handler
}

func (r *Router) POST(path string, handler http.HandlerFunc) {
	r.handlers[combine(http.MethodPost, path)] = handler
}

func (r *Router) DELETE(path string, handler http.HandlerFunc) {
	r.handlers[combine(http.MethodDelete, path)] = handler
}

func (r *Router) PUT(path string, handler http.HandlerFunc) {
	r.handlers[combine(http.MethodPut, path)] = handler
}

func (r *Router) PATCH(path string, handler http.HandlerFunc) {
	r.handlers[combine(http.MethodPatch, path)] = handler
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Could not find path: %s", r.URL.Path)))
}

func ParseJSON(body io.ReadCloser, res interface{}) error {
	return json.NewDecoder(body).Decode(res)
}

func Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) error {
	if e, ok := data.(error); ok {
		tmp := struct{
			Status string `json:"status"`
			Error string `json:"error"`
		}{
			"error",
			e.Error(),
		}

		data = tmp
	}

	j, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(j)
	LogRequest(r, code)
	return nil
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	f, ok := r.handlers[combine(req.Method, req.URL.Path)]
	if !ok {
		notFound(w, req)
		return
	}

	f(w, req)
}

func (r *Router) Listen(port string) {
	session := &http.Server{
		Addr:           port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(session.ListenAndServe())
}

