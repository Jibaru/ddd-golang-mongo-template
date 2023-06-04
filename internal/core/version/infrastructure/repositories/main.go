package repositories

import (
	"gitub.com/jibaru/ddd-golang-mongodb-template/internal/core/version/domain/repositories"
)

func NewMongoVersionRepository() repositories.VersionRepository {
	return &mongoVersionRepository{}
}
