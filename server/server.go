package server

import (
	routes "libBackend/routes"
	"net/http"
	"time"
)

func Start() {
	mux := http.NewServeMux()
	mux.Handle("/users", &routes.UserHandler{})
	mux.Handle("/books", &routes.BookRoutes{})

	server := http.Server{
		Addr:         ":4000",
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		Handler:      mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
