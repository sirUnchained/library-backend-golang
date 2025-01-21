package routes

import (
	controller "libBackend/controller"
	"net/http"
)

type UserHandler struct{}

func (u *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet:
		controller.GetAllUsers(w, r)
	case r.Method == http.MethodPost:
		controller.Register(w, r)
	}
}
