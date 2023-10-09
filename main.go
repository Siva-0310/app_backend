package main

import (
	"net/http"
	"nfc/m/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.AllowContentType("application/json"))
	r.NotFound(NotFound)
	r.MethodNotAllowed(MethodNotAllowed)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! World"))
	})
	sign_up := routes.SignUp()
	r.Mount("/sign_up", sign_up)
	server := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: r,
	}
	server.ListenAndServe()
}