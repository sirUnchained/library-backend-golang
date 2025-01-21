package routes

import (
	"libBackend/controller"
	"net/http"
)

type BookRoutes struct{}

func (b *BookRoutes) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet:
		controller.GetAllBooks(w, r)
	case r.Method == http.MethodPost:
		controller.CreateBook(w, r)
	}
}
