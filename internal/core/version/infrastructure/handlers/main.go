package handlers

import (
	"gitub.com/jibaru/ddd-golang-mongodb-template/internal/core/version/application/usecases"
	"net/http"
)

type GetVersionHandler interface {
	Execute(response http.ResponseWriter, request *http.Request)
}

func NewGetVersionHandler() GetVersionHandler {
	return &getVersionHandler{
		getVersionUseCase: usecases.NewGetVersionUseCase(),
	}
}
