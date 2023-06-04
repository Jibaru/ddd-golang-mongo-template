package routes

import (
	"gitub.com/jibaru/ddd-golang-mongodb-template/internal/core/version/infrastructure/handlers"
	"net/http"
)

func Use(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/api/version", handlers.NewGetVersionHandler().Execute)
	return mux
}
