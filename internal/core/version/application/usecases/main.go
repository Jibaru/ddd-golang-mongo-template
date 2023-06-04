package usecases

import (
	"gitub.com/jibaru/ddd-golang-mongodb-template/internal/core/version/domain/entities"
	"gitub.com/jibaru/ddd-golang-mongodb-template/internal/core/version/infrastructure/repositories"
)

type GetVersionUseCase interface {
	Execute() (*entities.Version, error)
}

func NewGetVersionUseCase() GetVersionUseCase {
	return &getVersionUseCase{
		versionRepository: repositories.NewMongoVersionRepository(),
	}
}
