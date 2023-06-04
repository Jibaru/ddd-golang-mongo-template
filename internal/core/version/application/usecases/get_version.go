package usecases

import (
	"gitub.com/jibaru/ddd-golang-mongodb-template/internal/core/version/domain/entities"
	"gitub.com/jibaru/ddd-golang-mongodb-template/internal/core/version/domain/repositories"
)

type getVersionUseCase struct {
	versionRepository repositories.VersionRepository
}

func (v *getVersionUseCase) Execute() (*entities.Version, error) {
	return v.versionRepository.FindLatest()
}
