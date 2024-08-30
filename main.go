package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router:= chi.NewRouter()

	router.Use(middleware.Logger)

	server := &http.Server{
		Addr: ":3000",
		Handler: router,
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	err := server.ListenAndServe()

	if (err != nil) {
		fmt.Println("Failed to start server", err)
	}

}