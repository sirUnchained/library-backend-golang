package server

import (
	"net/http"
	"time"
)

func Start() {
	server := http.Server{
		Addr:         ":4000",
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
