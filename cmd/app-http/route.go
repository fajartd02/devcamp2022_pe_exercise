package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func newRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.MethodFunc(http.MethodGet, "/v1/book", HelloServer)

	return router
}
