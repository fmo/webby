package main

import (
	"net/http"

	"github.com/fmo/webb"
)

func main() {
	router := webb.NewRouter()

	router.Get("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}))

	router.Post("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("there"))
	}))

	http.ListenAndServe(":8080", router)
}
