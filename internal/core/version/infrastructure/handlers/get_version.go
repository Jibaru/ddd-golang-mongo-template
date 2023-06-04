package handlers

import (
	"encoding/json"
	"gitub.com/jibaru/ddd-golang-mongodb-template/internal/core/version/application/usecases"
	"gitub.com/jibaru/ddd-golang-mongodb-template/internal/core/version/infrastructure/dtos"
	"gitub.com/jibaru/ddd-golang-mongodb-template/pkg/logger"
	"log"
	"net/http"
)

type getVersionHandler struct {
	getVersionUseCase usecases.GetVersionUseCase
}

func (h *getVersionHandler) Execute(response http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		_, err := response.Write([]byte("unsupported"))
		if err != nil {
			log.Print(err)
		}
	}

	version, err := h.getVersionUseCase.Execute()
	if err != nil {
		_, err = response.Write([]byte("cannot get version"))
		if err != nil {
			log.Print(err)
		}
	}

	versionDto := dtos.NewVersionDto(version.Value())

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(200)
	err = json.NewEncoder(response).Encode(versionDto)
	if err != nil {
		logger.Log(err.Error())
	}
}
