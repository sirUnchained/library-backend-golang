package routes

import (
	"libBackend/controller"
	"net/http"
)

type BookUserRoutes struct{}

func (bu *BookUserRoutes) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet:
		controller.GetAll(w, r)
	case r.Method == http.MethodPost:
		controller.Create(w, r)
	}
}
