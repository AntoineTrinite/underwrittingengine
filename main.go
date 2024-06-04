package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	router := chi.NewMux()
	router.Get("/test", HandleTest)

	http.ListenAndServe(":3000", router)
}

func HandleTest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello test"))
}
