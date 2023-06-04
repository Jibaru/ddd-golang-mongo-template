package repositories

import (
	"gitub.com/jibaru/ddd-golang-mongodb-template/internal/core/version/domain/entities"
)

type VersionRepository interface {
	FindLatest() (*entities.Version, error)
}
