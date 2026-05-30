package webb

import (
	"fmt"
	"net/http"
)

type Router struct {
	handlers map[string]http.Handler
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]http.Handler),
	}
}

func (rt Router) Get(path string, handler http.Handler) {
	fullHandler := fmt.Sprintf("GET-%s", path)
	rt.handlers[fullHandler] = handler
}

func (rt Router) Post(path string, handler http.Handler) {
	fullHandler := fmt.Sprintf("POST-%s", path)
	rt.handlers[fullHandler] = handler
}

func (rt Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	fullHandler := fmt.Sprintf("%s-%s", r.Method, path)

	if handler, ok := rt.handlers[fullHandler]; ok {
		handler.ServeHTTP(w, r)
	} else {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
}
