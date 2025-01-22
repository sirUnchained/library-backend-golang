package routes

import "net/http"

type BookUserRoutes struct{}

func (bu *BookUserRoutes) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
