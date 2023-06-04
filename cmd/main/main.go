package main

import (
	"gitub.com/jibaru/ddd-golang-mongodb-template/internal/routes"
	"gitub.com/jibaru/ddd-golang-mongodb-template/pkg/logger"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	routes.Use(mux)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		logger.Log(err.Error())
	}
}
