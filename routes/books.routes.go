package routes

import "net/http"

type BookRoutes struct{}

func (b *BookRoutes) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
